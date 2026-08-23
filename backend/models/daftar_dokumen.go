package models

import (
	"isdc-api/config"
)

type DaftarDokumen struct {
	ID           int    `json:"id"`
	TglEntri     string `json:"tgl_entri"`
	JenisDokumen int    `json:"jenis_dokumen"`
	NamaDokumen  string `json:"nama_dokumen"`
	TglExp       string `json:"tgl_exp"`
	Lokasi       string `json:"lokasi"`
	Ket          string `json:"ket"`
	Area         string `json:"area"`
	Organisasi   string `json:"organisasi"`
	Aset         int    `json:"aset"`
	Hari         int    `json:"hari"`
	Email        string `json:"email"`
	Jam          int    `json:"jam"`
	Menit        int    `json:"menit"`
	File         string `json:"file"`
}

// GetAllDaftarDokumen returns all document entries
func GetAllDaftarDokumen() ([]DaftarDokumen, error) {
	query := "SELECT id, tgl_entri, jenis_dokumen, nama_dokumen, tgl_exp, lokasi, ket, area, organisasi, aset, hari, email, jam, menit, file FROM tb_daftar_dokumen ORDER BY tgl_entri DESC"
	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []DaftarDokumen
	for rows.Next() {
		var d DaftarDokumen
		err := rows.Scan(&d.ID, &d.TglEntri, &d.JenisDokumen, &d.NamaDokumen, &d.TglExp, &d.Lokasi, &d.Ket, &d.Area, &d.Organisasi, &d.Aset, &d.Hari, &d.Email, &d.Jam, &d.Menit, &d.File)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

// CreateDaftarDokumen creates a new document entry
func CreateDaftarDokumen(d *DaftarDokumen) error {
	query := "INSERT INTO tb_daftar_dokumen (tgl_entri, jenis_dokumen, nama_dokumen, tgl_exp, lokasi, ket, area, organisasi, aset, hari, email, jam, menit, file) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	_, err := config.DB.Exec(query, d.TglEntri, d.JenisDokumen, d.NamaDokumen, d.TglExp, d.Lokasi, d.Ket, d.Area, d.Organisasi, d.Aset, d.Hari, d.Email, d.Jam, d.Menit, d.File)
	return err
}

// UpdateDaftarDokumen updates a document entry
func UpdateDaftarDokumen(id int, d *DaftarDokumen) error {
	query := "UPDATE tb_daftar_dokumen SET tgl_entri=?, jenis_dokumen=?, nama_dokumen=?, tgl_exp=?, lokasi=?, ket=?, area=?, organisasi=?, aset=?, hari=?, email=?, jam=?, menit=?, file=? WHERE id=?"
	_, err := config.DB.Exec(query, d.TglEntri, d.JenisDokumen, d.NamaDokumen, d.TglExp, d.Lokasi, d.Ket, d.Area, d.Organisasi, d.Aset, d.Hari, d.Email, d.Jam, d.Menit, d.File, id)
	return err
}

// DeleteDaftarDokumen deletes a document entry
func DeleteDaftarDokumen(id int) error {
	_, err := config.DB.Exec("DELETE FROM tb_daftar_dokumen WHERE id=?", id)
	return err
}
