-- Adminer 4.6.2 MySQL dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

DROP VIEW IF EXISTS `v_stock_opname`;
CREATE TABLE `v_stock_opname` (`so_tgl` date, `area_kode` varchar(2), `area_nama` varchar(100), `bulan` varchar(5000), `so_tahun` smallint(6), `so_id` int(11), `so_gudang` int(11), `so_bulan` smallint(6), `so_time_insert` timestamp, `so_user_insert` int(11), `so_lasttime_update` timestamp, `so_lastuser_update` int(11), `so_sinkron` tinyint(4), `sinkron` varchar(5));


DROP TABLE IF EXISTS `v_stock_opname`;
CREATE ALGORITHM=UNDEFINED DEFINER=`root`@`localhost` SQL SECURITY DEFINER VIEW `v_stock_opname` AS select `stock_opname`.`so_tgl` AS `so_tgl`,`mt_area`.`area_kode` AS `area_kode`,`mt_area`.`area_nama` AS `area_nama`,`getBulanTeks`(`stock_opname`.`so_bulan`) AS `bulan`,`stock_opname`.`so_tahun` AS `so_tahun`,`stock_opname`.`so_id` AS `so_id`,`stock_opname`.`so_gudang` AS `so_gudang`,`stock_opname`.`so_bulan` AS `so_bulan`,`stock_opname`.`so_time_insert` AS `so_time_insert`,`stock_opname`.`so_user_insert` AS `so_user_insert`,`stock_opname`.`so_lasttime_update` AS `so_lasttime_update`,`stock_opname`.`so_lastuser_update` AS `so_lastuser_update`,`stock_opname`.`so_sinkron` AS `so_sinkron`,if((`stock_opname`.`so_sinkron` = 1),'Sudah','Belum') AS `sinkron` from (`stock_opname` join `mt_area` on((`stock_opname`.`so_gudang` = `mt_area`.`area_id`)));

-- 2018-05-24 07:27:47
