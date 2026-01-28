package main

import (
	"crud-category/database"
	"crud-category/middleware"
	"crud-category/routes"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Port   string `mapstructure:"PORT"`
	DBConn string `mapstructure:"DB_CONN"`
}

func main() {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	config := Config{
		Port:   viper.GetString("PORT"),
		DBConn: viper.GetString("DB_CONN"),
	}

	db, err := database.InitDB(config.DBConn)
	if err != nil {
		fmt.Println("Failed to initialize database:", err)
		return
	}
	defer db.Close()

	router := routes.RegisterRoutes(db)
	loggedRouter := middleware.LoggingMiddleware(router)

	addr := "0.0.0.0:" + config.Port
	fmt.Println("Server started on port", addr)
	err = http.ListenAndServe(addr, loggedRouter)
	if err != nil {
		fmt.Println("Server failed to run", err)
	}
}
