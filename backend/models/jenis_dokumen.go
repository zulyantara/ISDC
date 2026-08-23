package models

import (
	"jsdc-api/config"
)

type JenisDokumen struct {
	ID   int    `json:"id"`
	Nama string `json:"nama"`
	Ket  string `json:"ket"`
}

// GetAllJenisDokumen returns all document types
func GetAllJenisDokumen() ([]JenisDokumen, error) {
	query := "SELECT id, nama, ket FROM tb_jenis_dokumen ORDER BY id"
	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []JenisDokumen
	for rows.Next() {
		var j JenisDokumen
		err := rows.Scan(&j.ID, &j.Nama, &j.Ket)
		if err != nil {
			return nil, err
		}
		list = append(list, j)
	}
	return list, nil
}

// GetJenisDokumenByID returns a single document type
func GetJenisDokumenByID(id int) (*JenisDokumen, error) {
	query := "SELECT id, nama, ket FROM tb_jenis_dokumen WHERE id = ?"
	var j JenisDokumen
	err := config.DB.QueryRow(query, id).Scan(&j.ID, &j.Nama, &j.Ket)
	if err != nil {
		return nil, err
	}
	return &j, nil
}

// CreateJenisDokumen creates a new document type
func CreateJenisDokumen(j *JenisDokumen) error {
	query := "INSERT INTO tb_jenis_dokumen (nama, ket) VALUES (?, ?)"
	_, err := config.DB.Exec(query, j.Nama, j.Ket)
	return err
}

// UpdateJenisDokumen updates a document type
func UpdateJenisDokumen(id int, j *JenisDokumen) error {
	query := "UPDATE tb_jenis_dokumen SET nama=?, ket=? WHERE id=?"
	_, err := config.DB.Exec(query, j.Nama, j.Ket, id)
	return err
}

// DeleteJenisDokumen deletes a document type
func DeleteJenisDokumen(id int) error {
	_, err := config.DB.Exec("DELETE FROM tb_jenis_dokumen WHERE id=?", id)
	return err
}
