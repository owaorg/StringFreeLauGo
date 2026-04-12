package module

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

func GenProxyMethod() {

	white := color.New(color.FgHiWhite)
	//cyan := color.New(color.FgHiCyan)

	rand.Seed(time.Now().UnixNano())
	fmt.Println("")
	for i := 0; i < 60; i++ {

		if i < 60 {
			white.Printf(fmt.Sprint(i)+" "+"Прокси: %d.%d.%d.%d:%d\n",
				rand.Intn(256),
				rand.Intn(256),
				rand.Intn(256),
				rand.Intn(256),
				rand.Intn(65536))
		}

	}
}
