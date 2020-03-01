package main_test

import (
	"github.com/stretchr/testify/mock"
	"go-restapi/models"
	"testing"
)

type opTest struct {
	input    int
	expected int
}

var opTestCases = []opTest{
	{
		1,
		2,
	},

	{
		10,
		11,
	},
}

func TestSample(t *testing.T) {
	for _, test := range opTestCases {
		if output := addOne(test.input); output != test.expected {
			t.Error("AddOne Failed: Input: ", test.input, " Expected:", test.expected, " Output:", output)
		}
	}
}

func addOne(n int) int {
	return n + 1
}

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

func TestLoadSource(t *testing.T) {
	mockRepo := new(MockRepoSource)
	mockRepo.On("Find").Return(
		models.User{
			Email: "ldavid@curb.com",
			Name:  "Larry David",
			Age:   75,
		},
		nil)
}
