package types

const (
	TaskTypeLuckydrawResult = "handle_luckydraw_result"
)

type LuckydrawResultMsg struct {
	IsWin  bool  //是否中奖
	UserId int64 //中奖人id
}
