package datatransfers

import ("time")

type PostCreate struct {
	Title string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Body string `json:"body" binding:"required"`
}

type PostInfor struct {
	Title string `json:"title"`
	Description string `json:"description"`
	Body string `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}