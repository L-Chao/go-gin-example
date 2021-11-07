package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Article struct {
	Model

	TagID int `json:"tag_id" gorm:"index"`
	Tag   Tag `json:"tag"`

	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CreatedBy     string `json:"created_by"`
	ModifiedBy    string `json:"modified_by"`
	State         int    `json:"state"`
	CoverImageUrl string `json:"cover_image_url"`
}

func ExistArticleByID(id int) (bool, error) {
	var article Article
	err := db.Select("id").Where("id = ?", id).First(&article)
	if err != nil {
		return false, nil
	}
	return article.ID > 0, nil
}

func GetArticleTotal(maps interface{}) (int, error) {
	var count int
	err := db.Model(&Article{}).Where(maps).Count(&count).Error
	return count, err
}

func GetArticles(pageNum int, pageSize int, maps interface{}) ([]Article, error) {
	var articles []Article
	err := db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles).Error
	return articles, err
}

func GetArticle(id int) (*Article, error) {
	var article Article
	err := db.Where("id = ?", id).First(&article).Error
	if err != nil {
		return nil, err
	}
	err = db.Model(&Tag{}).Related(&article.Tag).Error
	return &article, err
}

func EditArticle(id int, data interface{}) error {
	return db.Model(&Article{}).Where("id = ?", id).Update(data).Error
}

func AddArticle(data map[string]interface{}) error {
	err := db.Create(&Article{
		TagID:         data["tag_id"].(int),
		Title:         data["title"].(string),
		Desc:          data["desc"].(string),
		Content:       data["content"].(string),
		CreatedBy:     data["created_by"].(string),
		State:         data["state"].(int),
		CoverImageUrl: data["cover_image_url"].(string),
	}).Error
	return err
}

func DeleteArticle(id int) error {
	return db.Where("id = ?", id).Delete(&Article{}).Error
}

func (a *Article) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("CreatedOn", time.Now().Unix())
}

func (a *Article) BeforeUpdate(scop *gorm.Scope) error {
	return scop.SetColumn("ModifiedOn", time.Now().Unix())
}

func DeleteAllArticle() bool {
	db.Unscoped().Where("deleted_on != ", 0).Delete(&Article{})
	return true
}
