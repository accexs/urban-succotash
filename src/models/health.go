package models

type Health struct {
	Status      string `json:"status" example:"ok"`
	Environment string `json:"environment" example:"dev"`
}
