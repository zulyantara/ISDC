-- ============================================================
-- FIX: Membuat trigger tb_daftar_before_insert yang hilang
-- Trigger ini otomatis copy data dari tb_daftar ke tb_peserta
-- saat ada data baru diinsert ke tb_daftar
-- ============================================================

DROP TRIGGER IF EXISTS `tb_daftar_before_insert`;

DELIMITER ;;
CREATE TRIGGER `tb_daftar_before_insert` BEFORE INSERT ON `tb_daftar`
FOR EACH ROW
BEGIN

  SET NEW.`tgl_input` = NOW();

  INSERT INTO tb_peserta (
    `peserta_id`, `tgl_daftar`, `nama`, `kelamin_id`, `kelas_id`,
    `biaya`, `counter`, `tgl_input`, `tgl_update`, `user_id`
  ) VALUES (
    NEW.`peserta_id`, NEW.`tgl_daftar`, NEW.`nama`, NEW.`kelamin_id`, NEW.`kelas_id`,
    NEW.`biaya`, NEW.`counter`, NEW.`tgl_input`, NEW.`tgl_update`, NEW.`user_id`
  );

END ;;
DELIMITER ;
