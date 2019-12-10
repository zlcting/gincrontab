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
	//var cron *Crontab
	var crontabfeild []models.CrontabTable

	crontabfeild = models.QueryCronWightCon()
	// var scheduleTable map[string]*Crontab
	// scheduleTable = make(map[string]*Crontab)

	for k, v := range crontabfeild {
		fmt.Println(k)
		fmt.Println(v)
	}

	//Timg(cron)
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
func Timg(cron *Crontab) {
	fmt.Println(cron)
	var expr *cronexpr.Expression
	var err error
	// 当前时间
	now := time.Now()
	cron = &Crontab{
		expr:     expr,
		Command:  "",
		NextTime: expr.Next(now),
	}

	if expr, err = cronexpr.Parse("* 1 1 * * * *"); err != nil {
		fmt.Println(err)
		return
	}
	//expr = cronexpr.MustParse("* 1 1 * * * *")

	fmt.Println("当前时间：", now)
	// 下次调度时间
	nextTime := expr.Next(now)
	fmt.Println("下次时间：", nextTime)
	// 等待这个定时器超时
	time.AfterFunc(nextTime.Sub(now), func() {
		Run()
		now := time.Now()
		nextTime := expr.Next(now)
		fmt.Println("下次时间：", nextTime)
	}) // 下次时间减去当前时间

}
