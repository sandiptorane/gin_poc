package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sandiptorane/gin_poc/v1/handler"
	"log"
	"net/http"
	"time"
)

func Logger(userId string) gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		// Set userId variable
		c.Set("user_id", userId)

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}

func main() {
	server := new(handler.BookStore)
	r := gin.Default()
	r.Use(Logger("user1"))
	v1 := r.Group("/v1")
	{
		v1.POST("/create", server.CreateBook)
		v1.GET("/getbook/:id", server.GetBook)
		v1.GET("/getbooks", server.GetBooks)
		v1.PUT("/update/:id", server.UpdateBook)
		v1.DELETE("/delete/:id", server.DeleteBook)
	}
	v2 := r.Group("/v2")
	{
		//sample async process api
		v2.GET("/long_async", func(c *gin.Context) {
			// create copy to be used inside the goroutine
			cCp := c.Copy()
			go func() {
				time.Sleep(5 * time.Second)

				// note that you are using the copied context "cCp", IMPORTANT
				log.Println("Done! in path " + cCp.Request.URL.Path)
			}()
			c.JSON(http.StatusOK, gin.H{"msg": "process is started"})
		})
	}

	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
