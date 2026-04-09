package main

import (
	"bufio"
	"os"
	"strings"

	banner "str-free/utils/cli"
	welcomeMethod "str-free/utils/cli"
	ai "str-free/utils/module"
	clear "str-free/utils/module"
	infinity "str-free/utils/module"
	proxygen "str-free/utils/module"
	search "str-free/utils/module"
	title "str-free/utils/module"

	//nickname "str-free/utils/module/osint"

	"github.com/fatih/color"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin) // Инициализируем сканер для чтения пользовательского ввода
	white := color.New(color.FgHiWhite)   // Инициализируем цвет для белого текста

	title.TitleMethod() // Я так и не понял работает оно или нет, но пусть будет

	welcomeMethod.WelcomeBannerTextMethod() // Выводим приветственный баннер

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

		case "4":
			proxygen.GenProxyMethod()

			white.Println("\n[String-Free] Нажмите Enter для продолжения...")
			scanner.Scan()

		case "5":

			var method string = "number"

			white.Print("[String-Free] Введите Номер телефона > ")

			scanner.Scan()
			reg := scanner.Text()

			infinity.SearchApiPrivateMethod(method, reg)

			white.Println("\n[String-Free] Нажмите Enter для продолжения...")
			scanner.Scan()

		case "7":

			var method string = "email"

			white.Print("[String-Free] Введите Почту > ")

			scanner.Scan()
			reg := scanner.Text()

			infinity.SearchApiPrivateMethod(method, reg)

			white.Println("\n[String-Free] Нажмите Enter для продолжения...")
			scanner.Scan()

		case "6":

			var method string = "fio"

			white.Print("[String-Free] Введите ФИО > ")

			scanner.Scan()
			reg := scanner.Text()

			infinity.SearchApiPrivateMethod(method, reg)

			white.Println("\n[String-Free] Нажмите Enter для продолжения...")
			scanner.Scan()

		default:
			if input != "" {
				white.Print("[String-Free] Неверный запрос, попробуйте снова...")
				scanner.Scan()
			}
		}
	}
}
