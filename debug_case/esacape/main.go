package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

type JR struct {
	Name string          `json:"name"`
	Data json.RawMessage `json:"data"`
}

func main() {
	fmt.Println("vim-go")

	var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
	var memprofile = flag.String("memprofile", "", "write cpu profile to `file`")

	flag.Parse()

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create Cpu profile:", err)
		}

		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start Cpu profile:", err)
		}

		defer pprof.StopCPUProfile()
	}

	var cm = make(map[string]JR)
	var tmp string
	for i := 0; i < 100000; i++ {
		tmp = fmt.Sprintf("user_%d", i)
		cm[tmp] = JR{Name: tmp, Data: []byte(`{"kk":9}`)}
	}
	fmt.Println(len(cm))

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}

		runtime.GC()

		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile:", err)
		}
	}
}
