package types

import "gorm.io/gorm"

type User struct {
	Id       string `json:"id" gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Location string `json:"location"`
}

type UserFriend struct {
	UserId   string `json:"userId"`
	FriendId string `json:"friendId"`
}

type UserFriendResponse struct {
	ErrorMsg string `json:"errorMsg"`
	UserFriend
}

type Seed struct {
	Name string
	Run  func(*gorm.DB) (*User, error)
}

type Post struct {
	Id      string `json:"id" gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	UserId  string `json:"userId"`
	Message string `json:"message"`
}

type Like struct {
	Id     string `json:"id" gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	UserId string `json:"userId"`
	PostId string `json:"postId"`
}

type SharePost struct {
	Id     string `json:"id" gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	UserId string `json:"userId"`
	PostId string `json:"postId"`
}
