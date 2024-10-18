package consumer

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/mmfshirokan/apod/internal/model"
	log "github.com/sirupsen/logrus"
)

type InfoAdder interface {
	Add(ctx context.Context, info model.ImageInfo) error
}
type ImageAdder interface {
	Add(image io.Reader, name string) error
}

type Consumer struct {
	inf InfoAdder
	img ImageAdder
}

func New(inf InfoAdder, img ImageAdder) *Consumer {
	return &Consumer{
		inf: inf,
		img: img,
	}
}

func (c *Consumer) Consume(ctx context.Context, target, key string) {
	req, err := http.NewRequest(http.MethodGet, target+"?api_key="+key, nil)
	if err != nil {
		log.Fatal("Wrong http request: ", err)
	}

	reqAndSave := func() {
		// Sending http request
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatal("Can't obtain http response, uttempt failed: ", err)
		}

		// Parsing JSON
		ii := model.ImageInfo{}
		if err = json.NewDecoder(resp.Body).Decode(&ii); err != nil {
			log.Fatal("Can't decode JSON: ", err)
		}
		resp.Body.Close()

		// Adding image info
		if err = c.inf.Add(ctx, ii); err != nil {
			log.Error("Can't Add in DB: ", err)
		}

		// Creating Image request
		iReq, err := http.NewRequest(http.MethodGet, target+"?api_key="+key, nil)
		if err != nil {
			log.Error("Wrong image http request: ", err)
		}

		// Sending http request for image
		iResp, err := http.DefaultClient.Do(iReq)
		if err != nil {
			log.Error("Can't obtain http response, uttempt failed: ", err)
		}

		// Adding image
		if err = c.img.Add(iResp.Body.(io.Reader), ii.Date); err != nil {
			log.Error("Can't Add in DB: ", err)
		} else {
			log.Info("Data successfully saved on ", time.Now().UTC())
		}
		defer iResp.Body.Close()
	}

	reqAndSave()

	tiker := time.NewTicker(tillNextDayUTC())

	for {
		select {
		case <-ctx.Done():
			{
				tiker.Stop()
				return
			}
		case <-tiker.C:
			{
				reqAndSave()
				tiker.Reset(tillNextDayUTC())
			}
		}
	}
}

func tillNextDayUTC() time.Duration {
	return time.Date(
		time.Now().Year(),
		time.Now().Month(),
		time.Now().Day()+1,
		0, 0, 1, 0,
		time.UTC,
	).Sub(time.Now().UTC())
}
