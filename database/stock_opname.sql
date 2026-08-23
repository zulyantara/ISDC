-- Adminer 4.6.2 MySQL dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

DROP TABLE IF EXISTS `stock_opname`;
CREATE TABLE `stock_opname` (
  `so_id` int(11) NOT NULL AUTO_INCREMENT,
  `so_tgl` date NOT NULL,
  `so_gudang` int(11) NOT NULL,
  `so_bulan` smallint(6) NOT NULL,
  `so_tahun` smallint(6) NOT NULL,
  `so_sinkron` tinyint(4) NOT NULL DEFAULT '0' COMMENT '0: belum; 1: sudah',
  `so_time_insert` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `so_user_insert` int(11) NOT NULL,
  `so_lasttime_update` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `so_lastuser_update` int(11) NOT NULL,
  PRIMARY KEY (`so_id`),
  UNIQUE KEY `so_gudang_bln_thn` (`so_gudang`,`so_bulan`,`so_tahun`),
  CONSTRAINT `stock_opname_ibfk_1` FOREIGN KEY (`so_gudang`) REFERENCES `mt_area` (`area_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


DELIMITER ;;

CREATE TRIGGER `after_insert_stock_opname` AFTER INSERT ON `stock_opname` FOR EACH ROW
BEGIN

DECLARE var1 INT;
DECLARE var2 FLOAT;
DECLARE var3 INT;
DECLARE done INT DEFAULT FALSE;

DECLARE curs CURSOR FOR SELECT brg_id, stok_qty, brg_satuan_kecil FROM mt_barang LEFT JOIN stok_barang ON brg_id = stok_barang WHERE brg_aset = 0 AND stok_area = new.so_gudang;
DECLARE CONTINUE HANDLER FOR NOT FOUND SET done = TRUE;

OPEN curs;

read_loop: LOOP
	FETCH NEXT FROM curs INTO var1, var2, var3;

	IF done THEN
		LEAVE read_loop;
	END IF;

	INSERT INTO stock_opname_detil (sod_header, sod_barang, sod_qty_db, sod_satuan, sod_user_insert, sod_lastuser_update) 
VALUES (new.so_id, var1, COALESCE(var2, 0), var3, new.so_user_insert, new.so_lastuser_update);		
		
END LOOP;
		
  CLOSE curs;
END;;

CREATE TRIGGER `SYM_ON_I_FOR_STCK_PNM_TRG_CBNG` AFTER INSERT ON `stock_opname` FOR EACH ROW
begin                                                                                                                                                     
                                   
                                  if 1=1 and @sync_triggers_disabled is null then                                                                                                 
                                    insert into `db_jsdc`.sym_data (table_name, event_type, trigger_hist_id, row_data, channel_id, transaction_id, source_node_id, external_data, create_time)
                                    values(                                                                                                                                                            
                                      'stock_opname',                                                                                                                                            
                                      'I',                                                                                                                                                             
                                      34,                                                                                                                                             
                                      concat(
          if(new.`so_id` is null,'',concat('"',cast(new.`so_id` as char),'"')),',',
          if(new.`so_tgl` is null,'',concat('"',cast(new.`so_tgl` as char),'"')),',',
          if(new.`so_gudang` is null,'',concat('"',cast(new.`so_gudang` as char),'"')),',',
          if(new.`so_bulan` is null,'',concat('"',cast(new.`so_bulan` as char),'"')),',',
          if(new.`so_tahun` is null,'',concat('"',cast(new.`so_tahun` as char),'"')),',',
          if(new.`so_sinkron` is null,'',concat('"',cast(new.`so_sinkron` as char),'"')),',',
          if(new.`so_time_insert` is null,'',concat('"',cast(new.`so_time_insert` as char),'"')),',',
          if(new.`so_user_insert` is null,'',concat('"',cast(new.`so_user_insert` as char),'"')),',',
          if(new.`so_lasttime_update` is null,'',concat('"',cast(new.`so_lasttime_update` as char),'"')),',',
          if(new.`so_lastuser_update` is null,'',concat('"',cast(new.`so_lastuser_update` as char),'"'))                                                                                                                                                
                                       ),                                                                                                                                                              
                                      'sale_transaction', `db_jsdc`.sym_transaction_id_post_5_7_6(), @sync_node_disabled,                                                                                                        
                                      null,                                                                                                                                               
                                      CURRENT_TIMESTAMP                                                                                                                                                
                                    );                                                                                                                                                                 
                                  end if;                                                                                                                                                              
                                                                                                                                                                                  
                                end;;

CREATE TRIGGER `after_update_stock_opname` AFTER UPDATE ON `stock_opname` FOR EACH ROW
BEGIN

DECLARE var1 INT;
DECLARE var2 FLOAT;

DECLARE idstok INT(11) default 0;
DECLARE qtyrow INT;
DECLARE done INT DEFAULT FALSE;

DECLARE curs CURSOR FOR SELECT sod_barang, sod_qty_riil FROM stock_opname_detil WHERE sod_header = old.so_id;
DECLARE CONTINUE HANDLER FOR NOT FOUND SET done = TRUE;

IF new.so_sinkron = 1 AND old.so_sinkron = 0 THEN
	OPEN curs;

	read_loop: LOOP
		FETCH NEXT FROM curs INTO var1, var2;

		IF done THEN
      LEAVE read_loop;
    END IF;
		
		SELECT COUNT(*) INTO qtyrow FROM stok_barang WHERE stok_barang.stok_area = old.so_gudang AND stok_barang.stok_barang = var1;

		IF qtyrow > 0 THEN
			SELECT stok_barang.stok_id INTO idstok FROM stok_barang WHERE stok_barang.stok_area = old.so_gudang AND stok_barang.stok_barang = var1;

			UPDATE stok_barang SET stok_barang.stok_qty = var2 WHERE stok_barang.stok_id = idstok;
		ELSE
			INSERT INTO stok_barang (stok_barang.stok_area, stok_barang.stok_barang, stok_barang.stok_qty, stok_barang.stok_time_insert, stok_barang.stok_user_insert, stok_barang.stok_lasttime_update, stok_barang.stok_lastuser_update) 
VALUES (old.so_gudang, var1, var2, NOW(), new.so_user_insert, NOW(), new.so_lastuser_update);
		END IF;
			
  END LOOP;
		
  CLOSE curs;
END IF;
END;;

CREATE TRIGGER `SYM_ON_U_FOR_STCK_PNM_TRG_CBNG` AFTER UPDATE ON `stock_opname` FOR EACH ROW
begin                                                                                                                                                     
                                  DECLARE var_row_data mediumtext character set utf8;                                                                                                                                      
                                  DECLARE var_old_data mediumtext character set utf8;                                                                                                                                     
                                   
                                  if 1=1 and @sync_triggers_disabled is null then                                                                                                 
                                   set var_row_data = concat(
          if(new.`so_id` is null,'',concat('"',cast(new.`so_id` as char),'"')),',',
          if(new.`so_tgl` is null,'',concat('"',cast(new.`so_tgl` as char),'"')),',',
          if(new.`so_gudang` is null,'',concat('"',cast(new.`so_gudang` as char),'"')),',',
          if(new.`so_bulan` is null,'',concat('"',cast(new.`so_bulan` as char),'"')),',',
          if(new.`so_tahun` is null,'',concat('"',cast(new.`so_tahun` as char),'"')),',',
          if(new.`so_sinkron` is null,'',concat('"',cast(new.`so_sinkron` as char),'"')),',',
          if(new.`so_time_insert` is null,'',concat('"',cast(new.`so_time_insert` as char),'"')),',',
          if(new.`so_user_insert` is null,'',concat('"',cast(new.`so_user_insert` as char),'"')),',',
          if(new.`so_lasttime_update` is null,'',concat('"',cast(new.`so_lasttime_update` as char),'"')),',',
          if(new.`so_lastuser_update` is null,'',concat('"',cast(new.`so_lastuser_update` as char),'"')));                                                                                                                              
                                   set var_old_data = concat(
          if(old.`so_id` is null,'',concat('"',cast(old.`so_id` as char),'"')),',',
          if(old.`so_tgl` is null,'',concat('"',cast(old.`so_tgl` as char),'"')),',',
          if(old.`so_gudang` is null,'',concat('"',cast(old.`so_gudang` as char),'"')),',',
          if(old.`so_bulan` is null,'',concat('"',cast(old.`so_bulan` as char),'"')),',',
          if(old.`so_tahun` is null,'',concat('"',cast(old.`so_tahun` as char),'"')),',',
          if(old.`so_sinkron` is null,'',concat('"',cast(old.`so_sinkron` as char),'"')),',',
          if(old.`so_time_insert` is null,'',concat('"',cast(old.`so_time_insert` as char),'"')),',',
          if(old.`so_user_insert` is null,'',concat('"',cast(old.`so_user_insert` as char),'"')),',',
          if(old.`so_lasttime_update` is null,'',concat('"',cast(old.`so_lasttime_update` as char),'"')),',',
          if(old.`so_lastuser_update` is null,'',concat('"',cast(old.`so_lastuser_update` as char),'"')));                                                                                                                           
                                   if 1=1 then                                                                                                                                  
	                                    insert into `db_jsdc`.sym_data (table_name, event_type, trigger_hist_id, pk_data, row_data, old_data, channel_id, transaction_id, source_node_id, external_data, create_time)
	                                    values(                                                                                                                                                           
	                                      'stock_opname',                                                                                                                                           
	                                      'U',                                                                                                                                                            
	                                      34,                                                                                                                                            
	                                      concat(
          if(old.`so_id` is null,'',concat('"',cast(old.`so_id` as char),'"'))                                                                                                                                               
	                                       ),                                                                                                                                                             
	                                      var_row_data,                                                                                                                                                   
	                                      var_old_data,                                                                                                                                                   
	                                      'sale_transaction', `db_jsdc`.sym_transaction_id_post_5_7_6(), @sync_node_disabled,                                                                                                       
	                                      null,                                                                                                                                              
	                                      CURRENT_TIMESTAMP                                                                                                                                               
	                                    );                                                                                                                                                                
	                                end if;                                                                                                                                                               
                                  end if;                                                                                                                                                                
                                                                                                                                                                                    
                                end;;

CREATE TRIGGER `SYM_ON_D_FOR_STCK_PNM_TRG_CBNG` AFTER DELETE ON `stock_opname` FOR EACH ROW
begin                                                                                                                                                     
                                   
                                  if 1=1 and @sync_triggers_disabled is null then                                                                                                 
                                    insert into `db_jsdc`.sym_data (table_name, event_type, trigger_hist_id, pk_data, old_data, channel_id, transaction_id, source_node_id, external_data, create_time)
                                    values(                                                                                                                                                            
                                      'stock_opname',                                                                                                                                            
                                      'D',                                                                                                                                                             
                                      34,                                                                                                                                             
                                      concat(
          if(old.`so_id` is null,'',concat('"',cast(old.`so_id` as char),'"'))                                                                                                                                                
                                       ),                                                                                                                                                              
                                       concat(
          if(old.`so_id` is null,'',concat('"',cast(old.`so_id` as char),'"')),',',
          if(old.`so_tgl` is null,'',concat('"',cast(old.`so_tgl` as char),'"')),',',
          if(old.`so_gudang` is null,'',concat('"',cast(old.`so_gudang` as char),'"')),',',
          if(old.`so_bulan` is null,'',concat('"',cast(old.`so_bulan` as char),'"')),',',
          if(old.`so_tahun` is null,'',concat('"',cast(old.`so_tahun` as char),'"')),',',
          if(old.`so_sinkron` is null,'',concat('"',cast(old.`so_sinkron` as char),'"')),',',
          if(old.`so_time_insert` is null,'',concat('"',cast(old.`so_time_insert` as char),'"')),',',
          if(old.`so_user_insert` is null,'',concat('"',cast(old.`so_user_insert` as char),'"')),',',
          if(old.`so_lasttime_update` is null,'',concat('"',cast(old.`so_lasttime_update` as char),'"')),',',
          if(old.`so_lastuser_update` is null,'',concat('"',cast(old.`so_lastuser_update` as char),'"'))                                                                                                                                            
                                       ),                                                                                                                                                              
                                      'sale_transaction', `db_jsdc`.sym_transaction_id_post_5_7_6(), @sync_node_disabled,                                                                                                        
                                      null,                                                                                                                                               
                                      CURRENT_TIMESTAMP                                                                                                                                                
                                    );                                                                                                                                                                 
                                  end if;                                                                                                                                                              
                                                                                                                                                                                  
                                end;;

DELIMITER ;

-- 2018-05-24 07:24:21
