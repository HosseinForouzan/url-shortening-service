package shorten

import (
	"fmt"
	"time"

	"github.com/HosseinForouzan/url-shortening-service/entity"
)

const LengthOfShortCode = 7

type Repository interface {
	CheckExistenceOfShortCode(shortCode string)(bool, error)
	Create(sh entity.ShortURL) (entity.ShortURL, error)
	Read(shortCode string) (entity.ShortURL, error)
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
	ShortURL entity.ShortURL `json:"short_url"`
}

func (s Shorten) CreateService(req ShortenRequest) (ShortenResponse, error) {
	shortCode, err := randomShortCode(LengthOfShortCode)
	if err != nil {
		return ShortenResponse{}, fmt.Errorf("can't create shortCode %w", err)
	}

	retry := 5
	for retry >= 0 {
		isUnique, err := s.repo.CheckExistenceOfShortCode(shortCode)
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

type RetireveRequest struct {
	ShortCode string `json:"short_code"`
}

type RetireveResponse struct {
	ShortURL entity.ShortURL `json:"short_url"`
	Message string `json:"message"`
}

func (s Shorten) RetrieveService(req RetireveRequest) (RetireveResponse, error) {
	//the function CheckExistenceOfShortCode return true if short code doesn't exist 
	if doesShortCodeExist, _ := s.repo.CheckExistenceOfShortCode(req.ShortCode); doesShortCodeExist {
		return RetireveResponse{Message: "This shortcode doesn't exist"}, nil
	}

	RetrievedData, err := s.repo.Read(req.ShortCode)
	if err != nil {
		return RetireveResponse{}, fmt.Errorf("can't retrieve data %w", err)
	}

	shortUrl := entity.ShortURL {
		ID: RetrievedData.ID,
		URL: RetrievedData.URL,
		ShortCode: RetrievedData.ShortCode,
		CreatedAt: RetrievedData.CreatedAt,
		UpdatedAt: RetrievedData.UpdatedAt,
	}

	return RetireveResponse{ShortURL: shortUrl}, nil
}