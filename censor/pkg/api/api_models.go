package api

type CensorReqest struct {
	Content string `json:"content"`
}

type BadResponse struct {
	Success   bool   `json:"success"`
	Error     string `json:"error"`
	RequestID string `json:"request_id"`
}

type CensorResponse struct {
	Success    bool   `json:"success"`
	IsCensored bool   `json:"is_censored"`
	RequestID  string `json:"request_id"`
}
