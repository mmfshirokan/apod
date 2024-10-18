package repository

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mmfshirokan/apod/internal/model"
	"github.com/ory/dockertest"
	"github.com/ory/dockertest/docker"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

var (
	conn *Postgres

	testInfo = []model.ImageInfo{
		{
			Copyright:      "\nBrennan Gilmore\n",
			Date:           "2024-10-14",
			Explanation:    "Go outside at sunset tonight and see a comet! C/2023 A3 (Tsuchinshan–ATLAS) has become visible in the early evening sky in northern locations to the unaided eye. To see the comet, look west through a sky with a low horizon. If the sky is clear and dark enough, you will not even need binoculars -- the faint tail of the comet should be visible just above the horizon for about an hour. Pictured, Comet Tsuchinshan-ATLAS was captured two nights ago over the Lincoln Memorial monument in Washington, DC, USA. With each passing day at sunset, the comet and its changing tail should be higher and higher in the sky, although exactly how bright and how long its tails will be can only be guessed. Growing Gallery: Comet Tsuchinsan-ATLAS in 2024",
			UrlHD:          "https://apod.nasa.gov/apod/image/2410/CometA3Dc_Gilmore_1080.jpg",
			MediaType:      "image",
			ServiceVersion: "v1",
			Title:          "Comet Tsuchinshan-ATLAS Over the Lincoln Memorial",
			Url:            "https://apod.nasa.gov/apod/image/2410/CometA3Dc_Gilmore_1080.jpg",
		},
		{
			Copyright:      "\nMaria Johnson\n",
			Date:           "2024-10-15",
			Explanation:    "Tonight, gaze at the Moon and witness a stunning lunar eclipse! The partial eclipse will begin at 8:15 PM EDT, with the total phase starting at 9:30 PM EDT. This spectacular event will be visible across much of North America. In this image, the eclipsed Moon is shown rising over the Grand Canyon, illuminating the iconic landscape with a soft, eerie glow. Don't miss this rare opportunity to see the celestial dance!  Growing Gallery: Lunar Eclipses of 2024",
			UrlHD:          "https://apod.nasa.gov/apod/image/2410/LunarEclipseGrandCanyon_Johnson_1080.jpg",
			MediaType:      "image",
			ServiceVersion: "v1",
			Title:          "Lunar Eclipse Over the Grand Canyon",
			Url:            "https://apod.nasa.gov/apod/image/2410/LunarEclipseGrandCanyon_Johnson_1080.jpg",
		},
		{
			Copyright:      "\nAlex Chen\n",
			Date:           "2024-10-16",
			Explanation:    "Prepare for an enchanting night as the Orionids meteor shower peaks tonight! This annual event, caused by debris from Halley's Comet, promises to light up the sky with up to 20 meteors per hour. Best viewing will be after midnight in dark areas away from city lights. The photo showcases a previous year’s shower over a serene lake, capturing the beauty of nature and the cosmos. Grab a blanket and enjoy the celestial show!  Growing Gallery: Meteor Showers of 2024",
			UrlHD:          "https://apod.nasa.gov/apod/image/2410/OrionidsLake_Chen_1080.jpg",
			MediaType:      "image",
			ServiceVersion: "v1",
			Title:          "Orionids Meteor Shower Over a Tranquil Lake",
			Url:            "https://apod.nasa.gov/apod/image/2410/OrionidsLake_Chen_1080.jpg",
		},
		{
			Copyright:      "\nSofia Martinez\n",
			Date:           "2024-10-17",
			Explanation:    "Tonight, experience the beauty of the Pleiades star cluster, also known as the Seven Sisters. Best viewed in the late evening, this open star cluster in the constellation Taurus is a stunning sight through binoculars or a telescope. The photo captures the cluster shining brightly above a field of wildflowers, showcasing its ethereal glow. Join us in appreciating this celestial gem and the wonders of our universe!  Growing Gallery: Star Clusters in 2024",
			UrlHD:          "https://apod.nasa.gov/apod/image/2410/PleiadesWildflowers_Martinez_1080.jpg",
			MediaType:      "image",
			ServiceVersion: "v1",
			Title:          "The Pleiades Star Cluster Over Wildflowers",
			Url:            "https://apod.nasa.gov/apod/image/2410/PleiadesWildflowers_Martinez_1080.jpg",
		},
	}
)

func TestMain(m *testing.M) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Error("Could not construct pool: ", err)
		return
	}

	err = pool.Client.Ping()
	if err != nil {
		log.Error("Could not connect to docker: ", err)
		return
	}

	pg, err := pool.RunWithOptions(&dockertest.RunOptions{
		Hostname:   "postgres_test",
		Repository: "postgres",
		Tag:        "latest",
		Env: []string{
			"POSTGRES_PASSWORD=password",
			"POSTGRES_USER=user",
			"POSTGRES_DB=db",
			"listen_addresses = '*'",
		},
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	if err != nil {
		log.Error("Could not start resource: ", err)
		return
	}

	postgresHostAndPort := pg.GetHostPort("5432/tcp")
	postgresUrl := fmt.Sprintf("postgres://user:password@%s/db?sslmode=disable", postgresHostAndPort)

	log.Info("Connecting to database on url: ", postgresUrl)

	var dbpool *pgxpool.Pool
	if err = pool.Retry(func() error {
		dbpool, err = pgxpool.New(ctx, postgresUrl)
		if err != nil {
			dbpool.Close()
			log.Error("can't connect to the pgxpool: %w", err)
		}
		return dbpool.Ping(ctx)
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	commandArr := []string{
		"-url=jdbc:postgresql://" + postgresHostAndPort + "/db",
		"-user=user",
		"-password=password",
		"-locations=filesystem:../../migrations/",
		"-schemas=image",
		"-connectRetries=60",
		"migrate",
	}

	cmd := exec.Command("flyway", commandArr[:]...)

	err = cmd.Run()
	if err != nil {
		log.Error(fmt.Print("error: ", err))
	}

	pool.MaxWait = 120 * time.Second
	conn = NewInfo(dbpool)

	code := m.Run()

	if err := pool.Purge(pg); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

func TestAdd(t *testing.T) {
	ctx := context.Background()

	for i, tt := range testInfo {
		err := conn.Add(ctx, tt)
		assert.Nil(t, err, "test case: %d", i)
	}

	t.Log("TestAdd finished!")
}

func TestGet(t *testing.T) {
	for i, tt := range testInfo {
		ii, err := conn.Get(context.Background(), tt.Date)
		assert.Nil(t, err, "test case: %d", i)
		assert.Equal(t, tt, ii)
	}

	t.Log("TestGet finished!")
}

func TestGetAll(t *testing.T) {
	ctx := context.Background()

	all, err := conn.GetAll(ctx)

	assert.ElementsMatch(t, testInfo, all)
	assert.Nil(t, err)

	t.Log("TestGetAll finished!")
}
