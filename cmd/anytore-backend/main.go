package main

import (
	"context"
	"fmt"

	"github.com/openkrafter/anytore-backend/auth"
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

	if len(users) == 0 {
		fmt.Println("No users found")
		return
	}

	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s, Email: %s, Password: %s\n",
			user.Id, user.Name, user.Email, user.Password)

		// fmt.Println("password")
		// err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte("password"))
		// if err != nil {
		// 	fmt.Println("Password Error:", err)
		// 	return
		// }

		// fmt.Println("passpass")
		// err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte("passpass"))
		// if err != nil {
		// 	fmt.Println("Password Error:", err)
		// 	return
		// }

	}

}

func TmpCreateUser() {
	ctx := context.Background()

	user := &model.User{
		Name:     "test",
		Email:    "user1@example.com",
		Password: "password",
	}

	err := service.CreateUser(ctx, user)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("User created successfully")
}

func TmpUpdateUser() {
	ctx := context.Background()

	user := &model.User{
		Id:       2,
		Name:     "test2",
		Email:    "user2@example.com",
		Password: "password2",
	}
	// user := &model.User{}

	err := service.UpdateUser(ctx, user)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("User updated successfully")
}

func TmpDeleteUser() {
	ctx := context.Background()

	userId := 2
	err := service.DeleteUser(ctx, userId)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("User deleted successfully")
}

func main() {
	logger.InitLogger()
	config.InitConfig()
	dynamodb.InitDynamoDbClient()

	err := sqldb.InitSQLDBClient()
	if err != nil {
		logger.Logger.Error("Failed to initialize SQLDB client", logger.ErrAttr(err))
		return
	}

	auth.InitPassHasher()

	// controller.Run()

	TmpGetUsers()
	// TmpCreateUser()
	// TmpUpdateUser()
	// TmpDeleteUser()
}
