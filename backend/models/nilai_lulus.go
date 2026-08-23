package models

import (
	"jsdc-api/config"
)

type NilaiLulus struct {
	NLID    int     `json:"nl_id"`
	NLNilai float64 `json:"nl_nilai"`
}

// GetAllNilaiLulus returns all passing scores
func GetAllNilaiLulus() ([]NilaiLulus, error) {
	query := "SELECT nl_id, nl_nilai FROM mt_nilai_lulus ORDER BY nl_id"
	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []NilaiLulus
	for rows.Next() {
		var n NilaiLulus
		err := rows.Scan(&n.NLID, &n.NLNilai)
		if err != nil {
			return nil, err
		}
		list = append(list, n)
	}
	return list, nil
}

// GetNilaiLulusByID returns a single passing score
func GetNilaiLulusByID(id int) (*NilaiLulus, error) {
	query := "SELECT nl_id, nl_nilai FROM mt_nilai_lulus WHERE nl_id = ?"
	var n NilaiLulus
	err := config.DB.QueryRow(query, id).Scan(&n.NLID, &n.NLNilai)
	if err != nil {
		return nil, err
	}
	return &n, nil
}

// CreateNilaiLulus creates a new passing score
func CreateNilaiLulus(n *NilaiLulus) error {
	query := "INSERT INTO mt_nilai_lulus (nl_nilai) VALUES (?)"
	_, err := config.DB.Exec(query, n.NLNilai)
	return err
}

// UpdateNilaiLulus updates a passing score
func UpdateNilaiLulus(id int, n *NilaiLulus) error {
	query := "UPDATE mt_nilai_lulus SET nl_nilai=? WHERE nl_id=?"
	_, err := config.DB.Exec(query, n.NLNilai, id)
	return err
}

// DeleteNilaiLulus deletes a passing score
func DeleteNilaiLulus(id int) error {
	query := "DELETE FROM mt_nilai_lulus WHERE nl_id=?"
	_, err := config.DB.Exec(query, id)
	return err
}

// GetMaxNilaiLulus returns the maximum passing score
func GetMaxNilaiLulus() (float64, error) {
	var nilai float64
	err := config.DB.QueryRow("SELECT COALESCE(MAX(nl_nilai), 0) FROM mt_nilai_lulus").Scan(&nilai)
	return nilai, err
}
