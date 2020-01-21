package monitor

import (
	"os"
	"runtime"
	"runtime/pprof"
)

func StartCpuProf() {
	f, err := os.Create("cpu.prof")
	if err != nil {
		panic(err)
	}

	if err := pprof.StartCPUProfile(f); err != nil {
		panic(err)
	}
}

func StopCpuProf() {
	pprof.StopCPUProfile()
}

func StartGc() {
	runtime.GC()
}

func StartMemProf() {
	f, err := os.Create("memory.prof")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := pprof.WriteHeapProfile(f); err != nil {
		panic(err)
	}
}