package x

import (
	"testing"
)

// func Test_xLog(t *testing.T) {
// 	var err error
//	SetNagle(true)
// 	defer FiniX()

// 	quit := make(chan int)
// 	if err != nil {
// 		return
// 	}

// 	go func() {
// 		for i := 0; i < 100; i++ {
// 			LogInfoF("show me the money:%d", i)
// 			str := "show me the info"
// 			LogInfo(str)
// 			LogInfoF("skdsd%s", "sdfsfdk")
// 			LogErr(str)
// 			DumpInfoF("show me the dump:%d", i)

// 		}
// 		quit <- 0
// 	}()

// 	<-quit
// }

func Test_xLogNoFiniX(t *testing.T) {
	var err error

	quit := make(chan int)
	if err != nil {
		return
	}

	go func() {
		for i := 0; i < 100; i++ {
			LogInfoF("show me the money:%d", i)
			str := "show me the info"
			LogInfo(str)
			LogInfoF("skdsd%s", "sdfsfdk")
			LogErr(str)

		}
		quit <- 0
	}()

	<-quit
}
