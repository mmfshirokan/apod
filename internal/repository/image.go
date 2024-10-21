package repository

import (
	"io"
	"os"
)

type Image struct {
	location string
}

func NewImage(location string) Image {
	return Image{
		location: location,
	}
}

func (m Image) Add(image io.Reader, name string) error {
	f, err := os.Create(m.location + "/" + name + ".jpg")
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, image)

	return err
}
