package models

type Client struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	AccountNum string  `json:"account_num"`
	Balance    float64 `json:"balance"`
}
