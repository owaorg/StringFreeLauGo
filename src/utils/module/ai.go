package module

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/fatih/color"
)

type AIResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func AIRequestMethod(prompt string, data interface{}) {

	cyan := color.New(color.FgHiCyan)

	payload, _ := json.Marshal(map[string]interface{}{
		"model": "deepseek/deepseek-chat",
		"messages": []map[string]string{
			{"role": "user", "content": fmt.Sprintf("%s %v", prompt, data)},
		},
	})

	url := "https://openrouter.ai/api/v1/chat/completions"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("[String-Error] Ошибка создания запроса:", err)
		return
	}

	req.Header.Set("Authorization", "Bearer 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")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("[String-Error] Ошибка сети:", err)
		return
	}
	defer resp.Body.Close()

	var result AIResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		fmt.Println("[String-Error] Ошибка декодирования ответа:", err)
		return
	}

	if len(result.Choices) > 0 {
		content := result.Choices[0].Message.Content
		cyan.Print("\n[AI]: ", content)
		fmt.Println()
	} else {
		fmt.Println("[String-Error] AI прислал пустой ответ")
	}
}
