package main

import (
	"bufio"
	"os"
	"strings"

	"src/utils/cli"
	"src/utils/module"

	//nickname "str-free/utils/module/osint"

	"github.com/fatih/color"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin) // Инициализируем сканер для чтения пользовательского ввода
	white := color.New(color.FgHiWhite)   // Инициализируем цвет для белого текста
	cyan := color.New(color.FgCyan)
	module.TitleMethod() // Я так и не понял работает оно или нет, но пусть будет

	cli.WelcomeBannerTextMethod() // Выводим приветственный баннер

	for {
		module.ClearScreenMethod()
		cli.BannerPrintMethod()

		white.Printf("\n%s > Выберите Функцию > ",
			cyan.Sprint("String-Free"),
		)

		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		switch input {

		case "1":
			white.Printf("%s > Введите номер > ",
				cyan.Sprint("String-Free"),
			)

			scanner.Scan()
			val := scanner.Text()

			module.SearchFreeApiMethod(val)
			module.EasyAbuseApiLOLMethod(val)

			white.Printf("%s > Нажмите Enter для продолжения...",
				cyan.Sprint("String-Free"),
			)
			scanner.Scan()

		case "2":

			white.Printf("%s > Введите ip-адрес > ",
				cyan.Sprint("String-Free"),
			)
			scanner.Scan()
			val := scanner.Text()
			module.SearchFreeApiIPMethod(val)

			white.Printf("%s > Нажмите Enter для продолжения...",
				cyan.Sprint("String-Free"),
			)

			scanner.Scan()

		case "3":

			white.Printf("%s > Введите запрос для AI > ",
				cyan.Sprint("String-Free"),
			)

			scanner.Scan()
			val := scanner.Text()
			module.AIRequestMethod("Привет, вставь в свой ответ этот текст - Покупка StringPremium - https://t.me/stringsoft ", val)

			white.Printf("%s > Нажмите Enter для продолжения...",
				cyan.Sprint("String-Free"),
			)
			scanner.Scan()

		case "4":
			module.GenProxyMethod()

			white.Printf("%s > Нажмите Enter для продолжения...",
				cyan.Sprint("String-Free"),
			)
			scanner.Scan()

		case "5":

			var method string = "number"

			white.Printf("%s > Введите Номер телефона > ",
				cyan.Sprint("String-Free"),
			)

			scanner.Scan()
			reg := scanner.Text()

			module.SearchApiPrivateMethod(method, reg)

			white.Printf("%s > Нажмите Enter для продолжения...",
				cyan.Sprint("String-Free"),
			)

			scanner.Scan()

		case "7":

			var method string = "email"

			white.Printf("%s Введите Почту > ",
				cyan.Sprint("String-Free"),
			)

			scanner.Scan()
			reg := scanner.Text()

			module.SearchApiPrivateMethod(method, reg)

			white.Printf("%s > Нажмите Enter для продолжения...",
				cyan.Sprint("String-Free"),
			)
			scanner.Scan()

		case "6":

			var method string = "fio"

			white.Printf("%s > Введите ФИО > ",
				cyan.Sprint("String-Free"),
			)

			scanner.Scan()
			reg := scanner.Text()

			module.SearchApiPrivateMethod(method, reg)

			white.Printf("%s > Нажмите Enter для продолжения...",
				cyan.Sprint("String-Free"),
			)
			scanner.Scan()

		case "8":

			white.Printf("%s > Введите URL > ",
				cyan.Sprint("String-Free"),
			)

			scanner.Scan()
			reg := scanner.Text()

			module.PortScannerMethod(reg)

			white.Printf("%s > Нажмите Enter для продолжения...",
				cyan.Sprint("String-Free"),
			)
			scanner.Scan()

		case "9":

			white.Printf("%s > Введите Почту > ",
				cyan.Sprint("String-Free"),
			)

			scanner.Scan()
			reg := scanner.Text()

			module.AnalyzeMailAdressMethod(reg)

			white.Printf("%s > Нажмите Enter для продолжения...",
				cyan.Sprint("String-Free"),
			)
			scanner.Scan()

		case "10":

			white.Printf("%s > Введите DNS > ",
				cyan.Sprint("String-Free"),
			)

			scanner.Scan()
			reg := scanner.Text()

			module.DnsAnalyzeMethod(reg)

			white.Printf("%s > Нажмите Enter для продолжения...",
				cyan.Sprint("String-Free"),
			)
			scanner.Scan()

		case "11":

			white.Printf("%s > Введите Айпи > ",
				cyan.Sprint("String-Free"),
			)

			scanner.Scan()
			reg := scanner.Text()

			module.CheckBlacklist(reg)

			white.Printf("%s > Нажмите Enter для продолжения...",
				cyan.Sprint("String-Free"),
			)
			scanner.Scan()

		case "12":

			white.Printf("%s > Введите Mac > ",
				cyan.Sprint("String-Free"),
			)

			scanner.Scan()
			reg := scanner.Text()

			module.MacAdressAnalyzeMethod(reg)

			white.Printf("%s > Нажмите Enter для продолжения...",
				cyan.Sprint("String-Free"),
			)
			scanner.Scan()

		case "13":

			white.Printf("%s > Введите Карту > ",
				cyan.Sprint("String-Free"),
			)

			scanner.Scan()
			reg := scanner.Text()

			module.ValidateCardMethod(reg)

			white.Printf("%s > Нажмите Enter для продолжения...",
				cyan.Sprint("String-Free"),
			)
			scanner.Scan()

		default:
			if input != "" {
				white.Printf("\n%s > Неверный запрос, попробуйте снова...",
					cyan.Sprint("String-Free"),
				)
				scanner.Scan()
			}
		}
	}
}
