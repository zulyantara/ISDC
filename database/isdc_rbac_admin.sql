-- ============================================
-- ISDC RBAC: All Menus + Full Access for ADMIN
-- Run AFTER isdc_clean.sql
-- ============================================

SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

-- ============================================
-- MENUS (matches Sidebar.jsx menuUrl values)
-- ============================================

INSERT INTO `menu` (`menu_id`, `menu_ket`, `menu_parent`, `menu_url`, `menu_order`) VALUES
-- Root menus
(1,  'Dashboard',        0,  'welcome',       1),
(2,  'Pendaftaran',      0,  'pendaftaran',   2),
(3,  'Peserta',          0,  'petugas',       3),
(4,  'Sertifikat',       0,  'sertifikat',    4),
(5,  'Master Data',      0,  '#',             5),
-- Master Data children
(6,  'Kelas',            5,  'mt_kelas',      6),
(7,  'User',             5,  'mt_user',       7),
(8,  'Area',             5,  'mt_area',       8),
(9,  'Nilai Lulus',      5,  'mt_nilai_lulus',9),
(10, 'Soal',             5,  'ujian',         10),
(11, 'RBAC / Hak Akses', 5,  'rbac',          11);

-- ============================================
-- HAK AKSES: Full access for ADMIN (level_id=1)
-- view=1, insert=1, update=1, delete=1, proses=1
-- ============================================

INSERT INTO `hak_akses` (`ha_menu`, `ha_ur`, `ha_view`, `ha_insert`, `ha_update`, `ha_delete`, `ha_proses`) VALUES
-- Dashboard
(1,  1, 1, 1, 1, 1, 1),
-- Pendaftaran
(2,  1, 1, 1, 1, 1, 1),
-- Peserta
(3,  1, 1, 1, 1, 1, 1),
-- Sertifikat
(4,  1, 1, 1, 1, 1, 1),
-- Kelas
(6,  1, 1, 1, 1, 1, 1),
-- User
(7,  1, 1, 1, 1, 1, 1),
-- Area
(8,  1, 1, 1, 1, 1, 1),
-- Nilai Lulus
(9,  1, 1, 1, 1, 1, 1),
-- Soal
(10, 1, 1, 1, 1, 1, 1),
-- RBAC
(11, 1, 1, 1, 1, 1, 1);
