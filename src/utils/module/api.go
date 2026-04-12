package module

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/fatih/color"
)

type SortInfinityResponseMethod struct {
	Fio      string `json:"fio"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	Fullname string `json:"fullname"`
	Inn      string `json:"inn"`
	Bdate    string `json:"bdate"`
	Snils    string `json:"snils"`
	Gender   string `json:"gender"`
	Id       string `json:"id"`
}

type InfinityResponse struct {
	Result []SortInfinityResponseMethod `json:"results"`
}

func SearchApiPrivateMethod(method string, reg string) {

	//cyan := color.New(color.FgCyan)
	white := color.New(color.FgWhite)

	var lol string = "сюда ключ типа да?"
	var api_url string // Переменная для хранения URL API
	var data InfinityResponse

	if method == "number" {
		api_url = "эта хз что" + fmt.Sprint(reg) + "&token=" + fmt.Sprint(lol)
	}
	if method == "email" {
		api_url = "эта хз что" + fmt.Sprint(reg) + "&token=" + fmt.Sprint(lol)
	}

	if method == "fio" {
		api_url = "эта хз что" + fmt.Sprint(reg) + "&token=" + fmt.Sprint(lol)
	}

	response, err := http.Get(api_url)

	if err != nil {

		if os.IsTimeout(err) {
			fmt.Println("Привышенно ожидание от сервера InfinityAPI.")
		} else {
			fmt.Println("[String-Free] Ошибка при выполнении запроса:")
		}
	}

	defer response.Body.Close()

	if response.StatusCode == 200 {

		body, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Println("[String-Free] Ошибка чтения ответа от сервера:")
		}

		//run := []rune(full)

		json.Unmarshal(body, &data)

		for _, res := range data.Result {
			if res.Fio != "" {
				white.Printf("ФИО: %s\n", res.Fio)
			}

			if res.Email != "" && len(res.Email) >= 5 {
				mask := res.Email[:1] + "*" + res.Email[2:5] + "*" + res.Email[5:]
				white.Printf("Email: %s\n", mask)
			}

			if res.Address != "" && len(res.Address) >= 5 {
				mask := res.Address[:1] + "*" + res.Address[2:5] + "*" + res.Address[5:]
				white.Printf("Адрес: %s\n", mask)
			}

			if res.Fullname != "" && len(res.Fullname) >= 5 {
				mask := res.Fullname[:1] + "*" + res.Fullname[2:5] + "*" + res.Fullname[5:]
				white.Printf("Полное ФИО: %s\n", mask)
			}

			if res.Gender != "" {
				white.Printf("Гендер: %s\n", res.Gender)
			}

			if res.Id != "" {
				white.Printf("Id: %s\n", res.Id)
			}

			if res.Inn != "" && len(res.Inn) >= 5 {
				mask := res.Inn[:1] + "*" + res.Inn[2:5] + "*" + res.Inn[5:]
				white.Printf("ИНН: %s\n", mask)
			}

			if res.Snils != "" && len(res.Snils) >= 5 {
				mask := res.Snils[:1] + "*" + res.Snils[2:5] + "*" + res.Snils[5:]
				white.Printf("Снилс: %s\n", mask)
			}
		}
	}
}
