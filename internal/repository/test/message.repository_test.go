package test

import (
	"github.com/raamaj/chat-app/internal/entity"
	mock_repository "github.com/raamaj/chat-app/internal/repository/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestCreateMessageSuccess(t *testing.T) {
	db := SetubDb()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	now := time.Now()

	result := &entity.Message{
		Conversation: entity.Conversation{
			ID:        1,
			CreatedAt: now,
		},
		Sender: entity.User{
			ID:        1,
			Username:  "johndoe",
			Email:     "johndoe@example.com",
			Password:  "$2a$12$e8WdmKP/NgBvXRak0xoDJeZH5f2OG2RPX7cIev62xvjPj7PTvCHlS",
			CreatedAt: now,
		},
		Content: "Hello",
		SendAt:  now,
	}

	expected := &entity.Message{
		ID: 1,
		Conversation: entity.Conversation{
			ID:        1,
			CreatedAt: now,
		},
		Sender: entity.User{
			ID:        1,
			Username:  "johndoe",
			Email:     "johndoe@example.com",
			Password:  "$2a$12$e8WdmKP/NgBvXRak0xoDJeZH5f2OG2RPX7cIev62xvjPj7PTvCHlS",
			CreatedAt: now,
		},
		Content: "Hello",
		SendAt:  now,
	}

	mock_repo := mock_repository.NewMockMessageRepository(ctrl)

	mock_repo.EXPECT().Create(db, result).DoAndReturn(func(_ *gorm.DB, message *entity.Message) error {
		*message = *expected
		return nil
	})

	err := mock_repo.Create(db, result)
	assert.Nil(t, err)
	assert.Equal(t, expected.ID, result.ID)
}

func TestListMessageSuccess(t *testing.T) {
	db := SetubDb()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	now := time.Now()

	conversation := entity.Conversation{
		ID:        1,
		CreatedAt: now,
	}

	sender1 := entity.User{
		ID:        1,
		Username:  "johndoe",
		Email:     "johndoe@example.com",
		Password:  "$2a$12$e8WdmKP/NgBvXRak0xoDJeZH5f2OG2RPX7cIev62xvjPj7PTvCHlS",
		CreatedAt: now,
	}

	sender2 := entity.User{
		ID:        1,
		Username:  "marryjoe",
		Email:     "marryjoe@example.com",
		Password:  "$2a$12$e8WdmKP/NgBvXRak0xoDJeZH5f2OG2RPX7cIev62xvjPj7PTvCHlS",
		CreatedAt: now,
	}

	expectedResult := &[]entity.Message{
		{ID: int64(1), Conversation: conversation, Sender: sender1, Content: "Hello, How Are You?", SendAt: now},
		{ID: int64(2), Conversation: conversation, Sender: sender2, Content: "Hello, I'm fine", SendAt: now},
	}

	mock_repo := mock_repository.NewMockMessageRepository(ctrl)
	mock_repo.EXPECT().List(db, int64(1)).Return(expectedResult, nil)

	result, err := mock_repo.List(db, int64(1))
	assert.Nil(t, err)
	assert.Len(t, *result, 2)
	assert.Equal(t, len(*expectedResult), len(*result))
}
