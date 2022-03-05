package handlers

import (
	"fmt"
	"log"

	"github.com/xuanvan229/blog-core/datatransfers"
	"github.com/xuanvan229/blog-core/models"
)

func (m *module) CreatePost(id uint, credentials datatransfers.PostCreate) (err error) {
	var user models.User
	fmt.Println("user id", id)
	if user, err = m.db.userOrmer.GetOneByID(id); err != nil {
		return fmt.Errorf("cannot find user with username %s", id)
	}

	if _, err = m.db.postOrmer.InsertPost(models.Post{
		Title:       credentials.Title,
		Description: credentials.Description,
		Body:        credentials.Body,
		User:        user,
	}); err != nil {
		log.Print(err)
		return fmt.Errorf("error inserting user. %v", err)
	}
	return
}

func (m *module) GetAllPost(id uint) (list_post []datatransfers.PostInfor, err error) {
	var posts []models.Post
	if posts, err = m.db.postOrmer.GetAll(id); err != nil {
		log.Print(err)
		return []datatransfers.PostInfor{}, fmt.Errorf("error inserting user. %v", err)
	}

	for _, post := range posts {
		list_post = append(list_post, datatransfers.PostInfor{
			Title:       post.Title,
			Description: post.Description,
			Body:        post.Body,
			CreatedAt:   post.CreatedAt,
		})
	}
	return list_post, err
}
