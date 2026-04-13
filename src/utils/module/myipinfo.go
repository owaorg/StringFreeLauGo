package module

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fatih/color"
)

type MyIpInfo struct {
	Ip           string  `json:"ip"`
	City         string  `json:"city"`
	Region       string  `json:"region"`
	Country      string  `json:"country"`
	Country_code string  `json:"country_code"`
	AssHmmmXD    string  `json:"as"`
	Isp          string  `json:"isp"`
	Org          string  `json:"org"`
	Lat          float64 `json:"lat"`
	Lon          float64 `json:"lon"`
	Timezone     string  `json:"timezone"`
	Is_hosting   bool    `json:"is_hosting"`
	Is_mobile    bool    `json:"is_mobile"`
	Is_proxy     bool    `json:"is_proxy"`

	Proxy struct {
		Is_proxy2 bool `json:"is_proxy"`
		Is_vpn    bool `json:"is_vpn"`
		Is_tor    bool `json:"is_tor"`
	} `json:"proxy"`
}

func InfoMyIPMethod() {

	white := color.New(color.FgWhite)
	red := color.New(color.FgRed)
	cyan := color.New(color.FgCyan)

	var url string = "https://heimdal.techviral.video/ip/my-ip"
	response, err := http.Get(url)
	if err != nil {
		white.Printf("%s Ошибка отправки запроса на сервис \n",
			red.Sprint("StringError"),
		)
		return
	}

	defer response.Body.Close()
	var sort MyIpInfo

	err = json.NewDecoder(response.Body).Decode(&sort)
	if err != nil {
		white.Printf("%s Ошибка парсинга JSON \n",
			red.Sprint("StringError"),
		)
		return
	}

	fmt.Printf("\n┃ %s > Информация о Моем IP.\n", cyan.Sprint("StringFree"))
	cyan.Printf("╭ IP: %s\n", sort.Ip)
	cyan.Printf("┃ Откуда человек выходит в сеть: %s | %s | %s\n", sort.City, sort.Region, sort.Country)
	cyan.Printf("┃ Короткий код страны: %s\n", sort.Country_code)
	cyan.Printf("┃ Номер автономной системы: %s\n", sort.AssHmmmXD)
	cyan.Printf("┃ Твой провайдер: %s\n", sort.Isp)
	cyan.Printf("┃ Организация: %s\n", sort.Org)
	cyan.Printf("┃ Широта: %f\n", sort.Lat)
	cyan.Printf("┃ Долгота: %f\n", sort.Lon)
	cyan.Printf("┃ Часовой пояс: %s \n", sort.Timezone)
	cyan.Printf("┃ IP принадлежит дата-центру или серверу?: %t\n", sort.Is_hosting)
	cyan.Printf("┃ юзер сидит через мобильный интернет?: %t\n", sort.Is_mobile)
	cyan.Printf("┃ Прокси?: %t\n", sort.Is_proxy)
	cyan.Printf("┃ Впн?: %t\n", sort.Proxy.Is_vpn)
	cyan.Printf("┃ Tor Browser?: %t\n", sort.Proxy.Is_tor)
	cyan.Printf("╰ Вторая проверка на Прокси: %t\n", sort.Proxy.Is_proxy2)

	fmt.Println()

}

/* {
  "ip": "82.41.30.6",
  "city": "Almaty",
  "region": "Almaty",
  "country": "Kazakhstan",
  "country_code": "KZ",
  "as": "AS49791 Newserverlife LLC",
  "isp": "Newserverlife LLC",
  "org": "3HCloud LLC",
  "lat": 43.2525,
  "lon": 76.9115,
  "timezone": "Asia/Almaty",
  "is_hosting": false,
  "is_mobile": false,
  "is_proxy": false,
  "proxy": {
    "is_proxy": false,
    "is_vpn": false,
    "is_tor": false
  }
} */
