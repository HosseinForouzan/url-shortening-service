package shorten

import (
	"fmt"
	"time"

	"github.com/HosseinForouzan/url-shortening-service/entity"
)

const LengthOfShortCode = 7

type Repository interface {
	IsShortCodeUnique(shortCode string) (bool, error)
	Create(ShortURL entity.ShortURL) (entity.ShortURL, error)
}

type Shorten struct {
	repo Repository
}

func New(repo Repository) Shorten {
	return Shorten{repo: repo}
}

type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	ShortURL entity.ShortURL `json:"resp"`
}

func (s Shorten) CreateService(req ShortenRequest) (ShortenResponse, error) {
	shortCode, err := randomShortCode(LengthOfShortCode)
	if err != nil {
		return ShortenResponse{}, fmt.Errorf("can't create shortCode %w", err)
	}

	retry := 5
	for retry >= 0 {
		isUnique, err := s.repo.IsShortCodeUnique(shortCode)
		if err != nil {
			return ShortenResponse{}, fmt.Errorf("can't get uniquness of short code %w", err)
		}

		if isUnique {
			break
		}else {
			retry--
		}	
	}

	shortURL := entity.ShortURL{
		ID: 0,
		URL: req.URL,
		ShortCode: shortCode,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	createdShortURL, err := s.repo.Create(shortURL)
	if err != nil {
		return ShortenResponse{}, fmt.Errorf("can't create short url %w", err)
	}

	return ShortenResponse{createdShortURL}, nil

}