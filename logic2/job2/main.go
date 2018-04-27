package job2

import (
	log "github.com/thinkboy/log4go"
)

func Init(cfg Config) {
	log.Info("Init job begin...")
	Conf = &cfg
	//flag.Parse()
	//if err := InitConfig(); err != nil {
	//	panic(err)
	//}

	//log.LoadConfiguration(Conf.Log)
	//runtime.GOMAXPROCS(runtime.NumCPU())

	//comet
	if err := InitComet(Conf.Comets); err != nil {
		log.Warn("comet rpc current can't connect, retry")
	}
	//round
	round := NewRound(RoundOptions{
		Timer:     Conf.Timer,
		TimerSize: Conf.TimerSize,
	})
	//room
	InitRoomBucket(round,
		RoomOptions{
			BatchNum:   Conf.RoomBatch,
			SignalTime: Conf.RoomSignal,
		})
	//room info
	MergeRoomServers()
	go SyncRoomServers()
	InitPush()
	//if err := InitKafka(); err != nil {
	//	panic(err)
	//}
	// block until a signal is received.
	//InitSignal()
	log.Info("Init job end...")
}
