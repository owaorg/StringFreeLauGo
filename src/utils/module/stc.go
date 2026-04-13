package module

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fatih/color"
)

type SortResponseBlaAPI struct {
	Domain_requested string `json:"domain_requested"`
	Port_requested   int    `json:"port_requested"`

	Subject struct {
		CommonName1 string `json:"commonName"`
	} `json:"subject"`

	Inssuer struct {
		CointryName      string `json:"countryName"`
		OrganizationName string `json:"organizationName"`
		CommonName2      string `json:"commonName"`
	} `json:"issuer"`

	Version               int    `json:"version"`
	Serial_number         string `json:"serial_number"`
	ValidFrom             string `json:"valid_from"`
	ValidTo               string `json:"valid_to"`
	Days_until_expiration int    `json:"days_until_expiration"`
	Is_valid              bool   `json:"is_valid"`
	Tls_version           string `json:"tls_version"`
	Cipher_suite          string `json:"cipher_suite"`
	Cipher_strength_bits  int    `json:"cipher_strength_bits"`
	Cipher_security       string `json:"cipher_security"`
	Signature_algorithm   string `json:"signature_algorithm"`
	Public_key_info       string `json:"public_key_info"`
}

func SSLorTSLCheckMethod(owaqx string) {

	white := color.New(color.FgWhite)
	red := color.New(color.FgRed)
	cyan := color.New(color.FgCyan)

	var url string = "https://heimdal.techviral.video/tools/ssl/check"
	var json4ik string = fmt.Sprintf(`{"domain": "%s", "port": 443}`, owaqx)

	response, err := http.Post(url, "application/json", bytes.NewBufferString(json4ik))
	if err != nil {
		white.Printf("%s Ошибка отправки запроса на сервис \n",
			red.Sprint("StringError"),
		)
		return
	}

	defer response.Body.Close()
	var sort SortResponseBlaAPI

	err = json.NewDecoder(response.Body).Decode(&sort)
	if err != nil {
		white.Printf("%s Ошибка парсинга JSON \n",
			red.Sprint("StringError"),
		)
		return
	}

	fmt.Printf("\n┃ %s > Информация о SSL/TLS.\n", cyan.Sprint("StringFree"))
	cyan.Printf("╭ Домен: %s\n", sort.Domain_requested)
	cyan.Printf("┃ На кого выдан сертификат: %s\n", sort.Subject.CommonName1)
	cyan.Printf("┃ Страна центра сертификации: %s\n", sort.Inssuer.CointryName)
	cyan.Printf("┃ Название компании: %s\n", sort.Inssuer.OrganizationName)
	cyan.Printf("┃ Техническое имя сервера выдачи: %s\n", sort.Inssuer.CommonName2)
	cyan.Printf("┃ Версия формата сертификата: %d\n", sort.Version)
	cyan.Printf("┃ Уникальный «паспортный номер» сертификата: %s\n", sort.Serial_number)
	cyan.Printf("┃ Даты начала и окончания действия: %s\n", sort.ValidFrom)
	cyan.Printf("┃ Даты начала и окончания действия: %s | %s\n", sort.ValidFrom, sort.ValidTo)
	cyan.Printf("┃ Cертификат просрочен или отозван: %t\n", sort.Is_valid)
	cyan.Printf("┃ Протокол защиты: %s\n", sort.Tls_version)
	cyan.Printf("┃ Набор алгоритмов шифрования: %s\n", sort.Cipher_suite)
	cyan.Printf("┃ Стойкость шифра: %d\n", sort.Cipher_strength_bits)
	cyan.Printf("┃ Оценка безопасности: %s\n", sort.Cipher_security)
	cyan.Printf("┃ Алгоритм подписи: %s\n", sort.Signature_algorithm)
	cyan.Printf("╰ Информация о публичном ключе: %s\n", sort.Public_key_info)

	fmt.Println()

}

/*

Пример ответа от их API

{
  "domain_requested": "string.surf",
  "port_requested": 443,
  "subject": {
    "commonName": "string.surf"
  },
  "issuer": {
    "countryName": "US",
    "organizationName": "Let's Encrypt",
    "commonName": "E8"
  },
  "version": 3,
  "serial_number": "068C3B4FD37DE864EF0775A4E36F221E484E",
  "valid_from": "2026-04-06T15:57:40",
  "valid_to": "2026-07-05T15:57:39",
  "days_until_expiration": 83,
  "is_valid": true,
  "subject_alternative_names": [
    "*.string.surf",
    "string.surf"
  ],
  "ocsp_uris": [],
  "ca_issuers_uris": [
    "http://e8.i.lencr.org/"
  ],
  "crl_distribution_points": [
    "http://e8.c.lencr.org/79.crl"
  ],
  "tls_version": "TLSv1.3",
  "cipher_suite": "TLS_AES_256_GCM_SHA384",
  "cipher_strength_bits": 256,
  "cipher_security": "Strong",
  "signature_algorithm": "ecdsa-with-SHA384",
  "public_key_info": "ECC, Curve: secp256r1, Key Size: 256 bits"
} */
