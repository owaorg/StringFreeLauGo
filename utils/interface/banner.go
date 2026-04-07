package banner

import (
	_ "embed"

	"github.com/fatih/color"
)

//go:embed banner.string
var banner string

func BannerPrintMethod() {
	//Я вообще не знаю как работает эта штука
	color.New(color.FgHiWhite).Println(banner)
	cyan := color.New(color.FgHiCyan)

	println()
	println("Телеграм канал  >" + cyan.Sprint("    https://t.me/stringsoft"))
	println("Разработчик     > " + cyan.Sprint("   @owaqx"))
	println("Репозиторий     > " + cyan.Sprint("   https://github/owaorg/StringFreeLauGo"))

	println()

	println("[1]" + cyan.Sprint(" Поиск по номеру"))
	println("[2]" + cyan.Sprint(" Поиск по ip-адресу"))
	println("[3]" + cyan.Sprint(" AI-Запрос"))

}
