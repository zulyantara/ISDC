-- Adminer 4.6.2 MySQL dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

DROP VIEW IF EXISTS `v_perawatan_barang`;
CREATE TABLE `v_perawatan_barang` (`rb_tgl` date, `barang` text, `brg_kode` varchar(255), `brg_nama` varchar(255), `gudang` varchar(105), `area_kode` varchar(2), `area_nama` varchar(100), `rb_jenis` varchar(255), `rb_biaya` float, `rb_id` int(11), `rb_barang` int(11), `rb_gudang` int(11), `rb_time_insert` timestamp, `rb_user_insert` int(11), `rb_lasttime_update` timestamp, `rb_lastuser_update` int(11));


DROP TABLE IF EXISTS `v_perawatan_barang`;
CREATE ALGORITHM=UNDEFINED DEFINER=`root`@`localhost` SQL SECURITY DEFINER VIEW `v_perawatan_barang` AS select `perawatan_barang`.`rb_tgl` AS `rb_tgl`,concat(`mt_barang`.`brg_kode`,' - ',`mt_barang`.`brg_nama`) AS `barang`,`mt_barang`.`brg_kode` AS `brg_kode`,`mt_barang`.`brg_nama` AS `brg_nama`,concat(`mt_area`.`area_kode`,' - ',`mt_area`.`area_nama`) AS `gudang`,`mt_area`.`area_kode` AS `area_kode`,`mt_area`.`area_nama` AS `area_nama`,`perawatan_barang`.`rb_jenis` AS `rb_jenis`,`perawatan_barang`.`rb_biaya` AS `rb_biaya`,`perawatan_barang`.`rb_id` AS `rb_id`,`perawatan_barang`.`rb_barang` AS `rb_barang`,`perawatan_barang`.`rb_gudang` AS `rb_gudang`,`perawatan_barang`.`rb_time_insert` AS `rb_time_insert`,`perawatan_barang`.`rb_user_insert` AS `rb_user_insert`,`perawatan_barang`.`rb_lasttime_update` AS `rb_lasttime_update`,`perawatan_barang`.`rb_lastuser_update` AS `rb_lastuser_update` from ((`perawatan_barang` join `mt_barang` on((`perawatan_barang`.`rb_barang` = `mt_barang`.`brg_id`))) join `mt_area` on((`perawatan_barang`.`rb_gudang` = `mt_area`.`area_id`)));

-- 2018-05-24 07:26:41
