package repository

import (
	"context"

	"github.com/Calmantara/go-kominfo-2024/project-structure/internal/infrastructure"
	"github.com/Calmantara/go-kominfo-2024/project-structure/internal/model"
)

type UserQuery interface {
	GetUsers(ctx context.Context) ([]model.User, error)
	GetUsersByID(ctx context.Context, id uint64) (model.User, error)
}

type UserCommand interface {
	CreateUser(ctx context.Context, user model.User) (model.User, error)
}

type userQueryImpl struct {
	db infrastructure.GormPostgres
}

type userQueryMongoImpl struct{}

func NewUserQueryMongo() UserQuery {
	return userQueryMongoImpl{}
}

func (u userQueryMongoImpl) GetUsers(ctx context.Context) ([]model.User, error) {
	return nil, nil
}
func (u userQueryMongoImpl) GetUsersByID(ctx context.Context, id uint64) (model.User, error) {
	return model.User{}, nil
}

func NewUserQuery(db infrastructure.GormPostgres) UserQuery {
	return &userQueryImpl{db: db}
}

func (u *userQueryImpl) GetUsers(ctx context.Context) ([]model.User, error) {
	db := u.db.GetConnection()
	users := []model.User{}
	if err := db.WithContext(ctx).Table("users").Find(&users).Error; err != nil {
		return nil, nil
	}
	return users, nil
}

func (u *userQueryImpl) GetUsersByID(ctx context.Context, id uint64) (model.User, error) {
	db := u.db.GetConnection()
	users := model.User{}
	if err := db.WithContext(ctx).Table("users").Where("id = ?", id).Find(&users).Error; err != nil {
		return model.User{}, nil
	}
	return users, nil
}
