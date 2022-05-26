package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/samwhf/backendTest/api"
	"github.com/samwhf/backendTest/config"
	"github.com/samwhf/backendTest/database/postgres"
	us "github.com/samwhf/backendTest/services/user"
)

var (
	version                                               bool
	envFile                                               string
	Program, Version, CommitID, LastCommitTime, BuildTime string
)

// @title User Routing Service
// @Golang API REST
// @version 1.0
// @description API REST in Golang with Gin Framework

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @host localhost:8080
// @BasePath /api/v1
// @query.collection.format multi

func main() {
	log.Println("Starting the Rest-API Service!")

	flag.BoolVar(&version, "version", false, "Program version")
	flag.StringVar(&envFile, "env", "", "Environment Variable File")
	flag.Parse()
	if version {
		fmt.Printf("Program: %s \nVersion: %s \nCommitID: %s \nLastCommitTime: %s \nBuildTime: %s \nGo Version: %s\nGo OS/ARCH: %s %s\n",
			Program, Version, CommitID, LastCommitTime, BuildTime, runtime.Version(), runtime.GOOS, runtime.GOARCH)
		os.Exit(0)
	}
	if envFile == "" {
		log.Fatal("no ENV file provided")
	}

	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatal("error: ", err)
	}

	initContext := &gin.Context{}

	// configuration
	configuration, err := config.LoadEnvFile()
	if err != nil {
		log.Println("error: ", err)
	}
	log.Println("Configuration Variables:", configuration)
	// db
	log.Println("Connecting to postgres")
	database := postgres.New(*configuration)
	defer database.Close()
	err = us.CreateTable(initContext)
	if err != nil {
		log.Fatal("error: ", err)
	}

	// Configurations for CORS
	confg := cors.DefaultConfig()
	confg.AllowAllOrigins = true
	confg.AllowHeaders = []string{"Authorization", "Origin", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host"}
	confg.ExposeHeaders = []string{"Content-Length"}
	confg.AllowMethods = []string{"GET", "POST", "OPTIONS", "PUT", "PATCH", "DELETE"}
	confg.AllowCredentials = true

	router := gin.New()
	// Middlewares
	router.Use(cors.New(confg))
	router.Use(requestid.New())
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// setting up routes
	api.SetUpRoutes(router)

	srv := &http.Server{
		Addr:    configuration.Address,
		Handler: router,
	}

	go func() {
		// 在goroutine中初始化服务器，以使其不会阻止下面的正常关闭处理
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	// 等待中断信号以优雅地关闭服务器
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit //阻塞等待
	log.Println("Shutting Down GIN Server...")
	// 加个超时ctx，防止存活的连接处理时间超长（设置 5 秒的超时时间）
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// Shutdown处理存活的连接，阻止新连接
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Shutting Down GIN Server err:", err) //context deadline exceeded
	}
	log.Println("GIN Server exit")
}
