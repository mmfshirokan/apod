package model

type ImageInfo struct {
	Copyright      string `json:"copyright"`
	Date           string `json:"date"`
	Explanation    string `json:"explanation"`
	UrlHD          string `json:"hdurl"`
	MediaType      string `json:"media_type"`
	ServiceVersion string `json:"service_version"`
	Title          string `json:"title"`
	Url            string `json:"url"`
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
