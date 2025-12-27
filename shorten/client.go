package shorten

import (
	"fmt"
	"time"
)

func (s Shorten) GetURL(shortCode string) (string, error) {
	//the function CheckExistenceOfShortCode return true if short code doesn't exist
	if doesShortCodeExist, _ := s.repo.CheckExistenceOfShortCode(shortCode); doesShortCodeExist {
		return "", fmt.Errorf("This shortcode doesn't exist")	
	}


	var url string

	url, err := s.cachRepo.CachGet(shortCode)
	fmt.Println(url, err)
	if err != nil {
		resp , err := s.repo.Read(shortCode)
		if err != nil {
			return "", fmt.Errorf("can't read data and get url %w", err)
		}

		url = resp.URL

		_, err = s.cachRepo.CachSet(shortCode, url, 5 * time.Hour)
	}

	err = s.repo.IncrementVisit(shortCode)
	if err != nil {
		fmt.Errorf("errrr", err)
	}

	_, err = s.cachRepo.CachSet(shortCode, url, 5 * time.Hour)

	return url, nil
}