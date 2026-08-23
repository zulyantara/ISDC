-- Adminer 4.6.2 MySQL dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

DROP TABLE IF EXISTS `stock_opname_detil`;
CREATE TABLE `stock_opname_detil` (
  `sod_id` int(11) NOT NULL AUTO_INCREMENT,
  `sod_header` int(11) NOT NULL,
  `sod_barang` int(11) NOT NULL,
  `sod_qty_db` float NOT NULL,
  `sod_qty_riil` float DEFAULT NULL,
  `sod_satuan` int(11) DEFAULT NULL,
  `sod_time_insert` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `sod_user_insert` int(11) NOT NULL,
  `sod_lasttime_update` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `sod_lastuser_update` int(11) NOT NULL,
  PRIMARY KEY (`sod_id`),
  KEY `sod_header` (`sod_header`),
  KEY `sod_barang` (`sod_barang`),
  KEY `sod_satuan` (`sod_satuan`),
  CONSTRAINT `stock_opname_detil_ibfk_1` FOREIGN KEY (`sod_header`) REFERENCES `stock_opname` (`so_id`),
  CONSTRAINT `stock_opname_detil_ibfk_2` FOREIGN KEY (`sod_barang`) REFERENCES `mt_barang` (`brg_id`),
  CONSTRAINT `stock_opname_detil_ibfk_3` FOREIGN KEY (`sod_satuan`) REFERENCES `mt_satuan_barang` (`sb_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


DELIMITER ;;

CREATE TRIGGER `SYM_ON_I_FOR_STCK_PNM_DTL_TRG_CBNG` AFTER INSERT ON `stock_opname_detil` FOR EACH ROW
begin                                                                                                                                                     
                                   
                                  if 1=1 and @sync_triggers_disabled is null then                                                                                                 
                                    insert into `db_jsdc`.sym_data (table_name, event_type, trigger_hist_id, row_data, channel_id, transaction_id, source_node_id, external_data, create_time)
                                    values(                                                                                                                                                            
                                      'stock_opname_detil',                                                                                                                                            
                                      'I',                                                                                                                                                             
                                      5,                                                                                                                                             
                                      concat(
          if(new.`sod_id` is null,'',concat('"',cast(new.`sod_id` as char),'"')),',',
          if(new.`sod_header` is null,'',concat('"',cast(new.`sod_header` as char),'"')),',',
          if(new.`sod_barang` is null,'',concat('"',cast(new.`sod_barang` as char),'"')),',',
          if(new.`sod_qty_db` is null,'',concat('"',cast(new.`sod_qty_db` as char),'"')),',',
          if(new.`sod_qty_riil` is null,'',concat('"',cast(new.`sod_qty_riil` as char),'"')),',',
          if(new.`sod_satuan` is null,'',concat('"',cast(new.`sod_satuan` as char),'"')),',',
          if(new.`sod_time_insert` is null,'',concat('"',cast(new.`sod_time_insert` as char),'"')),',',
          if(new.`sod_user_insert` is null,'',concat('"',cast(new.`sod_user_insert` as char),'"')),',',
          if(new.`sod_lasttime_update` is null,'',concat('"',cast(new.`sod_lasttime_update` as char),'"')),',',
          if(new.`sod_lastuser_update` is null,'',concat('"',cast(new.`sod_lastuser_update` as char),'"'))                                                                                                                                                
                                       ),                                                                                                                                                              
                                      'sale_transaction', `db_jsdc`.sym_transaction_id_post_5_7_6(), @sync_node_disabled,                                                                                                        
                                      null,                                                                                                                                               
                                      CURRENT_TIMESTAMP                                                                                                                                                
                                    );                                                                                                                                                                 
                                  end if;                                                                                                                                                              
                                                                                                                                                                                  
                                end;;

CREATE TRIGGER `SYM_ON_U_FOR_STCK_PNM_DTL_TRG_CBNG` AFTER UPDATE ON `stock_opname_detil` FOR EACH ROW
begin                                                                                                                                                     
                                  DECLARE var_row_data mediumtext character set utf8;                                                                                                                                      
                                  DECLARE var_old_data mediumtext character set utf8;                                                                                                                                     
                                   
                                  if 1=1 and @sync_triggers_disabled is null then                                                                                                 
                                   set var_row_data = concat(
          if(new.`sod_id` is null,'',concat('"',cast(new.`sod_id` as char),'"')),',',
          if(new.`sod_header` is null,'',concat('"',cast(new.`sod_header` as char),'"')),',',
          if(new.`sod_barang` is null,'',concat('"',cast(new.`sod_barang` as char),'"')),',',
          if(new.`sod_qty_db` is null,'',concat('"',cast(new.`sod_qty_db` as char),'"')),',',
          if(new.`sod_qty_riil` is null,'',concat('"',cast(new.`sod_qty_riil` as char),'"')),',',
          if(new.`sod_satuan` is null,'',concat('"',cast(new.`sod_satuan` as char),'"')),',',
          if(new.`sod_time_insert` is null,'',concat('"',cast(new.`sod_time_insert` as char),'"')),',',
          if(new.`sod_user_insert` is null,'',concat('"',cast(new.`sod_user_insert` as char),'"')),',',
          if(new.`sod_lasttime_update` is null,'',concat('"',cast(new.`sod_lasttime_update` as char),'"')),',',
          if(new.`sod_lastuser_update` is null,'',concat('"',cast(new.`sod_lastuser_update` as char),'"')));                                                                                                                              
                                   set var_old_data = concat(
          if(old.`sod_id` is null,'',concat('"',cast(old.`sod_id` as char),'"')),',',
          if(old.`sod_header` is null,'',concat('"',cast(old.`sod_header` as char),'"')),',',
          if(old.`sod_barang` is null,'',concat('"',cast(old.`sod_barang` as char),'"')),',',
          if(old.`sod_qty_db` is null,'',concat('"',cast(old.`sod_qty_db` as char),'"')),',',
          if(old.`sod_qty_riil` is null,'',concat('"',cast(old.`sod_qty_riil` as char),'"')),',',
          if(old.`sod_satuan` is null,'',concat('"',cast(old.`sod_satuan` as char),'"')),',',
          if(old.`sod_time_insert` is null,'',concat('"',cast(old.`sod_time_insert` as char),'"')),',',
          if(old.`sod_user_insert` is null,'',concat('"',cast(old.`sod_user_insert` as char),'"')),',',
          if(old.`sod_lasttime_update` is null,'',concat('"',cast(old.`sod_lasttime_update` as char),'"')),',',
          if(old.`sod_lastuser_update` is null,'',concat('"',cast(old.`sod_lastuser_update` as char),'"')));                                                                                                                           
                                   if 1=1 then                                                                                                                                  
	                                    insert into `db_jsdc`.sym_data (table_name, event_type, trigger_hist_id, pk_data, row_data, old_data, channel_id, transaction_id, source_node_id, external_data, create_time)
	                                    values(                                                                                                                                                           
	                                      'stock_opname_detil',                                                                                                                                           
	                                      'U',                                                                                                                                                            
	                                      5,                                                                                                                                            
	                                      concat(
          if(old.`sod_id` is null,'',concat('"',cast(old.`sod_id` as char),'"'))                                                                                                                                               
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

CREATE TRIGGER `SYM_ON_D_FOR_STCK_PNM_DTL_TRG_CBNG` AFTER DELETE ON `stock_opname_detil` FOR EACH ROW
begin                                                                                                                                                     
                                   
                                  if 1=1 and @sync_triggers_disabled is null then                                                                                                 
                                    insert into `db_jsdc`.sym_data (table_name, event_type, trigger_hist_id, pk_data, old_data, channel_id, transaction_id, source_node_id, external_data, create_time)
                                    values(                                                                                                                                                            
                                      'stock_opname_detil',                                                                                                                                            
                                      'D',                                                                                                                                                             
                                      5,                                                                                                                                             
                                      concat(
          if(old.`sod_id` is null,'',concat('"',cast(old.`sod_id` as char),'"'))                                                                                                                                                
                                       ),                                                                                                                                                              
                                       concat(
          if(old.`sod_id` is null,'',concat('"',cast(old.`sod_id` as char),'"')),',',
          if(old.`sod_header` is null,'',concat('"',cast(old.`sod_header` as char),'"')),',',
          if(old.`sod_barang` is null,'',concat('"',cast(old.`sod_barang` as char),'"')),',',
          if(old.`sod_qty_db` is null,'',concat('"',cast(old.`sod_qty_db` as char),'"')),',',
          if(old.`sod_qty_riil` is null,'',concat('"',cast(old.`sod_qty_riil` as char),'"')),',',
          if(old.`sod_satuan` is null,'',concat('"',cast(old.`sod_satuan` as char),'"')),',',
          if(old.`sod_time_insert` is null,'',concat('"',cast(old.`sod_time_insert` as char),'"')),',',
          if(old.`sod_user_insert` is null,'',concat('"',cast(old.`sod_user_insert` as char),'"')),',',
          if(old.`sod_lasttime_update` is null,'',concat('"',cast(old.`sod_lasttime_update` as char),'"')),',',
          if(old.`sod_lastuser_update` is null,'',concat('"',cast(old.`sod_lastuser_update` as char),'"'))                                                                                                                                            
                                       ),                                                                                                                                                              
                                      'sale_transaction', `db_jsdc`.sym_transaction_id_post_5_7_6(), @sync_node_disabled,                                                                                                        
                                      null,                                                                                                                                               
                                      CURRENT_TIMESTAMP                                                                                                                                                
                                    );                                                                                                                                                                 
                                  end if;                                                                                                                                                              
                                                                                                                                                                                  
                                end;;

DELIMITER ;

-- 2018-05-24 07:24:29
