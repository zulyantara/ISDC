package models

import (
	"isdc-api/config"
)

type Peserta struct {
	PesertaID    string `json:"peserta_id"`
	TglDaftar    string `json:"tgl_daftar"`
	Nama         string `json:"nama"`
	KelaminID    int    `json:"kelamin_id"`
	KelasID      int    `json:"kelas_id"`
	Biaya        int    `json:"biaya"`
	TeoriHadir   string `json:"teori_hadir"`
	TeoriTgl     string `json:"teori_tgl"`
	TeoriKode    string `json:"teori_kode"`
	TeoriNilai   int    `json:"teori_nilai"`
	TeoriHasil   string `json:"teori_hasil"`
	TeoriCetak   string `json:"teori_cetak"`
	OperatorID   int    `json:"operator_id"`
	ClientID     int    `json:"client_id"`
	AKunci       string `json:"a_kunci"`
	AJawab       string `json:"a_jawab"`
	Penguji      string `json:"penguji"`
	SertifNomor  string `json:"sertif_nomor"`
	SertifTanggal string `json:"sertif_tanggal"`
	SertifCetak  string `json:"sertif_cetak"`
	TglInput     string `json:"tgl_input"`
	TglUpdate    string `json:"tgl_update"`
	UserID       string `json:"user_id"`
	AreaID       int    `json:"area_id"`
	PraktekHasil string `json:"praktek_hasil"`
	PraktekID    int    `json:"praktek_id"`

	// Joined fields
	KelasNama string `json:"kelas_nama,omitempty"`
	TeoriID   int    `json:"teori_id,omitempty"`
}

// GetAllPeserta returns all participants with joins
func GetAllPeserta(areaID int) ([]Peserta, error) {
	query := `SELECT p.peserta_id, COALESCE(p.tgl_daftar,''), p.nama, p.kelamin_id, p.kelas_id, 
		p.biaya, COALESCE(p.teori_hadir,'T'), COALESCE(p.teori_tgl,''), COALESCE(p.teori_kode,''), 
		p.teori_nilai, COALESCE(p.teori_hasil,'T'), COALESCE(p.teori_cetak,'T'),
		p.operator_id, p.client_id, COALESCE(p.a_kunci,''), COALESCE(p.a_jawab,''), 
		COALESCE(p.penguji,''), COALESCE(p.sertif_nomor,''), COALESCE(p.sertif_tanggal,''), 
		COALESCE(p.sertif_cetak,'T'), COALESCE(p.tgl_input,''), COALESCE(p.tgl_update,''),
		COALESCE(p.user_id,''), COALESCE(p.area_id, 0), COALESCE(p.praktek_hasil,''), 
		COALESCE(p.praktek_id, 0), COALESCE(k.kelas,''), COALESCE(k.teori_id, 0)
		FROM tb_peserta p
		LEFT JOIN mt_kelas k ON p.kelas_id = k.kelas_id AND p.area_id = k.area_id`

	var args []interface{}
	if areaID > 0 {
		query += " WHERE p.area_id=?"
		args = append(args, areaID)
	}
	query += " ORDER BY p.tgl_input DESC"

	rows, err := config.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []Peserta
	for rows.Next() {
		var p Peserta
		err := rows.Scan(&p.PesertaID, &p.TglDaftar, &p.Nama, &p.KelaminID, &p.KelasID,
			&p.Biaya, &p.TeoriHadir, &p.TeoriTgl, &p.TeoriKode,
			&p.TeoriNilai, &p.TeoriHasil, &p.TeoriCetak,
			&p.OperatorID, &p.ClientID, &p.AKunci, &p.AJawab,
			&p.Penguji, &p.SertifNomor, &p.SertifTanggal,
			&p.SertifCetak, &p.TglInput, &p.TglUpdate,
			&p.UserID, &p.AreaID, &p.PraktekHasil,
			&p.PraktekID, &p.KelasNama, &p.TeoriID)
		if err != nil {
			return nil, err
		}
		list = append(list, p)
	}
	return list, nil
}

// GetPesertaByID returns a single participant
func GetPesertaByID(pesertaID string) (*Peserta, error) {
	query := `SELECT p.peserta_id, COALESCE(p.tgl_daftar,''), p.nama, p.kelamin_id, p.kelas_id, 
		p.biaya, COALESCE(p.teori_hadir,'T'), COALESCE(p.teori_tgl,''), COALESCE(p.teori_kode,''), 
		p.teori_nilai, COALESCE(p.teori_hasil,'T'), COALESCE(p.teori_cetak,'T'),
		p.operator_id, p.client_id, COALESCE(p.a_kunci,''), COALESCE(p.a_jawab,''), 
		COALESCE(p.penguji,''), COALESCE(p.sertif_nomor,''), COALESCE(p.sertif_tanggal,''), 
		COALESCE(p.sertif_cetak,'T'), COALESCE(p.tgl_input,''), COALESCE(p.tgl_update,''),
		COALESCE(p.user_id,''), COALESCE(p.area_id, 0), COALESCE(p.praktek_hasil,''), 
		COALESCE(p.praktek_id, 0), COALESCE(k.kelas,''), COALESCE(k.teori_id, 0)
		FROM tb_peserta p
		LEFT JOIN mt_kelas k ON p.kelas_id = k.kelas_id AND p.area_id = k.area_id
		WHERE p.peserta_id = ?`

	var p Peserta
	err := config.DB.QueryRow(query, pesertaID).Scan(&p.PesertaID, &p.TglDaftar, &p.Nama, &p.KelaminID, &p.KelasID,
		&p.Biaya, &p.TeoriHadir, &p.TeoriTgl, &p.TeoriKode,
		&p.TeoriNilai, &p.TeoriHasil, &p.TeoriCetak,
		&p.OperatorID, &p.ClientID, &p.AKunci, &p.AJawab,
		&p.Penguji, &p.SertifNomor, &p.SertifTanggal,
		&p.SertifCetak, &p.TglInput, &p.TglUpdate,
		&p.UserID, &p.AreaID, &p.PraktekHasil,
		&p.PraktekID, &p.KelasNama, &p.TeoriID)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// CheckPesertaExists checks if a participant has already been examined
func CheckPesertaExists(pesertaID string) (bool, error) {
	var count int64
	err := config.DB.QueryRow("SELECT COUNT(*) FROM tb_peserta WHERE peserta_id=?", pesertaID).Scan(&count)
	return count > 0, err
}

// UpdatePraktekHasil updates the practical exam result for a participant
func UpdatePraktekHasil(pesertaID string, hasil string) error {
	query := "UPDATE tb_peserta SET praktek_hasil=?, tgl_update=NOW() WHERE peserta_id=?"
	_, err := config.DB.Exec(query, hasil, pesertaID)
	return err
}

// UpdateTeori updates theory exam data for a participant
func UpdateTeori(pesertaID string, hadir, kode string, nilai int, hasil string) error {
	query := "UPDATE tb_peserta SET teori_hadir=?, teori_tgl=CURDATE(), teori_kode=?, teori_nilai=?, teori_hasil=?, tgl_update=NOW() WHERE peserta_id=?"
	_, err := config.DB.Exec(query, hadir, kode, nilai, hasil, pesertaID)
	return err
}

// UpdateSertif updates sertifikat data for a participant
func UpdateSertif(pesertaID, nomor, tanggal string) error {
	query := "UPDATE tb_peserta SET sertif_nomor=?, sertif_tanggal=?, sertif_cetak='Y', tgl_update=NOW() WHERE peserta_id=?"
	_, err := config.DB.Exec(query, nomor, tanggal, pesertaID)
	return err
}

// SearchPeserta searches participants by name or ID
func SearchPeserta(keyword string, areaID int) ([]Peserta, error) {
	query := `SELECT p.peserta_id, COALESCE(p.tgl_daftar,''), p.nama, p.kelamin_id, p.kelas_id, 
		p.biaya, COALESCE(p.teori_hadir,'T'), COALESCE(p.teori_tgl,''), COALESCE(p.teori_kode,''), 
		p.teori_nilai, COALESCE(p.teori_hasil,'T'), COALESCE(p.teori_cetak,'T'),
		p.operator_id, p.client_id, COALESCE(p.a_kunci,''), COALESCE(p.a_jawab,''), 
		COALESCE(p.penguji,''), COALESCE(p.sertif_nomor,''), COALESCE(p.sertif_tanggal,''), 
		COALESCE(p.sertif_cetak,'T'), COALESCE(p.tgl_input,''), COALESCE(p.tgl_update,''),
		COALESCE(p.user_id,''), COALESCE(p.area_id, 0), COALESCE(p.praktek_hasil,''), 
		COALESCE(p.praktek_id, 0), COALESCE(k.kelas,''), COALESCE(k.teori_id, 0)
		FROM tb_peserta p
		LEFT JOIN mt_kelas k ON p.kelas_id = k.kelas_id AND p.area_id = k.area_id
		WHERE (p.peserta_id LIKE ? OR p.nama LIKE ?)`

	args := []interface{}{"%" + keyword + "%", "%" + keyword + "%"}
	if areaID > 0 {
		query += " AND p.area_id=?"
		args = append(args, areaID)
	}
	query += " ORDER BY p.tgl_input DESC LIMIT 50"

	rows, err := config.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []Peserta
	for rows.Next() {
		var p Peserta
		err := rows.Scan(&p.PesertaID, &p.TglDaftar, &p.Nama, &p.KelaminID, &p.KelasID,
			&p.Biaya, &p.TeoriHadir, &p.TeoriTgl, &p.TeoriKode,
			&p.TeoriNilai, &p.TeoriHasil, &p.TeoriCetak,
			&p.OperatorID, &p.ClientID, &p.AKunci, &p.AJawab,
			&p.Penguji, &p.SertifNomor, &p.SertifTanggal,
			&p.SertifCetak, &p.TglInput, &p.TglUpdate,
			&p.UserID, &p.AreaID, &p.PraktekHasil,
			&p.PraktekID, &p.KelasNama, &p.TeoriID)
		if err != nil {
			return nil, err
		}
		list = append(list, p)
	}
	return list, nil
}
