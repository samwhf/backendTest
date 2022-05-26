package user

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	pg "github.com/samwhf/backendTest/database/postgres"
	"github.com/samwhf/backendTest/objects"
)

// post
func Create(ctx context.Context, user *objects.User) (string, error) {
	user.ID = uuid.New().String()
	user.CreatedAt = time.Now()
	if err := pg.GetClient().DB.Create(&user).Error; err != nil {
		return "", err
	}
	return user.ID, nil
}

// get by id
func Get(ctx context.Context, id string) (*objects.User, error) {
	var user objects.User
	return &user, pg.GetClient().DB.Set("gorm:auto_preload", true).Where("id = ?", id).Find(&user).Error
}

// put
func Update(ctx context.Context, user *objects.User) error {
	if user.ID == "" {
		return errors.New("user id is invalid")
	}
	return pg.GetClient().DB.Save(&user).Error
}

// delete
func Delete(ctx context.Context, id string) error {
	var user objects.User
	return pg.GetClient().DB.Where("id = ?", id).Delete(&user).Error
}

// get by name
func FindUserByName(ctx context.Context, name string) (*objects.User, error) {
	p := &objects.User{}
	return p, pg.GetClient().DB.Set("gorm:auto_preload", true).Find(p, "name = ?", name).Error
}

func CreateTable(ctx context.Context) error {
	return pg.GetClient().DB.AutoMigrate(&objects.User{}).Error
}
