package module

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/fatih/color"
)

type SortInfo struct {
	Operator string `json:"operator"`
	Region   string `json:"region"`
	FullNum  string `json:"full_num"`
}

func SearchFreeApiMethod(number string) {

	cyan := color.New(color.FgHiCyan)
	//lightBlack := color.New(color.FgHiBlack)

	response, _ := http.Get("http://num.voxlink.ru/get/?num=+" + fmt.Sprint(number))

	body, _ := io.ReadAll(response.Body)

	var info SortInfo

	err := json.Unmarshal(body, &info)
	if err != nil {
		fmt.Println("Ошибка парсинга JSON:", err)
		return
	}
	//cyan.Print(string(body))

	cyan.Println("\n/./ Информация о номере /./")
	cyan.Printf("Оператор: %s\n", info.Operator)
	cyan.Printf("Регион: %s\n", info.Region)
	cyan.Printf("Полный номер: %s\n", info.FullNum)

}
