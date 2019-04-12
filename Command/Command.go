package Command

import (
	"SORA/Base"
	"SORA/Server"
	"os"
	"os/signal"
	"runtime/trace"
	"syscall"
)

func Start() {
	stopArg := []interface{}{}
	cherr := make(chan error)
	Base.LogInit(cherr)
	if err := Base.InitMongoDB(); err != nil {
		Base.Error.Panicln(err)
	}
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
	trace.Stop()
}
