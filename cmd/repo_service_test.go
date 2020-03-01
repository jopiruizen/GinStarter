package main_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go-restapi/models"
	"go-restapi/services"
	"testing"
)

/*
 * With Mock Repo Source
 */

type MockRepoSource struct {
	mock.Mock
	usersList []models.User
}

/*
 * on test it should populate repoSource userList
 */

/* REPOSITORY TESTS */

func (mock *MockRepoSource) LoadSource() {
	args := mock.Called()
	result := args.Get(0)
	mock.usersList = result.([]models.User)
}

/*
 * on test it should find a User with email and equal to supplied email and non empty name and age
 */
func (mock *MockRepoSource) Find(email string) (models.User, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(models.User), args.Error(1)
}

func TestFind(t *testing.T) {
	mockRepo := new(MockRepoSource)

	//set up expecations
	mockRepo.On("Find").Return(
		models.User{
			Email: "ldavid@curb.com",
			Name:  "Larry David",
			Age:   75,
		},
		nil)

	testService := services.NewService(mockRepo)
	testResult, _ := testService.Find("ldavid@curb.com")
	mockRepo.AssertExpectations(t)
	assert.Equal(t, "Larry David", testResult.Name)
	assert.Equal(t, "ldavid@curb.com", testResult.Email)
	assert.Equal(t, 75, testResult.Age)
}

func TestFindErrorNoRecordFound(t *testing.T) {
	mockRepo := new(MockRepoSource)

	//set up expecations
	mockRepo.On("Find").Return(
		models.User{
			Email: "ldavid@curb.com",
			Name:  "Larry David",
			Age:   75,
		},
		models.ErrNoRecordFound)

	testService := services.NewService(mockRepo)
	_, err := testService.Find("ldavid@curb.com")
	mockRepo.AssertExpectations(t)
	assert.Equal(t, models.ErrNoRecordFound.Error(), err.Error())
}
