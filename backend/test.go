package main

import(
	"github.com/gin-gonic/gin"
    "github.com/golang/glog"
	"net/http"
	"fmt"
)
func main() {
	router := gin.New()
	router.Use(glog)
	v1 := router.Group("/v1")
	// This handler will match /user/john but will not match /user/ or /user
	v1.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// However, this one will match /user/john/ and also /user/john/send
	// If no other routers match /user/john, it will redirect to /user/john/
	v1.GET("/user/:name/:action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	// For each matched request Context will hold the route definition
	v1.POST("/user/:name/*action", func(c *gin.Context) {
		c.String(http.StatusOK, c.FullPath())
	})

	v2 := router.Group("/v2")
	// This handler will match /user/john but will not match /user/ or /user
	v2.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})
	v2.POST("/upload", func(c *gin.Context) {
		// single file
		file, _ := c.FormFile("file")
		println(file.Filename)

		// Upload the file to specific dst.
		c.SaveUploadedFile(file, "./upload.txt")

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	}) 

	router.Run(":8080")
}
