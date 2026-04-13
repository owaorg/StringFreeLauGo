package cli

import (
	_ "embed"

	"github.com/fatih/color"
)

//go:embed banner.string
var banner string

func BannerPrintMethod() {
	//Я вообще не знаю как работает эта штука
	color.New(color.FgHiBlack).Println(banner)
	cyan := color.New(color.FgHiBlack)
	white := color.New(color.FgWhite)

	println()
	println("Telegram Channel  > " + cyan.Sprint("   https://t.me/stringsoft"))
	println("Devloper String   > " + cyan.Sprint("   @owaqx / string.surf"))
	println("GitHub Repository > " + cyan.Sprint("   https://github/owaorg/StringFreeLauGo"))
	println("Language          > " + cyan.Sprint("   Golang. Full SRC"))

	println()

	white.Println("[ 1 ]" + cyan.Sprint(" Поиск по номеру                    [2x]"+white.Sprint("	[ 10 ]"+cyan.Sprint(" Анализ DNS                      [1x]"))))
	white.Println("[ 2 ]" + cyan.Sprint(" Поиск по ip-адресу                 [1x]"+white.Sprint("   [ 11 ]"+cyan.Sprint(" Проверка IP на BlackList        [1x]"))))
	white.Println("[ 3 ]" + cyan.Sprint(" AI-Запрос (DeepSeek)               [1x]"+white.Sprint("	[ 12 ]"+cyan.Sprint(" Анализ MAC-Адреса               [1x]"))))
	white.Println("[ 4 ]" + cyan.Sprint(" Генерация прокси                   [1x]"+white.Sprint("	[ 13 ]"+cyan.Sprint(" Анализ Карты                    [1x]"))))
	white.Println("[ 5 ]" + cyan.Sprint(" Поиск по Номеру (InfinityAPI)  [NoWork]"+white.Sprint("	[ 14 ]"+cyan.Sprint(" SSL Анализ                      [1x]"))))
	white.Println("[ 6 ]" + cyan.Sprint(" Поиск по ФИО    (InfinityAPI)  [NoWork]"+white.Sprint("	[ 15 ]"+cyan.Sprint(" Информация о моем IP            [1x]"))))
	white.Println("[ 7 ]" + cyan.Sprint(" Поиск по Почте  (InfinityAPI)  [NoWork]"+white.Sprint("	[ 16 ]"+cyan.Sprint(" Декодер координат	       [1x]"))))
	white.Println("[ 8 ]" + cyan.Sprint(" Порт Сканнер                       [1x]"+white.Sprint("	[ 17 ]"+cyan.Sprint(" Поиск по ФИО		       [1x]"))))
	white.Println("[ 9 ]" + cyan.Sprint(" Анализ почты                       [1x]"))

}
