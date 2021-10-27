package types

type ErrorResponse struct {
	Http    int    `json:"http"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}
