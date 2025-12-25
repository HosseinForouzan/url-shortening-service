package psql

import (
	"context"
	"fmt"
	"time"

	"github.com/HosseinForouzan/url-shortening-service/entity"
	"github.com/jackc/pgx/v5"
)

func (p *psqlDB) CheckExistenceOfShortCode(shortCode string) (bool, error) {
	var id int
	err := p.db.QueryRow(context.Background(), "SELECT id FROM urls WHERE short_code = $1", shortCode).Scan(&id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return true, nil
		}
		return false, fmt.Errorf("can't get uniquness of shortcode %w", err)
	}

	return false, nil
}



func (p *psqlDB) Create(sh entity.ShortURL) (entity.ShortURL, error) {
	var id int
	err := p.db.QueryRow(context.Background(),
	"INSERT INTO urls(long_url, short_code, created_at, updated_at) VALUES($1, $2, $3, $4) RETURNING id ",
											sh.URL, sh.ShortCode, sh.CreatedAt, sh.UpdatedAt).Scan(&id)
	if err != nil {
		return entity.ShortURL{}, fmt.Errorf("can't execute query %w", err)
	}

	sh.ID = id

	return sh, nil
	
}

func (p *psqlDB) Read(shortCode string) (entity.ShortURL, error) {
	var shortUrl entity.ShortURL
	err := p.db.QueryRow(context.Background(), "SELECT * FROM urls WHERE short_code = $1", shortCode).Scan(
		&shortUrl.ID, &shortUrl.URL, &shortUrl.ShortCode, &shortUrl.CreatedAt, &shortUrl.UpdatedAt)
	if err != nil {
		return entity.ShortURL{}, fmt.Errorf("can't read data %w", err)
	}

	return shortUrl, nil
	
}

func (p *psqlDB) Update(shortCode, url string) (entity.ShortURL, error) {
	var shortUrl entity.ShortURL
	err := p.db.QueryRow(context.Background(), "UPDATE urls SET long_url=$1, updated_at=$2 WHERE short_code =$3 RETURNING id, long_url, short_code, created_at, updated_at" ,
	 url, time.Now(),shortCode ).Scan(&shortUrl.ID, &shortUrl.URL, &shortUrl.ShortCode, &shortUrl.CreatedAt, &shortUrl.UpdatedAt)
	 	if err != nil {
		return entity.ShortURL{}, fmt.Errorf("can't update data %w", err)
	}

	return shortUrl, nil

}