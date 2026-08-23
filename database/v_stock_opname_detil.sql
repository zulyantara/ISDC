-- Adminer 4.6.2 MySQL dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

DROP VIEW IF EXISTS `v_stock_opname_detil`;
CREATE TABLE `v_stock_opname_detil` (`barang` text, `brg_kode` varchar(255), `brg_nama` varchar(255), `sod_qty_db` float, `sod_qty_riil` float, `sod_id` int(11), `sod_header` int(11), `sod_barang` int(11), `sod_time_insert` timestamp, `sod_user_insert` int(11), `sod_lasttime_update` timestamp, `sod_lastuser_update` int(11), `sb_nama` varchar(255), `sod_satuan` int(11));


DROP TABLE IF EXISTS `v_stock_opname_detil`;
CREATE ALGORITHM=UNDEFINED DEFINER=`root`@`localhost` SQL SECURITY DEFINER VIEW `v_stock_opname_detil` AS select concat(`mt_barang`.`brg_kode`,' - ',`mt_barang`.`brg_nama`) AS `barang`,`mt_barang`.`brg_kode` AS `brg_kode`,`mt_barang`.`brg_nama` AS `brg_nama`,`stock_opname_detil`.`sod_qty_db` AS `sod_qty_db`,`stock_opname_detil`.`sod_qty_riil` AS `sod_qty_riil`,`stock_opname_detil`.`sod_id` AS `sod_id`,`stock_opname_detil`.`sod_header` AS `sod_header`,`stock_opname_detil`.`sod_barang` AS `sod_barang`,`stock_opname_detil`.`sod_time_insert` AS `sod_time_insert`,`stock_opname_detil`.`sod_user_insert` AS `sod_user_insert`,`stock_opname_detil`.`sod_lasttime_update` AS `sod_lasttime_update`,`stock_opname_detil`.`sod_lastuser_update` AS `sod_lastuser_update`,`mt_satuan_barang`.`sb_nama` AS `sb_nama`,`stock_opname_detil`.`sod_satuan` AS `sod_satuan` from ((`stock_opname_detil` join `mt_barang` on((`stock_opname_detil`.`sod_barang` = `mt_barang`.`brg_id`))) left join `mt_satuan_barang` on((`stock_opname_detil`.`sod_satuan` = `mt_satuan_barang`.`sb_id`)));

-- 2018-05-24 07:27:59
