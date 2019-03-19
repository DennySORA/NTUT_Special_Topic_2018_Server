package Command

import (
	"SORA/Base"
	"SORA/Server"
	"os"
	"os/signal"
	"syscall"
)

func Start() {
	stopArg := []interface{}{}
	cherr := make(chan error)
	Base.LogInit(cherr)
	stopArg = append(stopArg, Base.TraceInit(cherr))
	go Server.StartGraphQLServer()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, os.Kill, syscall.SIGTERM)
	for {
		select {
		case <-stop:
			Stopfunc()
			panic(nil)
		case returnerr := <-cherr:
			Base.Error.Panic(returnerr)
		}
	}
}

func Stopfunc(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case *os.File:
			arg.(*os.File).Close()
		}
	}
}
