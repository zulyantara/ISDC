package models

import (
	"jsdc-api/config"
)

type Kelas struct {
	KelasID int    `json:"kelasid"`      // auto increment PK
	KelasKd int    `json:"kelas_id"`     // kelas kode
	Kelas   string `json:"kelas"`        // nama kelas
	Tarif   int    `json:"tarif"`        // tarif
	TeoriID int    `json:"teori_id"`     // kategori teori (1=R2, 2=R4, 3=truck)
	AreaID  int    `json:"area_id"`      // area
}

// GetAllKelas returns all kelas
func GetAllKelas() ([]Kelas, error) {
	query := "SELECT kelasid, kelas_id, kelas, tarif, teori_id, area_id FROM mt_kelas ORDER BY kelas_id"
	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var kelasList []Kelas
	for rows.Next() {
		var k Kelas
		err := rows.Scan(&k.KelasID, &k.KelasKd, &k.Kelas, &k.Tarif, &k.TeoriID, &k.AreaID)
		if err != nil {
			return nil, err
		}
		kelasList = append(kelasList, k)
	}
	return kelasList, nil
}

// GetKelasByID returns a single kelas
func GetKelasByID(kelasID int) (*Kelas, error) {
	query := "SELECT kelasid, kelas_id, kelas, tarif, teori_id, area_id FROM mt_kelas WHERE kelasid = ?"
	var k Kelas
	err := config.DB.QueryRow(query, kelasID).Scan(&k.KelasID, &k.KelasKd, &k.Kelas, &k.Tarif, &k.TeoriID, &k.AreaID)
	if err != nil {
		return nil, err
	}
	return &k, nil
}

// GetKelasByKode returns kelas by kelas_id (kode)
func GetKelasByKode(kelasKd int) (*Kelas, error) {
	query := "SELECT kelasid, kelas_id, kelas, tarif, teori_id, area_id FROM mt_kelas WHERE kelas_id = ?"
	var k Kelas
	err := config.DB.QueryRow(query, kelasKd).Scan(&k.KelasID, &k.KelasKd, &k.Kelas, &k.Tarif, &k.TeoriID, &k.AreaID)
	if err != nil {
		return nil, err
	}
	return &k, nil
}

// GetKelasByArea returns kelas filtered by area
func GetKelasByArea(areaID int) ([]Kelas, error) {
	query := "SELECT kelasid, kelas_id, kelas, tarif, teori_id, area_id FROM mt_kelas WHERE area_id=? ORDER BY kelas_id"
	rows, err := config.DB.Query(query, areaID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var kelasList []Kelas
	for rows.Next() {
		var k Kelas
		err := rows.Scan(&k.KelasID, &k.KelasKd, &k.Kelas, &k.Tarif, &k.TeoriID, &k.AreaID)
		if err != nil {
			return nil, err
		}
		kelasList = append(kelasList, k)
	}
	return kelasList, nil
}

// CreateKelas creates a new kelas
func CreateKelas(k *Kelas) error {
	query := "INSERT INTO mt_kelas (kelas_id, kelas, tarif, teori_id, area_id) VALUES (?, ?, ?, ?, ?)"
	_, err := config.DB.Exec(query, k.KelasKd, k.Kelas, k.Tarif, k.TeoriID, k.AreaID)
	return err
}

// UpdateKelas updates kelas data
func UpdateKelas(kelasID int, k *Kelas) error {
	query := "UPDATE mt_kelas SET kelas_id=?, kelas=?, tarif=?, teori_id=?, area_id=? WHERE kelasid=?"
	_, err := config.DB.Exec(query, k.KelasKd, k.Kelas, k.Tarif, k.TeoriID, k.AreaID, kelasID)
	return err
}

// DeleteKelas deletes a kelas
func DeleteKelas(kelasID int) error {
	query := "DELETE FROM mt_kelas WHERE kelasid=?"
	_, err := config.DB.Exec(query, kelasID)
	return err
}

// GetTeoriID returns teori_id (category) from kelas
func GetTeoriID(kelasKd int) (int, error) {
	var teoriID int
	err := config.DB.QueryRow("SELECT teori_id FROM mt_kelas WHERE kelas_id=?", kelasKd).Scan(&teoriID)
	return teoriID, err
}
