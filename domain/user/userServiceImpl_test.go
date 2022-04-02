package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (mr *MockRepository) SaveUser(u *User) (*User, error) {
	args := mr.Called(u)
	result := args.Get(0)
	return result.(*User), args.Error(1)
}

func (mr *MockRepository) GetAllUsers() ([]*User, error) {
	args := mr.Called()
	result := args.Get(0)
	return result.([]*User), args.Error(1)
}

func (mr *MockRepository) GetUserByField(field string, fieldValue string) (*User, error) {
	args := mr.Called(field, fieldValue)
	result := args.Get(0)
	if result != nil {
		return result.(*User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (mr *MockRepository) GetUserById(userId string) (*User, error) {
	args := mr.Called()
	result := args.Get(0)
	return result.(*User), args.Error(1)
}

func TestCreateUserWithExistingEmail(t *testing.T) {
	mockRepo := new(MockRepository)

	userInput := &User{
		Username: "Ammce",
		Password: "12345678",
		Email:    "amcenp@gmail.com",
	}

	expectedUser := &User{
		ID:       "1",
		Username: "Ammce",
		Password: "12345678",
		Email:    "amcenp@gmail.com",
	}

	mockRepo.On("GetUserByField", "email", userInput.Email).Return(nil, nil)
	mockRepo.On("SaveUser", userInput).Return(expectedUser, nil)

	testUserService := NewUserServiceImpl(mockRepo)

	result, _ := testUserService.CreateUser(userInput)

	mockRepo.AssertExpectations(t)

	assert.Equal(t, userInput.Email, result.Email)
	assert.Equal(t, userInput.Username, result.Username)
	assert.NotEqual(t, userInput.Password, result.Password)

}
