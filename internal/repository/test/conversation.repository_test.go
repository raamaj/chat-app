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

func TestFindConversationByIdSuccess(t *testing.T) {
	db := SetubDb()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock_repo := mock_repository.NewMockConversationRepository(ctrl)

	result := &entity.Conversation{}

	expectedResult := &entity.Conversation{
		ID:        1,
		CreatedAt: time.Now(),
	}

	mock_repo.EXPECT().FindById(db, result, int64(1)).DoAndReturn(func(_ *gorm.DB, conversation *entity.Conversation, id int64) error {
		*conversation = *expectedResult
		return nil
	})

	err := mock_repo.FindById(db, result, int64(1))

	assert.Nil(t, err)
	assert.Equal(t, expectedResult.ID, result.ID)
}

func TestCreateConversationSuccess(t *testing.T) {
	db := SetubDb()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	now := time.Now()

	mock_repo := mock_repository.NewMockConversationRepository(ctrl)
	result := &entity.Conversation{
		CreatedAt: now,
	}

	expectedResult := &entity.Conversation{
		ID:        1,
		CreatedAt: now,
	}

	mock_repo.EXPECT().Create(db, result).DoAndReturn(func(_ *gorm.DB, conversation *entity.Conversation) error {
		*conversation = *expectedResult
		return nil
	})

	err := mock_repo.Create(db, result)
	assert.Nil(t, err)
	assert.Equal(t, expectedResult.ID, result.ID)
}
