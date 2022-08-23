package main

import (
	"os"
	"sirka-be-test/pkg/common/db"
	"sirka-be-test/pkg/users"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func getEnv(key, defaultValue string) string {
    value := os.Getenv(key)
    if len(value) == 0 {
        return defaultValue
    }
    return value
}

func main() {
    var port string
    var dbUrl string

    isOnHeroku := os.Getenv("DATABASE_URL") != ""
    isOnDocker := os.Getenv("DOCKER_DATABASE_URL") != ""

    if isOnHeroku {
        port = ":" + os.Getenv("DOCKER_PORT")
        dbUrl = os.Getenv("DATABASE_URL")
    } else if isOnDocker {
        port = os.Getenv("DOCKER_PORT")
        dbUrl = os.Getenv("DOCKER_DATABASE_URL")
    } else {
        viper.SetConfigFile("./pkg/common/env/.env")
        viper.ReadInConfig()

        port = viper.Get("PORT").(string)
        dbUrl = viper.Get("DB_URL").(string)
    }

	r := gin.Default()
	h := db.Init(dbUrl)

	users.RegisterRoutes(r, h)

	r.Run(port)
}
