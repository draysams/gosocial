package store

import (
	"context"
	"database/sql"
)

type Comment struct {
	ID        int64  `json:"id"`
	Content   string `json:"content"`
	PostID    int64  `json:"post_id"`
	UserID    int64  `json:"user_id"`
	CreatedAt string `json:"created_at"`
	User      User   `json:"user"`
}

type CommentStore struct {
	db *sql.DB
}

func (s *CommentStore) Create(ctx context.Context, comment *Comment) error {
	query := `
		INSERT INTO comments (content, post_id, user_id)
		VALUES ($1, $2, $3)
		RETURNING id, created_at
	`

	ctx2, cancel := context.WithTimeout(context.Background(), QueryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(ctx2, query, comment.Content, comment.PostID, comment.UserID).Scan(&comment.ID, &comment.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (s *CommentStore) GetCommentsByPostID(ctx context.Context, postID int64) ([]Comment, error) {
	query := `
		SELECT c.id, c.content, c.post_id, c.user_id ,users.username, users.id, c.created_at
		FROM comments c
		JOIN users ON c.user_id = users.id
		WHERE c.post_id = $1
		ORDER BY created_at DESC
	`
	rows, err := s.db.QueryContext(ctx, query, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	comments := []Comment{}
	for rows.Next() {
		var comment Comment
		err := rows.Scan(&comment.ID, &comment.Content, &comment.PostID, &comment.UserID, &comment.User.Username, &comment.User.ID, &comment.CreatedAt)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
}
