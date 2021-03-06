CREATE DATABASE `entry_task` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;


-- entry_task.`user` definition
CREATE TABLE `user_0000` (
     `uid` bigint unsigned NOT NULL AUTO_INCREMENT,
     `username` varchar(64) DEFAULT NULL COMMENT '用户姓名',
     `nickname` varchar(64) DEFAULT NULL COMMENT '昵称',
     `password` char(32) NOT NULL DEFAULT '',
     `pic_url` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
     `salt` char(8) NOT NULL DEFAULT '',
     `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
     `update_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
     PRIMARY KEY (`uid`),
     KEY `idx_username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

CREATE TABLE `user_0001` LIKE `user_0000`;
CREATE TABLE `user_0002` LIKE `user_0000`;
CREATE TABLE `user_0003` LIKE `user_0000`;
CREATE TABLE `user_0004` LIKE `user_0000`;
CREATE TABLE `user_0005` LIKE `user_0000`;
CREATE TABLE `user_0006` LIKE `user_0000`;
CREATE TABLE `user_0007` LIKE `user_0000`;
CREATE TABLE `user_0008` LIKE `user_0000`;
CREATE TABLE `user_0009` LIKE `user_0000`;
CREATE TABLE `user_0010` LIKE `user_0000`;
CREATE TABLE `user_0011` LIKE `user_0000`;
CREATE TABLE `user_0012` LIKE `user_0000`;
CREATE TABLE `user_0013` LIKE `user_0000`;
CREATE TABLE `user_0014` LIKE `user_0000`;
CREATE TABLE `user_0015` LIKE `user_0000`;
CREATE TABLE `user_0016` LIKE `user_0000`;
CREATE TABLE `user_0017` LIKE `user_0000`;
CREATE TABLE `user_0018` LIKE `user_0000`;
CREATE TABLE `user_0019` LIKE `user_0000`;
CREATE TABLE `user_0020` LIKE `user_0000`;
CREATE TABLE `user_0021` LIKE `user_0000`;
CREATE TABLE `user_0022` LIKE `user_0000`;
CREATE TABLE `user_0023` LIKE `user_0000`;
CREATE TABLE `user_0024` LIKE `user_0000`;
CREATE TABLE `user_0025` LIKE `user_0000`;
CREATE TABLE `user_0026` LIKE `user_0000`;
CREATE TABLE `user_0027` LIKE `user_0000`;
CREATE TABLE `user_0028` LIKE `user_0000`;
CREATE TABLE `user_0029` LIKE `user_0000`;
CREATE TABLE `user_0030` LIKE `user_0000`;
CREATE TABLE `user_0031` LIKE `user_0000`;
CREATE TABLE `user_0032` LIKE `user_0000`;
CREATE TABLE `user_0033` LIKE `user_0000`;
CREATE TABLE `user_0034` LIKE `user_0000`;
CREATE TABLE `user_0035` LIKE `user_0000`;
CREATE TABLE `user_0036` LIKE `user_0000`;
CREATE TABLE `user_0037` LIKE `user_0000`;
CREATE TABLE `user_0038` LIKE `user_0000`;
CREATE TABLE `user_0039` LIKE `user_0000`;
CREATE TABLE `user_0040` LIKE `user_0000`;
CREATE TABLE `user_0041` LIKE `user_0000`;
CREATE TABLE `user_0042` LIKE `user_0000`;
CREATE TABLE `user_0043` LIKE `user_0000`;
CREATE TABLE `user_0044` LIKE `user_0000`;
CREATE TABLE `user_0045` LIKE `user_0000`;
CREATE TABLE `user_0046` LIKE `user_0000`;
CREATE TABLE `user_0047` LIKE `user_0000`;
CREATE TABLE `user_0048` LIKE `user_0000`;
CREATE TABLE `user_0049` LIKE `user_0000`;
CREATE TABLE `user_0050` LIKE `user_0000`;
CREATE TABLE `user_0051` LIKE `user_0000`;
CREATE TABLE `user_0052` LIKE `user_0000`;
CREATE TABLE `user_0053` LIKE `user_0000`;
CREATE TABLE `user_0054` LIKE `user_0000`;
CREATE TABLE `user_0055` LIKE `user_0000`;
CREATE TABLE `user_0056` LIKE `user_0000`;
CREATE TABLE `user_0057` LIKE `user_0000`;
CREATE TABLE `user_0058` LIKE `user_0000`;
CREATE TABLE `user_0059` LIKE `user_0000`;
CREATE TABLE `user_0060` LIKE `user_0000`;
CREATE TABLE `user_0061` LIKE `user_0000`;
CREATE TABLE `user_0062` LIKE `user_0000`;
CREATE TABLE `user_0063` LIKE `user_0000`;
CREATE TABLE `user_0064` LIKE `user_0000`;
CREATE TABLE `user_0065` LIKE `user_0000`;
CREATE TABLE `user_0066` LIKE `user_0000`;
CREATE TABLE `user_0067` LIKE `user_0000`;
CREATE TABLE `user_0068` LIKE `user_0000`;
CREATE TABLE `user_0069` LIKE `user_0000`;
CREATE TABLE `user_0070` LIKE `user_0000`;
CREATE TABLE `user_0071` LIKE `user_0000`;
CREATE TABLE `user_0072` LIKE `user_0000`;
CREATE TABLE `user_0073` LIKE `user_0000`;
CREATE TABLE `user_0074` LIKE `user_0000`;
CREATE TABLE `user_0075` LIKE `user_0000`;
CREATE TABLE `user_0076` LIKE `user_0000`;
CREATE TABLE `user_0077` LIKE `user_0000`;
CREATE TABLE `user_0078` LIKE `user_0000`;
CREATE TABLE `user_0079` LIKE `user_0000`;
CREATE TABLE `user_0080` LIKE `user_0000`;
CREATE TABLE `user_0081` LIKE `user_0000`;
CREATE TABLE `user_0082` LIKE `user_0000`;
CREATE TABLE `user_0083` LIKE `user_0000`;
CREATE TABLE `user_0084` LIKE `user_0000`;
CREATE TABLE `user_0085` LIKE `user_0000`;
CREATE TABLE `user_0086` LIKE `user_0000`;
CREATE TABLE `user_0087` LIKE `user_0000`;
CREATE TABLE `user_0088` LIKE `user_0000`;
CREATE TABLE `user_0089` LIKE `user_0000`;
CREATE TABLE `user_0090` LIKE `user_0000`;
CREATE TABLE `user_0091` LIKE `user_0000`;
CREATE TABLE `user_0092` LIKE `user_0000`;
CREATE TABLE `user_0093` LIKE `user_0000`;
CREATE TABLE `user_0094` LIKE `user_0000`;
CREATE TABLE `user_0095` LIKE `user_0000`;
CREATE TABLE `user_0096` LIKE `user_0000`;
CREATE TABLE `user_0097` LIKE `user_0000`;
CREATE TABLE `user_0098` LIKE `user_0000`;
CREATE TABLE `user_0099` LIKE `user_0000`;
CREATE TABLE `user_0100` LIKE `user_0000`;
CREATE TABLE `user_0101` LIKE `user_0000`;
CREATE TABLE `user_0102` LIKE `user_0000`;
CREATE TABLE `user_0103` LIKE `user_0000`;
CREATE TABLE `user_0104` LIKE `user_0000`;
CREATE TABLE `user_0105` LIKE `user_0000`;
CREATE TABLE `user_0106` LIKE `user_0000`;
CREATE TABLE `user_0107` LIKE `user_0000`;
CREATE TABLE `user_0108` LIKE `user_0000`;
CREATE TABLE `user_0109` LIKE `user_0000`;
CREATE TABLE `user_0110` LIKE `user_0000`;
CREATE TABLE `user_0111` LIKE `user_0000`;
CREATE TABLE `user_0112` LIKE `user_0000`;
CREATE TABLE `user_0113` LIKE `user_0000`;
CREATE TABLE `user_0114` LIKE `user_0000`;
CREATE TABLE `user_0115` LIKE `user_0000`;
CREATE TABLE `user_0116` LIKE `user_0000`;
CREATE TABLE `user_0117` LIKE `user_0000`;
CREATE TABLE `user_0118` LIKE `user_0000`;
CREATE TABLE `user_0119` LIKE `user_0000`;
CREATE TABLE `user_0120` LIKE `user_0000`;
CREATE TABLE `user_0121` LIKE `user_0000`;
CREATE TABLE `user_0122` LIKE `user_0000`;
CREATE TABLE `user_0123` LIKE `user_0000`;
CREATE TABLE `user_0124` LIKE `user_0000`;
CREATE TABLE `user_0125` LIKE `user_0000`;
CREATE TABLE `user_0126` LIKE `user_0000`;
CREATE TABLE `user_0127` LIKE `user_0000`;


INSERT INTO entry_task.`user_0118` (username,nickname,password,pic_url,salt,create_time,update_time) VALUES
    ('admin','超级管理员','0b89f9065836d5e6163dc8c1d60b160e','','xLeuBm','2021-12-06 15:19:12','2021-12-09 16:13:57');
