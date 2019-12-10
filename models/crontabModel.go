package models

import "fmt"
import "gospiderkeeper/database"

type CrontabTable struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Command    string `json:"command"`
	Timing     string `json:"timing"`
	Createtime int64  `json:"createtime"`
}

func QueryCronWightCon() (crontabs []CrontabTable) {
	sql := fmt.Sprintf("select id,name,command,timing,createtime from crontab")
	fmt.Println(sql)
	row := database.QueryRowDB(sql)
	for row.Next() {
		var cronfelid CrontabTable
		row.Scan(&cronfelid.Id, &cronfelid.Name, &cronfelid.Command, &cronfelid.Timing, &cronfelid.Createtime)
		//将user添加到users中
		crontabs = append(crontabs, cronfelid)
	}
	return crontabs
}
