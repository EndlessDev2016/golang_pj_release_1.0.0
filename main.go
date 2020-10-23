package main

import (
	"flag"
	"fmt"
	"go-gin-pj/controller"
	"go-gin-pj/middleware"
	"os"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/ilyakaznacheev/cleanenv"
)

// Config is a application configuration structure
type Config struct {
	Database struct {
		Host        string `yaml:"host" env:"DB_HOST" env-description:"Database host"`
		Port        string `yaml:"port" env:"DB_PORT" env-description:"Database port"`
		Username    string `yaml:"username" env:"DB_USER" env-description:"Database user name"`
		Password    string `env:"DB_PASSWORD" env-description:"Database user password"`
		Name        string `yaml:"db-name" env:"DB_NAME" env-description:"Database name"`
		Connections int    `yaml:"connections" env:"DB_CONNECTIONS" env-description:"Total number of database connections"`
	} `yaml:"database"`
	Server struct {
		Host string `yaml:"host" env:"SRV_HOST,HOST" env-description:"Server host" env-default:"localhost"`
		Port string `yaml:"port" env:"SRV_PORT,PORT" env-description:"Server port" env-default:"8080"`
	} `yaml:"server"`
	Greeting string `env:"GREETING" env-description:"Greeting phrase" env-default:"Hello!"`
}

// Args command-line parameters
type Args struct {
	ConfigPath string
}

// ProcessArgs processes and handles CLI arguments
func ProcessArgs(cfg interface{}) Args {
	var a Args

	f := flag.NewFlagSet("Example server", 1)
	f.StringVar(&a.ConfigPath, "c", "config.yml", "Path to configuration file")

	fu := f.Usage
	f.Usage = func() {
		fu()
		envHelp, _ := cleanenv.GetDescription(cfg, nil)
		fmt.Fprintln(f.Output())
		fmt.Fprintln(f.Output(), envHelp)
	}

	f.Parse(os.Args[1:])
	return a
}

// Request -> Route Parser -> Middleware -> Route Handler -> Middleware -> Response
func main() {

	var cfg Config
	args := ProcessArgs(&cfg)
	// read configuration from the file and environment variables
	if err := cleanenv.ReadConfig(args.ConfigPath, &cfg); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	fmt.Println(cfg.Server.Host)

	router := gin.Default()

	// use the Middleware
	// https://github.com/gin-gonic/gin#using-middleware
	// router.Use(func(c *gin.Context) {
	// 	c.GetHeader("User-Agent")
	// 	c.Next()
	// })

	// use the Templates
	// router.LoadHTMLGlob("templates/*")
	// router.GET("/", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message":    "hello world",
	// 		"User-Agent": ua,
	// 	})
	// })

	// static file
	// router.Static("/assets", "./assets")
	// router.StaticFS("/more_static", http.Dir("my_file_system"))
	// router.StaticFile("/favicon.ico", "./resources/favicon.ico")

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	// Set Middleware
	router.Use(middleware.WebTestUaAndTime)

	// Setup route group for the API
	api := router.Group("/api")
	{
		{
			v1 := api.Group("/v1")
			v1.GET("/test", controller.WebTest)
		}
	}

	router.Run(":3000")
}
