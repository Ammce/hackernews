package user

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

var userInput = &User{
	Username: "Ammce",
	Password: "12345678",
	Email:    "amcenp@gmail.com",
}

var expectedUser = &User{
	ID:       "1",
	Username: "Ammce",
	Password: "12345678",
	Email:    "amcenp@gmail.com",
}

func (mr *MockRepository) SaveUser(u *User) (*User, error) {
	args := mr.Called(u)
	result := args.Get(0)
	if result != nil {
		return result.(*User), args.Error(1)
	}
	return nil, args.Error(1)
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
	args := mr.Called(userId)
	result := args.Get(0)
	return result.(*User), args.Error(1)
}

func TestCreateUser(t *testing.T) {
	mockRepo := new(MockRepository)

	mockRepo.On("GetUserByField", "email", userInput.Email).Return(nil, nil)
	mockRepo.On("SaveUser", userInput).Return(expectedUser, nil)

	testUserService := NewUserServiceImpl(mockRepo)

	result, _ := testUserService.CreateUser(userInput)

	mockRepo.AssertExpectations(t)

	assert.Equal(t, userInput.Email, result.Email)
	assert.Equal(t, userInput.Username, result.Username)
	assert.NotEqual(t, userInput.Password, result.Password)
}

func TestCreateUserWithSavingError(t *testing.T) {
	mockRepo := new(MockRepository)

	mockRepo.On("GetUserByField", "email", userInput.Email).Return(nil, nil)
	mockRepo.On("SaveUser", userInput).Return(nil, errors.New("invalid query execution"))

	testUserService := NewUserServiceImpl(mockRepo)

	_, err := testUserService.CreateUser(userInput)

	mockRepo.AssertExpectations(t)

	assert.Equal(t, "invalid query execution", err.Error())
}

func TestCreateUserWithExistingEmail(t *testing.T) {
	mockRepo := new(MockRepository)

	mockRepo.On("GetUserByField", "email", userInput.Email).Return(expectedUser, nil)

	testUserService := NewUserServiceImpl(mockRepo)

	createdUser, err := testUserService.CreateUser(userInput)

	mockRepo.AssertExpectations(t)

	fmt.Println(err.Error())
	assert.Error(t, err, "User with email amcenp@gmail.com already exist")

	assert.Nil(t, createdUser, "Email")
}

func TestGetUser(t *testing.T) {
	mockRepo := new(MockRepository)
	mockRepo.On("GetUserById", "1").Return(expectedUser, nil)
	testUserService := NewUserServiceImpl(mockRepo)
	user, _ := testUserService.GetUser("1")
	mockRepo.AssertExpectations(t)
	assert.Equal(t, "1", user.ID)
}

func TestGetUsers(t *testing.T) {
	mockRepo := new(MockRepository)
	mockRepo.On("GetAllUsers").Return([]*User{expectedUser}, nil)
	testUserService := NewUserServiceImpl(mockRepo)
	users, _ := testUserService.GetUsers()
	mockRepo.AssertExpectations(t)
	assert.Equal(t, 1, len(users))
}
