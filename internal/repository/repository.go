package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mmfshirokan/apod/internal/model"
)

type Postgres struct {
	db *pgxpool.Pool
}

func NewInfo(db *pgxpool.Pool) *Postgres {
	return &Postgres{
		db: db,
	}
}

func (p *Postgres) Add(ctx context.Context, ii model.ImageInfo) error {
	_, err := p.db.Exec(ctx, "INSERT INTO image.info (copyright, image_date, explanation, hd_url, media_type, service_version, title, url) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		ii.Copyright,
		ii.Date,
		ii.Explanation,
		ii.UrlHD,
		ii.MediaType,
		ii.ServiceVersion,
		ii.Title,
		ii.Url,
	)

	return err
}

func (p *Postgres) Get(ctx context.Context, date string) (model.ImageInfo, error) {
	var (
		result model.ImageInfo
		tm     time.Time
	)

	err := p.db.QueryRow(ctx, "SELECT copyright, image_date, explanation, hd_url, media_type, service_version, title, url FROM image.info WHERE image_date = $1", date).Scan(
		&result.Copyright,
		&tm,
		&result.Explanation,
		&result.UrlHD,
		&result.MediaType,
		&result.ServiceVersion,
		&result.Title,
		&result.Url,
	)

	result.Date = tm.Format(time.DateOnly)

	return result, err
}

func (p *Postgres) GetAll(ctx context.Context) ([]model.ImageInfo, error) {
	var (
		result []model.ImageInfo
		tm     time.Time
	)

	rows, err := p.db.Query(ctx, "SELECT copyright, image_date, explanation, hd_url, media_type, service_version, title, url FROM image.info")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var ii model.ImageInfo

		if err = rows.Scan(
			&ii.Copyright,
			&tm,
			&ii.Explanation,
			&ii.UrlHD,
			&ii.MediaType,
			&ii.ServiceVersion,
			&ii.Title,
			&ii.Url,
		); err != nil {
			return nil, err
		}

		ii.Date = tm.Format(time.DateOnly)
		result = append(result, ii)
	}

	return result, nil
}
