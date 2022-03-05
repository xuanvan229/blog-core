package models

import (
	"time"

	"gorm.io/gorm"
)

type postOrm struct {
	db *gorm.DB
}

type Post struct {
	ID          uint      `gorm:"primaryKey" json:"-"`
	Title       string    `json:"-"`
	Description string    `json:"-"`
	Body        string    `json:"-"`
	UserID      uint      `json:"-"`
	User        User      `gorm:"foreignKey:UserID"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"-"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"-"`
}

type PostOrmer interface {
	GetOneByID(id uint) (post Post, err error)
	GetAll(id uint) (posts []Post, err error)
	InsertPost(post Post) (id uint, err error)
	UpdatePost(post Post) (err error)
}

func NewPostOrmer(db *gorm.DB) PostOrmer {
	_ = db.AutoMigrate(&Post{}) // builds table when enabled
	return &postOrm{db}
}

func (o *postOrm) GetOneByID(id uint) (post Post, err error) {
	result := o.db.Model(&Post{}).Where("id = ?", id).First(&post)
	return post, result.Error
}

func (o *postOrm) GetAll(id uint) (posts []Post, err error) {
	result := o.db.Model(&Post{}).Where("user_id = ?", id).Find(&posts)
	return posts, result.Error
}

func (o *postOrm) InsertPost(post Post) (id uint, err error) {
	result := o.db.Model(&Post{}).Create(&post)
	return post.ID, result.Error
}

func (o *postOrm) UpdatePost(post Post) (err error) {
	// By default, only non-empty fields are updated. See https://gorm.io/docs/update.html#Updates-multiple-columns
	result := o.db.Model(&Post{}).Model(&post).Updates(&post)
	return result.Error
}
