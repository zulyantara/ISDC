package models

import (
	"isdc-api/config"
)

type Area struct {
	AreaID    int    `json:"area_id"`
	AreaKode  string `json:"area_kode"`
	AreaNama  string `json:"area_nama"`
	AreaAlamat string `json:"area_alamat"`
	AreaTelp  string `json:"area_telp"`
}

// GetAllArea returns all areas
func GetAllArea() ([]Area, error) {
	query := "SELECT area_id, area_kode, area_nama, area_alamat, area_telp FROM mt_area ORDER BY area_id"
	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var areas []Area
	for rows.Next() {
		var a Area
		err := rows.Scan(&a.AreaID, &a.AreaKode, &a.AreaNama, &a.AreaAlamat, &a.AreaTelp)
		if err != nil {
			return nil, err
		}
		areas = append(areas, a)
	}
	return areas, nil
}

// GetAreaByID returns a single area
func GetAreaByID(areaID int) (*Area, error) {
	query := "SELECT area_id, area_kode, area_nama, area_alamat, area_telp FROM mt_area WHERE area_id = ?"
	var a Area
	err := config.DB.QueryRow(query, areaID).Scan(&a.AreaID, &a.AreaKode, &a.AreaNama, &a.AreaAlamat, &a.AreaTelp)
	if err != nil {
		return nil, err
	}
	return &a, nil
}

// GetAreaByKode returns area by kode
func GetAreaByKode(kode string) (*Area, error) {
	query := "SELECT area_id, area_kode, area_nama, area_alamat, area_telp FROM mt_area WHERE area_kode = ?"
	var a Area
	err := config.DB.QueryRow(query, kode).Scan(&a.AreaID, &a.AreaKode, &a.AreaNama, &a.AreaAlamat, &a.AreaTelp)
	if err != nil {
		return nil, err
	}
	return &a, nil
}

// CreateArea creates a new area
func CreateArea(a *Area) error {
	query := "INSERT INTO mt_area (area_kode, area_nama, area_alamat, area_telp) VALUES (?, ?, ?, ?)"
	_, err := config.DB.Exec(query, a.AreaKode, a.AreaNama, a.AreaAlamat, a.AreaTelp)
	return err
}

// UpdateArea updates area data
func UpdateArea(areaID int, a *Area) error {
	query := "UPDATE mt_area SET area_kode=?, area_nama=?, area_alamat=?, area_telp=? WHERE area_id=?"
	_, err := config.DB.Exec(query, a.AreaKode, a.AreaNama, a.AreaAlamat, a.AreaTelp, areaID)
	return err
}

// DeleteArea deletes an area
func DeleteArea(areaID int) error {
	query := "DELETE FROM mt_area WHERE area_id=?"
	_, err := config.DB.Exec(query, areaID)
	return err
}
