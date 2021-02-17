package db

import (
	"errors"

	"github.com/c-major/blog/caller"
	"github.com/c-major/blog/common"
	"github.com/c-major/blog/model"
	"gorm.io/gorm"
)

// ArticleUpdateFilter .
type ArticleUpdateFilter struct {
	Title   *string
	Content *string
	Status  *int8
}

// ArticleQueryFilter .
type ArticleQueryFilter struct {
	Title  *string
	Status *int8
}

// CreateArticleByEntity .
func CreateArticleByEntity(entity *model.BlogArticle) error {
	return caller.DBWrite.Debug().Model(&model.BlogArticle{}).Create(entity).Error
}

// UpdateArticleByFilter .
func UpdateArticleByFilter(id uint64, filter *ArticleUpdateFilter) error {
	db := caller.DBWrite.Debug().Model(&model.BlogArticle{}).Where("id = ?", id)
	fieldsUpdated := make(map[string]interface{})

	if filter.Title != nil {
		fieldsUpdated["title"] = *filter.Title
	}

	if filter.Content != nil {
		fieldsUpdated["content"] = *filter.Content
	}

	if filter.Status != nil {
		fieldsUpdated["status"] = *filter.Status
	}

	return db.Updates(fieldsUpdated).Error
}

// GetArticleByID .
func GetArticleByID(id uint64) (*model.BlogArticle, error) {
	var res model.BlogArticle
	db := caller.DBRead.Debug().Model(&model.BlogArticle{}).Where("id = ?", id).First(&res)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			common.TextLog.WithField("id", id).Warn("[GetArticleByID] record not found")
			return nil, nil
		}

		common.TextLog.WithField("id", id).Error("[GetArticleByID] failed to get article")
		return nil, db.Error
	}

	return &res, nil
}

// GetArticlesByAuthorID .
func GetArticlesByAuthorID(authorID uint64) ([]*model.BlogArticle, error) {
	var res []*model.BlogArticle
	db := caller.DBRead.Debug().Model(&model.BlogArticle{}).Where("author_id = ?", authorID).Find(&res)
	if db.Error != nil {
		common.TextLog.WithField("authorID", authorID).Error("[GetArticlesByAuthorID] failed to get articles")
		return nil, db.Error
	}

	return res, nil
}

// GetArticlesByFilter .
func GetArticlesByFilter(filter *ArticleQueryFilter) ([]*model.BlogArticle, error) {
	var res []*model.BlogArticle
	db := caller.DBRead.Debug().Model(&model.BlogArticle{})

	if filter.Title != nil {
		db = db.Where("title LIKE ?", "%"+*filter.Title+"%")
	}

	if filter.Status != nil {
		db = db.Where("status = ?", *filter.Status)
	}

	db.Find(&res)
	if db.Error != nil {
		common.TextLog.WithField("filter", *filter).Error("[GetArticlesByFilter] failed to get articles")
		return nil, db.Error
	}

	return res, nil
}

// GetAllArticles .
func GetAllArticles() ([]*model.BlogArticle, error) {
	var res []*model.BlogArticle
	db := caller.DBRead.Debug().Model(&model.BlogArticle{}).Find(&res)
	if db.Error != nil {
		common.TextLog.Error("[GetAllArticles] failed to get all articles")
		return nil, db.Error
	}

	return res, db.Error
}
