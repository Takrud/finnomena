package model

import "time"

type Fund struct {
	Name        string    `json:"thailand_fund_code" `
	RankOfFund  float32   `json:"avg_return" `
	UpdatedDate time.Time `json:"nav_date" `
	Performance float32   `json:"nav_return" `
	Price       float32   `json:"nav" `
}
