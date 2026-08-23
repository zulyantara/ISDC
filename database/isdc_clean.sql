-- ============================================
-- ISDC Clean Database Schema
-- All tables empty except: mt_kelas, mt_user (ADMIN only)
-- Default password: password321!*
-- ============================================

SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';
SET NAMES utf8mb4;

-- ============================================
-- MASTER TABLES
-- ============================================

DROP TABLE IF EXISTS `mt_app`;
CREATE TABLE `mt_app` (
  `app_id` tinyint(1) unsigned DEFAULT '0',
  `app` varchar(50) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=latin1 ROW_FORMAT=FIXED COMMENT='App Configuration';

INSERT INTO `mt_app` (`app_id`, `app`) VALUES
(0, 'ADMIN'),
(1, 'KASIR 1'),
(2, 'KASIR 2'),
(3, 'KASIR 3'),
(100, 'ADMIN'),
(101, 'INSTRUKTUR 1'),
(102, 'INSTRUKTUR 2');

DROP TABLE IF EXISTS `mt_area`;
CREATE TABLE `mt_area` (
  `area_id` int(11) NOT NULL AUTO_INCREMENT,
  `area_kode` varchar(5) NOT NULL,
  `area_nama` varchar(100) NOT NULL,
  `area_alamat` varchar(100) DEFAULT NULL,
  `area_telp` varchar(50) DEFAULT NULL,
  `area_insert_date` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`area_id`),
  KEY `area_kode` (`area_kode`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO `mt_area` (`area_id`, `area_kode`, `area_nama`, `area_alamat`, `area_telp`) VALUES
(1, '01', 'JAKARTA', 'Daan Mogot Jakarta Barat', '');

DROP TABLE IF EXISTS `mt_kelas`;
CREATE TABLE `mt_kelas` (
  `kelas_id` smallint(5) unsigned NOT NULL,
  `kelas` varchar(50) DEFAULT NULL,
  `tarif` int(10) unsigned NOT NULL DEFAULT '0',
  `teori_id` tinyint(3) unsigned NOT NULL DEFAULT '0',
  `area_id` int(11) NOT NULL,
  PRIMARY KEY (`kelas_id`) USING BTREE,
  KEY `area_id` (`area_id`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1 ROW_FORMAT=FIXED;

INSERT INTO `mt_kelas` (`kelas_id`, `kelas`, `tarif`, `teori_id`, `area_id`) VALUES
(101, 'BASIC SAFETY RIDING', 500000, 1, 0),
(102, 'ADVANCE SAFETY RIDING', 2000000, 1, 0),
(103, 'GREEN DRIVER RIDING', 1500000, 1, 0),
(201, 'BASIC SAFETY DRIVING', 500000, 2, 0),
(202, 'GREEN DRIVER DRIVING', 1200000, 2, 0),
(203, 'ADVANCE SAFETY DRIVING', 2000000, 2, 0),
(302, 'BASIC SAFETY DRIVING TRUCK', 750000, 3, 0),
(303, 'ADVANCE SAFETY DRIVING TRUCK', 2000000, 3, 0),
(401, 'TK', 30000, 0, 0),
(402, 'SD', 40000, 0, 0),
(403, 'SMP', 50000, 0, 0),
(404, 'SMA', 60000, 0, 0),
(501, 'CORPORATION SAFETY RIDING', 2000000, 1, 0),
(502, 'CORPORATION SAFETY DRIVING', 2000000, 2, 0),
(503, 'CORPORATION SAFETY DRIVING TRUCK', 2000000, 3, 0),
(601, 'TNI,POLRI,REKANAN R2', 0, 1, 0),
(602, 'TNI POLRI REKANAN R4', 0, 2, 0),
(603, 'TNI POLRI REKANAN R6+', 0, 3, 0),
(808, 'DEFENSIVE DRIVING BAHASA INGGRIS RODA 4', 1500000, 0, 0),
(809, 'DEFENSIVE RIDING BAHASA INGGRIS R2', 1500000, 0, 0);

DROP TABLE IF EXISTS `mt_level`;
CREATE TABLE `mt_level` (
  `level_id` int(11) NOT NULL AUTO_INCREMENT,
  `level_desc` varchar(50) NOT NULL,
  PRIMARY KEY (`level_id`),
  KEY `level_desc` (`level_desc`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO `mt_level` (`level_id`, `level_desc`) VALUES
(1, 'ADMIN'),
(2, 'BOD'),
(3, 'USER');

DROP TABLE IF EXISTS `mt_user`;
CREATE TABLE `mt_user` (
  `user_id` varchar(50) NOT NULL,
  `user_name` varchar(50) NOT NULL,
  `user_pwd` varchar(255) NOT NULL,
  `user_level` tinyint(1) unsigned NOT NULL DEFAULT '0',
  `app_id` tinyint(3) unsigned NOT NULL DEFAULT '0',
  `flag` tinyint(1) NOT NULL DEFAULT '0',
  `aktif` tinyint(1) NOT NULL DEFAULT '-1',
  `area_id` int(11) NOT NULL,
  PRIMARY KEY (`user_id`),
  KEY `area_id` (`area_id`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1 ROW_FORMAT=FIXED;

-- ADMIN user only, password: password321!*
INSERT INTO `mt_user` (`user_id`, `user_name`, `user_pwd`, `user_level`, `app_id`, `flag`, `aktif`, `area_id`) VALUES
('ADMIN', 'ADMINISTRATOR', 'a6c993aceacc57e783d3379f06ac8305', 1, 0, 0, -1, 0);

-- ============================================
-- RBAC TABLES
-- ============================================

DROP TABLE IF EXISTS `menu`;
CREATE TABLE `menu` (
  `menu_id` int(11) NOT NULL AUTO_INCREMENT,
  `menu_ket` varchar(50) NOT NULL,
  `menu_parent` int(11) NOT NULL,
  `menu_url` varchar(100) NOT NULL,
  `menu_order` int(11) NOT NULL,
  PRIMARY KEY (`menu_id`),
  KEY `menu_parent` (`menu_parent`),
  KEY `menu_ket` (`menu_ket`) USING BTREE,
  KEY `menu_url` (`menu_url`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

DROP TABLE IF EXISTS `hak_akses`;
CREATE TABLE `hak_akses` (
  `ha_id` int(11) NOT NULL AUTO_INCREMENT,
  `ha_menu` int(11) NOT NULL,
  `ha_ur` int(11) NOT NULL,
  `ha_view` int(11) DEFAULT '0',
  `ha_insert` int(11) DEFAULT '0',
  `ha_update` int(11) DEFAULT '0',
  `ha_delete` int(11) DEFAULT '0',
  `ha_proses` int(11) DEFAULT '0',
  PRIMARY KEY (`ha_id`),
  KEY `ha_menu` (`ha_menu`),
  KEY `ha_ur` (`ha_ur`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- ============================================
-- APPLICATION TABLES
-- ============================================

DROP TABLE IF EXISTS `mt_nilai_lulus`;
CREATE TABLE `mt_nilai_lulus` (
  `nl_id` int(11) NOT NULL AUTO_INCREMENT,
  `nl_nilai` float NOT NULL,
  `nl_nilai_teori` float NOT NULL,
  `nl_insert_date` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`nl_id`),
  KEY `nl_nilai` (`nl_nilai`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

DROP TABLE IF EXISTS `tb_daftar`;
CREATE TABLE `tb_daftar` (
  `peserta_id` varchar(15) NOT NULL,
  `tgl_daftar` date DEFAULT NULL,
  `ref_id` varchar(10) DEFAULT NULL,
  `nm_depan` varchar(30) DEFAULT NULL,
  `nm_tengah` varchar(30) DEFAULT NULL,
  `nm_belakang` varchar(30) DEFAULT NULL,
  `nama` varchar(50) DEFAULT NULL,
  `kelamin_id` tinyint(1) unsigned NOT NULL DEFAULT '0',
  `tempat_lahir` varchar(30) DEFAULT NULL,
  `tgl_lahir` date DEFAULT NULL,
  `alamat1` varchar(50) DEFAULT NULL,
  `alamat2` varchar(50) DEFAULT NULL,
  `kota` varchar(30) DEFAULT NULL,
  `kelas_id` smallint(5) unsigned NOT NULL DEFAULT '0',
  `harga` int(10) unsigned NOT NULL DEFAULT '0',
  `diskon` int(10) unsigned NOT NULL DEFAULT '0',
  `biaya` int(10) unsigned NOT NULL DEFAULT '0',
  `cetak` varchar(1) NOT NULL DEFAULT 'N',
  `counter` smallint(5) unsigned zerofill NOT NULL DEFAULT '00000',
  `tgl_input` datetime DEFAULT NULL,
  `tgl_update` datetime DEFAULT NULL,
  `user_id` varchar(50) DEFAULT NULL,
  `log` text,
  `area_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`peserta_id`),
  KEY `area_id` (`area_id`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1 ROW_FORMAT=FIXED COMMENT='Pendaftaran Peserta';

DROP TABLE IF EXISTS `tb_daftar_log`;
CREATE TABLE `tb_daftar_log` (
  `log_id` int(15) NOT NULL AUTO_INCREMENT,
  `peserta_id` varchar(15) NOT NULL,
  `tgl_daftar` date DEFAULT NULL,
  `ref_id` varchar(10) DEFAULT NULL,
  `nm_depan` varchar(30) DEFAULT NULL,
  `nm_tengah` varchar(30) DEFAULT NULL,
  `nm_belakang` varchar(30) DEFAULT NULL,
  `nama` varchar(50) DEFAULT NULL,
  `kelamin_id` tinyint(1) unsigned NOT NULL DEFAULT '0',
  `tempat_lahir` varchar(30) DEFAULT NULL,
  `tgl_lahir` date DEFAULT NULL,
  `alamat1` varchar(50) DEFAULT NULL,
  `alamat2` varchar(50) DEFAULT NULL,
  `kota` varchar(30) DEFAULT NULL,
  `kelas_id` smallint(5) unsigned NOT NULL DEFAULT '0',
  `harga` int(10) unsigned NOT NULL DEFAULT '0',
  `diskon` int(10) unsigned NOT NULL DEFAULT '0',
  `biaya` int(10) unsigned NOT NULL DEFAULT '0',
  `cetak` varchar(1) NOT NULL DEFAULT 'N',
  `counter` smallint(5) unsigned zerofill NOT NULL DEFAULT '00000',
  `tgl_input` datetime DEFAULT NULL,
  `tgl_update` datetime DEFAULT NULL,
  `user_id` varchar(50) DEFAULT NULL,
  `log` text,
  `area_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`log_id`),
  KEY `area_id` (`area_id`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1 ROW_FORMAT=FIXED COMMENT='Log Pendaftaran Peserta';

DROP TABLE IF EXISTS `tb_peserta`;
CREATE TABLE `tb_peserta` (
  `peserta_id` varchar(15) NOT NULL,
  `tgl_daftar` date DEFAULT NULL,
  `nama` varchar(50) DEFAULT NULL,
  `kelamin_id` tinyint(1) unsigned NOT NULL DEFAULT '0',
  `kelas_id` smallint(5) unsigned NOT NULL DEFAULT '0',
  `biaya` int(10) unsigned NOT NULL DEFAULT '0',
  `teori_hadir` varchar(1) NOT NULL DEFAULT 'T',
  `teori_tgl` date DEFAULT NULL,
  `teori_kode` varchar(15) DEFAULT NULL,
  `teori_nilai` smallint(5) unsigned NOT NULL DEFAULT '0',
  `teori_hasil` varchar(1) NOT NULL DEFAULT 'T',
  `teori_cetak` varchar(1) NOT NULL DEFAULT 'T',
  `teori_log` text,
  `operator_id` tinyint(3) unsigned NOT NULL DEFAULT '0',
  `client_id` tinyint(3) unsigned NOT NULL DEFAULT '0',
  `a_kunci` varchar(50) DEFAULT NULL,
  `a_jawab` varchar(50) DEFAULT NULL,
  `penguji` varchar(50) DEFAULT NULL,
  `sertif_nomor` varchar(30) DEFAULT NULL,
  `sertif_tanggal` date DEFAULT NULL,
  `sertif_cetak` varchar(1) NOT NULL DEFAULT 'T',
  `counter` smallint(5) unsigned zerofill NOT NULL DEFAULT '00000',
  `tgl_input` datetime DEFAULT NULL,
  `tgl_update` datetime DEFAULT NULL,
  `user_id` varchar(50) DEFAULT NULL,
  `area_id` int(11) DEFAULT NULL,
  `praktek_hasil` varchar(3) DEFAULT NULL,
  `praktek_id` int(5) DEFAULT NULL,
  PRIMARY KEY (`peserta_id`),
  KEY `area_id` (`area_id`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1 ROW_FORMAT=FIXED COMMENT='Daftar peserta pelatihan';

DROP TABLE IF EXISTS `tb_soal`;
CREATE TABLE `tb_soal` (
  `ujian_id` int(5) NOT NULL AUTO_INCREMENT,
  `sesi` int(2) NOT NULL,
  `nomor` int(3) NOT NULL,
  `category` int(3) NOT NULL,
  `soal` varchar(255) NOT NULL,
  PRIMARY KEY (`ujian_id`),
  KEY `idx_soal` (`category`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

DROP TABLE IF EXISTS `tb_uji_praktek`;
CREATE TABLE `tb_uji_praktek` (
  `praktek_id` int(5) NOT NULL AUTO_INCREMENT,
  `soal_id` int(5) NOT NULL,
  `peserta_id` varchar(20) NOT NULL,
  `hasil` int(5) NOT NULL,
  `tanggal` datetime NOT NULL,
  `modified` datetime NOT NULL,
  `platform` varchar(20) NOT NULL,
  PRIMARY KEY (`praktek_id`),
  UNIQUE KEY `uk_praktek` (`soal_id`, `peserta_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

DROP TABLE IF EXISTS `tb_comments`;
CREATE TABLE `tb_comments` (
  `comment_id` int(5) NOT NULL AUTO_INCREMENT,
  `peserta_id` varchar(20) NOT NULL,
  `pengetahuan` text NOT NULL,
  `teknik` text NOT NULL,
  `perilaku` text NOT NULL,
  PRIMARY KEY (`comment_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

DROP TABLE IF EXISTS `tb_ujian`;
CREATE TABLE `tb_ujian` (
  `client_id` tinyint(3) unsigned NOT NULL,
  `operator_id` tinyint(3) unsigned NOT NULL DEFAULT '0',
  `peserta_id` varchar(15) NOT NULL,
  `teori_tgl` datetime DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=latin1 COMMENT='Daftar peserta pelatihan';

-- ============================================
-- DOKUMEN TABLES
-- ============================================

DROP TABLE IF EXISTS `tb_jenis_dokumen`;
CREATE TABLE `tb_jenis_dokumen` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `nama` varchar(25) NOT NULL,
  `ket` varchar(100) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

DROP TABLE IF EXISTS `tb_daftar_dokumen`;
CREATE TABLE `tb_daftar_dokumen` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `tgl_entri` date NOT NULL,
  `jenis_dokumen` int(11) DEFAULT NULL,
  `nama_dokumen` varchar(500) NOT NULL,
  `tgl_exp` date NOT NULL,
  `lokasi` varchar(35) NOT NULL,
  `ket` varchar(50) NOT NULL,
  `area` varchar(35) NOT NULL,
  `organisasi` varchar(35) NOT NULL,
  `aset` int(11) NOT NULL,
  `hari` int(11) NOT NULL,
  `email` varchar(45) NOT NULL,
  `jam` int(11) NOT NULL,
  `menit` int(11) NOT NULL,
  `file` varchar(100) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- ============================================
-- VIEWS (read-only, no data)
-- ============================================

DROP VIEW IF EXISTS `v_daftar`;
CREATE VIEW `v_daftar` AS
SELECT d.peserta_id, d.tgl_daftar, d.ref_id, d.nm_depan, d.nm_tengah, d.nm_belakang,
  d.nama, d.kelamin_id, d.tempat_lahir, d.tgl_lahir, d.alamat1, d.alamat2, d.kota,
  d.kelas_id, k.kelas, d.harga, d.diskon, d.biaya, d.counter, d.tgl_input, d.tgl_update,
  d.user_id, u.user_name, d.log, d.area_id, a.area_kode, a.area_nama
FROM tb_daftar d
LEFT JOIN mt_kelas k ON d.kelas_id = k.kelas_id
LEFT JOIN mt_user u ON d.user_id = u.user_id
LEFT JOIN mt_area a ON d.area_id = a.area_id;

DROP VIEW IF EXISTS `v_peserta`;
CREATE VIEW `v_peserta` AS
SELECT p.peserta_id, p.tgl_daftar, p.nama, p.kelamin_id, p.kelas_id, k.kelas, k.teori_id,
  p.biaya, p.teori_hadir, p.teori_tgl, p.teori_kode, p.teori_nilai, p.teori_hasil,
  p.teori_cetak, p.teori_log, p.operator_id, p.client_id, p.a_kunci, p.a_jawab,
  p.penguji, p.sertif_nomor, p.sertif_tanggal, p.sertif_cetak, p.counter,
  p.tgl_input, p.tgl_update, p.user_id, u.user_name, a.app,
  YEAR(p.tgl_daftar) AS tahun, MONTH(p.tgl_daftar) AS bulan, DAY(p.tgl_daftar) AS tgl,
  p.praktek_hasil, p.praktek_id
FROM tb_peserta p
LEFT JOIN mt_kelas k ON p.kelas_id = k.kelas_id
LEFT JOIN mt_user u ON p.user_id = u.user_id
LEFT JOIN mt_app a ON u.app_id = a.app_id;

DROP VIEW IF EXISTS `v_user`;
CREATE VIEW `v_user` AS
SELECT u.user_id, u.user_name, u.user_pwd, u.user_level, u.app_id, u.flag, u.aktif,
  CASE u.user_level WHEN 0 THEN 'USER' WHEN 1 THEN 'ADMIN' ELSE 'BOD' END AS level,
  a.app
FROM mt_user u
LEFT JOIN mt_app a ON u.app_id = a.app_id;

-- ============================================
-- TRIGGERS: daftar <-> peserta sync
-- NO triggers on mt_user (Go backend handles bcrypt)
-- ============================================

DELIMITER ;;

CREATE TRIGGER `tb_daftar_before_insert` BEFORE INSERT ON `tb_daftar` FOR EACH ROW
BEGIN
  SET NEW.`tgl_input` = NOW();
  INSERT INTO tb_peserta (
    `peserta_id`, `tgl_daftar`, `nama`, `kelamin_id`, `kelas_id`,
    `biaya`, `counter`, `tgl_input`, `tgl_update`, `user_id`, `area_id`
  ) VALUES (
    NEW.`peserta_id`, NEW.`tgl_daftar`, NEW.`nama`, NEW.`kelamin_id`, NEW.`kelas_id`,
    NEW.`biaya`, NEW.`counter`, NEW.`tgl_input`, NEW.`tgl_update`, NEW.`user_id`, NEW.`area_id`
  );
END;;

CREATE TRIGGER `tb_daftar_before_update` BEFORE UPDATE ON `tb_daftar` FOR EACH ROW
BEGIN
  SET NEW.`tgl_update` = NOW();
  UPDATE tb_peserta SET
    `peserta_id` = NEW.`peserta_id`, `tgl_daftar` = NEW.`tgl_daftar`,
    `nama` = NEW.`nama`, `kelamin_id` = NEW.`kelamin_id`, `kelas_id` = NEW.`kelas_id`,
    `biaya` = NEW.`biaya`, `counter` = NEW.`counter`, `tgl_input` = NEW.`tgl_input`,
    `tgl_update` = NEW.`tgl_update`, `user_id` = NEW.`user_id`, `area_id` = NEW.`area_id`
  WHERE `peserta_id` = OLD.`peserta_id`;
END;;

CREATE TRIGGER `tb_daftar_after_update` AFTER UPDATE ON `tb_daftar` FOR EACH ROW
BEGIN
  IF (NEW.`cetak` = OLD.`cetak`) THEN
    INSERT INTO tb_daftar_log (
      `peserta_id`, `tgl_daftar`, `ref_id`, `nm_depan`, `nm_tengah`,
      `nm_belakang`, `nama`, `kelamin_id`, `tempat_lahir`, `tgl_lahir`,
      `alamat1`, `alamat2`, `kota`, `kelas_id`, `harga`,
      `diskon`, `biaya`, `cetak`, `counter`, `tgl_input`, `tgl_update`,
      `user_id`, `log`, `area_id`
    ) VALUES (
      OLD.`peserta_id`, OLD.`tgl_daftar`, OLD.`ref_id`, OLD.`nm_depan`, OLD.`nm_tengah`,
      OLD.`nm_belakang`, OLD.`nama`, OLD.`kelamin_id`, OLD.`tempat_lahir`, OLD.`tgl_lahir`,
      OLD.`alamat1`, OLD.`alamat2`, OLD.`kota`, OLD.`kelas_id`, OLD.`harga`,
      OLD.`diskon`, OLD.`biaya`, OLD.`cetak`, OLD.`counter`, OLD.`tgl_input`, OLD.`tgl_update`,
      OLD.`user_id`, OLD.`log`, OLD.`area_id`
    );
  END IF;
END;;

CREATE TRIGGER `tb_daftar_after_delete` AFTER DELETE ON `tb_daftar` FOR EACH ROW
BEGIN
  DELETE FROM tb_peserta WHERE peserta_id = OLD.`peserta_id`;
END;;

DELIMITER ;
