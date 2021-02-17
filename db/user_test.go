package db

import (
	"os"
	"testing"

	"github.com/c-major/blog/caller"
	"github.com/c-major/blog/common"
	"github.com/c-major/blog/model"
	"github.com/stretchr/testify/assert"
)

func init() {
	common.InitLog()

	rootDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	os.Setenv("IS_TEST_ENV", "1")
	config, err := common.GetConfig(rootDir, "../", "conf")
	if err != nil {
		common.TextLog.Error("failed to initialize config")
	}

	err = caller.InitCaller(config)
	if err != nil {
		common.TextLog.Error("failed to initialize caller")
	}
}

func TestCreateUserByEntity(t *testing.T) {
	User := model.BlogUser{
		Email:    "test2@test.com",
		Password: "test2",
		Name:     "test2",
		Status:   0,
	}

	err := CreateUserByEntity(&User)
	assert.Nil(t, err)
}

func TestUpdateUserByFilter(t *testing.T) {
	var (
		id      uint64 = 1
		newName string = "new name"
	)

	filter := UserUpdateFilter{
		Name: &newName,
	}
	err := UpdateUserByFilter(id, &filter)
	assert.Nil(t, err)
}

func TestGetUserByID(t *testing.T) {
	var id uint64 = 1

	User, err := GetUserByID(id)
	assert.Nil(t, err)
	t.Logf("user: %v", User)
}

func TestGetUserByFilter(t *testing.T) {
	var (
		// email  string = "test@test.com"
		name   string = "test"
		status int8   = 0
	)

	filter := UserQueryFilter{
		Email:  nil,
		Name:   &name,
		Status: &status,
	}

	UserList, err := GetUsersByFilter(&filter)
	assert.Nil(t, err)
	for _, User := range UserList {
		t.Logf("user: %v", User)
	}
}

func TestGetAllUsers(t *testing.T) {
	userList, err := GetAllUsers()
	assert.Nil(t, err)
	for _, user := range userList {
		t.Logf("user: %v", user)
	}
}
