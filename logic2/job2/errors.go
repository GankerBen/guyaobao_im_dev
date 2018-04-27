package job2

import (
	"errors"
)

var (
	// comet
	ErrComet = errors.New("comet rpc is not available")
	// room
	ErrRoomFull = errors.New("room proto chan full")
)
