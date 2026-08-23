package models

import (
	"jsdc-api/config"
)

type Level struct {
	LevelID   int    `json:"level_id"`
	LevelDesc string `json:"level_desc"`
}

type Menu struct {
	MenuID     int    `json:"menu_id"`
	MenuKet    string `json:"menu_ket"`
	MenuParent int    `json:"menu_parent"`
	MenuURL    string `json:"menu_url"`
	MenuOrder  int    `json:"menu_order"`
}

type HakAkses struct {
	HAID     int `json:"ha_id"`
	HAMenu   int `json:"ha_menu"`
	HAUR     int `json:"ha_ur"` // user_role / level_id
	HAView   int `json:"ha_view"`
	HAInsert int `json:"ha_insert"`
	HAUpdate int `json:"ha_update"`
	HADelete int `json:"ha_delete"`
	HAProses int `json:"ha_proses"`
}

// Permission represents a combined menu + permission for a role
type Permission struct {
	MenuID   int    `json:"menu_id"`
	MenuKet  string `json:"menu_ket"`
	MenuURL  string `json:"menu_url"`
	MenuOrder int   `json:"menu_order"`
	MenuParent int  `json:"menu_parent"`
	View     int    `json:"view"`
	Insert   int    `json:"insert"`
	Update   int    `json:"update"`
	Delete   int    `json:"delete"`
	Proses   int    `json:"proses"`
}

// GetAllLevels returns all user levels
func GetAllLevels() ([]Level, error) {
	rows, err := config.DB.Query("SELECT level_id, level_desc FROM mt_level ORDER BY level_id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []Level
	for rows.Next() {
		var l Level
		if err := rows.Scan(&l.LevelID, &l.LevelDesc); err != nil {
			return nil, err
		}
		list = append(list, l)
	}
	return list, nil
}

// GetAllMenus returns all menus
func GetAllMenus() ([]Menu, error) {
	rows, err := config.DB.Query("SELECT menu_id, menu_ket, menu_parent, menu_url, menu_order FROM menu ORDER BY menu_order")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []Menu
	for rows.Next() {
		var m Menu
		if err := rows.Scan(&m.MenuID, &m.MenuKet, &m.MenuParent, &m.MenuURL, &m.MenuOrder); err != nil {
			return nil, err
		}
		list = append(list, m)
	}
	return list, nil
}

// GetPermissionsByRole returns all menu permissions for a given role (level_id)
func GetPermissionsByRole(levelID int) ([]Permission, error) {
	query := `SELECT m.menu_id, m.menu_ket, m.menu_url, m.menu_order, m.menu_parent,
		COALESCE(h.ha_view,0), COALESCE(h.ha_insert,0), COALESCE(h.ha_update,0),
		COALESCE(h.ha_delete,0), COALESCE(h.ha_proses,0)
		FROM menu m
		LEFT JOIN hak_akses h ON m.menu_id = h.ha_menu AND h.ha_ur = ?
		WHERE h.ha_id IS NOT NULL
		ORDER BY m.menu_order`

	rows, err := config.DB.Query(query, levelID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []Permission
	for rows.Next() {
		var p Permission
		if err := rows.Scan(&p.MenuID, &p.MenuKet, &p.MenuURL, &p.MenuOrder, &p.MenuParent,
			&p.View, &p.Insert, &p.Update, &p.Delete, &p.Proses); err != nil {
			return nil, err
		}
		list = append(list, p)
	}
	return list, nil
}

// GetAllHakAkses returns all permissions
func GetAllHakAkses() ([]HakAkses, error) {
	rows, err := config.DB.Query("SELECT ha_id, ha_menu, ha_ur, ha_view, ha_insert, ha_update, ha_delete, ha_proses FROM hak_akses ORDER BY ha_ur, ha_menu")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []HakAkses
	for rows.Next() {
		var h HakAkses
		if err := rows.Scan(&h.HAID, &h.HAMenu, &h.HAUR, &h.HAView, &h.HAInsert, &h.HAUpdate, &h.HADelete, &h.HAProses); err != nil {
			return nil, err
		}
		list = append(list, h)
	}
	return list, nil
}

// GetHakAksesByRole returns permissions for a specific role
func GetHakAksesByRole(levelID int) ([]HakAkses, error) {
	rows, err := config.DB.Query("SELECT ha_id, ha_menu, ha_ur, ha_view, ha_insert, ha_update, ha_delete, ha_proses FROM hak_akses WHERE ha_ur = ? ORDER BY ha_menu", levelID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []HakAkses
	for rows.Next() {
		var h HakAkses
		if err := rows.Scan(&h.HAID, &h.HAMenu, &h.HAUR, &h.HAView, &h.HAInsert, &h.HAUpdate, &h.HADelete, &h.HAProses); err != nil {
			return nil, err
		}
		list = append(list, h)
	}
	return list, nil
}

// CreateHakAkses creates a new permission entry
func CreateHakAkses(h *HakAkses) error {
	query := "INSERT INTO hak_akses (ha_menu, ha_ur, ha_view, ha_insert, ha_update, ha_delete, ha_proses) VALUES (?, ?, ?, ?, ?, ?, ?)"
	_, err := config.DB.Exec(query, h.HAMenu, h.HAUR, h.HAView, h.HAInsert, h.HAUpdate, h.HADelete, h.HAProses)
	return err
}

// UpdateHakAkses updates an existing permission
func UpdateHakAkses(id int, h *HakAkses) error {
	query := "UPDATE hak_akSES SET ha_menu=?, ha_ur=?, ha_view=?, ha_insert=?, ha_update=?, ha_delete=?, ha_proses=? WHERE ha_id=?"
	_, err := config.DB.Exec(query, h.HAMenu, h.HAUR, h.HAView, h.HAInsert, h.HAUpdate, h.HADelete, h.HAProses, id)
	return err
}

// DeleteHakAkses deletes a permission
func DeleteHakAkses(id int) error {
	_, err := config.DB.Exec("DELETE FROM hak_akses WHERE ha_id=?", id)
	return err
}

// UpsertHakAksesForRole replaces all permissions for a role (bulk save)
func UpsertHakAksesForRole(levelID int, permissions []HakAkses) error {
	tx, err := config.DB.Begin()
	if err != nil {
		return err
	}

	// Delete existing permissions for this role
	_, err = tx.Exec("DELETE FROM hak_akses WHERE ha_ur=?", levelID)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Insert new permissions
	stmt, err := tx.Prepare("INSERT INTO hak_akses (ha_menu, ha_ur, ha_view, ha_insert, ha_update, ha_delete, ha_proses) VALUES (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt.Close()

	for _, h := range permissions {
		_, err = stmt.Exec(h.HAMenu, levelID, h.HAView, h.HAInsert, h.HAUpdate, h.HADelete, h.HAProses)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

// CreateLevel creates a new level/role
func CreateLevel(l *Level) error {
	_, err := config.DB.Exec("INSERT INTO mt_level (level_desc) VALUES (?)", l.LevelDesc)
	return err
}

// UpdateLevel updates a level/role
func UpdateLevel(id int, l *Level) error {
	_, err := config.DB.Exec("UPDATE mt_level SET level_desc=? WHERE level_id=?", l.LevelDesc, id)
	return err
}

// DeleteLevel deletes a level/role
func DeleteLevel(id int) error {
	_, err := config.DB.Exec("DELETE FROM mt_level WHERE level_id=?", id)
	return err
}

// CreateMenu creates a new menu
func CreateMenu(m *Menu) error {
	_, err := config.DB.Exec("INSERT INTO menu (menu_ket, menu_parent, menu_url, menu_order) VALUES (?, ?, ?, ?)", m.MenuKet, m.MenuParent, m.MenuURL, m.MenuOrder)
	return err
}

// UpdateMenu updates a menu
func UpdateMenu(id int, m *Menu) error {
	_, err := config.DB.Exec("UPDATE menu SET menu_ket=?, menu_parent=?, menu_url=?, menu_order=? WHERE menu_id=?", m.MenuKet, m.MenuParent, m.MenuURL, m.MenuOrder, id)
	return err
}

// DeleteMenu deletes a menu
func DeleteMenu(id int) error {
	_, err := config.DB.Exec("DELETE FROM menu WHERE menu_id=?", id)
	return err
}
