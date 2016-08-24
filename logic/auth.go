package main

import (
	//"goim/libs/define"
	"encoding/json"
	//log "github.com/thinkboy/log4go"
	//"strconv"
	//"strings"
)

// developer could implement "Auth" interface for decide how get userId, or roomId
type Auther interface {
	Auth(token string) (userId int64, roomId int32)
}

type DefaultAuther struct {
	Rid int32 `json:"rid"`
	Uid int64 `json:"uid"`
}

func NewDefaultAuther() *DefaultAuther {
	return &DefaultAuther{}
}

// rid_uid
func (a *DefaultAuther) Auth(token string) (userId int64, roomId int32) {
	var (
		err error
	)

	if err = json.Unmarshal([]byte(token), a); err != nil {
		return
	}

	return a.Uid, a.Rid
}
