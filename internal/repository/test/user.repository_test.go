package test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/raamaj/chat-app/internal/entity"
	mock_repository "github.com/raamaj/chat-app/internal/repository/mock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
	"time"
)

func SetubDb() *gorm.DB {
	mockDb, _, err := sqlmock.New()
	if err != nil {
		logrus.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDb.Close()

	dialector := mysql.New(mysql.Config{
		Conn:                      mockDb,
		SkipInitializeWithVersion: true,
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})

	return db
}

func TestFindUserByUsernameSuccess(t *testing.T) {
	db := SetubDb()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockUserRepository(ctrl)

	mockRepo.EXPECT().CountByUsername(db, "johndoe").Return(int64(1), nil)

	totalUser, err := mockRepo.CountByUsername(db, "johndoe")

	assert.Nil(t, err)
	assert.Equal(t, totalUser, int64(1))
}

func TestFinUserByIdSuccess(t *testing.T) {
	db := SetubDb()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockUserRepository(ctrl)

	userModel := &entity.User{}

	expectedUser := &entity.User{
		ID:        1,
		Username:  "johndoe",
		Email:     "johndoe@mail.com",
		Password:  "$2a$12$e8WdmKP/NgBvXRak0xoDJeZH5f2OG2RPX7cIev62xvjPj7PTvCHlS",
		CreatedAt: time.Now(),
	}

	mockRepo.EXPECT().FindById(db, userModel, int64(1)).DoAndReturn(func(_ *gorm.DB, user *entity.User, id int64) error {
		*user = *expectedUser // Manually set the expected user data
		return nil
	})

	err := mockRepo.FindById(db, userModel, int64(1))
	assert.Nil(t, err)
	assert.Equal(t, expectedUser.ID, userModel.ID)
}

func TestCreateUserSuccess(t *testing.T) {
	db := SetubDb()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockUserRepository(ctrl)

	now := time.Now()

	userModel := &entity.User{
		Username:  "johndoe",
		Email:     "johndoe@mail.com",
		Password:  "$2a$12$e8WdmKP/NgBvXRak0xoDJeZH5f2OG2RPX7cIev62xvjPj7PTvCHlS",
		CreatedAt: now,
	}

	expectedUser := &entity.User{
		ID:        1,
		Username:  "johndoe",
		Email:     "johndoe@mail.com",
		Password:  "$2a$12$e8WdmKP/NgBvXRak0xoDJeZH5f2OG2RPX7cIev62xvjPj7PTvCHlS",
		CreatedAt: now,
	}

	mockRepo.EXPECT().Create(db, userModel).DoAndReturn(func(_ *gorm.DB, user *entity.User) error {
		user.ID = expectedUser.ID
		return nil
	})
	err := mockRepo.Create(db, userModel)
	assert.Nil(t, err)
	assert.Equal(t, expectedUser.ID, userModel.ID)
}
