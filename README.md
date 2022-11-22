
# Friend Book: 
## This is a sample project and not meant for project production use.

## How to run project
* Create a database having name "friend_book" in mysql.

* Create .env file and give appropriate username and password as given in .env_example file as shown below.

```
    MYSQL_USERNAME=username

    MYSQL_PASS=your_password
```
  & run below command 

```bash 
    $ go run .
```

Initially when program starts, hardcore data of users inserted into users tables by using seeding function.

## Tasks and explanations

1.User can create their own profile.

##### Explanation: 
 The endpoint takes user’s details such as name, email, password and location and stores it in the users table in the database.

    End point for own profile: http://localhost:8000/createProfile

2.A user can add other users to his friend list.

##### Explanation: 
 Users can search other users from the users table by using the searchFriends function. This end point receives other user’s userId. And then user can add other users as friends using the AddFriend function in userFriend table in the database.

    End point for add friend: http://localhost:8000/addFriend

3.Users can post messages to their timelines.

##### Explanation: 
Users can post messages using the PostMessage function in the posts table in the database.
message and postId are stored in the posts table.

    End point for post message: http://localhost:8000/postMessage

4.The system should display posts of friends on the display board/timeline.

##### Explanation: 
This endpoint displays posts of friends on the display timeline. It takes the userId of the user and displays posts of the user and user's friends.

    End point for display post: http://localhost:8000/displayPost/4a1f02ce-4e6c-4ed1-a56b-f5f018aa5a68

5.People can like a post.

##### Explanation: 
Users can like posts using the LikePost function in the like table in the database.
userId and postId are stored in the like table.

    End point for like post:  http://localhost:8000/likePost

6.People can share their friend's posts on their own display board/timeline.

##### Explanation: 
This endpoint shares posts of friends on the display timeline. It takes the userId and postId of the user and stores it in the SharePost table.

    End point for share post: http://localhost:8000/sharePost















