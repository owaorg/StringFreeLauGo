package cli

import (
	_ "embed"

	"github.com/fatih/color"
)

//go:embed banner.string
var banner string

func BannerPrintMethod() {
	//Я вообще не знаю как работает эта штука
	color.New(color.FgCyan).Println(banner)
	cyan := color.New(color.FgCyan)

	println()
	println("Telegram Channel  > " + cyan.Sprint("   https://t.me/stringsoft"))
	println("Devloper String   > " + cyan.Sprint("   @owaqx / string.surf"))
	println("GitHub Repository > " + cyan.Sprint("   https://github/owaorg/StringFreeLauGo"))
	println("Language          > " + cyan.Sprint("   Golang. Full SRC"))

	println()

	println("[ 1 ]" + cyan.Sprint(" Поиск по номеру                    [2x]"))
	println("[ 2 ]" + cyan.Sprint(" Поиск по ip-адресу                 [1x]"))
	println("[ 3 ]" + cyan.Sprint(" AI-Запрос (DeepSeek)               [1x]"))
	println("[ 4 ]" + cyan.Sprint(" Генерация прокси                   [1x]"))
	println("[ 5 ]" + cyan.Sprint(" Поиск по Номеру (InfinityAPI)  [NoWork]"))
	println("[ 6 ]" + cyan.Sprint(" Поиск по ФИО    (InfinityAPI)  [NoWork]"))
	println("[ 7 ]" + cyan.Sprint(" Поиск по Почте  (InfinityAPI)  [NoWork]"))
	println("[ 8 ]" + cyan.Sprint(" Порт Сканнер                       [1x]"))
	println("[ 9 ]" + cyan.Sprint(" Анализ почты                       [1x]"))
	println("[ 10 ]" + cyan.Sprint(" DNS Чекер                         [1x]"))
	println("[ 11 ]" + cyan.Sprint(" Проверка IP на BlackList          [1x]"))
	println("[ 12 ]" + cyan.Sprint(" Анализ MAC-Адреса                 [1x]"))
	println("[ 13 ]" + cyan.Sprint(" Анализ Карты                      [1x]"))
}
