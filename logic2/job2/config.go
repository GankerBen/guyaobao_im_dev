package job2

import (
	"time"

	"github.com/Terry-Mao/goconf"
)

var (
	gconf *goconf.Config
	Conf  *Config
)

type Config struct {
	//Log string `goconf:"base:log"`
	//ZKAddrs      []string         `goconf:"kafka:zookeeper.list:,"`
	//ZKRoot       string           `goconf:"kafka:zkroot"`
	//KafkaTopic   string           `goconf:"kafka:topic"`
	Comets       map[int32]string `goconf:"-"`
	PushChan     int              `goconf:"push:chan"`
	PushChanSize int              `goconf:"push:chan.size"`
	// timer
	Timer     int `goconf:"timer:num"`
	TimerSize int `goconf:"timer:size"`
	// room
	RoomBatch  int           `goconf:"room:batch"`
	RoomSignal time.Duration `goconf:"room:signal:time"`
}

func NewConfig() *Config { return &Config{} }
