package shorten

import "fmt"

func (s Shorten) GetURL(shortCode string) (string, error) {
	//the function CheckExistenceOfShortCode return true if short code doesn't exist

	if doesShortCodeExist, _ := s.repo.CheckExistenceOfShortCode(shortCode); doesShortCodeExist {
		return "", fmt.Errorf("This shortcode doesn't exist")	
	}

	resp, err := s.repo.Read(shortCode)
	if err != nil {
		return "", fmt.Errorf("can't read data and get url %w", err)
	}

	url := resp.URL

	err = s.repo.IncrementVisit(shortCode)
	if err != nil {
		fmt.Errorf("errrr", err)
	}

	return url, nil
}