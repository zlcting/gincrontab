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

var CrontabStartFlag chan int

var CrontabStop chan bool

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
			Command:  v.Command,
			NextTime: expr.Next(now),
		}

		scheduleTable[v.Name] = &cron

	}
	CrontabStartFlag = make(chan int, 1)
	CrontabStartFlag <- 1
	go Timg(scheduleTable)

	//返回html
	c.HTML(http.StatusOK, "home.html", gin.H{"title": ""})
}

func CrontabList(c *gin.Context) {

	var crontabfeild []models.CrontabTable

	crontabfeild = models.QueryCronWightCon()
	c.HTML(http.StatusOK, "crontablist.html", gin.H{"Content": crontabfeild, "StartFlag": len(CrontabStartFlag)})
}

func CrontabStopAction(c *gin.Context) {
	CrontabStop = make(chan bool, 1)
	CrontabStop <- true
}

func Run(command string) {
	fmt.Printf("combined out：")
	cmd := exec.Command(command)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("combined out:\n%s\n", string(out))
}

//nihao de
func Timg(scheduleTable map[string]*Crontab) {
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
					fmt.Println("执行任务名", jobName)
					fmt.Println("执行命令", cron.Command)
					Run(cron.Command)
				}(jobName)

				cron.NextTime = cron.expr.Next(now)
				fmt.Println(jobName, "下次执行时间：", cron.NextTime)
			}
		}
		//等待100毫秒
		select {
		case <-time.NewTimer(100 * time.Millisecond).C:

		}
		// go stop := <-CrontabStop

		// if stop {
		// 	break
		// }

	}
	fmt.Println("crontab任务退出")
}
