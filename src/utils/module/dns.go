package module

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fatih/color"
)

type DNSRecord struct {
	Type     string `json:"type"`
	Value    string `json:"value"`
	TTL      int    `json:"ttl"`
	Hostname string `json:"hostname,omitempty"`
}

func DnsAnalyzeMethod(domen string) {
	cyan := color.New(color.FgHiBlack)
	red := color.New(color.FgRed)

	url := "https://heimdal.techviral.video/dns/records"
	payload := map[string]interface{}{
		"domain":       domen,
		"record_types": []string{"A", "MX", "NS", "TXT", "AAAA"},
	}
	body, _ := json.Marshal(payload)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))

	if err != nil {
		fmt.Printf("┃ %s > Ошибка при запросе на сервис.\n", red.Sprint("StringError"))
		return
	}

	defer resp.Body.Close()

	var records []DNSRecord

	json.NewDecoder(resp.Body).Decode(&records)
	cyan.Printf("\n╭ %s > Результаты:\n", cyan.Sprint("StringFree"))

	for _, rec := range records {
		cyan.Printf("┃ [%s] -> %s\n", rec.Type, rec.Value)
	}
	cyan.Println("╰ Завершено.")
}
