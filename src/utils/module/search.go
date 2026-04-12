package module

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

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

	cyan := color.New(color.FgCyan)
	//lightBlack := color.New(color.FgHiBlack)

	response, _ := http.Get("http://num.voxlink.ru/get/?num=+" + fmt.Sprint(number))

	body, _ := io.ReadAll(response.Body)

	var info SortInfoNumber

	err := json.Unmarshal(body, &info)
	if err != nil {
		fmt.Println("[String-Free] Ошибка парсинга JSON:", err)

	}
	//cyan.Print(string(body))

	fmt.Println()

	fmt.Printf("┃ %s > Информация о номере.",
		cyan.Sprint("StringFree"),
	)

	fmt.Println()

	cyan.Printf("╭ Оператор: %s\n", info.Operator)
	cyan.Printf("┃ Регион: %s\n", info.Region)
	cyan.Printf("╰ Полный номер: %s\n", info.FullNum)

	fmt.Println()

}

func SearchFreeApiIPMethod(ip string) {

	cyan := color.New(color.FgCyan)

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

	fmt.Println()

	fmt.Printf("┃ %s > Информация об IP-Адресе.",
		cyan.Sprint("StringFree"),
	)

	cyan.Printf("\n╭ Страна: %s\n", info.Country)
	cyan.Printf("┃ Код страны: %s\n", info.CoCode)
	cyan.Printf("╰ Регион: %s\n", info.Region)
	cyan.Println("\nХочешь больше информации? Покупай StringPremium. @owaqx")

}

type SortScum struct {
	Country       string `json:"country"`
	Types         string `json:"type"`
	Ip_address    string `json:"ip_address"`
	Operator      string `json:"operator"`
	ValidikxD     bool   `json:"valid"`
	Is_voip       bool   `json:"is_voip"`
	Operator_link string `json:"operator_link"`
}

type SortScums struct {
	Data SortScum `json:"0"`
}

// Их апи почти никакую защиту не использовало от абуза
// Никак не гарантирую что проработает долго

func EasyAbuseApiLOLMethod(number string) {
	cyan := color.New(color.FgCyan)
	red := color.New(color.FgRed)

	url := "https://getscam.com/api/get-phone-free-data/"
	response, err := http.Get(url + number)

	if err != nil {
		fmt.Printf("┃ %s > Ошибка при запросе: %v\n", red.Sprint("StringError"), err)
		return
	}
	defer response.Body.Close()
	body, _ := io.ReadAll(response.Body)

	var fullResponse SortScums
	err = json.Unmarshal(body, &fullResponse)

	if err != nil {
		fmt.Printf("┃ %s > Ошибка парсинга JSON.\n", red.Sprint("StringError"))
		return
	}
	sorts := fullResponse.Data

	var Oper string
	var ValidokXDDDbyOWAQXBla string

	if sorts.Operator == "false" {
		Oper = "Нет"
	} else {
		Oper = "Да"
	}

	if sorts.ValidikxD == false {
		ValidokXDDDbyOWAQXBla = "Нет"
	} else {
		ValidokXDDDbyOWAQXBla = "Да"
	}

	fmt.Printf("\n┃ %s > Информация о номере.\n", cyan.Sprint("StringFree"))
	cyan.Printf("╭ Страна: %s\n", sorts.Country)
	cyan.Printf("┃ Тип: %s\n", sorts.Types)
	cyan.Printf("┃ Айпи: %s\n", sorts.Ip_address)
	cyan.Printf("┃ Оператор: %s\n", Oper)
	cyan.Printf("┃ Валидный?: %s\n", ValidokXDDDbyOWAQXBla)
	cyan.Printf("┃ Виртуальный?: %t\n", sorts.Is_voip)
	cyan.Printf("╰ Ссылка на оператора: %s\n", sorts.Operator_link)

	fmt.Println("")
}

type SortMailSearch struct {
	Deliverability      string `json:"deliverability"`      //Готова ли почта принимать письма
	Is_valid_format     bool   `json:"is_valid_format"`     //Валидный ли формат
	Is_disposable_email bool   `json:"is_disposable_email"` //Является ли этот email одноразовым:
}

func AnalyzeMailAdressMethod(mail string) {

	// Еще один сайт нашел, у них так-же нету никакой защиты.
	cyan := color.New(color.FgCyan)
	red := color.New(color.FgRed)

	var url string = "https://claritycheck.com/api/search/email/information"

	pay_blat := strings.NewReader(fmt.Sprintf(`{"email": "%s"}`, mail))
	resp, err := http.Post(url, "application/json", pay_blat)

	if err != nil {
		fmt.Printf("┃ %s > Ошибка при запросе на сервис.\n", red.Sprint("StringError"))
		return
	}

	defer resp.Body.Close()
	var sort SortMailSearch

	err = json.NewDecoder(resp.Body).Decode(&sort)
	if err != nil {
		fmt.Printf("┃ %s > Ошибка парсинга JSON: %v\n", red.Sprint("StringError"), err)
		return
	}

	var DELIVERABLE string

	if sort.Deliverability == "DELIVERABLE" {
		DELIVERABLE = "Да"
	} else {
		DELIVERABLE = "Нет"
	}

	fmt.Printf("\n┃ %s > Информация о Почте.\n", cyan.Sprint("StringFree"))
	cyan.Printf("╭ Готова ли почта принимать письма: %s\n", DELIVERABLE)
	cyan.Printf("┃ Валидный ли формат: %t\n", sort.Is_valid_format)
	cyan.Printf("╰ Является ли этот email одноразовым: %t\n", sort.Is_disposable_email)

	fmt.Println("")
}
