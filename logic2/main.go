package main

import (
	"flag"
	"goim/libs/perf"
	"goim/logic2/job2"
	"runtime"

	log "github.com/thinkboy/log4go"
)

func main() {
	flag.Parse()
	if err := InitConfig(); err != nil {
		panic(err)
	}
	runtime.GOMAXPROCS(Conf.MaxProc)
	log.LoadConfiguration(Conf.Log)
	defer log.Close()
	log.Info("logic[%s] start: %#v", Ver, Conf)
	perf.Init(Conf.PprofAddrs)

	/**
	Comets       map[int32]string `goconf:"-"`
	PushChan     int              `goconf:"push:chan"`
	PushChanSize int              `goconf:"push:chan.size"`
	// timer
	Timer     int `goconf:"timer:num"`
	TimerSize int `goconf:"timer:size"`
	// room
	RoomBatch  int           `goconf:"room:batch"`
	RoomSignal time.Duration `goconf:"room:signal:time"`
	*/
	jobConf := job2.Config{}
	jobConf.Comets = Conf.Comets
	jobConf.PushChan = Conf.PushChan
	jobConf.PushChanSize = Conf.PushChanSize
	jobConf.Timer = Conf.Timer
	jobConf.TimerSize = Conf.TimerSize
	jobConf.RoomBatch = Conf.RoomBatch
	jobConf.RoomSignal = Conf.RoomSignal

	go job2.Init(jobConf)

	// router rpc
	if err := InitRouter(); err != nil {
		log.Warn("router rpc current can't connect, retry")
	}
	MergeCount()
	go SyncCount()
	// logic rpc
	if err := InitRPC(NewDefaultAuther()); err != nil {
		panic(err)
	}
	if err := InitHTTP(); err != nil {
		panic(err)
	}
	//if err := InitKafka(Conf.KafkaAddrs); err != nil {
	//	panic(err)
	//}
	// block until a signal is received.
	InitSignal()
}
