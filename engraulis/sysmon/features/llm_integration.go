package sysmon
/* 
import (
	"fmt"
	"net/http"
	"strings"
)

func LlmIntegration(provider string, token string, model string) string {
	providerList := [7]string{"openai", "chatgpt", "anthropic", "deepseek", "google", "gemini", "openrouter"}
	rawProviderName := strings.ToLower(provider)
	switch {
	case rawProviderName == providerList[4] || rawProviderName == providerList[5]:
		url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta/models/%s:generateContent?key=%s", model, token)
		headers, err := http.NewRequest("POST", url, nil)
		if err != nil {
			fmt.Println(err)
		}
		req, err := http.Post(url, "application/json")
	}
}
*/