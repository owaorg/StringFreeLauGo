package cli

import (
	"fmt"
	"time"

	clear "str-free/utils/module" //мОДУЛЬ ДЛЯ ОЧИСТКИ ЭКРАНА
)

func WelcomeBannerTextMethod() {
	var WelcomeBannerText string = `
// STRING FREE EDITION
// Disclaimer: ПО предоставляется "как есть" (AS IS). Стабильная работа не гарантируется.

// Технический стек: Golang. Инструмент находится в стадии активной разработки и обучения.
// Нашли баг? PR welcome: https://github.com/owaorg/StringFreeLauGo
// Автор не несет обязательств по исправлению ошибок в бесплатной версии.

// Использование: Инструмент взаимодействует со сторонними API (Public/Private).
`
	fmt.Println(WelcomeBannerText) // Выводим баннер в консоль
	time.Sleep(5 * time.Second)    // Задержка 5 секунд для прочтения баннера
	clear.ClearScreenMethod()      // Очищаем экран после отображения баннера
}
