package main

import (
	"flag"
	"github.com/golang/glog"
    "github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	flag.Parse()
}

func main() {
	router := gin.New()

    /**用户管理接口**/
	user := router.Group("/user")
	user.POST("/register", func(c *gin.Context) {
		c.String(http.StatusOK, "register user")
	})

	user.POST("/login", func(c *gin.Context) {
		c.String(http.StatusOK, "user login")
	})

    user.POST("/logout", func(c *gin.Context) {
		c.String(http.StatusOK, "user logout")
	})
    user.POST("/query", func(c *gin.Context) {
		c.String(http.StatusOK, "query user by userId")
	})

    /**专家管理接口**/
	expert := router.Group("/expert")
    export.POST("/add", func(c *gin.Context) {
		c.String(http.StatusOK, "add expert by userId")
	})
    export.POST("/delete", func(c *gin.Context) {
		c.String(http.StatusOK, "delete expert by userId")
	})
    export.POST("/search", func(c *gin.Context) {
		c.String(http.StatusOK, "delete expert by userId")
	})



    /**任务管理接口**/
	task := router.Group("/task")
	task.POST("/add", func(c *gin.Context) {
		c.String(http.StatusOK, "add task")
	})

	task.POST("/search", func(c *gin.Context) {
		c.String(http.StatusOK, "search task")
	})

    task.POST("/query", func(c *gin.Context) {
		c.String(http.StatusOK, "query task")
	})

    task.POST("/update", func(c *gin.Context) {
		c.String(http.StatusOK, "update task")
	})

    /**方案管理接口**/
	solution := router.Group("/solution")
    solution.POST("/build/:taskId", func(c *gin.Context) {
		c.String(http.StatusOK, "build solution")
	})
    solution.POST("/update/:solutionId", func(c *gin.Context) {
		c.String(http.StatusOK, "update solution")
	})
    solution.POST("/search", func(c *gin.Context) {
		c.String(http.StatusOK, "search solution")
	})
    solution.POST("/query/:solutionId", func(c *gin.Context) {
		c.String(http.StatusOK, "query solution")
	})

    /**消息管理接口**/
	message := router.Group("/message")
    message.GET("/query/:messageId", func(c *gin.Context) {
		c.String(http.StatusOK, "query message")
	})
    message.POST("/add", func(c *gin.Context) {
		c.String(http.StatusOK, "add message")
	})
    message.POST("/delete/:messageId", func(c *gin.Context) {
		c.String(http.StatusOK, "delete message")
	})

    glog.Info("start server")
	router.Run(":8000")
}
