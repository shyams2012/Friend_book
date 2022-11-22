package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/shyams2012/friend_book/db"
	"github.com/shyams2012/friend_book/types"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Create user profile
func CreateProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		db := db.DbConn()
		var user = types.User{}

		body, _ := ioutil.ReadAll(r.Body)
		if err := json.Unmarshal(body, &user); err != nil {
			fmt.Println(err)
		}
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

		userProfile, err := CreateUser(db, uuid.NewString(), user.Email, string(hashedPassword), user.Name, user.Location)

		if err != nil {
			fmt.Println(err)
		}

		json.NewEncoder(w).Encode(userProfile)

	}
}

// Create user in DB
func CreateUser(db *gorm.DB, id, email, password, name, location string) (*types.User, error) {
	var users = []types.User{}
	// Check if user already exists
	if err := db.Where("email = ?", email).Find(&users).Error; err != nil {
		fmt.Println("Error getting user. Error :", err)
	}
	// Check length of users to avoid duplication
	if len(users) > 0 {
		return nil, nil
	}

	user := &types.User{
		Id:       id,
		Email:    email,
		Password: password,
		Name:     name,
		Location: location,
	}

	if err := db.Create(user).Error; err != nil {
		fmt.Print("error while creating user", err)
	}
	return user, nil
}

// Search friends
func SearchFriends(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		db := db.DbConn()
		vars := mux.Vars(r)
		id := vars["id"]
		var users []types.User

		if err := db.Table("users").Select("users.id, users.name").Where("users.id != ?", id).Find(&users).Error; err != nil {
			fmt.Print("error while searching friend", err)
		}
		json.NewEncoder(w).Encode(users)

	}
}

// Add friends
func AddFriends(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		db := db.DbConn()
		var userFriend types.UserFriend
		var userFriendResponse types.UserFriendResponse
		var responseErr error

		body, _ := ioutil.ReadAll(r.Body)
		if err := json.Unmarshal(body, &userFriend); err != nil {
			responseErr = err
		}

		// Check if userId already exists
		if tx := db.First(&types.User{}, "id = ?", userFriend.UserId); tx.Error != nil {
			responseErr = tx.Error
		}

		// Check if FriendId already exists
		if tx := db.First(&types.User{}, "id = ?", userFriend.FriendId); tx.Error != nil {
			responseErr = tx.Error
		}

		// Create user friend only if user and friend exist
		if responseErr == nil {
			userFriend = types.UserFriend{
				UserId:   userFriend.UserId,
				FriendId: userFriend.FriendId,
			}

			if err := db.Create(userFriend).Error; err != nil {
				fmt.Print("error while creating userFriend", err)
			}

			userFriendResponse = types.UserFriendResponse{
				UserFriend: userFriend,
				ErrorMsg:   "",
			}
		} else {
			userFriendResponse = types.UserFriendResponse{
				ErrorMsg: "either user or friend or both do not exist",
			}
		}
		json.NewEncoder(w).Encode(userFriendResponse)
	}

}

// Post message
func PostMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var post types.Post

		body, _ := ioutil.ReadAll(r.Body)
		if err := json.Unmarshal(body, &post); err != nil {
			fmt.Println(err)
		}

		db := db.DbConn()
		post = types.Post{
			Id:      uuid.NewString(),
			UserId:  post.UserId,
			Message: post.Message,
		}

		if err := db.Create(post).Error; err != nil {
			fmt.Print("error while creating post", err)
		}
		json.NewEncoder(w).Encode(post)
	}
}

// Display posts
func DisplayPosts(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		db := db.DbConn()
		var post []types.Post

		vars := mux.Vars(r)
		userId := vars["userId"]

		// Query to get posts from user and user's friend posts.
		db.Raw("? UNION ?",
			db.Select("posts.user_id, posts.message").Joins("JOIN user_friends ON user_friends.friend_id = posts.user_id AND user_friends.user_id = ?", userId).Find(&post),
			db.Table("posts").Select("posts.user_id, posts.message").Where("posts.user_id = ?", userId).Find(&post),
		).Scan(&post)
		json.NewEncoder(w).Encode(post)
	}
}

// Like posts
func LikePosts(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var like types.Like
		db := db.DbConn()

		body, _ := ioutil.ReadAll(r.Body)
		if err := json.Unmarshal(body, &like); err != nil {
			fmt.Println(err)
		}

		like = types.Like{
			Id:     uuid.NewString(),
			UserId: like.UserId,
			PostId: like.PostId,
		}

		if err := db.Create(like).Error; err != nil {
			fmt.Print("error while liking post", err)
		}
		json.NewEncoder(w).Encode(like)
	}
}

// Share posts
func SharePosts(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var sharePosts types.SharePost
		db := db.DbConn()

		body, _ := ioutil.ReadAll(r.Body)
		if err := json.Unmarshal(body, &sharePosts); err != nil {
			fmt.Println(err)
		}

		sharePosts = types.SharePost{
			Id:     uuid.NewString(),
			UserId: sharePosts.UserId,
			PostId: sharePosts.PostId,
		}

		if err := db.Create(sharePosts).Error; err != nil {
			fmt.Print("error while liking post", err)
		}
		json.NewEncoder(w).Encode(sharePosts)
	}
}
