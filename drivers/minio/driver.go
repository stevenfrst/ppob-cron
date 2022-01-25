package storagedriver

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
)

type MinioService struct {
	Host     string
	Username string
	Secret   string
}

// NewClient create new client
func (m *MinioService) NewClient() *minio.Client {
	s3Client, err := minio.New(m.Host, &minio.Options{
		//Creds:  credentials.NewStaticV4(m.Username, m.Secret, ""),
		Creds:  credentials.NewStaticV4("stevenhumam", "jambu123", ""),
		Secure: false,
	})
	if err != nil {
		log.Fatal(err)
	}
	return s3Client
}
