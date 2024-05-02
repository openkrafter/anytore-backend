package main

import (
	"context"
	"fmt"

	"github.com/openkrafter/anytore-backend/config"
	"github.com/openkrafter/anytore-backend/logger"
	"github.com/openkrafter/anytore-backend/model"
	"github.com/openkrafter/anytore-backend/service"

	"github.com/openkrafter/anytore-backend/database/dynamodb"
	"github.com/openkrafter/anytore-backend/database/sqldb"

	_ "github.com/go-sql-driver/mysql"
)

func TmpGetUsers() {
	ctx := context.Background()
	users, err := service.GetUsers(ctx)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s, Email: %s\n", user.ID, user.Name, user.Email)
	}

}

func TmpCreateUser() {
	ctx := context.Background()

	user := &model.User{
		Name:     "test",
		Email:    "user1@example.com",
		Password: "password",
		Salt:     "salt",
	}

	err := service.CreateUser(ctx, user)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("User created successfully")
}

func main() {
	logger.InitLogger()
	config.InitConfig()
	dynamodb.InitDynamoDbClient()
	sqldb.InitSQLDBClient()

	// controller.Run()
	TmpGetUsers()
	// TmpCreateUser()
}
