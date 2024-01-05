package main

import (
	"log"

	"github.com/Robi-Nurhidayat/simple-bank-gorm/account"
	"github.com/Robi-Nurhidayat/simple-bank-gorm/handler"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := "host=localhost user=root password=secret dbname=simple_bank port=5000 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("failed connect to database", err)
	}

	accountRepository := account.NewAccountRepository(db)
	accountService := account.NewAccountService(accountRepository)
	accountHandler := handler.NewAccountHandler(accountService)

	r := gin.Default()

	api := r.Group("/api/v1")

	api.POST("/account", accountHandler.CreateNewAccount)
	api.GET("/accounts", accountHandler.ListAccount)
	api.PATCH("/account/:id", accountHandler.UpdateAccount)
	api.DELETE("/account/:id", accountHandler.DeleteById)
	api.GET("/account/:id", accountHandler.GetAccountByID)
	r.Run()
}
