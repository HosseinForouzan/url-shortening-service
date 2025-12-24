package psql

import (
	"context"
	"fmt"

	"github.com/HosseinForouzan/url-shortening-service/entity"
	"github.com/jackc/pgx/v5"
)

func (p *psqlDB) IsShortCodeUnique(shortCode string) (bool, error) {
	var id int
	err := p.db.QueryRow(context.Background(), "SELECT id FROM short_url WHERE short_code = $1", shortCode).Scan(&id)
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
	err := p.db.QueryRow(context.Background(),"INSERT INTO short_url(long_url, short_code) VALUES($1, $2) RETURNING id ",
											sh.URL, sh.ShortCode).Scan(&id)
	if err != nil {
		return entity.ShortURL{}, fmt.Errorf("can't execute query %w", err)
	}

	sh.ID = id

	return sh, nil
	
}

