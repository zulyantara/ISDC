package models

import (
	"jsdc-api/config"
)

type Daftar struct {
	PesertaID   string `json:"peserta_id"`
	TglDaftar   string `json:"tgl_daftar"`
	RefID       string `json:"ref_id"`
	NmDepan     string `json:"nm_depan"`
	NmTengah    string `json:"nm_tengah"`
	NmBelakang  string `json:"nm_belakang"`
	Nama        string `json:"nama"`
	KelaminID   int    `json:"kelamin_id"`
	TempatLahir string `json:"tempat_lahir"`
	TglLahir    string `json:"tgl_lahir"`
	Alamat1     string `json:"alamat1"`
	Alamat2     string `json:"alamat2"`
	Kota        string `json:"kota"`
	KelasID     int    `json:"kelas_id"`
	Harga       int    `json:"harga"`
	Diskon      int    `json:"diskon"`
	Biaya       int    `json:"biaya"`
	TglInput    string `json:"tgl_input"`
	TglUpdate   string `json:"tgl_update"`
	UserID      string `json:"user_id"`
	AreaID      int    `json:"area_id"`

	// Joined fields
	KelasNama   string `json:"kelas_nama,omitempty"`
	UserName    string `json:"user_name,omitempty"`
	AreaNama    string `json:"area_nama,omitempty"`
}

// DaftarView is used for the v_daftar view
type DaftarView struct {
	Daftar
	Kelas     string `json:"kelas"`
	UserNama  string `json:"user_name"`
	App       string `json:"app"`
	Tahun     int    `json:"tahun"`
	Bulan     int    `json:"bulan"`
	Tgl       int    `json:"tgl"`
	AreaKode  string `json:"area_kode"`
	AreaNama  string `json:"area_nama"`
}

// GetAllDaftar returns all registrations (with joins)
func GetAllDaftar(areaID int) ([]DaftarView, error) {
	query := `SELECT d.peserta_id, COALESCE(d.tgl_daftar,''), COALESCE(d.ref_id,''), COALESCE(d.nm_depan,''), COALESCE(d.nm_tengah,''), COALESCE(d.nm_belakang,''), 
		d.nama, d.kelamin_id, COALESCE(d.tempat_lahir,''), COALESCE(d.tgl_lahir,''), COALESCE(d.alamat1,''), COALESCE(d.alamat2,''), COALESCE(d.kota,''), 
		d.kelas_id, COALESCE(k.kelas,''), d.harga, d.diskon, d.biaya, COALESCE(d.tgl_input,''), COALESCE(d.tgl_update,''), 
		d.user_id, COALESCE(u.user_name,''), d.area_id, COALESCE(a.area_kode,''), COALESCE(a.area_nama,''),
		YEAR(COALESCE(d.tgl_daftar,'2000-01-01')) as tahun, MONTH(COALESCE(d.tgl_daftar,'2000-01-01')) as bulan, DAY(COALESCE(d.tgl_daftar,'2000-01-01')) as tgl
		FROM tb_daftar d
		LEFT JOIN mt_kelas k ON d.kelas_id = k.kelas_id AND d.area_id = k.area_id
		LEFT JOIN mt_user u ON d.user_id = u.user_id
		LEFT JOIN mt_area a ON d.area_id = a.area_id`

	var args []interface{}
	if areaID > 0 {
		query += " WHERE d.area_id=?"
		args = append(args, areaID)
	}
	query += " ORDER BY d.tgl_input DESC"

	rows, err := config.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []DaftarView
	for rows.Next() {
		var d DaftarView
		err := rows.Scan(&d.PesertaID, &d.TglDaftar, &d.RefID, &d.NmDepan, &d.NmTengah, &d.NmBelakang,
			&d.Nama, &d.KelaminID, &d.TempatLahir, &d.TglLahir, &d.Alamat1, &d.Alamat2, &d.Kota,
			&d.KelasID, &d.Kelas, &d.Harga, &d.Diskon, &d.Biaya, &d.TglInput, &d.TglUpdate,
			&d.UserID, &d.UserName, &d.AreaID, &d.AreaKode, &d.AreaNama,
			&d.Tahun, &d.Bulan, &d.Tgl)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

// GetDaftarByID returns a single registration
func GetDaftarByID(pesertaID string) (*DaftarView, error) {
	query := `SELECT d.peserta_id, COALESCE(d.tgl_daftar,''), COALESCE(d.ref_id,''), COALESCE(d.nm_depan,''), COALESCE(d.nm_tengah,''), COALESCE(d.nm_belakang,''), 
		d.nama, d.kelamin_id, COALESCE(d.tempat_lahir,''), COALESCE(d.tgl_lahir,''), COALESCE(d.alamat1,''), COALESCE(d.alamat2,''), COALESCE(d.kota,''), 
		d.kelas_id, COALESCE(k.kelas,''), d.harga, d.diskon, d.biaya, COALESCE(d.tgl_input,''), COALESCE(d.tgl_update,''), 
		d.user_id, COALESCE(u.user_name,''), d.area_id, COALESCE(a.area_kode,''), COALESCE(a.area_nama,''),
		YEAR(COALESCE(d.tgl_daftar,'2000-01-01')) as tahun, MONTH(COALESCE(d.tgl_daftar,'2000-01-01')) as bulan, DAY(COALESCE(d.tgl_daftar,'2000-01-01')) as tgl
		FROM tb_daftar d
		LEFT JOIN mt_kelas k ON d.kelas_id = k.kelas_id AND d.area_id = k.area_id
		LEFT JOIN mt_user u ON d.user_id = u.user_id
		LEFT JOIN mt_area a ON d.area_id = a.area_id
		WHERE d.peserta_id = ?`

	var d DaftarView
	err := config.DB.QueryRow(query, pesertaID).Scan(&d.PesertaID, &d.TglDaftar, &d.RefID, &d.NmDepan, &d.NmTengah, &d.NmBelakang,
		&d.Nama, &d.KelaminID, &d.TempatLahir, &d.TglLahir, &d.Alamat1, &d.Alamat2, &d.Kota,
		&d.KelasID, &d.Kelas, &d.Harga, &d.Diskon, &d.Biaya, &d.TglInput, &d.TglUpdate,
		&d.UserID, &d.UserName, &d.AreaID, &d.AreaKode, &d.AreaNama,
		&d.Tahun, &d.Bulan, &d.Tgl)
	if err != nil {
		return nil, err
	}
	return &d, nil
}

// CreateDaftar creates a new registration
func CreateDaftar(d *Daftar) error {
	query := `INSERT INTO tb_daftar (peserta_id, tgl_daftar, ref_id, nm_depan, nm_tengah, nm_belakang, 
		nama, kelamin_id, tempat_lahir, tgl_lahir, alamat1, alamat2, kota, kelas_id, 
		harga, diskon, biaya, tgl_input, user_id, area_id) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, NOW(), ?, ?)`
	_, err := config.DB.Exec(query, d.PesertaID, d.TglDaftar, d.RefID, d.NmDepan, d.NmTengah, d.NmBelakang,
		d.Nama, d.KelaminID, d.TempatLahir, d.TglLahir, d.Alamat1, d.Alamat2, d.Kota, d.KelasID,
		d.Harga, d.Diskon, d.Biaya, d.UserID, d.AreaID)
	return err
}

// UpdateDaftar updates a registration
func UpdateDaftar(pesertaID string, d *Daftar) error {
	query := `UPDATE tb_daftar SET nm_depan=?, nm_tengah=?, nm_belakang=?, nama=?, 
		kelamin_id=?, tempat_lahir=?, tgl_lahir=?, alamat1=?, alamat2=?, kota=?, 
		kelas_id=?, harga=?, diskon=?, biaya=?, tgl_update=NOW() 
		WHERE peserta_id=?`
	_, err := config.DB.Exec(query, d.NmDepan, d.NmTengah, d.NmBelakang, d.Nama,
		d.KelaminID, d.TempatLahir, d.TglLahir, d.Alamat1, d.Alamat2, d.Kota,
		d.KelasID, d.Harga, d.Diskon, d.Biaya, pesertaID)
	return err
}

// DeleteDaftar deletes a registration
func DeleteDaftar(pesertaID string) error {
	_, err := config.DB.Exec("DELETE FROM tb_daftar WHERE peserta_id=?", pesertaID)
	return err
}

// CountDaftar returns total registration count
func CountDaftar(areaID int) (int64, error) {
	var count int64
	if areaID > 0 {
		err := config.DB.QueryRow("SELECT COUNT(*) FROM tb_daftar WHERE area_id=?", areaID).Scan(&count)
		return count, err
	}
	err := config.DB.QueryRow("SELECT COUNT(*) FROM tb_daftar").Scan(&count)
	return count, err
}

// GetLastPesertaID gets the last auto-generated peserta_id for a given area and month
func GetLastPesertaID(areaKode, yearMonth string) (string, error) {
	query := "SELECT peserta_id FROM tb_daftar WHERE SUBSTRING(peserta_id,1,7)=? AND area_id IN (SELECT area_id FROM mt_area WHERE area_kode=?) ORDER BY peserta_id DESC LIMIT 1"
	var pesertaID string
	err := config.DB.QueryRow(query, yearMonth, areaKode).Scan(&pesertaID)
	return pesertaID, err
}

// GetCountDaftarToday returns today's registration count for a user
func GetCountDaftarToday(userID string) (int64, error) {
	var count int64
	query := "SELECT COUNT(*) FROM tb_daftar WHERE user_id=? AND DATE(tgl_input)=CURDATE()"
	err := config.DB.QueryRow(query, userID).Scan(&count)
	return count, err
}
