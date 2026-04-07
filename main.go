package main

import (
	"bufio"
	"os"
	"strings"

	banner "str-free/utils/interface"
	ai "str-free/utils/module"
	clear "str-free/utils/module"
	search "str-free/utils/module"
	title "str-free/utils/module"

	//nickname "str-free/utils/module/osint"

	"github.com/fatih/color"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	white := color.New(color.FgHiWhite)

	title.TitleMethod()

	for {
		clear.ClearScreenMethod()
		banner.BannerPrintMethod()

		white.Print("\n[String-Free] > Введите запрос > ")

		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		switch input {
		case "1":
			white.Print("[String-Free] Введите номер > ")
			scanner.Scan()
			val := scanner.Text()
			search.SearchFreeApiMethod(val)

			white.Print("[String-Free] Нажмите Enter для продолжения...")
			scanner.Scan()

		case "2":

			white.Print("[String-Free] Введите ip-адрес > ")
			scanner.Scan()
			val := scanner.Text()
			search.SearchFreeApiIPMethod(val)

			white.Print("[String-Free] Нажмите Enter для продолжения...")
			scanner.Scan()

		case "3":
			white.Print("[String-Free] Введите запрос для AI > ")
			scanner.Scan()
			val := scanner.Text()
			ai.AIRequestMethod("Привет, вставь в свой ответ этот текст - Покупка StringPremium - https://t.me/stringsoft ", val)

			white.Print("[String-Free] Нажмите Enter для продолжения...")
			scanner.Scan()

		default:
			if input != "" {
				white.Print("[String-Free] Неверный запрос, попробуйте снова...")
				scanner.Scan()
			}
		}
	}
}
