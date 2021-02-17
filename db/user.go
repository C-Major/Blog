package db

import (
	"errors"

	"github.com/c-major/blog/caller"
	"github.com/c-major/blog/common"
	"github.com/c-major/blog/model"
	"gorm.io/gorm"
)

// UserUpdateFilter .
type UserUpdateFilter struct {
	Email    *string
	Password *string
	Name     *string
	Status   *int8
}

// UserQueryFilter .
type UserQueryFilter struct {
	Email  *string
	Name   *string
	Status *int8
}

// CreateUserByEntity .
func CreateUserByEntity(entity *model.BlogUser) error {
	return caller.DBWrite.Debug().Model(&model.BlogUser{}).Create(&entity).Error
}

// UpdateUserByFilter .
func UpdateUserByFilter(id uint64, filter *UserUpdateFilter) error {
	db := caller.DBWrite.Debug().Model(&model.BlogUser{ID: id})
	fieldsUpdated := make(map[string]interface{})

	if filter.Email != nil {
		fieldsUpdated["email"] = *filter.Email
	}

	if filter.Password != nil {
		fieldsUpdated["password"] = *filter.Password
	}

	if filter.Name != nil {
		fieldsUpdated["name"] = *filter.Name
	}

	if filter.Status != nil {
		fieldsUpdated["status"] = *filter.Status
	}

	return db.Updates(fieldsUpdated).Error
}

// GetUserByID .
func GetUserByID(id uint64) (*model.BlogUser, error) {
	var res model.BlogUser
	db := caller.DBRead.Debug().Model(&model.BlogUser{}).Where("id = ?", id).First(&res)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			common.TextLog.WithField("id", id).Warn("[GetUserByID] record not found")
			return nil, nil
		}

		common.TextLog.WithField("id", id).Error("[GetUserByID] failed to get user")
		return nil, db.Error
	}

	return &res, nil
}

// GetUsersByFilter .
func GetUsersByFilter(filter *UserQueryFilter) ([]*model.BlogUser, error) {
	var res []*model.BlogUser
	db := caller.DBRead.Debug().Model(&model.BlogUser{})

	if filter.Email != nil {
		db = db.Where("email = ?", *filter.Email)
	}

	if filter.Name != nil {
		db = db.Where("name LIKE ?", "%"+*filter.Name+"%")
	}

	if filter.Status != nil {
		db = db.Where("status = ?", *filter.Status)
	}

	db.Find(&res)
	if db.Error != nil {
		common.TextLog.WithField("filter", *filter).Error("[GetUsersByFilter] failed to get users")
		return nil, db.Error
	}

	return res, nil
}

// GetAllUsers .
func GetAllUsers() ([]*model.BlogUser, error) {
	var res []*model.BlogUser
	db := caller.DBRead.Debug().Model(&model.BlogUser{}).Find(&res)
	if db.Error != nil {
		common.TextLog.Error("[GetAllUsers] failed to get all users")
		return nil, db.Error
	}

	return res, db.Error
}
