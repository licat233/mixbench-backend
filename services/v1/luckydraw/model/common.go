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
	PrizePool  map[int64]*PrizePool
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

	LuckyModel      LuckyModel
	PrizePoolModel  PrizePoolModel
	UserPoolModel   UserPoolModel
	PrizeGradeModel PrizeGradeModel

	//临时map，用于缓存已经查询过的数据，避免频繁调用redis和Mysql查询
	storage *ObjMaping
}

type LuckyDetail struct {
	Lucky *Lucky //中奖信息

	User *User //中奖用户

	PrizePool *PrizePool //奖品池的信息

	Prize *Prize //奖品信息

	PrizeGrade *PrizeGrade //奖品级别

	Activity *Activity //活动信息

	Project *Project //项目信息
}

type LuckyDetailOptions struct {
	PageSize int64
	Page     int64
	Keyword  string
	//lucky表字段
	Lucky *Lucky
	//other表字段
	UserName     string
	ProjectId    int64
	ActivityId   int64
	ActivityName string
	PrizeId      int64
	PrizeName    string
	PrizeGradeId int64
}

func NewLuckyDetailOptions() *LuckyDetailOptions {
	return &LuckyDetailOptions{
		PageSize:     -1,
		Page:         -1,
		Keyword:      "",
		Lucky:        &Lucky{Id: -1, Uuid: "", UserId: -1, PrizePoolId: -1, Status: -1, LuckDate: time.Time{}, Ip: ""},
		UserName:     "",
		ProjectId:    -1,
		ActivityId:   -1,
		ActivityName: "",
		PrizeId:      -1,
		PrizeName:    "",
		PrizeGradeId: -1,
	}
}

// FindsLuckyDetails 查询订单详细，不支持keyword查询
func (m *CommonModel) FindsLuckyDetails(ctx context.Context, options *LuckyDetailOptions) (resp []*LuckyDetail, total int64, err error) {
	if options.PageSize == 0 {
		return
	}
	//临时map，用于缓存已经查询过的数据，避免频繁调用redis和Mysql查询
	if m.storage == nil {
		m.storage = &ObjMaping{
			User:       map[int64]*User{},
			PrizePool:  map[int64]*PrizePool{},
			Prize:      map[int64]*Prize{},
			PrizeGrade: map[int64]*PrizeGrade{},
			Activity:   map[int64]*Activity{},
			Project:    map[int64]*Project{},
		}
	}

	// fmt.Printf("options：%+v", options.Lucky)

	var luckyList []*Lucky
	from := fmt.Sprintf("%s AS l", m.LuckyModel.TableName())
	diyLuckyRowsArr := strings.Split(luckyRows, ",")
	for i := range diyLuckyRowsArr {
		diyLuckyRowsArr[i] = fmt.Sprintf("l.%s", diyLuckyRowsArr[i])
	}
	diyLuckyRows := strings.Join(diyLuckyRowsArr, ",")
	luckyQuery := sq.Select(diyLuckyRows).From(from)
	luckyCount := sq.Select("COUNT(*)").From(from)

	var eq = make(sq.Eq)
	if options == nil {
		luckyList, _, err = m.LuckyModel.FindList(ctx, -1, -1, "", nil)
		if err != nil {
			return
		}
	} else {
		//lucky表的条件查询构造-start：user_id、prize_pool_id、status、luck_date、ip
		if options.Lucky != nil {
			//user_id
			if utils.IdValid(options.Lucky.UserId) { //user_id与user_name只能取一
				eq["l.user_id"] = options.Lucky.UserId
			}

			//prize_pool_id
			if utils.IdValid(options.Lucky.PrizePoolId) {
				eq["l.prize_pool_id"] = options.Lucky.PrizePoolId
			}

			//status
			if options.Lucky.Status != -1 {
				eq["l.status"] = options.Lucky.Status
			}

			//luck_date
			if !options.Lucky.LuckDate.IsZero() {
				eq["DATE_FORMAT(l.`luck_date`,'%Y-%m-%d')"] = options.Lucky.LuckDate.Format("2006-01-02")
			}

			//ip
			if options.Lucky.Ip != "" {
				eq["l.ip"] = options.Lucky.Ip
			}
		}

		//user_name
		options.UserName = strings.TrimSpace(options.UserName)
		if utils.NameValid(options.UserName) {
			if options.Lucky == nil || (options.Lucky != nil && !utils.IdValid(options.Lucky.UserId)) {
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

		//lucky表的条件查询构造-end

		//other表的条件查询构造-start：project_id、activity_id、prize_id、prize_grade_id
		if utils.IdValid(options.ProjectId) || utils.IdValid(options.ActivityId) || utils.IdValid(options.PrizeId) || utils.IdValid(options.PrizeGradeId) || options.PrizeName != "" || options.ActivityName != "" {
			joinPrizePool := fmt.Sprintf("%s AS pp ON l.prize_pool_id = pp.id", m.PrizePoolModel.TableName())
			luckyQuery = luckyQuery.LeftJoin(joinPrizePool)
			luckyCount = luckyCount.LeftJoin(joinPrizePool)

			options.ActivityName = strings.TrimSpace(options.ActivityName)
			if utils.IdValid(options.ProjectId) || utils.IdValid(options.ActivityId) || options.ActivityName != "" {
				joinActivity := fmt.Sprintf("%s AS act ON pp.activity_id = act.id", m.ActivityModel.TableName())
				luckyQuery = luckyQuery.LeftJoin(joinActivity)
				luckyCount = luckyCount.LeftJoin(joinActivity)
				if utils.IdValid(options.ProjectId) { //某项目下
					joinProject := fmt.Sprintf("%s AS pjt ON act.project_id = pjt.id", m.ProjectModel.TableName())
					luckyQuery = luckyQuery.LeftJoin(joinProject)
					luckyCount = luckyCount.LeftJoin(joinProject)
					eq["pjt.id"] = options.ProjectId
				}
				if utils.IdValid(options.ActivityId) { //某活动下
					eq["act.id"] = options.ActivityId
				} else if options.ActivityName != "" { //某些活动下
					luckyQuery = luckyQuery.Where("act.name LIKE ?", fmt.Sprint("%", options.ActivityName, "%"))
					luckyCount = luckyCount.Where("act.name LIKE ?", fmt.Sprint("%", options.ActivityName, "%"))
				}
			}

			options.PrizeName = strings.TrimSpace(options.PrizeName)
			if utils.IdValid(options.PrizeId) || options.PrizeName != "" {
				joinPrize := fmt.Sprintf("%s AS pz ON pp.prize_id = pz.id", m.PrizeModel.TableName())
				luckyQuery = luckyQuery.LeftJoin(joinPrize)
				luckyCount = luckyCount.LeftJoin(joinPrize)
				if utils.IdValid(options.PrizeId) {
					eq["pz.id"] = options.PrizeId
				} else if options.PrizeName != "" {
					luckyQuery = luckyQuery.Where("pz.name LIKE ?", fmt.Sprint("%", options.PrizeName, "%"))
					luckyCount = luckyCount.Where("pz.name LIKE ?", fmt.Sprint("%", options.PrizeName, "%"))
				}
			}

			if utils.IdValid(options.PrizeGradeId) {
				joinPrizeGrade := fmt.Sprintf("%s AS pg ON pp.grade_id = pg.id", m.PrizeGradeModel.TableName())
				luckyQuery = luckyQuery.LeftJoin(joinPrizeGrade)
				luckyCount = luckyCount.LeftJoin(joinPrizeGrade)
				eq["pg.id"] = options.PrizeGradeId
			}
		}
		//other表的条件查询构造-end

		luckyQuery = luckyQuery.Where(eq)
		luckyCount = luckyCount.Where(eq)
	}

	isPageQuery := options.PageSize > 0 && options.Page > 0
	if isPageQuery {
		luckyQuery = luckyQuery.Offset(uint64((options.Page - 1) * options.PageSize)).Limit(uint64(options.PageSize))
		//总数量查询
		luckyCountSql, countArgs, e := luckyCount.ToSql()
		if e != nil {
			err = e
			return
		}
		err = m.LuckyModel.Conn().QueryRowNoCacheCtx(ctx, &total, luckyCountSql, countArgs...)
		if err != nil {
			return
		}
	}

	//record查询
	luckyQuerySql, queryArgs, e := luckyQuery.ToSql()
	if e != nil {
		err = e
		return
	}
	err = m.LuckyModel.Conn().QueryRowsNoCacheCtx(ctx, &luckyList, luckyQuerySql, queryArgs...)
	if err != nil {
		return
	}

	//如果不是页查询
	if !isPageQuery {
		total = int64(len(luckyList))
	}

	//临时仓储，避免频繁的查询

	//查询详细
	resp, err = m.QueryLuckyDetailsInfo2(ctx, luckyList)

	//释放内存
	m.storage = nil
	return
}

// QueryLuckyDetailsInfo2 查询lucky详细，使用了go内存做数据缓存，方案二
// 用空间换时间
func (m *CommonModel) QueryLuckyDetailsInfo2(ctx context.Context, luckyList []*Lucky) (resp []*LuckyDetail, err error) {
	for _, lucky := range luckyList {
		var user *User
		var prizesPool *PrizePool
		var activity *Activity
		var project *Project
		var prizeGrade *PrizeGrade
		var prize *Prize

		//查询用户
		if v, ok := m.storage.User[lucky.UserId]; ok && v != nil {
			user = v
		} else {
			user, err = m.UserModel.FindOne(ctx, lucky.UserId)
			if err != nil && err != ErrNotFound {
				return
			}
			if err != ErrNotFound && user != nil {
				m.storage.User[lucky.UserId] = user
			}
			// if err == ErrNotFound {
			// 	//不存在该用户，则跳过
			// 	continue //用户不存在，则不需要导出该记录，可以跳过
			// }
		}

		//查询奖品池
		if v, ok := m.storage.PrizePool[lucky.PrizePoolId]; ok && v != nil {
			prizesPool = v
		} else {
			prizesPool, err = m.PrizePoolModel.FindOne(ctx, lucky.PrizePoolId)
			if err != nil && err != ErrNotFound {
				return
			}
			if err != ErrNotFound && prizesPool != nil {
				m.storage.PrizePool[lucky.PrizePoolId] = prizesPool
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
		luckyDetailsInfo := &LuckyDetail{
			Lucky:      lucky,
			User:       user,
			PrizePool:  prizesPool,
			Prize:      prize,
			PrizeGrade: prizeGrade,
			Activity:   activity,
			Project:    project,
		}
		resp = append(resp, luckyDetailsInfo)
	}
	return
}

// QueryLuckyDetailsInfo1 查询lucky详细，方案一
func (m *CommonModel) QueryLuckyDetailsInfo1(ctx context.Context, luckyList []*Lucky) (resp []*LuckyDetail, err error) {
	for _, lucky := range luckyList {
		//查询用户
		user, e := m.UserModel.FindOne(ctx, lucky.UserId)
		if e != nil && e != ErrNotFound {
			err = e
			return
		}

		//查询奖品池
		prizesPool, e := m.PrizePoolModel.FindOne(ctx, lucky.PrizePoolId)
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
		luckyDetailsInfo := &LuckyDetail{
			Lucky:      lucky,
			User:       user,
			PrizePool:  prizesPool,
			Prize:      prize,
			PrizeGrade: prizeGrade,
			Activity:   activity,
			Project:    project,
		}
		resp = append(resp, luckyDetailsInfo)
	}
	return
}
