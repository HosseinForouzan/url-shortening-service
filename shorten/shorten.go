package shorten

import (
	"fmt"
	"time"

	"github.com/HosseinForouzan/url-shortening-service/entity"
)

const LengthOfShortCode = 7

type Repository interface {
	IncrementVisit(shortCode string) (error)
	CheckExistenceOfShortCode(shortCode string)(bool, error)
	Create(sh entity.ShortURL) (entity.ShortURL, error)
	Read(shortCode string) (entity.ShortURL, error)
	Update(short_code, url string) (entity.ShortURL, error)
	Delete(shortCode string) (error)
	GetStats(shortCode string) (entity.Stats, error)

}

type CachRepo interface {
	CachSet(key string, value interface{}, expiration time.Duration) (string, error)
	CachGet(key string) (string, error)
}

type Shorten struct {
	repo Repository
	cachRepo CachRepo
}

func New(repo Repository, cachRepo CachRepo) Shorten {
	return Shorten{repo: repo, cachRepo: cachRepo}
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
}

func (s Shorten) RetrieveService(req RetireveRequest) (RetireveResponse, error) {
	//the function CheckExistenceOfShortCode return true if short code doesn't exist 
	if doesShortCodeExist, _ := s.repo.CheckExistenceOfShortCode(req.ShortCode); doesShortCodeExist {
		return RetireveResponse{}, fmt.Errorf("This shortcode doesn't exist")
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

type UpdateRequest struct {
	URL string `json:"url"`
	ShortCode string `json:"short_code"`
}

type UpdateResponse struct {
	ShortURL entity.ShortURL `json:"short_url"`
}

func (s Shorten) UpdateService(req UpdateRequest) (UpdateResponse, error) {
	//the function CheckExistenceOfShortCode return true if short code doesn't exist 
	if doesShortCodeExist, _ := s.repo.CheckExistenceOfShortCode(req.ShortCode); doesShortCodeExist {
		return UpdateResponse{}, fmt.Errorf("This shortcode doesn't exist")
	}

	updatedShortUrl, err := s.repo.Update(req.ShortCode, req.URL)
	fmt.Println(req.ShortCode, req.URL)
	if err != nil{
		return UpdateResponse{}, fmt.Errorf("can't update record %w", err)
	}

	return UpdateResponse{ShortURL: updatedShortUrl}, nil

}

type DeleteRequest struct {
	ShortCode string `json:"short_code"`
}

type DeleteResponse struct {
	Message string `json:"message"`
}


func (s Shorten) DeleteService(req DeleteRequest) (DeleteResponse, error) {
	//the function CheckExistenceOfShortCode return true if short code doesn't exist 
	if doesShortCodeExist, _ := s.repo.CheckExistenceOfShortCode(req.ShortCode); doesShortCodeExist {
		return  DeleteResponse{} ,fmt.Errorf("This shortcode doesn't exist")
	}

	err := s.repo.Delete(req.ShortCode)
	if err != nil {
		return DeleteResponse{} ,fmt.Errorf("can't delete record %w", err)
	}

	return DeleteResponse{Message: "recored deleted"} ,nil


}

type StatsRequest struct {
	ShortCode string `json:"short_cde"`
}

type StatsResponse struct {
	Stats entity.Stats `json:"stats"`
}

func (s Shorten) GetStatsService(req StatsRequest) (StatsResponse, error) {
		//the function CheckExistenceOfShortCode return true if short code doesn't exist 
	if doesShortCodeExist, _ := s.repo.CheckExistenceOfShortCode(req.ShortCode); doesShortCodeExist {
		return StatsResponse{}, fmt.Errorf("This shortcode doesn't exist")
	}

	RetrievedData, err := s.repo.GetStats(req.ShortCode)
	if err != nil {
		return StatsResponse{}, fmt.Errorf("can't retrieve stats %w", err)
	}

	stats := entity.Stats {
		ShortURL: entity.ShortURL{
			ID: RetrievedData.ShortURL.ID,
			URL: RetrievedData.ShortURL.URL,
			ShortCode: RetrievedData.ShortURL.ShortCode,
			CreatedAt: RetrievedData.ShortURL.CreatedAt,
			UpdatedAt: RetrievedData.ShortURL.UpdatedAt,
		},
		Visits: RetrievedData.Visits,	
	}

	return StatsResponse{Stats: stats}, nil
}