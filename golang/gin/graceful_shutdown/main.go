package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		time.Sleep(200 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		// if err := router.Run(":8080"); err != nil && err != http.ErrServerClosed {
		// 	log.Fatalf("listen: %s\n", err)
		// }
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 信号知识点
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("shutdown server...")

	// context的WithTimeout超时知识点
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// http.Server优雅关停的知识点
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown ", err)
	}

	log.Println("Server existing")
}
