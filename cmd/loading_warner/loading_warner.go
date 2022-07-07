package main

import (
	"fmt"
	"time"

	ps "github.com/mitchellh/go-ps"
)

const (
	message = `



	W   W   W  A   II TTTTTT
	 W W W W  AAA  II   TT
	  W   W  A   A II   TT

	FFFFFF  OOO  RRRRR
	FFFF   O   O RR RR
	FF      OOO  RR  RR

	LL      OOO    A   DDDD  II NN  N GGGG
	LL     O   O  AAA  D   D II N N N G  GG
	LLLLLL  OOO  A   A DDDD  II N  NN GGGGG `
)

func main() {
	fmt.Printf(message)
	time.Sleep(20 * time.Second)
	processEnded := false
	for {
		if processEnded {
			break
		}
		processList, err := ps.Processes()
		if err != nil {
			fmt.Print("\n\nFatal with getting process names. This program is not operated in current version of Windows.")
			for {
				time.Sleep(time.Duration(1<<63 - 1))
			}
		}
		processEnded = true
		for _, val := range processList {
			if val.Executable() == "userinit.exe" {
				processEnded = false
			}
		}
		time.Sleep(2 * time.Second)
	}
}
