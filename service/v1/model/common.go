package model

import (
	"context"
	"errors"
	"fmt"
	"mixbench/common/utils"
	"strings"
	"time"

	sq "github.com/Masterminds/squirrel"
)

func VerifyUser(user *User) error {
	if user == nil {
		return errors.New("user is nil")
	}
	if user.Status == -1 {
		return errors.New("用户已被封禁")
	}
	if user.Status == 0 {
		return errors.New("用户未被授权")
	}
	if user.Status != 1 {
		return errors.New("用户状态异常")
	}
	return nil
}

func VerifyActivity(activity *Activity) error {
	if activity == nil {
		return errors.New("activity is nil")
	}
	if activity.Status == 0 {
		return errors.New("活动已关闭")
	}
	if activity.Status != 1 {
		return errors.New("活动状态异常")
	}
	return nil
}

type ObjMaping struct {
	User       map[int64]*User
	PrizesPool map[int64]*PrizesPool
	Prize      map[int64]*Prize
	PrizeGrade map[int64]*PrizeGrade
	Activity   map[int64]*Activity
	Project    map[int64]*Project
}

type CommonModel struct {
	ProjectModel  ProjectModel
	ActivityModel ActivityModel
	PrizeModel    PrizeModel
	UserModel     UserModel

	LuckerModel     LuckerModel
	PrizesPoolModel PrizesPoolModel
	UsersPoolModel  UsersPoolModel
	PrizeGradeModel PrizeGradeModel

	//临时map，用于缓存已经查询过的数据，避免频繁调用redis和Mysql查询
	storage *ObjMaping
}

type LuckerDetailsInfo struct {
	Lucker *Lucker //中奖信息

	User *User //中奖用户

	PrizesPool *PrizesPool //奖品池的信息

	Prize *Prize //奖品信息

	PrizeGrade *PrizeGrade //奖品级别

	Activity *Activity //活动信息

	Project *Project //项目信息
}

type LuckerDetailsOptions struct {
	PageSize int64
	Page     int64
	Keyword  string
	//lucker表字段
	Lucker *Lucker
	//other表字段
	UserName     string
	ProjectId    int64
	ActivityId   int64
	ActivityName string
	PrizeId      int64
	PrizeName    string
	PrizeGradeId int64
}

func NewLuckerDetailsOptions() *LuckerDetailsOptions {
	return &LuckerDetailsOptions{
		PageSize:     -1,
		Page:         -1,
		Keyword:      "",
		Lucker:       &Lucker{Id: -1, Uuid: "", UserId: -1, PrizesPoolId: -1, Status: -1, LuckDate: time.Time{}, Ip: ""},
		UserName:     "",
		ProjectId:    -1,
		ActivityId:   -1,
		ActivityName: "",
		PrizeId:      -1,
		PrizeName:    "",
		PrizeGradeId: -1,
	}
}

// FindsLuckerDetails 查询订单详细，不支持keyword查询
func (m *CommonModel) FindsLuckerDetails(ctx context.Context, options *LuckerDetailsOptions) (resp []*LuckerDetailsInfo, total int64, err error) {
	if options.PageSize == 0 {
		return
	}
	//临时map，用于缓存已经查询过的数据，避免频繁调用redis和Mysql查询
	if m.storage == nil {
		m.storage = &ObjMaping{
			User:       map[int64]*User{},
			PrizesPool: map[int64]*PrizesPool{},
			Prize:      map[int64]*Prize{},
			PrizeGrade: map[int64]*PrizeGrade{},
			Activity:   map[int64]*Activity{},
			Project:    map[int64]*Project{},
		}
	}

	// fmt.Printf("options：%+v", options.Lucker)

	var luckerList []*Lucker
	from := fmt.Sprintf("%s AS l", m.LuckerModel.TableName())
	diyLuckerRowsArr := strings.Split(luckerRows, ",")
	for i := range diyLuckerRowsArr {
		diyLuckerRowsArr[i] = fmt.Sprintf("l.%s", diyLuckerRowsArr[i])
	}
	diyLuckerRows := strings.Join(diyLuckerRowsArr, ",")
	luckerQuery := sq.Select(diyLuckerRows).From(from).Where("l.is_deleted = 0")
	luckerCount := sq.Select("COUNT(*)").From(from).Where("l.is_deleted = 0")

	var eq = make(sq.Eq)
	if options == nil {
		luckerList, _, err = m.LuckerModel.FindList(ctx, -1, -1, "", nil)
		if err != nil {
			return
		}
	} else {
		//lucker表的条件查询构造-start：user_id、prizes_pool_id、status、luck_date、ip
		if options.Lucker != nil {
			//user_id
			if utils.IdValid(options.Lucker.UserId) { //user_id与user_name只能取一
				eq["l.user_id"] = options.Lucker.UserId
			}

			//prizes_pool_id
			if utils.IdValid(options.Lucker.PrizesPoolId) {
				eq["l.prizes_pool_id"] = options.Lucker.PrizesPoolId
			}

			//status
			if options.Lucker.Status != -1 {
				eq["l.status"] = options.Lucker.Status
			}

			//luck_date
			if !options.Lucker.LuckDate.IsZero() {
				eq["DATE_FORMAT(l.`luck_date`,'%Y-%m-%d')"] = options.Lucker.LuckDate.Format("2006-01-02")
			}

			//ip
			if options.Lucker.Ip != "" {
				eq["l.ip"] = options.Lucker.Ip
			}
		}

		//user_name
		options.UserName = strings.TrimSpace(options.UserName)
		if utils.NameValid(options.UserName) {
			if options.Lucker == nil || (options.Lucker != nil && !utils.IdValid(options.Lucker.UserId)) {
				var userIds []int64
				userList, e := m.UserModel.FindsByName(ctx, options.UserName)
				if e != nil {
					err = e
					return
				}
				if len(userList) == 0 {
					//不存在该用户
					return
				}
				for _, user := range userList {
					userIds = append(userIds, user.Id)
					m.storage.User[user.Id] = user
				}
				eq["l.user_id"] = userIds
			}
		}

		//lucker表的条件查询构造-end

		//other表的条件查询构造-start：project_id、activity_id、prize_id、prize_grade_id
		if utils.IdValid(options.ProjectId) || utils.IdValid(options.ActivityId) || utils.IdValid(options.PrizeId) || utils.IdValid(options.PrizeGradeId) || options.PrizeName != "" || options.ActivityName != "" {
			joinPrizesPool := fmt.Sprintf("%s AS pp ON l.prizes_pool_id = pp.id", m.PrizesPoolModel.TableName())
			luckerQuery = luckerQuery.LeftJoin(joinPrizesPool).Where("pp.is_deleted = 0")
			luckerCount = luckerCount.LeftJoin(joinPrizesPool).Where("pp.is_deleted = 0")

			options.ActivityName = strings.TrimSpace(options.ActivityName)
			if utils.IdValid(options.ProjectId) || utils.IdValid(options.ActivityId) || options.ActivityName != "" {
				joinActivity := fmt.Sprintf("%s AS act ON pp.activity_id = act.id", m.ActivityModel.TableName())
				luckerQuery = luckerQuery.LeftJoin(joinActivity).Where("act.is_deleted = 0")
				luckerCount = luckerCount.LeftJoin(joinActivity).Where("act.is_deleted = 0")
				if utils.IdValid(options.ProjectId) { //某项目下
					joinProject := fmt.Sprintf("%s AS pjt ON act.project_id = pjt.id", m.ProjectModel.TableName())
					luckerQuery = luckerQuery.LeftJoin(joinProject).Where("pjt.is_deleted = 0")
					luckerCount = luckerCount.LeftJoin(joinProject).Where("pjt.is_deleted = 0")
					eq["pjt.id"] = options.ProjectId
				}
				if utils.IdValid(options.ActivityId) { //某活动下
					eq["act.id"] = options.ActivityId
				} else if options.ActivityName != "" { //某些活动下
					luckerQuery = luckerQuery.Where("act.name LIKE ?", fmt.Sprint("%", options.ActivityName, "%"))
					luckerCount = luckerCount.Where("act.name LIKE ?", fmt.Sprint("%", options.ActivityName, "%"))
				}
			}

			options.PrizeName = strings.TrimSpace(options.PrizeName)
			if utils.IdValid(options.PrizeId) || options.PrizeName != "" {
				joinPrize := fmt.Sprintf("%s AS pz ON pp.prize_id = pz.id", m.PrizeModel.TableName())
				luckerQuery = luckerQuery.LeftJoin(joinPrize).Where("pz.is_deleted = 0")
				luckerCount = luckerCount.LeftJoin(joinPrize).Where("pz.is_deleted = 0")
				if utils.IdValid(options.PrizeId) {
					eq["pz.id"] = options.PrizeId
				} else if options.PrizeName != "" {
					luckerQuery = luckerQuery.Where("pz.name LIKE ?", fmt.Sprint("%", options.PrizeName, "%"))
					luckerCount = luckerCount.Where("pz.name LIKE ?", fmt.Sprint("%", options.PrizeName, "%"))
				}
			}

			if utils.IdValid(options.PrizeGradeId) {
				joinPrizeGrade := fmt.Sprintf("%s AS pg ON pp.grade_id = pg.id", m.PrizeGradeModel.TableName())
				luckerQuery = luckerQuery.LeftJoin(joinPrizeGrade).Where("pg.is_deleted = 0")
				luckerCount = luckerCount.LeftJoin(joinPrizeGrade).Where("pg.is_deleted = 0")
				eq["pg.id"] = options.PrizeGradeId
			}
		}
		//other表的条件查询构造-end

		luckerQuery = luckerQuery.Where(eq)
		luckerCount = luckerCount.Where(eq)
	}

	isPageQuery := options.PageSize > 0 && options.Page > 0
	if isPageQuery {
		luckerQuery = luckerQuery.Offset(uint64((options.Page - 1) * options.PageSize)).Limit(uint64(options.PageSize))
		//总数量查询
		luckerCountSql, countArgs, e := luckerCount.ToSql()
		if e != nil {
			err = e
			return
		}
		err = m.LuckerModel.Conn().QueryRowNoCacheCtx(ctx, &total, luckerCountSql, countArgs...)
		if err != nil {
			return
		}
	}

	//record查询
	luckerQuerySql, queryArgs, e := luckerQuery.ToSql()
	if e != nil {
		err = e
		return
	}
	err = m.LuckerModel.Conn().QueryRowsNoCacheCtx(ctx, &luckerList, luckerQuerySql, queryArgs...)
	if err != nil {
		return
	}

	//如果不是页查询
	if !isPageQuery {
		total = int64(len(luckerList))
	}

	//临时仓储，避免频繁的查询

	//查询详细
	resp, err = m.QueryLuckerDetailsInfo2(ctx, luckerList)

	//释放内存
	m.storage = nil
	return
}

// QueryLuckerDetailsInfo2 查询lucker详细，使用了go内存做数据缓存，方案二
// 用空间换时间
func (m *CommonModel) QueryLuckerDetailsInfo2(ctx context.Context, luckerList []*Lucker) (resp []*LuckerDetailsInfo, err error) {
	for _, lucker := range luckerList {
		var user *User
		var prizesPool *PrizesPool
		var activity *Activity
		var project *Project
		var prizeGrade *PrizeGrade
		var prize *Prize

		//查询用户
		if v, ok := m.storage.User[lucker.UserId]; ok && v != nil {
			user = v
		} else {
			user, err = m.UserModel.FindOne(ctx, lucker.UserId)
			if err != nil && err != ErrNotFound {
				return
			}
			if err != ErrNotFound && user != nil {
				m.storage.User[lucker.UserId] = user
			}
			// if err == ErrNotFound {
			// 	//不存在该用户，则跳过
			// 	continue //用户不存在，则不需要导出该记录，可以跳过
			// }
		}

		//查询奖品池
		if v, ok := m.storage.PrizesPool[lucker.PrizesPoolId]; ok && v != nil {
			prizesPool = v
		} else {
			prizesPool, err = m.PrizesPoolModel.QueryOne(ctx, lucker.PrizesPoolId)
			if err != nil && err != ErrNotFound {
				return
			}
			if err != ErrNotFound && prizesPool != nil {
				m.storage.PrizesPool[lucker.PrizesPoolId] = prizesPool
			}
			if err == ErrNotFound {
				// //奖品池不存在，则无法查询剩余信息，直接跳过
				// continue
				goto SAVE
			}
		}

		//通过奖品池查询获奖等级
		if v, ok := m.storage.PrizeGrade[prizesPool.GradeId]; ok && v != nil {
			prizeGrade = v
		} else {
			prizeGrade, err = m.PrizeGradeModel.FindOne(ctx, prizesPool.GradeId)
			if err != nil && err != ErrNotFound {
				return
			}
			if err != ErrNotFound && prizeGrade != nil {
				m.storage.PrizeGrade[prizesPool.GradeId] = prizeGrade
			}
		}

		//通过奖品池查询奖品
		if v, ok := m.storage.Prize[prizesPool.PrizeId]; ok && v != nil {
			prize = v
		} else {
			prize, err = m.PrizeModel.FindOne(ctx, prizesPool.PrizeId)
			if err != nil && err != ErrNotFound {
				return
			}
			if err != ErrNotFound && prize != nil {
				m.storage.Prize[prizesPool.PrizeId] = prize
			}
		}

		//通过奖品池查询活动
		if v, ok := m.storage.Activity[prizesPool.ActivityId]; ok && v != nil {
			activity = v
		} else {
			activity, err = m.ActivityModel.FindOne(ctx, prizesPool.ActivityId)
			if err != nil && err != ErrNotFound {
				return
			}
			if err != ErrNotFound && activity != nil {
				m.storage.Activity[prizesPool.ActivityId] = activity
			}
			if err == ErrNotFound {
				goto SAVE
			}
		}

		//通过活动查询项目
		if v, ok := m.storage.Project[activity.ProjectId]; ok && v != nil {
			project = v
		} else {
			project, err = m.ProjectModel.FindOne(ctx, activity.ProjectId)
			if err != nil && err != ErrNotFound {
				return
			}
			if err != ErrNotFound && project != nil {
				m.storage.Project[activity.ProjectId] = project
			}
		}

	SAVE:
		//构建保存
		luckerDetailsInfo := &LuckerDetailsInfo{
			Lucker:     lucker,
			User:       user,
			PrizesPool: prizesPool,
			Prize:      prize,
			PrizeGrade: prizeGrade,
			Activity:   activity,
			Project:    project,
		}
		resp = append(resp, luckerDetailsInfo)
	}
	return
}

// QueryLuckerDetailsInfo1 查询lucker详细，方案一
func (m *CommonModel) QueryLuckerDetailsInfo1(ctx context.Context, luckerList []*Lucker) (resp []*LuckerDetailsInfo, err error) {
	for _, lucker := range luckerList {
		//查询用户
		user, e := m.UserModel.FindOne(ctx, lucker.UserId)
		if e != nil && e != ErrNotFound {
			err = e
			return
		}

		//查询奖品池
		prizesPool, e := m.PrizesPoolModel.FindOne(ctx, lucker.PrizesPoolId)
		if e != nil && e != ErrNotFound {
			err = e
			return
		}

		//查询活动
		activity, e := m.ActivityModel.FindOne(ctx, prizesPool.ActivityId)
		if e != nil && e != ErrNotFound {
			err = e
			return
		}

		//查询奖品
		prize, e := m.PrizeModel.FindOne(ctx, prizesPool.PrizeId)
		if e != nil && e != ErrNotFound {
			err = e
			return
		}

		//查询获奖等级
		prizeGrade, e := m.PrizeGradeModel.FindOne(ctx, prizesPool.GradeId)
		if e != nil && e != ErrNotFound {
			err = e
			return
		}

		//查询项目
		project, e := m.ProjectModel.FindOne(ctx, activity.ProjectId)
		if e != nil && e != ErrNotFound {
			err = e
			return
		}

		//构建保存
		luckerDetailsInfo := &LuckerDetailsInfo{
			Lucker:     lucker,
			User:       user,
			PrizesPool: prizesPool,
			Prize:      prize,
			PrizeGrade: prizeGrade,
			Activity:   activity,
			Project:    project,
		}
		resp = append(resp, luckerDetailsInfo)
	}
	return
}
