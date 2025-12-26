package entity

import (
	"time"
)

type ShortURL struct {
    ID        int   `json:"id"`
    URL       string    `json:"url"`
    ShortCode string    `json:"shortCode"`
    CreatedAt time.Time `json:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt"`
}

type Stats struct {
    ShortURL ShortURL
    Visits int
}