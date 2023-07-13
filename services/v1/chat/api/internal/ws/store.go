package ws

import "sync"

// 存储在线的客服和游客
var OnlineKefu = NewOnlineUser()
var OnlineVisitor = NewOnlineUser()

type OnlineUser struct {
	Store   map[int64]int //map[id]count每个用户的好友数
	RwMutex *sync.RWMutex
}

func NewOnlineUser() *OnlineUser {
	return &OnlineUser{
		Store:   map[int64]int{},
		RwMutex: &sync.RWMutex{},
	}
}

func (l *OnlineUser) Ids() []int64 {
	l.RwMutex.RLock()
	defer l.RwMutex.RUnlock()
	var res []int64
	for i := range l.Store {
		res = append(res, i)
	}
	return res
}

func (l *OnlineUser) Set(key int64, val int) {
	l.RwMutex.Lock()
	defer l.RwMutex.Unlock()
	l.Store[key] = val
}

func (l *OnlineUser) Del(id int64) {
	l.RwMutex.Lock()
	defer l.RwMutex.Unlock()
	delete(l.Store, id)
}

func (l *OnlineUser) Get(id int64) int {
	l.RwMutex.RLock()
	defer l.RwMutex.RUnlock()
	return l.Store[id]
}

func (l *OnlineUser) Exist(id int64) bool {
	l.RwMutex.RLock()
	defer l.RwMutex.RUnlock()
	_, ok := l.Store[id]
	return ok
}

func (l *OnlineUser) Len() int {
	l.RwMutex.RLock()
	defer l.RwMutex.RUnlock()
	return len(l.Store)
}

// 分配一个负载最小的用户出来
func (l *OnlineUser) AssignKefu() int64 {
	// 前提是有客服在线
	l.RwMutex.RLock()
	defer l.RwMutex.RUnlock()
	if len(l.Store) == 0 {
		// 没有客服在线，表示未分配状态，应该执行消息自动回复机制
		return 0
	}
	// 获取最小值
	var minCount int
	var kefuId int64
	for id, count := range l.Store {
		if minCount >= count {
			minCount = count
			kefuId = id
		}
	}
	return kefuId
}
