package main

import (
	"fmt"

	banner "str-free/utils/interface"
	clear "str-free/utils/module"
	kryto "str-free/utils/module"
	search "str-free/utils/module"
	title "str-free/utils/module"

	"github.com/fatih/color"
)

func main() {

	var input string
	//var pause string

	black := color.New(color.FgHiWhite)
	title.TitleMethod() //Установка названия окна

	for {
		clear.ClearScreenMethod() //Очистка консоли

		banner.BannerPrintMethod() //Вывод баннера

		black.Print("\nString-Free > Введите запрос > ")

		fmt.Scan(&input)

		if input == "free-sub" {
			kryto.KrytoMethod() //Hello World
		} else {
			search.SearchFreeApiMethod(input)

			black.Println("\nНажмите Enter для продолжения...")
			fmt.Scanln()
			fmt.Scanln()

		}
	}
}
