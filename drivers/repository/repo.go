package repositories

import (
	"cron-service/models"
	"github.com/golang-module/carbon"
	"github.com/gomodule/redigo/redis"
	"github.com/thanhpk/randstr"
	"gorm.io/gorm"
	"log"
)

type Repository struct {
	db    *gorm.DB
	redis redis.Conn
}

func NewRepository(db *gorm.DB, cache redis.Conn) Repository {
	return Repository{
		db:    db,
		redis: cache,
	}
}

func (r *Repository) GetAdmin() string {
	var repoModel models.User
	r.db.Where("role = ?", "admin").First(&repoModel)
	return repoModel.Email
}

func (r *Repository) GetTodayTx() []models.CSVModels {
	var repoModel []models.Transaction
	r.db.Preload("DetailTransaction").Where("updated_at BETWEEN ? AND ?", carbon.Now().StartOfDay(), carbon.Now().EndOfDay()).Find(&repoModel)
	log.Println(repoModel)
	return models.CsvModelList(repoModel)
}

func (r *Repository) RandomizeUser() string {
	var repoModel models.User
	r.db.Where("role = ?", "user").Take(&repoModel)
	return repoModel.Email
}

func (r *Repository) CreateRandomVoucher() string {
	var repoModel models.Voucher
	token := randstr.String(16)
	repoModel = models.Voucher{
		Code:  token,
		Value: 5000,
		Valid: carbon.Tomorrow().Time,
	}
	r.db.Save(&repoModel)
	return token
}
