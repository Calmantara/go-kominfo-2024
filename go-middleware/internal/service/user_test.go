package service

import (
	"context"
	"errors"
	"testing"

	"github.com/Calmantara/go-kominfo-2024/go-middleware/internal/model"
	"github.com/Calmantara/go-kominfo-2024/go-middleware/internal/repository/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetUsers(t *testing.T) {
	t.Parallel()
	t.Run("error call repo get users", func(t *testing.T) {
		repoMock := mocks.NewUserQuery(t)

		svc := userServiceImpl{
			repo: repoMock,
		}
		repoMock.On("GetUsers", context.Background()).Return([]model.User{}, errors.New("some error"))

		// call method
		usr, err := svc.GetUsers(context.Background())
		assert.NotNil(t, err)
		assert.Equal(t, 0, len(usr))
	})
	t.Run("success call repo get users", func(t *testing.T) {
		repoMock := mocks.NewUserQuery(t)

		svc := userServiceImpl{
			repo: repoMock,
		}
		repoMock.On("GetUsers", context.Background()).Return([]model.User{{ID: 1, Username: "user1"}}, nil)

		// call method
		usr, err := svc.GetUsers(context.Background())
		assert.Nil(t, err)
		assert.Equal(t, 1, len(usr))
	})
}

func TestGetUserById(t *testing.T) {
	type input struct {
		ctx context.Context
		id  uint64
	}

	type output struct {
		user model.User
		err  error
	}

	testCases := []struct {
		desc   string
		in     input
		out    output
		doMock func() *mocks.UserQuery
	}{
		{
			desc: "error get user by id repo",
			in: input{
				ctx: context.Background(),
				id:  100,
			},
			out: output{
				err:  errors.New("some error"),
				user: model.User{},
			},
			doMock: func() *mocks.UserQuery {
				repoMock := mocks.NewUserQuery(t)
				repoMock.On("GetUsersByID", context.Background(), uint64(100)).Return(model.User{}, errors.New("some error"))
				return repoMock
			},
		},
		{
			desc: "success get user by id repo",
			in: input{
				ctx: context.Background(),
				id:  100,
			},
			out: output{
				err:  nil,
				user: model.User{ID: 1, Username: "user1"},
			},
			doMock: func() *mocks.UserQuery {
				repoMock := mocks.NewUserQuery(t)
				repoMock.On("GetUsersByID", context.Background(), uint64(100)).Return(model.User{ID: 1, Username: "user1"}, nil)
				return repoMock
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			repoMock := tC.doMock()
			svc := userServiceImpl{repo: repoMock}
			usr, err := svc.GetUsersById(tC.in.ctx, tC.in.id)
			if tC.out.err != nil {
				assert.EqualError(t, err, tC.out.err.Error())
			} else {
				assert.Nil(t, err)
			}
			assert.Equal(t, tC.out.user, usr)
		})
	}
}
