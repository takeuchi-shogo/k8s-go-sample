package models

type Comments struct {
	ID        int    `json:"id"`
	TweetID   int    `json:"tweet_id"`
	CommentID int    `json:"comment_id"`
	UserID    int    `json:"user_id"`
	Text      string `json:"text"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	DeletedAt *int64 `json:"deleted_at"`
}
