package models

type Credentials struct {
	IDInstance       string `json:"idInstance"`
	APITokenInstance string `json:"apiTokenInstance"`
}

type SendMessageRequest struct {
	Credentials
	PhoneNumber string `json:"phoneNumber"`
	Message     string `json:"message"`
}

type SendFileByURLRequest struct {
	Credentials
	PhoneNumber string `json:"phoneNumber"`
	FileURL     string `json:"fileUrl"`
	FileName    string `json:"fileName"`
	Caption     string `json:"caption,omitempty"`
}

type GreenAPISendMessage struct {
	ChatID  string `json:"chatId"`
	Message string `json:"message"`
}

type GreenAPISendFileByURL struct {
	ChatID   string `json:"chatId"`
	URLFile  string `json:"urlFile"`
	FileName string `json:"fileName"`
	Caption  string `json:"caption,omitempty"`
}

type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}
