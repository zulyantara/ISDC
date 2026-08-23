package models

import (
	"isdc-api/config"
)

type Comment struct {
	CommentID   int    `json:"comment_id"`
	PesertaID   string `json:"peserta_id"`
	Pengetahuan string `json:"pengetahuan"`
	Teknik      string `json:"teknik"`
	Perilaku    string `json:"perilaku"`
}

// GetCommentByPeserta returns comments for a participant
func GetCommentByPeserta(pesertaID string) (*Comment, error) {
	query := "SELECT comment_id, peserta_id, pengetahuan, teknik, perilaku FROM tb_comments WHERE peserta_id=?"
	var c Comment
	err := config.DB.QueryRow(query, pesertaID).Scan(&c.CommentID, &c.PesertaID, &c.Pengetahuan, &c.Teknik, &c.Perilaku)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

// CreateComment creates a new comment
func CreateComment(c *Comment) error {
	query := "INSERT INTO tb_comments (peserta_id, pengetahuan, teknik, perilaku) VALUES (?, ?, ?, ?)"
	_, err := config.DB.Exec(query, c.PesertaID, c.Pengetahuan, c.Teknik, c.Perilaku)
	return err
}

// UpdateComment updates a comment
func UpdateComment(commentID int, c *Comment) error {
	query := "UPDATE tb_comments SET pengetahuan=?, teknik=?, perilaku=? WHERE comment_id=?"
	_, err := config.DB.Exec(query, c.Pengetahuan, c.Teknik, c.Perilaku, commentID)
	return err
}

// UpsertComment creates or updates a comment
func UpsertComment(c *Comment) error {
	// Check if comment exists for this peserta
	var count int64
	err := config.DB.QueryRow("SELECT COUNT(*) FROM tb_comments WHERE peserta_id=?", c.PesertaID).Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		// Update existing
		query := "UPDATE tb_comments SET pengetahuan=?, teknik=?, perilaku=? WHERE peserta_id=?"
		_, err = config.DB.Exec(query, c.Pengetahuan, c.Teknik, c.Perilaku, c.PesertaID)
		return err
	}

	// Create new
	return CreateComment(c)
}

// DeleteComment deletes a comment
func DeleteComment(commentID int) error {
	_, err := config.DB.Exec("DELETE FROM tb_comments WHERE comment_id=?", commentID)
	return err
}
