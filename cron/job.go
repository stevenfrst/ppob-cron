package cronjob

import (
	"cron-service/usecase"
	"github.com/robfig/cron/v3"
	"time"
)

type CronCase struct {
	uc usecase.Usecase
}

func NewCronCase(uc usecase.Usecase) CronCase {
	return CronCase{uc: uc}
}

func (c CronCase) DoSchedule() *cron.Cron {
	jakartaTime, _ := time.LoadLocation("Asia/Jakarta")
	scheduler := cron.New(cron.WithLocation(jakartaTime))

	scheduler.AddFunc("0 21 * * *", c.uc.PushEmail)
	//scheduler.AddFunc("* * * * *", c.uc.PushEmail)
	//scheduler.AddFunc("* * * * *", c.uc.LuckyVoucher)

	return scheduler
}
