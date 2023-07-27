package main

import (
	"account-app/internal/initializers"
	"account-app/internal/model/repository"
	"database/sql"
	"log"
)

func main() {
	config := initializers.LoadConfig("config.yml")
	db := initializers.Initialize(config)

	sqlDB, err := db.DB()
	defer func(sqlDB *sql.DB) {
		err := sqlDB.Close()
		if err != nil {
			log.Printf("Error on closing DB %v", err)
		}
	}(sqlDB)

	if err != nil {
		log.Fatalf("Can't get sql connection %v", err)
	}

	migrator := db.Migrator()

	migrator.DropTable(&repository.User{}, &repository.Account{}, &repository.AccountType{}, &repository.Plan{}, &repository.SubscriptionPlan{})

	err = db.AutoMigrate(&repository.User{}, &repository.Account{}, &repository.AccountType{}, &repository.Plan{}, &repository.SubscriptionPlan{})
	if err != nil {
		log.Fatalf("Error during migration %v", err)
	}
}
