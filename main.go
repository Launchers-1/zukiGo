package zukiGO

import (
	"github.com/Launchers-1/zukiGo/MainMods"
)

type ZukiCall struct {
	apiKey       string
	apiKeyBackup string
	model        string
	systemPrompt string
	temperature  float64
	zukiChat     *MainMods.ZukiChat
	zukiImage    *MainMods.ZukiImage
}

func NewZukiCall(apiKey, apiKeyBackup, model, systemPrompt string, temperature float64) *ZukiCall {
	zukiChat := MainMods.NewZukiChat(apiKey, apiKeyBackup, model, systemPrompt, temperature)
	return &ZukiCall{
		apiKey:       apiKey,
		apiKeyBackup: apiKeyBackup,
		model:        model,
		systemPrompt: systemPrompt,
		temperature:  temperature,
		zukiChat:     zukiChat,
	}
}
