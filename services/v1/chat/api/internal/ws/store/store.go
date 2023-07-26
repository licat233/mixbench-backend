package store

import (
	"sync"
)

// 存储在线的客服和游客

type onlineKefu struct {
	*OnlineUser
}

func NewOnlineKefu() *onlineKefu {
	return &onlineKefu{
		NewOnlineUser(),
	}
}

type onlineVisitor struct {
	*OnlineUser
}

func (l *onlineVisitor) Set(visitorId int64, kefuId int64) {
	l.OnlineUser.Set(visitorId, kefuId)
}

func (l *onlineVisitor) FindsByVal(kefuId int64) (ids []int64) {
	return l.OnlineUser.FindsByVal(kefuId)
}

func NewOnlineVisitor() *onlineVisitor {
	return &onlineVisitor{
		NewOnlineUser(),
	}
}

type OnlineUser struct {
	//对于kefu来说：map[kefuId]count 客服ID-游客数
	//对于visitor来说：map[visitorId]any 游客ID-待定
	Store   map[int64]int64
	RwMutex *sync.RWMutex
}

func NewOnlineUser() *OnlineUser {
	return &OnlineUser{
		Store:   map[int64]int64{},
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

func (l *OnlineUser) Set(key int64, val int64) {
	l.RwMutex.Lock()
	defer l.RwMutex.Unlock()
	l.Store[key] = val
}

func (l *OnlineUser) Del(id int64) {
	l.RwMutex.Lock()
	defer l.RwMutex.Unlock()
	delete(l.Store, id)
}

func (l *OnlineUser) Get(id int64) int64 {
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

func (l *OnlineUser) FindsByVal(val int64) (ids []int64) {
	l.RwMutex.RLock()
	defer l.RwMutex.RUnlock()
	for id, v := range l.Store {
		if v == val {
			ids = append(ids, id)
		}
	}
	return
}
