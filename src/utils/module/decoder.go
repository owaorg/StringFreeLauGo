package module

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fatih/color"
)

type SortInfoDecoder struct {
	Display_name string `json:"display_name"`
	Place_rank   int    `json:"place_rank"`
}

func DecoderLanLonMethod(lans float64, lons float64) {
	white := color.New(color.FgWhite)
	red := color.New(color.FgRed)
	cyan := color.New(color.FgHiBlack)

	urls := fmt.Sprintf("https://nominatim.openstreetmap.org/reverse?format=jsonv2&lat=%f&lon=%f", lans, lons)

	client := &http.Client{}
	req, err := http.NewRequest("GET", urls, nil)
	if err != nil {
		return
	}
	req.Header.Set("User-Agent", "StringLoLLLLL))))))))))))))))))))))))))))))))))))))))))")

	response, err := client.Do(req)
	if err != nil {
		white.Printf("%s Ошибка отправки запроса \n", red.Sprint("StringError"))
		return
	}
	defer response.Body.Close()

	var sort SortInfoDecoder
	err = json.NewDecoder(response.Body).Decode(&sort)
	if err != nil {
		white.Printf("%s Ошибка парсинга JSON \n", red.Sprint("StringError"))
		return
	}

	fmt.Printf("\n┃ %s > Информация после Декодирования.\n", cyan.Sprint("StringFree"))
	cyan.Printf("╭ Широта, Долгота: %f | %f\n", lans, lons)
	cyan.Printf("┃ Ранк точности: %v \n", sort.Place_rank)
	cyan.Printf("╰ Адрес: %s\n", sort.Display_name)
	fmt.Println()
}
