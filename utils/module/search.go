package module

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/fatih/color"
)

type SortInfoNumber struct {
	Operator string `json:"operator"`
	Region   string `json:"region"`
	FullNum  string `json:"full_num"`
}

type SortInfoIP struct {
	Country string `json:"country"`
	CoCode  string `json:"countryCode"`
	Region  string `json:"region"`
}

func SearchFreeApiMethod(number string) {

	cyan := color.New(color.FgHiCyan)
	//lightBlack := color.New(color.FgHiBlack)

	response, _ := http.Get("http://num.voxlink.ru/get/?num=+" + fmt.Sprint(number))

	body, _ := io.ReadAll(response.Body)

	var info SortInfoNumber

	err := json.Unmarshal(body, &info)
	if err != nil {
		fmt.Println("[String-Free] Ошибка парсинга JSON:", err)

	}
	cyan.Print(string(body))

	cyan.Println("\n/./ Информация о номере /./")
	cyan.Printf("Оператор: %s\n", info.Operator)
	cyan.Printf("Регион: %s\n", info.Region)
	cyan.Printf("Полный номер: %s\n", info.FullNum)

}

func SearchFreeApiIPMethod(ip string) {

	cyan := color.New(color.FgHiCyan)

	response, err := http.Get("http://ip-api.com/json/" + fmt.Sprint(ip))
	if err != nil {
		fmt.Println("[String-Free] Ошибка при выполнении запроса:", err)

	}

	defer response.Body.Close()

	var info SortInfoIP
	body, _ := io.ReadAll(response.Body)
	err = json.Unmarshal(body, &info)

	if err != nil {
		fmt.Println("[String-Free] Ошибка парсинга JSON:", err)
	}

	cyan.Println("\n/./ Информация о IP-адресе /./")
	cyan.Printf("Страна: %s\n", info.Country)
	cyan.Printf("Код страны: %s\n", info.CoCode)
	cyan.Printf("Регион: %s\n", info.Region)
	cyan.Println("\nХочешь больше информации? Покупай StringPremium. @owaqx")

}
