package main1

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/gin-gonic/gin"
)

type Crontab struct {
	Id   int    `json:"id"`
	Name string `json:"Name"`
}

type Login struct {
	User     string `form:"username" json:"user" uri:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("before middleware")
		//设置request变量到Context的Key中,通过Get等函数可以取得
		c.Set("request", "client_request")
		//发送request之前
		c.Next()

		//发送requst之后

		// 这个c.Write是ResponseWriter,我们可以获得状态等信息
		status := c.Writer.Status()
		fmt.Println("after middleware,", status)
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
}
func main() {
	router := gin.Default()

	router.Use(MiddleWare())

	router.GET("/middleware", func(c *gin.Context) {
		//获取gin上下文中的变量
		request := c.MustGet("request").(string)
		req, _ := c.Get("request")
		fmt.Println("request:", request)
		c.JSON(http.StatusOK, gin.H{
			"middile_request": request,
			"request":         req,
		})
	})
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})

	router.GET("/login", func(c *gin.Context) {
		c.String(http.StatusOK, "登录页面")
	})

	router.POST("/loginJSON", func(c *gin.Context) {
		var json Login
		//其实就是将request中的Body中的数据按照JSON格式解析到json变量中
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if json.User != "hanru" || json.Password != "hanru123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	//正则路由
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	//承接表单post
	router.POST("/form", func(c *gin.Context) {
		type1 := c.DefaultPostForm("type", "alert") //可设置默认值
		username := c.PostForm("username")
		password := c.PostForm("password")
		hobbys := c.PostFormArray("hobby")

		c.String(http.StatusOK, fmt.Sprintf("type is %s, username is %s, password is %s,hobby is %v", type1, username, password, hobbys))

	})

	//承接get参数
	router.GET("/welcome", func(c *gin.Context) {
		name := c.DefaultQuery("name", "Guest") //可设置默认值
		//nickname := c.Query("nickname") // 是 c.Request.URL.Query().Get("nickname") 的简写
		c.String(http.StatusOK, fmt.Sprintf("Hello %s ", name))
	})

	//加载模板
	router.LoadHTMLGlob("templates/*")
	router.GET("/index", func(c *gin.Context) {
		//根据完整文件名渲染模板，并传递参数
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})

	router.GET("/redirect", func(c *gin.Context) {
		//支持内部和外部的重定向
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com/")
	})

	//异步
	router.GET("/long_async", func(c *gin.Context) {
		// goroutine 中只能使用只读的上下文 c.Copy()
		cCp := c.Copy()
		go func() {
			time.Sleep(5 * time.Second)
			// 注意使用只读上下文
			log.Println("Done! in path " + cCp.Request.URL.Path)
		}()
	})

	router.GET("/user", func(c *gin.Context) {
		users, err := getAll()
		if err != nil {
			log.Fatal(err)
		}
		//H is a shortcut for map[string]interface{}
		c.JSON(http.StatusOK, gin.H{
			"result": users,
			"count":  len(users),
		})

	})

	router.Run(":8000")

}

//定义一个getALL函数用于回去全部的信息
func getAll() (crontabs []Crontab, err error) {

	//1.操作数据库
	db, _ := sql.Open("mysql", "crawler_bch_user:m0FeBjVNMDQps@tcp(123.59.190.202:58111)/bch_crawler?charset=utf8")
	//错误检查
	if err != nil {
		log.Fatal(err.Error())
	}
	//推迟数据库连接的关闭
	defer db.Close()

	//2.查询
	rows, err := db.Query("SELECT id, name FROM crontab")
	if err != nil {
		log.Fatal(err.Error())
	}

	for rows.Next() {
		//var crontab Crontab
		//遍历表中所有行的信息
		rows.Scan(&crontab.Id, &crontab.Name)
		//将user添加到users中
		crontabs = append(crontabs, crontab)
	}
	//最后关闭连接
	defer rows.Close()
	return
}
