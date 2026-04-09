package module

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

func GenProxyMethod() {

	white := color.New(color.FgHiWhite)
	cyan := color.New(color.FgHiCyan)

	rand.Seed(time.Now().UnixNano())
	fmt.Println("")
	for i := 0; i < 6; i++ {
		if i == 5 {
			fmt.Print("\n\n")
			fmt.Printf("%s %s",
				cyan.Sprint("[Купи Подписку!]"),
				white.Sprint("Генерация прокси завершена, максимальное количество для генерации 5."),
			)

		} else if i < 5 {
			white.Printf(fmt.Sprint(i)+" "+"Прокси: %d.%d.%d.%d:%d\n",
				rand.Intn(256),
				rand.Intn(256),
				rand.Intn(256),
				rand.Intn(256),
				rand.Intn(65536))
		}

	}
}
