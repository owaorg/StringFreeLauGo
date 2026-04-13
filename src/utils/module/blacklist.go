package module

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/fatih/color"
)

type SortResponseApiCheckBlackListGudamFuckPaketLOX struct {
	// пакет лошарик в жопе шарик (рилрол мамент)

	Ip                       string `json:"ip"`                       //айпи
	Total_blacklists_checked int    `json:"total_blacklists_checked"` // типо сколько апи проверил сервисов хз
	Status                   string `json:"status"`                   // статус обычно clean
	Message                  string `json:"message"`                  // сообщение от их апи
	Is_blacklisted           bool   `json:"is_blacklisted"`
}

func CheckBlacklist(ip string) {
	cyan := color.New(color.FgHiBlack)
	red := color.New(color.FgRed)
	white := color.New(color.FgWhite)

	var url string = "https://heimdal.techviral.video/blacklist/check"
	pay_owaqx := strings.NewReader(fmt.Sprintf(`{"ip": "%s"}`, ip))

	response, err := http.Post(url, "application/json", pay_owaqx)

	if err != nil {
		white.Printf("%s Ошибка отправки запроса на сервис",
			red.Sprint("StringError"),
		)
	}

	defer response.Body.Close()
	var soriTING SortResponseApiCheckBlackListGudamFuckPaketLOX

	err = json.NewDecoder(response.Body).Decode(&soriTING)

	if err != nil {
		white.Printf("%s Ошибка парсинга JSON",
			red.Sprint("StringError"),
		)
	}

	var is_blacklistedFalse string

	if soriTING.Is_blacklisted == false {
		is_blacklistedFalse = "Нету"
	} else {
		is_blacklistedFalse = "Есть"
	}

	fmt.Printf("\n┃ %s > Информация о BlackList IP.\n", cyan.Sprint("StringFree"))
	cyan.Printf("╭ IP: %s\n", soriTING.Ip)
	cyan.Printf("┃ Есть ли в черных списках?: %s\n", is_blacklistedFalse)
	cyan.Printf("┃ Проверок было: %d\n", soriTING.Total_blacklists_checked)
	cyan.Printf("┃ Статус: %s\n", soriTING.Status)
	cyan.Printf("╰ Сообщение: %s\n", soriTING.Message)

	fmt.Println()
}

/*

Пример ответа от API

{
    "ip":  "82.41.30.6",
    "is_blacklisted":  false,
    "blacklisted_servers":  [ <-

                            ], <-
    "total_blacklists_checked":  45,
    "blacklists_found":  0,
    "status":  "clean",
    "message":  "IP is clean and not found on any blacklists"
}

*/
