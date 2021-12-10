CREATE DATABASE `entry_task` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;


-- entry_task.`user` definition

CREATE TABLE `user` (
                        `uid` bigint unsigned NOT NULL AUTO_INCREMENT,
                        `username` varchar(64) DEFAULT NULL COMMENT '用户姓名',
                        `nickname` varchar(64) DEFAULT NULL COMMENT '昵称',
                        `password` char(32) NOT NULL DEFAULT '',
                        `salt` char(8) NOT NULL DEFAULT '',
                        `pic_url` varchar(256) DEFAULT NULL,
                        `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                        `update_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
                        PRIMARY KEY (`uid`),
                        KEY `idx_username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

INSERT INTO entry_task.`user` (username,nickname,password,pic_url,salt,create_time,update_time) VALUES
    ('admin','超级管理员','0b89f9065836d5e6163dc8c1d60b160e','','xLeuBm','2021-12-06 15:19:12','2021-12-09 16:13:57');
