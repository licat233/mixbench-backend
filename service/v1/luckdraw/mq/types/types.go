package types

const (
	TypeLuckdrawResult = "handle_luckdraw_result"
)

type LuckdrawResultMsg struct {
	IsWin  bool  //是否中奖
	UserId int64 //中奖人id
}
