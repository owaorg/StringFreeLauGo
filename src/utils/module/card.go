package module

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/fatih/color"
)

type sort struct {
	Valid bool   `json:"valid"`
	Type1 string `json:"type"`
}

func ValidateCardMethod(card string) {

	var turl string = "https://www.freecodeformat.com/credit-card-validator.php"

	cyan := color.New(color.FgCyan)
	red := color.New(color.FgRed)
	white := color.New(color.FgWhite)

	bases := url.Values{}
	bases.Set("action", "validate")
	bases.Set("card_number", card)

	response, err := http.PostForm(turl, bases)
	if err != nil {
		white.Printf("%s Ошибка отправки запроса на сервис",
			red.Sprint("StringError"),
		)
	}

	defer response.Body.Close()
	var loh sort
	err = json.NewDecoder(response.Body).Decode(&loh)

	var falseis string

	if loh.Valid == false {
		falseis = "Нет"
	} else {
		falseis = "Да"
	}

	fmt.Printf("\n┃ %s > Информация о Карте.\n", cyan.Sprint("StringFree"))
	cyan.Printf("╭ Валидный?: %s\n", falseis)
	cyan.Printf("╰ Сообщение: %s\n", loh.Type1)

	fmt.Println()
}
