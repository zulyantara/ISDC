-- Adminer 4.6.2 MySQL dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

DROP TABLE IF EXISTS `perawatan_barang`;
CREATE TABLE `perawatan_barang` (
  `rb_id` int(11) NOT NULL AUTO_INCREMENT,
  `rb_tgl` date NOT NULL,
  `rb_barang` int(11) NOT NULL,
  `rb_gudang` int(11) NOT NULL,
  `rb_jenis` varchar(255) NOT NULL,
  `rb_biaya` float NOT NULL,
  `rb_time_insert` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `rb_user_insert` int(11) NOT NULL,
  `rb_lasttime_update` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `rb_lastuser_update` int(11) NOT NULL,
  PRIMARY KEY (`rb_id`),
  KEY `rb_barang` (`rb_barang`),
  KEY `rb_gudang` (`rb_gudang`),
  CONSTRAINT `perawatan_barang_ibfk_1` FOREIGN KEY (`rb_barang`) REFERENCES `mt_barang` (`brg_id`),
  CONSTRAINT `perawatan_barang_ibfk_2` FOREIGN KEY (`rb_gudang`) REFERENCES `mt_area` (`area_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO `perawatan_barang` (`rb_id`, `rb_tgl`, `rb_barang`, `rb_gudang`, `rb_jenis`, `rb_biaya`, `rb_time_insert`, `rb_user_insert`, `rb_lasttime_update`, `rb_lastuser_update`) VALUES
(5,	'2018-04-12',	7,	2,	'Service sekaligus ganti tinta',	990888,	'2018-04-12 05:52:49',	3,	'2018-04-12 06:16:49',	3);

DELIMITER ;;

CREATE TRIGGER `SYM_ON_I_FOR_PRWTN_BRNG_TRG_CBNG` AFTER INSERT ON `perawatan_barang` FOR EACH ROW
begin                                                                                                                                                     
                                   
                                  if 1=1 and @sync_triggers_disabled is null then                                                                                                 
                                    insert into `db_jsdc`.sym_data (table_name, event_type, trigger_hist_id, row_data, channel_id, transaction_id, source_node_id, external_data, create_time)
                                    values(                                                                                                                                                            
                                      'perawatan_barang',                                                                                                                                            
                                      'I',                                                                                                                                                             
                                      15,                                                                                                                                             
                                      concat(
          if(new.`rb_id` is null,'',concat('"',cast(new.`rb_id` as char),'"')),',',
          if(new.`rb_tgl` is null,'',concat('"',cast(new.`rb_tgl` as char),'"')),',',
          if(new.`rb_barang` is null,'',concat('"',cast(new.`rb_barang` as char),'"')),',',
          if(new.`rb_gudang` is null,'',concat('"',cast(new.`rb_gudang` as char),'"')),',',
          cast(if(new.`rb_jenis` is null,'',concat('"',replace(replace(new.`rb_jenis`,'\','\\'),'"','\"'),'"')) as char),',',
          if(new.`rb_biaya` is null,'',concat('"',cast(new.`rb_biaya` as char),'"')),',',
          if(new.`rb_time_insert` is null,'',concat('"',cast(new.`rb_time_insert` as char),'"')),',',
          if(new.`rb_user_insert` is null,'',concat('"',cast(new.`rb_user_insert` as char),'"')),',',
          if(new.`rb_lasttime_update` is null,'',concat('"',cast(new.`rb_lasttime_update` as char),'"')),',',
          if(new.`rb_lastuser_update` is null,'',concat('"',cast(new.`rb_lastuser_update` as char),'"'))                                                                                                                                                
                                       ),                                                                                                                                                              
                                      'sale_transaction', `db_jsdc`.sym_transaction_id_post_5_7_6(), @sync_node_disabled,                                                                                                        
                                      null,                                                                                                                                               
                                      CURRENT_TIMESTAMP                                                                                                                                                
                                    );                                                                                                                                                                 
                                  end if;                                                                                                                                                              
                                                                                                                                                                                  
                                end;;

CREATE TRIGGER `SYM_ON_U_FOR_PRWTN_BRNG_TRG_CBNG` AFTER UPDATE ON `perawatan_barang` FOR EACH ROW
begin                                                                                                                                                     
                                  DECLARE var_row_data mediumtext character set utf8;                                                                                                                                      
                                  DECLARE var_old_data mediumtext character set utf8;                                                                                                                                     
                                   
                                  if 1=1 and @sync_triggers_disabled is null then                                                                                                 
                                   set var_row_data = concat(
          if(new.`rb_id` is null,'',concat('"',cast(new.`rb_id` as char),'"')),',',
          if(new.`rb_tgl` is null,'',concat('"',cast(new.`rb_tgl` as char),'"')),',',
          if(new.`rb_barang` is null,'',concat('"',cast(new.`rb_barang` as char),'"')),',',
          if(new.`rb_gudang` is null,'',concat('"',cast(new.`rb_gudang` as char),'"')),',',
          cast(if(new.`rb_jenis` is null,'',concat('"',replace(replace(new.`rb_jenis`,'\','\\'),'"','\"'),'"')) as char),',',
          if(new.`rb_biaya` is null,'',concat('"',cast(new.`rb_biaya` as char),'"')),',',
          if(new.`rb_time_insert` is null,'',concat('"',cast(new.`rb_time_insert` as char),'"')),',',
          if(new.`rb_user_insert` is null,'',concat('"',cast(new.`rb_user_insert` as char),'"')),',',
          if(new.`rb_lasttime_update` is null,'',concat('"',cast(new.`rb_lasttime_update` as char),'"')),',',
          if(new.`rb_lastuser_update` is null,'',concat('"',cast(new.`rb_lastuser_update` as char),'"')));                                                                                                                              
                                   set var_old_data = concat(
          if(old.`rb_id` is null,'',concat('"',cast(old.`rb_id` as char),'"')),',',
          if(old.`rb_tgl` is null,'',concat('"',cast(old.`rb_tgl` as char),'"')),',',
          if(old.`rb_barang` is null,'',concat('"',cast(old.`rb_barang` as char),'"')),',',
          if(old.`rb_gudang` is null,'',concat('"',cast(old.`rb_gudang` as char),'"')),',',
          cast(if(old.`rb_jenis` is null,'',concat('"',replace(replace(old.`rb_jenis`,'\','\\'),'"','\"'),'"')) as char),',',
          if(old.`rb_biaya` is null,'',concat('"',cast(old.`rb_biaya` as char),'"')),',',
          if(old.`rb_time_insert` is null,'',concat('"',cast(old.`rb_time_insert` as char),'"')),',',
          if(old.`rb_user_insert` is null,'',concat('"',cast(old.`rb_user_insert` as char),'"')),',',
          if(old.`rb_lasttime_update` is null,'',concat('"',cast(old.`rb_lasttime_update` as char),'"')),',',
          if(old.`rb_lastuser_update` is null,'',concat('"',cast(old.`rb_lastuser_update` as char),'"')));                                                                                                                           
                                   if 1=1 then                                                                                                                                  
	                                    insert into `db_jsdc`.sym_data (table_name, event_type, trigger_hist_id, pk_data, row_data, old_data, channel_id, transaction_id, source_node_id, external_data, create_time)
	                                    values(                                                                                                                                                           
	                                      'perawatan_barang',                                                                                                                                           
	                                      'U',                                                                                                                                                            
	                                      15,                                                                                                                                            
	                                      concat(
          if(old.`rb_id` is null,'',concat('"',cast(old.`rb_id` as char),'"'))                                                                                                                                               
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

CREATE TRIGGER `SYM_ON_D_FOR_PRWTN_BRNG_TRG_CBNG` AFTER DELETE ON `perawatan_barang` FOR EACH ROW
begin                                                                                                                                                     
                                   
                                  if 1=1 and @sync_triggers_disabled is null then                                                                                                 
                                    insert into `db_jsdc`.sym_data (table_name, event_type, trigger_hist_id, pk_data, old_data, channel_id, transaction_id, source_node_id, external_data, create_time)
                                    values(                                                                                                                                                            
                                      'perawatan_barang',                                                                                                                                            
                                      'D',                                                                                                                                                             
                                      15,                                                                                                                                             
                                      concat(
          if(old.`rb_id` is null,'',concat('"',cast(old.`rb_id` as char),'"'))                                                                                                                                                
                                       ),                                                                                                                                                              
                                       concat(
          if(old.`rb_id` is null,'',concat('"',cast(old.`rb_id` as char),'"')),',',
          if(old.`rb_tgl` is null,'',concat('"',cast(old.`rb_tgl` as char),'"')),',',
          if(old.`rb_barang` is null,'',concat('"',cast(old.`rb_barang` as char),'"')),',',
          if(old.`rb_gudang` is null,'',concat('"',cast(old.`rb_gudang` as char),'"')),',',
          cast(if(old.`rb_jenis` is null,'',concat('"',replace(replace(old.`rb_jenis`,'\','\\'),'"','\"'),'"')) as char),',',
          if(old.`rb_biaya` is null,'',concat('"',cast(old.`rb_biaya` as char),'"')),',',
          if(old.`rb_time_insert` is null,'',concat('"',cast(old.`rb_time_insert` as char),'"')),',',
          if(old.`rb_user_insert` is null,'',concat('"',cast(old.`rb_user_insert` as char),'"')),',',
          if(old.`rb_lasttime_update` is null,'',concat('"',cast(old.`rb_lasttime_update` as char),'"')),',',
          if(old.`rb_lastuser_update` is null,'',concat('"',cast(old.`rb_lastuser_update` as char),'"'))                                                                                                                                            
                                       ),                                                                                                                                                              
                                      'sale_transaction', `db_jsdc`.sym_transaction_id_post_5_7_6(), @sync_node_disabled,                                                                                                        
                                      null,                                                                                                                                               
                                      CURRENT_TIMESTAMP                                                                                                                                                
                                    );                                                                                                                                                                 
                                  end if;                                                                                                                                                              
                                                                                                                                                                                  
                                end;;

DELIMITER ;

-- 2018-05-24 07:22:27
