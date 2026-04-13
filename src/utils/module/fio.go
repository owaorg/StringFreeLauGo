package module

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fatih/color"
)

type SortFioResponseBlat struct {
	Data struct {
		СтрВсего int `json:"СтрВсего"`

		Записи []struct {
			ОГРН      string `json:"ОГРН"`
			Инн       string `json:"ИНН"`
			КПП       string `json:"КПП"`
			НаимСокр  string `json:"НаимСокр"`
			НаимПолн  string `json:"НаимПолн"`
			ДатаРег   string `json:"ДатаРег"`
			Статус    string `json:"Статус"`
			РегионКод string `json:"РегионКод"`
			ЮрАдрес   string `json:"ЮрАдрес"`
			ОКВЭД     string `json:"ОКВЭД"`

			Руковод []struct {
				ФИО       string `json:"ФИО"`
				ИНН       string `json:"ИНН"`
				ВидДолжн  string `json:"ВидДолжн"`
				НаимДолжн string `json:"НаимДолжн"`
				Недост    bool   `json:"Недост"`
			} `json:"Руковод"`
		} `json:"Записи"`
	} `json:"data"`
}

func SearchFioMethod(req string) {
	cyan := color.New(color.FgHiBlack)
	red := color.New(color.FgRed)
	white := color.New(color.FgWhite)
	var api_url string = fmt.Sprintf("https://api.ofdata.ru/v2/search?key=DiC9ALodH5T12BfR&by=leader-name&obj=org&query=%s", req)

	response, err := http.Get(api_url)
	if err != nil {
		white.Printf("%s Ошибка отправки запроса на сервис",
			red.Sprint("StringError"),
		)
	}

	defer response.Body.Close()
	var sort SortFioResponseBlat

	err = json.NewDecoder(response.Body).Decode(&sort)

	for _, запись := range sort.Data.Записи {
		cyan.Printf("╭ ОГРН: %s\n", запись.ОГРН)
		cyan.Printf("┃ ИНН: %s\n", запись.Инн)
		cyan.Printf("┃ КПП: %s\n", запись.КПП)
		cyan.Printf("┃ Сокращённое наименование: %s\n", запись.НаимСокр)
		cyan.Printf("┃ Полное наименование: %s\n", запись.НаимПолн)
		cyan.Printf("┃ Дата регистрации: %s\n", запись.ДатаРег)
		cyan.Printf("┃ Статус: %s\n", запись.Статус)
		cyan.Printf("┃ Код региона: %s\n", запись.РегионКод)
		cyan.Printf("┃ Юридический адрес: %s\n", запись.ЮрАдрес)
		cyan.Printf("╰ ОКВЭД: %s\n", запись.ОКВЭД)

		if len(запись.Руковод) > 0 {
			for j, руководитель := range запись.Руковод {
				cyan.Printf("     ╭ %d. ФИО: %s\n", j+1, руководитель.ФИО)
				cyan.Printf("     ┃ ИНН: %s\n", руководитель.ИНН)
				cyan.Printf("     ┃ Вид должности: %s\n", руководитель.ВидДолжн)
				cyan.Printf("     ┃ Наименование должности: %s\n", руководитель.НаимДолжн)
				cyan.Printf("     ╰ Недостоверность: %t\n", руководитель.Недост)
				cyan.Println()
			}
		}
	}
}
