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
		fmt.Println("[String-Free] Ошибка создания запроса:", err)
		return
	}

	req.Header.Set("Authorization", "Bearer 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")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("[String-Free] Ошибка сети:", err)
		return
	}
	defer resp.Body.Close()

	var result AIResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		fmt.Println("[String-Free] Ошибка декодирования ответа:", err)
		return
	}

	if len(result.Choices) > 0 {
		content := result.Choices[0].Message.Content
		cyan.Print("\n[AI]: ", content)
		fmt.Println()
	} else {
		fmt.Println("[String-Free] AI прислал пустой ответ")
	}
}
