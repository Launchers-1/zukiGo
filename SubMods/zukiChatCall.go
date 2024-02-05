package SubMods

import (
  "bytes"
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
)

type ZukiChatCall struct {
  ApiKey string
}

type ChatData struct {
  Model       string      `json:"model"`
  Messages    []Message   `json:"messages"`
  Temperature float64     `json:"temperature"`
}

type Message struct {
  Role    string `json:"role"`
  Content string `json:"content"`
}

func NewZukiChatCall(apiKey string) *ZukiChatCall {
  return &ZukiChatCall{ApiKey: apiKey}
}

func (z *ZukiChatCall) ChatData(userName, userMessage, requestedModel, systemPrompt string, currTemp float64) ChatData {
  return ChatData{
    Model: requestedModel,
    Messages: []Message{
      {
        Role:    "system",
        Content: systemPrompt,
      },
      {
        Role:    "user",
        Content: systemPrompt + "\n Here is a message a user called " + userName + " sent you: " + userMessage,
      },
    },
    Temperature: currTemp,
  }
}

func (z *ZukiChatCall) ChatCall(userName, userMessage, requestedModel, systemPrompt string, currTemp float64, endpoint string) (string, error) {
  chatData := z.ChatData(userName, userMessage, requestedModel, systemPrompt, currTemp)
  payload, err := json.Marshal(chatData)
  if err != nil {
    return "", err
  }

  req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(payload))
  if err != nil {
    return "", err
  }

  req.Header.Set("Content-Type", "application/json")
  req.Header.Set("Authorization", "Bearer "+z.ApiKey)

  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    return "", err
  }
  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return "", err
  }

  var result map[string]interface{}
  err = json.Unmarshal(body, &result)
  if err != nil {
    return "", err
  }

  return result["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string), nil
}
