-- Adminer 4.6.2 MySQL dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

DELIMITER ;;

DROP FUNCTION IF EXISTS `fromRoman`;;
CREATE DEFINER=`root`@`localhost` FUNCTION `fromRoman`(`inRoman` VARCHAR(15)) RETURNS int(11)
BEGIN

    DECLARE numeral CHAR(7) DEFAULT 'IVXLCDM';

    DECLARE digit TINYINT;
    DECLARE previous INT DEFAULT 0;
    DECLARE current INT;
    DECLARE sum INT DEFAULT 0;

    SET inRoman = UPPER(inRoman);

    WHILE LENGTH(inRoman) > 0 DO
        SET digit := LOCATE(RIGHT(inRoman, 1), numeral) - 1;
        SET current := POW(10, FLOOR(digit / 2)) * POW(5, MOD(digit, 2));
        SET sum := sum + POW(-1, current < previous) * current;
        SET previous := current;
        SET inRoman = LEFT(inRoman, LENGTH(inRoman) - 1);
    END WHILE;

    RETURN sum;
END;;

DROP FUNCTION IF EXISTS `getBulanTeks`;;
CREATE DEFINER=`root`@`localhost` FUNCTION `getBulanTeks`(`x` INT) RETURNS varchar(5000) CHARSET latin1
BEGIN
	
	DECLARE bulan TEXT DEFAULT '';

	CASE CAST(bulan AS UNSIGNED)
    WHEN 1 THEN SET bulan = 'Januari';
    WHEN 2 THEN SET bulan = 'Pebruari';
    WHEN 3 THEN SET bulan = 'Maret';
    WHEN 4 THEN SET bulan = 'April';
    WHEN 5 THEN SET bulan = "Mei";
    WHEN 6 THEN SET bulan = "Juni";
    WHEN 7 THEN SET bulan = "Juli";
    WHEN 8 THEN SET bulan = "Agustus";
    WHEN 9 THEN SET bulan = "September";
    WHEN 10 THEN SET bulan = "Oktober";
    WHEN 11 THEN SET bulan = "Nopember";
    WHEN 12 THEN SET bulan = "Desember";
   ELSE BEGIN END;
   END CASE;
 
	RETURN 	bulan;
END;;

DROP FUNCTION IF EXISTS `sym_transaction_id_post_5_7_6`;;
CREATE DEFINER=`root`@`localhost` FUNCTION `sym_transaction_id_post_5_7_6`() RETURNS varchar(50) CHARSET latin1
    READS SQL DATA
begin                                                                                                                           
    declare done int default 0;                                                                                                  
    declare comm_value varchar(50);                                                                                              
    declare comm_cur cursor for select TRX_ID from INFORMATION_SCHEMA.INNODB_TRX where TRX_MYSQL_THREAD_ID = CONNECTION_ID();    
    declare continue handler for not found set done = 1;                                                                         
    open comm_cur;                                                                                                               
    fetch comm_cur into comm_value;                                                                                              
    close comm_cur;                                                                                                              
    return concat(concat(connection_id(), '.'), comm_value);                                                                     
 end;;

DROP FUNCTION IF EXISTS `toHari`;;
CREATE DEFINER=`root`@`localhost` FUNCTION `toHari`(`tgl` DATE) RETURNS varchar(5000) CHARSET latin1
BEGIN
	
	DECLARE hari TEXT DEFAULT '';
	DECLARE dow INT;
	
	SELECT DATE_FORMAT(tgl,'%w') INTO dow;

	CASE CAST(dow AS UNSIGNED)
    WHEN 0 THEN SET hari = 'Minggu';
    WHEN 1 THEN SET hari = 'Senin';
    WHEN 2 THEN SET hari = 'Selasa';
    WHEN 3 THEN SET hari = 'Rabu';
    WHEN 4 THEN SET hari = 'Kamis';
    WHEN 5 THEN SET hari = "Jum'at";
    WHEN 6 THEN SET hari = 'Sabtu';
   ELSE BEGIN END;
   END CASE;
 
	RETURN 	hari;
END;;

DROP FUNCTION IF EXISTS `toRoman`;;
CREATE DEFINER=`root`@`localhost` FUNCTION `toRoman`(`number` INT) RETURNS text CHARSET latin1
BEGIN
 
	DECLARE basic_roman TEXT DEFAULT 'M,CM,D,CD,C,XC,L,XL,X,IX,V,IV,I';
	DECLARE basic_value TEXT DEFAULT '1000,900,500,400,100,90,50,40,10,9,5,4,1';
	DECLARE roman_string TEXT DEFAULT '';
	DECLARE i INT DEFAULT 1;
	DECLARE roman_symbol TEXT;
	DECLARE roman_value INT;
 
	SET roman_string = (SELECT IF(number<=0 OR number>=4000,NULL,roman_string));
 
	WHILE number > 0 AND number < 4000 DO
 
		SET roman_symbol = SUBSTRING_INDEX(SUBSTRING_INDEX(basic_roman,',',i),',',-1);
		SET roman_value  = SUBSTRING_INDEX(SUBSTRING_INDEX(basic_value,',',i),',',-1);
 
		IF number >= roman_value THEN
			SET roman_string = CONCAT(roman_string,roman_symbol);
			SET number = number - roman_value;
		ELSE
			SET i = i + 1;
		END IF;
 
	END WHILE;
 
	RETURN 	roman_string;
 
END;;

DROP FUNCTION IF EXISTS `xf_terbilang`;;
CREATE DEFINER=`root`@`localhost` FUNCTION `xf_terbilang`(`angka` BIGINT) RETURNS varchar(5000) CHARSET utf8
BEGIN
	
	DECLARE sString varchar(30);
 DECLARE Bil1 varchar(255);
 DECLARE Bil2 varchar(255);
 DECLARE STot varchar(255);
 DECLARE X int;
 DECLARE Y int;
 DECLARE Z int;
 DECLARE Urai varchar(5000);
 SET sString = CAST(angka as char);
 SET Urai = '';
 SET X = 0;
 SET Y = 0;
 WHILE X <>  LENGTH(sString) DO
SET X = X + 1;
SET sTot = MID(sString, X, 1);
SET Y = Y + CAST(sTot as UNSIGNED);
SET Z = LENGTH(sString) - X + 1;
CASE CAST(sTot as UNSIGNED)
WHEN 1 THEN
 BEGIN
  IF (Z = 1 OR Z = 7 OR Z = 10 OR Z = 13) THEN
   SET Bil1 = 'SATU ';
  ELSEIF (z = 4) THEN
   IF (x = 1) THEN
    SET Bil1 = 'SE';
   ELSE
    SET Bil1 = 'SATU';
   END IF;
  ELSEIF (Z = 2 OR Z = 5 OR Z = 8 OR Z = 11 OR Z = 14) THEN
   SET X = X + 1;
   SET sTot = MID(sString, X, 1);
   SET Z = LENGTH(sString) - X + 1;
   SET Bil2 = '';
   CASE CAST(sTot AS UNSIGNED)
    WHEN 0 THEN SET Bil1 = 'SEPULUH ';
    WHEN 1 THEN SET Bil1 = 'SEBELAS ';
    WHEN 2 THEN SET Bil1 = 'DUA BELAS ';
    WHEN 3 THEN SET Bil1 = 'TIGA BELAS ';
    WHEN 4 THEN SET Bil1 = 'EMPAT BELAS ';
    WHEN 5 THEN SET Bil1 = 'LIMA BELAS ';
    WHEN 6 THEN SET Bil1 = 'ENAM BELAS ';
    WHEN 7 THEN SET Bil1 = 'TUJUH BELAS ';
    WHEN 8 THEN SET Bil1 = 'DELAPAN BELAS ';
    WHEN 9 THEN SET Bil1 = 'SEMBILAN BELAS ';
   ELSE BEGIN END;
   END CASE;
  ELSE
   SET Bil1 = 'SE';
  END IF;
 END;
WHEN 2 THEN SET Bil1 = 'DUA ';
WHEN 3 THEN SET Bil1 = 'TIGA ';
WHEN 4 THEN SET Bil1 = 'EMPAT ';
WHEN 5 THEN SET Bil1 = 'LIMA ';
WHEN 6 THEN SET Bil1 = 'ENAM ';
WHEN 7 THEN SET Bil1 = 'TUJUH ';
WHEN 8 THEN SET Bil1 = 'DELAPAN ';
WHEN 9 THEN SET Bil1 = 'SEMBILAN ';
ELSE SET Bil1 = '';
END CASE;
IF CAST(sTot as UNSIGNED) > 0 THEN
IF (Z = 2 OR Z = 5 OR Z = 8 OR Z = 11 OR Z = 14) THEN
 SET Bil2 = 'PULUH ';
ELSEIF (Z = 3 OR Z = 6 OR Z = 9 OR Z = 12 OR Z = 15) THEN
 SET Bil2 = 'RATUS ';
ELSE
 SET Bil2 = '';
END IF;
ELSE
SET Bil2 = '';
END IF;
IF Y > 0 THEN
CASE Z
 WHEN 4 THEN BEGIN SET Bil2 = CONCAT(Bil2, 'RIBU '); SET Y = 0; END;
 WHEN 7 THEN BEGIN SET Bil2 = CONCAT(Bil2, 'JUTA '); SET Y = 0; END;
 WHEN 10 THEN BEGIN SET Bil2 = CONCAT(Bil2, 'MILYAR '); SET Y = 0; END;
 WHEN 13 THEN BEGIN SET Bil2 = CONCAT(Bil2, 'TRILYUN '); SET Y = 0; END;
 ELSE BEGIN END;
END CASE;
END IF;
SET Urai = CONCAT(Urai, Bil1, Bil2);
END WHILE;
RETURN Urai;
END;;

DELIMITER ;

DROP VIEW IF EXISTS `v_stok_barang`;
CREATE TABLE `v_stok_barang` (`area_id` int(11), `area_kode` varchar(2), `area_nama` varchar(100), `brg_id` int(11), `brg_kode` varchar(255), `brg_nama` varchar(255), `sb_nama` varchar(255), `stok_qty` float, `stok_limit` float, `stok_id` int(11), `stok_time_insert` timestamp, `stok_lasttime_update` timestamp, `stok_user_insert` int(11), `stok_lastuser_update` int(11));


DROP TABLE IF EXISTS `v_stok_barang`;
CREATE ALGORITHM=UNDEFINED SQL SECURITY DEFINER VIEW `v_stok_barang` AS select `t0`.`area_id` AS `area_id`,`t0`.`area_kode` AS `area_kode`,`t0`.`area_nama` AS `area_nama`,`t0`.`brg_id` AS `brg_id`,`t0`.`brg_kode` AS `brg_kode`,`t0`.`brg_nama` AS `brg_nama`,`t0`.`sb_nama` AS `sb_nama`,`isdc_db`.`stok_barang`.`stok_qty` AS `stok_qty`,`isdc_db`.`stok_barang`.`stok_limit` AS `stok_limit`,`isdc_db`.`stok_barang`.`stok_id` AS `stok_id`,`isdc_db`.`stok_barang`.`stok_time_insert` AS `stok_time_insert`,`isdc_db`.`stok_barang`.`stok_lasttime_update` AS `stok_lasttime_update`,`isdc_db`.`stok_barang`.`stok_user_insert` AS `stok_user_insert`,`isdc_db`.`stok_barang`.`stok_lastuser_update` AS `stok_lastuser_update` from (((select `isdc_db`.`mt_area`.`area_id` AS `area_id`,`isdc_db`.`mt_area`.`area_kode` AS `area_kode`,`isdc_db`.`mt_area`.`area_nama` AS `area_nama`,`isdc_db`.`mt_barang`.`brg_id` AS `brg_id`,`isdc_db`.`mt_barang`.`brg_kode` AS `brg_kode`,`isdc_db`.`mt_barang`.`brg_nama` AS `brg_nama`,`isdc_db`.`mt_satuan_barang`.`sb_nama` AS `sb_nama` from (`isdc_db`.`mt_area` join (`isdc_db`.`mt_barang` left join `isdc_db`.`mt_satuan_barang` on((`isdc_db`.`mt_satuan_barang`.`sb_id` = `isdc_db`.`mt_barang`.`brg_satuan_kecil`)))) where (`isdc_db`.`mt_barang`.`brg_aset` <> 1))) `t0` left join `isdc_db`.`stok_barang` on(((`isdc_db`.`stok_barang`.`stok_area` = `t0`.`area_id`) and (`isdc_db`.`stok_barang`.`stok_barang` = `t0`.`brg_id`))));

-- 2018-05-24 07:38:10
