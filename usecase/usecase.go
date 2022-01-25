package usecase

import (
	"context"
	repositories "cron-service/drivers/repository"
	"fmt"
	"github.com/golang-module/carbon"
	"github.com/jszwec/csvutil"
	"github.com/minio/minio-go/v7"
	"gopkg.in/gomail.v2"
	"log"
	"os"
)

type Usecase struct {
	repo repositories.Repository
	s3   *minio.Client
	mail gomail.Dialer
}

func NewUsecase(repo repositories.Repository, s3 *minio.Client, mail gomail.Dialer) Usecase {
	return Usecase{repo: repo, s3: s3, mail: mail}
}

func (u Usecase) LuckyVoucher() {
	resp := u.repo.RandomizeUser()

	codeVoucher := u.repo.CreateRandomVoucher()
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", u.mail.Username)
	mailer.SetHeader("To", resp)
	mailer.SetHeader("Subject", "Selemat Anda Mendapatkan Diskon Potongan")
	mailer.SetBody("text/html", fmt.Sprintf("Selamat Anda Mendapatkan Voucher Potongan Rp.5000 <b>%v<b> Berlaku Sampai %v", codeVoucher, carbon.Tomorrow()))

	err := u.mail.DialAndSend(mailer)
	if err != nil {
		log.Println(err)
		return
	}
}

func (u Usecase) PushEmail() {
	// Get Email
	resp := u.repo.GetAdmin()

	// GET Data Today
	todayTx := u.repo.GetTodayTx()
	log.Println(todayTx)

	// Format to CSV
	b, _ := csvutil.Marshal(todayTx)
	log.Println(string(b))

	// Create File
	emptyCsv, _ := os.Create("report.csv")
	emptyCsv.Write(b)

	if _, err := u.s3.FPutObject(context.Background(), "static", "report.csv", "report.csv", minio.PutObjectOptions{
		ContentType: "text/csv",
	}); err != nil {
		log.Println(err)
		return
	}
	emptyCsv.Close()
	defer func() {
		os.Remove("report.csv")
	}()
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", u.mail.Username)
	mailer.SetHeader("To", resp)
	mailer.SetHeader("Subject", "Laporan Penjualan Harian")
	mailer.SetBody("text/html", fmt.Sprintf("Halo <b>Admoon<b>, ini laporan penjualan pada %v cheers h3h3h3", carbon.Now().ToDateString()))
	mailer.Attach("report.csv")

	err := u.mail.DialAndSend(mailer)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Email Send")
}
