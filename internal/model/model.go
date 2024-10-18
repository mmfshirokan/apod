package model

var Example = map[string]string{
	"copyright":       "\nBrennan Gilmore\n",
	"date":            "2024-10-14",
	"explanation":     "Go outside at sunset tonight and see a comet!  C/2023 A3 (Tsuchinshan–ATLAS) has become visible in the early evening sky in northern locations to the unaided eye. To see the comet, look west through a sky with a low horizon. If the sky is clear and dark enough, you will not even need binoculars -- the faint tail of the comet should be visible just above the horizon for about an hour.  Pictured, Comet Tsuchinshan-ATLAS was captured two nights ago over the Lincoln Memorial monument in Washington, DC, USA.  With each passing day at sunset, the comet and its changing tail should be higher and higher in the sky, although exactly how bright and how long its tails will be can only be guessed.   Growing Gallery: Comet Tsuchinsan-ATLAS in 2024",
	"hdurl":           "https://apod.nasa.gov/apod/image/2410/CometA3Dc_Gilmore_1080.jpg",
	"media_type":      "image",
	"service_version": "v1",
	"title":           "Comet Tsuchinshan-ATLAS Over the Lincoln Memorial",
	"url":             "https://apod.nasa.gov/apod/image/2410/CometA3Dc_Gilmore_1080.jpg",
}

type ImageInfo struct {
	Copyright      string
	Date           string
	Explanation    string
	UrlHD          string
	MediaType      string
	ServiceVersion string
	Title          string
	Url            string
}

// swagger:model ImageInfoResponse
type ImageInfoResponse struct {
	// Copyright of the image
	Copyright string `json:"copyright" example:"\nBrennan Gilmore\n"`

	// Date wnen image became apod
	Date string `json:"date" example:"2024-10-14"`

	// Explanation of the image
	Explanation string `json:"explanation" example:" Go outside at sunset tonight and see a comet!  C/2023 A3 (Tsuchinshan–ATLAS) has become visible..."`

	// HD URL of the image (link to nasa website)
	UrlHD string `json:"hdurl" example:"https://apod.nasa.gov/apod/image/2410/CometA3Dc_Gilmore_1080.jpg"`

	// MediaType of the date provided
	MediaType string `json:"media_type" example:"image"`

	// Nasa service version
	ServiceVersion string `json:"service_version" example:"v1"`

	// Title of the image provided
	Title string `json:"title" example:"Comet Tsuchinshan-ATLAS Over the Lincoln Memorial"`

	// URL of the image (link to nasa website)
	Url string `json:"url" example:"https://apod.nasa.gov/apod/image/2410/CometA3Dc_Gilmore_1080.jpg"`

	// URL of the stored image (proxy image)
	ProxyURL string `json:"proxy_url" example:"http://localhost:8089/2024-10-14.jpg"`
}
