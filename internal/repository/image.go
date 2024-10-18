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

// func (m Image) addLoop(image io.Reader, name string) error {
// 	f, err := os.Create(m.location + "/" + name + ".jpg")
// 	if err != nil {
// 		return err
// 	}
// 	defer f.Close()

// 	for {
// 		i, err := f.ReadFrom(image)
// 		if err == io.EOF {
// 			return nil
// 		}
// 		if err != nil {
// 			return err
// 		}
// 		if i == 0 {
// 			return nil
// 		}
// 	}
// }

// type Minio struct {
// 	db         *minio.Client
// 	bucketName string
// }

// func NewImage(db *minio.Client, bucketName string) *Minio {
// 	return &Minio{
// 		db:         db,
// 		bucketName: bucketName,
// 	}
// }

// func (m *Minio) Add(ctx context.Context, image io.ReadCloser, name string) error {
// 	_, err := m.db.PutObject(ctx, m.bucketName, name, image, -1, minio.PutObjectOptions{})
// 	return err
// }
