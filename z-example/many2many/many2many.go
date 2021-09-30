package main

// import (
// 	"fmt"

// 	"gorm.io/gorm"
// )

// type User struct {
// 	gorm.Model
// 	Name          string
// 	Conversations []Conversation `gorm:"many2many:user_conversations;"`
// }

// type Conversation struct {
// 	gorm.Model
// 	Name  string
// 	Users []*User `gorm:"many2many:user_conversations;"`
// }

// func main() {

// 	db, err := gorm.Open(sqlite.Open("many2many.db"), &gorm.Config{})
// 	if err != nil {
// 		panic("failed to connect database")
// 	}

// 	// Migrate the schema
// 	err = db.AutoMigrate(&User{}, &Conversation{})
// 	if err != nil {
// 		fmt.Print(err)
// 	}

// 	userOne := User{
// 		Name: "User One",
// 	}
// 	userTwo := User{
// 		Name: "User Two",
// 	}
// 	// Create users
// 	db.Create(&userOne)
// 	db.Create(&userTwo)

// 	conversation := Conversation{
// 		Name: "Conversation One",
// 	}
// 	// Create conversation
// 	db.Create(&conversation)

// 	// Append users
// 	err = db.Model(&conversation).Association("Users").Append([]User{userOne, userTwo})

// 	if err != nil {
// 		fmt.Print(err)
// 	}

// 	for _, convUser := range conversation.Users {
// 		fmt.Println("Hello I am in the conversation: " + convUser.Name)
// 	}

// 	// Clean up database
// 	db.Delete(&userOne)
// 	db.Delete(&userTwo)
// 	db.Delete(&conversation)
// }
