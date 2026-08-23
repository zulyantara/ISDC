package models

import (
	"isdc-api/config"
)

type Soal struct {
	UjianID  int    `json:"ujian_id"`
	Sesi     int    `json:"sesi"`
	Nomor    int    `json:"nomor"`
	Category int    `json:"category"`
	Soal     string `json:"soal"`
}

// GetAllSoal returns all exam questions
func GetAllSoal() ([]Soal, error) {
	query := "SELECT ujian_id, sesi, nomor, category, soal FROM tb_soal ORDER BY category, sesi, nomor"
	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []Soal
	for rows.Next() {
		var s Soal
		err := rows.Scan(&s.UjianID, &s.Sesi, &s.Nomor, &s.Category, &s.Soal)
		if err != nil {
			return nil, err
		}
		list = append(list, s)
	}
	return list, nil
}

// GetSoalByCategory returns exam questions filtered by category (teori_id)
func GetSoalByCategory(category int) ([]Soal, error) {
	query := "SELECT ujian_id, sesi, nomor, category, soal FROM tb_soal WHERE category=? ORDER BY sesi, nomor"
	rows, err := config.DB.Query(query, category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []Soal
	for rows.Next() {
		var s Soal
		err := rows.Scan(&s.UjianID, &s.Sesi, &s.Nomor, &s.Category, &s.Soal)
		if err != nil {
			return nil, err
		}
		list = append(list, s)
	}
	return list, nil
}

// GetSoalByID returns a single question
func GetSoalByID(ujianID int) (*Soal, error) {
	query := "SELECT ujian_id, sesi, nomor, category, soal FROM tb_soal WHERE ujian_id = ?"
	var s Soal
	err := config.DB.QueryRow(query, ujianID).Scan(&s.UjianID, &s.Sesi, &s.Nomor, &s.Category, &s.Soal)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

// CreateSoal creates a new question
func CreateSoal(s *Soal) error {
	query := "INSERT INTO tb_soal (sesi, nomor, category, soal) VALUES (?, ?, ?, ?)"
	_, err := config.DB.Exec(query, s.Sesi, s.Nomor, s.Category, s.Soal)
	return err
}

// UpdateSoal updates a question
func UpdateSoal(ujianID int, s *Soal) error {
	query := "UPDATE tb_soal SET sesi=?, nomor=?, category=?, soal=? WHERE ujian_id=?"
	_, err := config.DB.Exec(query, s.Sesi, s.Nomor, s.Category, s.Soal, ujianID)
	return err
}

// DeleteSoal deletes a question
func DeleteSoal(ujianID int) error {
	_, err := config.DB.Exec("DELETE FROM tb_soal WHERE ujian_id=?", ujianID)
	return err
}

// GetSoalForPeserta returns questions for a specific participant based on their class category
func GetSoalForPeserta(pesertaID string) ([]Soal, error) {
	query := `SELECT s.ujian_id, s.sesi, s.nomor, s.category, s.soal 
		FROM tb_soal s 
		JOIN tb_peserta p ON s.category = (SELECT teori_id FROM mt_kelas WHERE kelas_id=p.kelas_id AND area_id=p.area_id)
		WHERE p.peserta_id = ?
		ORDER BY s.sesi, s.nomor`
	rows, err := config.DB.Query(query, pesertaID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []Soal
	for rows.Next() {
		var s Soal
		err := rows.Scan(&s.UjianID, &s.Sesi, &s.Nomor, &s.Category, &s.Soal)
		if err != nil {
			return nil, err
		}
		list = append(list, s)
	}
	return list, nil
}
