package sys

import (
	"sync"
)

var (
	g_wait_quit sync.WaitGroup
)

func AddRoutine() {
	g_wait_quit.Add(1)
}

func DelRoutine() {
	g_wait_quit.Done()
}

func WaitAllRoutineQuit() {
	g_wait_quit.Wait()
}
