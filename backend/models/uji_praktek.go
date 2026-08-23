package models

import (
	"isdc-api/config"
)

type UjiPraktek struct {
	PraktekID  int    `json:"praktek_id"`
	SoalID     int    `json:"soal_id"`
	PesertaID  string `json:"peserta_id"`
	Hasil      int    `json:"hasil"`
	Tanggal    string `json:"tanggal"`
	Modified   string `json:"modified"`
	Platform   string `json:"platform"`
}

// GetPraktekByPeserta returns practical exam results for a participant
func GetPraktekByPeserta(pesertaID string) ([]UjiPraktek, error) {
	query := "SELECT praktek_id, soal_id, peserta_id, hasil, tanggal, modified, platform FROM tb_uji_praktek WHERE peserta_id=? ORDER BY soal_id"
	rows, err := config.DB.Query(query, pesertaID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []UjiPraktek
	for rows.Next() {
		var u UjiPraktek
		err := rows.Scan(&u.PraktekID, &u.SoalID, &u.PesertaID, &u.Hasil, &u.Tanggal, &u.Modified, &u.Platform)
		if err != nil {
			return nil, err
		}
		list = append(list, u)
	}
	return list, nil
}

// GetPraktekByID returns a single practical exam result
func GetPraktekByID(praktekID int) (*UjiPraktek, error) {
	query := "SELECT praktek_id, soal_id, peserta_id, hasil, tanggal, modified, platform FROM tb_uji_praktek WHERE praktek_id=?"
	var u UjiPraktek
	err := config.DB.QueryRow(query, praktekID).Scan(&u.PraktekID, &u.SoalID, &u.PesertaID, &u.Hasil, &u.Tanggal, &u.Modified, &u.Platform)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

// CreatePraktek creates a new practical exam entry
func CreatePraktek(u *UjiPraktek) error {
	query := "INSERT INTO tb_uji_praktek (soal_id, peserta_id, hasil, tanggal, modified, platform) VALUES (?, ?, ?, NOW(), NOW(), ?)"
	_, err := config.DB.Exec(query, u.SoalID, u.PesertaID, u.Hasil, u.Platform)
	return err
}

// UpdatePraktekHasilByID updates just the hasil for a practical exam
func UpdatePraktekHasilByID(praktekID int, hasil int) error {
	query := "UPDATE tb_uji_praktek SET hasil=?, modified=NOW() WHERE praktek_id=?"
	_, err := config.DB.Exec(query, hasil, praktekID)
	return err
}

// UpsertPraktek creates or updates practical exam results in batch
func UpsertPraktek(results []UjiPraktek, platform string) error {
	tx, err := config.DB.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("INSERT INTO tb_uji_praktek (soal_id, peserta_id, hasil, tanggal, modified, platform) VALUES (?, ?, ?, NOW(), NOW(), ?) ON DUPLICATE KEY UPDATE hasil=?, modified=NOW()")
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt.Close()

	for _, r := range results {
		_, err := stmt.Exec(r.SoalID, r.PesertaID, r.Hasil, platform, r.Hasil)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

// GetRataRataPraktek returns average practical score for a participant
func GetRataRataPraktek(pesertaID string) (float64, error) {
	var avg float64
	err := config.DB.QueryRow("SELECT COALESCE(AVG(hasil), 0) FROM tb_uji_praktek WHERE peserta_id=?", pesertaID).Scan(&avg)
	return avg, err
}
