package main

import (
	"cron-service/app/config"
	cronjob "cron-service/cron"
	"cron-service/drivers/email"
	storagedriver "cron-service/drivers/minio"
	"cron-service/drivers/mysql"
	cache "cron-service/drivers/redis"
	repositories "cron-service/drivers/repository"
	"cron-service/usecase"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	getConfig := config.GetConfig()
	configdb := mysql.ConfigDB{
		DB_Username: getConfig.DB_USERNAME,
		DB_Password: getConfig.DB_PASSWORD,
		DB_Host:     getConfig.DB_HOST,
		DB_Port:     getConfig.DB_PORT,
		DB_Database: getConfig.DB_NAME,
	}
	db := configdb.InitialDb()

	configCache := cache.ConfigRedis{
		DB_Host: "207.148.77.251",
		DB_Port: "6379",
	}

	s3Config := storagedriver.MinioService{
		Host:     getConfig.STORAGE_URL,
		Username: getConfig.STORAGE_ID,
		Secret:   getConfig.STORAGE_SECRET,
	}

	s3 := s3Config.NewClient()

	gmail := email.SmtpConfig{
		CONFIG_SMTP_HOST:       getConfig.CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT:       getConfig.CONFIG_SMTP_PORT,
		CONFIG_SMTP_AUTH_EMAIL: getConfig.CONFIG_SMTP_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD:   getConfig.CONFIG_AUTH_PASSWORD,
		CONFIG_SENDER_NAME:     getConfig.CONFIG_SENDER_NAME,
	}

	dialer := email.NewGmailConfig(gmail)

	conn := configCache.InitRedis()
	repoStruct := repositories.NewRepository(db, conn)
	useCaseStruct := usecase.NewUsecase(repoStruct, s3, *dialer)
	cronApp := cronjob.NewCronCase(useCaseStruct)

	go cronApp.DoSchedule().Start()

	// trap SIGINT untuk trigger shutdown.
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}
