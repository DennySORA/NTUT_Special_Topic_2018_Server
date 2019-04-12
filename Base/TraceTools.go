package Base

import (
	"os"
	"runtime/trace"
)

func TraceInit(returnerr chan error) *os.File {
	if file, err := os.Create(".//log//trace.out"); err != nil {
		Error.Println(err)
		returnerr <- err
		return nil
	} else if err = trace.Start(file); err != nil {
		Error.Println(err)
		returnerr <- err
		return nil
	} else {
		return file
	}
}
