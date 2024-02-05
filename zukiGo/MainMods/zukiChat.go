package MainMods

import (
	"https://github.com/Launchers-1/zukiGo/SubMods"
	"errors"
	"strings"
)

const (
	defaultAPIEndpoint           = "https://zukijourney.xyzbot.net/v1/chat/completions"
	defaultAPIEndpointUnfiltered = "https://zukijourney.xyzbot.net/unf/chat/completions"
	defaultAPIEndpointBackup     = "https://thirdparty.webraft.in/v1/chat/completions"
	defaultModel                 = "gpt-3.5"
)

var modelsList = []string{"gpt-3.5-turbo", "gpt-3.5-turbo-instruct", "gpt-3.5-turbo-16k", "gpt-4", "gpt-4-32k", "gpt-4-1106-preview", "gpt-4-0125-preview", "gpt-4-vision-preview", "claude", "claude-2", "claude-2.1", "claude-instant-v1", "claude-instant-v1-100k", "pplx-70b-online", "palm-2", "bard", "gemini-pro", "gemini-pro-vision", "mixtral-8x7b", "mixtral-8x7b-instruct", "mistral-tiny", "mistral-small", "mistral-medium", "mistral-7b-instruct", "codellama-7b-instruct", "llama-2-7b", "llama-2-70b-chat", "mythomax-l2-13b-8k", "sheep-duck-llama", "goliath-120b", "nous-llama2", "yi-34b", "openchat", "solar10-7b", "pi"}

type ZukiChat struct {
	apiKey                string
	apiEndpoint           string
	apiEndpointUnfiltered string
	apiEndpointBackup     string
	systemPrompt          string
	model                 string
	apiKeyBackup          string
	temperature           float64
}

func NewZukiChat(apiKey, apiKeyBackup, model, systemPrompt string, temperature float64) (*ZukiChat, error) {
	if !contains(modelsList, model) {
		return nil, errors.New("Invalid model. Please choose from the following: " + strings.Join(modelsList, ", "))
	}
	return &ZukiChat{
		apiKey:                apiKey,
		apiEndpoint:           defaultAPIEndpoint,
		apiEndpointUnfiltered: defaultAPIEndpointUnfiltered,
		apiEndpointBackup:     defaultAPIEndpointBackup,
		systemPrompt:          systemPrompt,
		model:                 model,
		apiKeyBackup:          apiKeyBackup,
		temperature:           temperature,
	}, nil
}

func (z *ZukiChat) ChangeBackupEndpoint(endpoint string) {
	z.apiEndpointBackup = endpoint
}

func (z *ZukiChat) SetSystemPrompt(systemPrompt string) {
	z.systemPrompt = systemPrompt
}

func (z *ZukiChat) SetTemp(newTemp float64) {
	if newTemp >= 0 && newTemp <= 1 {
		z.temperature = newTemp
	}
}

func (z *ZukiChat) SendMessage(userName, userMessage string) (string, error) {
	return z.sendMessage(userName, userMessage, z.apiEndpoint)
}

func (z *ZukiChat) SendUnfilteredMessage(userName, userMessage string) (string, error) {
	return z.sendMessage(userName, userMessage, z.apiEndpointUnfiltered)
}

func (z *ZukiChat) SendBackupMessage(userName, userMessage string) (string, error) {
	return z.sendMessage(userName, userMessage, z.apiEndpointBackup)
}

func (z *ZukiChat) sendMessage(userName, userMessage, endpoint string) (string, error) {
	zukiChatCall := SubMods.NewZukiChatCall(z.apiKey)
	return zukiChatCall.ChatCall(userName, userMessage, z.model, z.systemPrompt, z.temperature, endpoint)
}

func contains(slice []string, item string) bool {
	for _, a := range slice {
		if a == item {
			return true
		}
	}
	return false
}
