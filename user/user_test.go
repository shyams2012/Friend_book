package user_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/shyams2012/friend_book/types"
	"github.com/shyams2012/friend_book/user"
)

func TestSearchFriend(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "/searchFriend?", nil)
	w := httptest.NewRecorder()
	user.SearchFriends(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		t.Errorf("error found Error: %v", err)
	}
	if len(data) < 1 {
		t.Fatal("Expected friends but got nothing")
	}
}

func TestCreateProfile(t *testing.T) {

	var profile = types.User{
		Name:     "Rajesh",
		Email:    "rajus@gmail.com",
		Location: "Ktm",
	}
	profileData, err := json.Marshal(profile)
	if err != nil {
		fmt.Print(err)
	}

	reader := strings.NewReader(string(profileData))

	req := httptest.NewRequest(http.MethodPost, "/createProfile?", reader)
	w := httptest.NewRecorder()
	user.CreateProfile(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)

	user := types.User{}

	if err := json.Unmarshal(data, &user); err != nil {
		fmt.Println(err)
	}

	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if user.Name != profile.Name && user.Email != profile.Email && user.Location != profile.Location {
		t.Fatal("Expected profile but got nothing")
	}
}

func TestAddFriend(t *testing.T) {
	var userFriend = types.UserFriend{
		UserId:   "4547fd4b-f28c-4868-8561-7409dede1c83",
		FriendId: "7ff2412d-0453-420a-9520-20eabc096006",
	}

	userFriendData, err := json.Marshal(userFriend)
	if err != nil {
		fmt.Print(err)
	}
	reader := strings.NewReader(string(userFriendData))

	req := httptest.NewRequest(http.MethodPost, "/addFriend?", reader)
	w := httptest.NewRecorder()
	user.AddFriends(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)

	user := types.UserFriend{}

	if err := json.Unmarshal(data, &user); err != nil {
		fmt.Println(err)
	}

	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if userFriend.UserId != user.UserId && userFriend.FriendId != user.FriendId {
		t.Fatal("Expected data but got nothing")
	}
}

func TestPostMessage(t *testing.T) {

	var post = types.Post{
		UserId:  "31718658-8849-42d3-8819-50ab42f1c9af",
		Message: "Ktm",
	}
	postData, err := json.Marshal(post)
	if err != nil {
		fmt.Print(err)
	}
	reader := strings.NewReader(string(postData))

	req := httptest.NewRequest(http.MethodPost, "/postMessage?", reader)
	w := httptest.NewRecorder()
	user.PostMessage(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)

	postResponse := types.Post{}

	if err := json.Unmarshal(data, &postResponse); err != nil {
		fmt.Println(err)
	}

	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if post.UserId != postResponse.UserId && post.Message != postResponse.Message {
		t.Fatal("Expected post but got nothing")
	}
}

func TestLikePosts(t *testing.T) {
	var like = types.Like{
		UserId: "31718658-8849-42d3-8819-50ab42f1c9af",
		PostId: "Ktm",
	}
	likeData, err := json.Marshal(like)
	if err != nil {
		fmt.Print(err)
	}

	reader := strings.NewReader(string(likeData))

	req := httptest.NewRequest(http.MethodPost, "/likePost?", reader)
	w := httptest.NewRecorder()
	user.LikePosts(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)

	likeResponse := types.Like{}

	if err := json.Unmarshal(data, &likeResponse); err != nil {
		fmt.Println(err)
	}

	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if like.PostId != likeResponse.PostId && like.UserId != likeResponse.UserId {
		t.Fatal("Expected like but got nothing")
	}
}

func TestSharePosts(t *testing.T) {
	var sharePost = types.SharePost{
		UserId: "31718658-8849-42d3-8819-50ab42f1c9af",
		PostId: "Ktm",
	}
	sharePostData, err := json.Marshal(sharePost)
	if err != nil {
		fmt.Print(err)
	}

	reader := strings.NewReader(string(sharePostData))

	req := httptest.NewRequest(http.MethodPost, "/sharePosts?", reader)
	w := httptest.NewRecorder()
	user.SharePosts(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)

	sharePostResponse := types.SharePost{}

	if err := json.Unmarshal(data, &sharePostResponse); err != nil {
		fmt.Println(err)
	}

	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if sharePost.PostId != sharePostResponse.PostId && sharePost.UserId != sharePostResponse.UserId {
		t.Fatal("Expected share post but got nothing")
	}
}

func TestDisplayPosts(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/displayPost/{userId}", nil)
	w := httptest.NewRecorder()
	user.DisplayPosts(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		t.Errorf("error found Error: %v", err)
	}
	if len(data) < 1 {
		t.Fatal("Expected posts but got nothing")
	}
}
