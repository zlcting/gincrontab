package controllers

import (
	"fmt"
	"gospiderkeeper/models"
	"log"
	"net/http"
	"os/exec"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorhill/cronexpr"
)

type Crontab struct {
	NextTime time.Time
	Command  string
	expr     *cronexpr.Expression
}

type CrontabTable struct {
	ContabID int
	Name     string
	Command  string
}

func Crontabrun(c *gin.Context) {
	var cron Crontab
	var crontabfeild []models.CrontabTable
	var expr *cronexpr.Expression
	crontabfeild = models.QueryCronWightCon()
	var scheduleTable map[string]*Crontab
	scheduleTable = make(map[string]*Crontab)
	now := time.Now()

	for k, v := range crontabfeild {
		fmt.Println(k)
		fmt.Println(v)
		expr = cronexpr.MustParse(v.Timing)
		cron = Crontab{
			expr:     expr,
			Command:  "",
			NextTime: expr.Next(now),
		}

		scheduleTable[v.Name] = &cron

	}

	Timg(scheduleTable)
	//返回html
	c.HTML(http.StatusOK, "home.html", gin.H{"title": ""})
}

func Run() {
	fmt.Printf("combined out：")
	cmd := exec.Command("ls", "-lah")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("combined out:\n%s\n", string(out))
}

//nihao de
func Timg(scheduleTable map[string]*Crontab) {
	go func() {
		var (
			jobName string
			cron    *Crontab
			now     time.Time
		)
		for {
			now = time.Now()

			for jobName, cron = range scheduleTable {

				if cron.NextTime.Before(now) || cron.NextTime.Equal(now) {
					go func(string) {
						fmt.Println("执行", jobName)
					}(jobName)

					cron.NextTime = cron.expr.Next(now)
					fmt.Println(jobName, "下次执行时间：", cron.NextTime)
				}
			}
			select {
			case <-time.NewTimer(100 * time.Millisecond).C:

			}
		}
	}()
}
