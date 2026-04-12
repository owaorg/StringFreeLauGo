package module

import (
	"fmt"
	"net"
	"time"

	"github.com/fatih/color"
)

func PortScannerMethod(url string) {

	green := color.New(color.FgGreen)
	red := color.New(color.FgRed)

	ports := []int{80, 443, 21, 22, 3306}

	for _, port := range ports {

		adress_full := net.JoinHostPort(url, fmt.Sprintf("%d", port))
		con, err := net.DialTimeout("tcp", adress_full, 2*time.Second)

		if err != nil {

			fmt.Printf(red.Sprint("\n[CLOSE PORT]")+" Порт закрыт. Порт: %d | Ссылка: %s",

				port,
				url)

		} else {

			fmt.Printf(green.Sprint("\n[OPEN PORT]")+" Порт открыт. Порт: %d | Ссылка: %s",

				port,
				url)

			con.Close()

		}
	}
	fmt.Println("\n")
}
