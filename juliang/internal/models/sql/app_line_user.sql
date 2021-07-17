-- MySQL dump 10.13  Distrib 5.6.50, for Linux (x86_64)
--
-- Host: localhost    Database: app_line_user
-- ------------------------------------------------------
-- Server version	5.6.50-log

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `service`
--

DROP TABLE IF EXISTS `service`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `service` (
  `id` bigint(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(128) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '服务器名',
  `path` varchar(128) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '游戏地址(服务器端请求地址)',
  `type` int(11) DEFAULT '0' COMMENT '区服类型(0:普通 1:新开 2:推荐 3:新开推荐)',
  `status` int(11) DEFAULT '0' COMMENT '服务器状态(0:关服 1:顺畅 2:拥挤 3:爆满,4:维护)',
  `channel` varchar(128) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '渠道来源',
  `zoneno` int(11) DEFAULT NULL COMMENT '区服编号(x区)',
  `times` int(11) DEFAULT '0' COMMENT '开服时间戳(秒)',
  `normal` int(11) DEFAULT '0' COMMENT '正常人数(区服玩家数)',
  `current` int(11) DEFAULT '0' COMMENT '当前人数',
  `signkey` varchar(128) COLLATE utf8_unicode_ci DEFAULT '' COMMENT '签名key值',
  `version` varchar(45) COLLATE utf8_unicode_ci DEFAULT '' COMMENT '版本号',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `source_juliang`
--

DROP TABLE IF EXISTS `source_juliang`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `source_juliang` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `open_id` varchar(128) COLLATE utf8_bin DEFAULT NULL COMMENT '微信openId',
  `oaid` varchar(64) COLLATE utf8_bin DEFAULT NULL COMMENT 'Android Q版本的oaid原值',
  `imei_md5` varchar(64) COLLATE utf8_bin NOT NULL COMMENT '安卓系统imei的md5摘要',
  `aid` varchar(64) COLLATE utf8_bin DEFAULT NULL COMMENT '广告计划id',
  `os` varchar(32) COLLATE utf8_bin DEFAULT NULL COMMENT '操作系统',
  `callback_url` varchar(256) COLLATE utf8_bin DEFAULT NULL COMMENT '回调地址',
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `open_id` (`open_id`),
  KEY `imei_md5` (`imei_md5`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user` (
  `id` int(20) NOT NULL AUTO_INCREMENT,
  `open_id` varchar(128) COLLATE utf8_unicode_ci NOT NULL COMMENT '微信openId',
  `password` varchar(128) COLLATE utf8_unicode_ci DEFAULT '' COMMENT '登录密码',
  `type` varchar(45) COLLATE utf8_unicode_ci DEFAULT '' COMMENT '来源',
  `union_id` varchar(128) COLLATE utf8_unicode_ci DEFAULT '' COMMENT '微信union_id',
  `server_id` int(11) DEFAULT '0' COMMENT '区服id',
  `created_at` datetime DEFAULT NULL COMMENT '注册时间',
  `login_at` datetime DEFAULT NULL COMMENT '登录时间',
  `notice_version_at` datetime DEFAULT NULL COMMENT '公告版本弹一次',
  `notice_daily_at` datetime DEFAULT NULL COMMENT '每日首次必弹',
  `is_white_type` int(11) DEFAULT '0' COMMENT '是否白名单类型(0:不是,1:是)',
  `ip` varchar(32) COLLATE utf8_unicode_ci DEFAULT '' COMMENT 'ip地址',
  PRIMARY KEY (`id`),
  UNIQUE KEY `open_id` (`open_id`),
  KEY `server_id` (`server_id`)
) ENGINE=InnoDB AUTO_INCREMENT=47608 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-07-08 18:08:07
