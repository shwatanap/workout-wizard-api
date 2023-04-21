package request

type ChatGPTRequest struct {
	Model    string            `json:"model"`
	Messages []*RequestMessage `json:"messages"`
}

func NewChatGPTRequest(modelID string, messages []*RequestMessage) *ChatGPTRequest {
	return &ChatGPTRequest{
		Model:    modelID,
		Messages: messages,
	}
}

type RequestMessage struct {
	Role    string `json:"role"` // user or assistant or system
	Content string `json:"content"`
}

func NewRequestMessage(role string, content string) *RequestMessage {
	return &RequestMessage{
		Role:    role,
		Content: content,
	}
}
