package seed

import (
	"github.com/google/uuid"
	"github.com/shyams2012/friend_book/db"
	"github.com/shyams2012/friend_book/types"
	"github.com/shyams2012/friend_book/user"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Migrate to DB
func Migrate() error {
	db := db.DbConn()
	return db.AutoMigrate(
		&types.User{},
		&types.UserFriend{},
		&types.Post{},
		&types.Like{},
		&types.SharePost{},
	)
}

// Seeding of users
func SeedUsers() (err error) {
	db := db.DbConn()

	for _, seed := range All() {
		seed.Run(db)
	}
	return
}

func All() []types.Seed {
	// Hash passwords
	hashedPassword1, _ := bcrypt.GenerateFromPassword([]byte("password1"), bcrypt.DefaultCost)
	hashedPassword2, _ := bcrypt.GenerateFromPassword([]byte("password2"), bcrypt.DefaultCost)
	hashedPassword3, _ := bcrypt.GenerateFromPassword([]byte("password3"), bcrypt.DefaultCost)
	hashedPassword4, _ := bcrypt.GenerateFromPassword([]byte("password4"), bcrypt.DefaultCost)
	hashedPassword5, _ := bcrypt.GenerateFromPassword([]byte("password5"), bcrypt.DefaultCost)

	// Set and return users for seeding
	return []types.Seed{
		types.Seed{
			Name: "CreateShyam",
			Run: func(db *gorm.DB) (*types.User, error) {
				return user.CreateUser(db, uuid.NewString(), "shyams2012@gmail.com", string(hashedPassword1), "Shyam", "Ktm")
			},
		},
		types.Seed{
			Name: "CreateAjay",
			Run: func(db *gorm.DB) (*types.User, error) {
				return user.CreateUser(db, uuid.NewString(), "ajay@gmail.com", string(hashedPassword2), "Ajay", "Ktm")
			},
		},
		types.Seed{
			Name: "CreateRam",
			Run: func(db *gorm.DB) (*types.User, error) {
				return user.CreateUser(db, uuid.New().String(), "ram@gmail.com", string(hashedPassword3), "Ram", "Ktm")
			},
		},
		types.Seed{
			Name: "CreateHari",
			Run: func(db *gorm.DB) (*types.User, error) {
				return user.CreateUser(db, uuid.New().String(), "hari@gmail.com", string(hashedPassword4), "Hari", "Ktm")
			},
		},
		types.Seed{
			Name: "CreateAakash",
			Run: func(db *gorm.DB) (*types.User, error) {
				return user.CreateUser(db, uuid.New().String(), "aakash@gmail.com", string(hashedPassword5), "Aakash", "Ktm")
			},
		},
	}
}
