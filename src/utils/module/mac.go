package module

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/fatih/color"
)

type MacInfo struct {
	Company   string `json:"company"`
	Country   string `json:"country"`
	AddressL1 string `json:"addressL1"`
	AddressL2 string `json:"addressL2"`
	AddressL3 string `json:"addressL3"`
	Type      string `json:"type"`
}

func MacAdressAnalyzeMethod(mac string) {
	cyan := color.New(color.FgCyan)
	red := color.New(color.FgRed)
	white := color.New(color.FgWhite)

	url := "https://www.macvendorlookup.com/api/v2/"
	response, err := http.Get(url + mac)
	if err != nil {
		white.Printf("%s Ошибка запроса на сервис\n", red.Sprint("[StringError]"))
		return
	}

	defer response.Body.Close()
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		white.Printf("%s Ошибка чтения данных\n", red.Sprint("[StringError]"))
		return
	}

	if len(bodyBytes) == 0 {
		white.Printf("\n%s Данные по этому MAC не найдены.",
			red.Sprint("[StringNoFound]"),
		)

		return
	}

	var sort []MacInfo
	err = json.Unmarshal(bodyBytes, &sort)
	if err != nil {
		white.Printf("%s Ошибка парсинга JSON\n", red.Sprint("[StringError]"))
		return
	}

	if len(sort) > 0 {
		fmt.Printf("\n┃ %s > Информация о MAC-адресе.\n", cyan.Sprint("StringFree"))
		cyan.Printf("╭ Компания: %s\n", sort[0].Company)
		cyan.Printf("┃ Страна: %s\n", sort[0].Country)
		cyan.Printf("┃ Адрес: %s, %s\n", sort[0].AddressL1, sort[0].AddressL2)
		cyan.Printf("╰ Тип: %s\n", sort[0].Type)
	} else {
		red.Println("\n┃ [StringNoFound] Данные по этому MAC не найдены.")

		fmt.Println("")
	}
}
