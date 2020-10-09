package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

func printCommand(p player) command {
	return command{
		execute: func() {
			fmt.Print(p.x)
		},

		undo: func() {
			fmt.Print("undo")
		},
	}
}

type player struct {
	x int
}

type command struct {
	execute, undo func()
}

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	f, err := os.Create(*cpuprofile)
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	reader := bufio.NewReader(os.Stdin)
	p := player{
		x: 8,
	}

	l:
	for {
		char, _, _ := reader.ReadRune()
		cmd := printCommand(p)
		switch char {
		case 'a':
			cmd.execute()
			break
		case 'b':
			cmd.undo()
			break
		default:
			fmt.Print("wrong input: " + string(char))
			break l
		}
	}
	defer pprof.StopCPUProfile()

}