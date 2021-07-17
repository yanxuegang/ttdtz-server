-- MySQL dump 10.13  Distrib 5.6.50, for Linux (x86_64)
--
-- Host: localhost    Database: app_line
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
-- Table structure for table `acl_users`
--

DROP TABLE IF EXISTS `acl_users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `acl_users` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `platform` varchar(10) NOT NULL DEFAULT '' COMMENT '平台',
  `username` varchar(64) NOT NULL COMMENT '用户名',
  `password` varchar(64) NOT NULL COMMENT '密码',
  `email` varchar(100) DEFAULT NULL COMMENT '邮箱',
  `is_admin` int(11) DEFAULT '0' COMMENT '是否管理员(0:不是,1:是)',
  `created_at` datetime DEFAULT NULL COMMENT '添加时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `username` (`username`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='用户表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `openid_map`
--

DROP TABLE IF EXISTS `openid_map`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `openid_map` (
  `openid` varchar(64) NOT NULL DEFAULT '',
  `system` tinyint(2) NOT NULL DEFAULT '0',
  `player_id` bigint(20) unsigned NOT NULL,
  PRIMARY KEY (`openid`,`system`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `chapter`
--

DROP TABLE IF EXISTS `chapter`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `chapter` (
  `id` int(20) NOT NULL AUTO_INCREMENT,
  `player_id` int(11) NOT NULL COMMENT '玩家表id',
  `chapter_id` int(11) DEFAULT '1' COMMENT '当前大章节',
  `type` int(11) DEFAULT '0' COMMENT '关卡类型',
  `chapter_sub_id` int(11) DEFAULT '0' COMMENT '打到最高关卡',
  `max_sub_id` int(11) DEFAULT '0' COMMENT '章节最高关卡',
  `number` int(11) DEFAULT '0' COMMENT '次数(星数)',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `player_id` (`player_id`,`chapter_id`),
  KEY `player_date` (`player_id`,`updated_at`)
) ENGINE=InnoDB AUTO_INCREMENT=42399 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `notice`
--

DROP TABLE IF EXISTS `notice`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `notice` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '标题',
  `type` int(11) NOT NULL DEFAULT '1' COMMENT '1 开服更新 2 停服更新',
  `notice_test` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '公告详情',
  `level` smallint(3) DEFAULT NULL COMMENT '水平格式 0左1中2右',
  `vertical` smallint(3) DEFAULT NULL COMMENT '垂直格式 0上1中2下',
  `created_at` datetime NOT NULL COMMENT '添加时间',
  `is_show` smallint(3) DEFAULT '1' COMMENT '1 显示 2删除',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=MyISAM AUTO_INCREMENT=27 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `payment`
--

DROP TABLE IF EXISTS `payment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `payment` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `player_id` int(11) NOT NULL COMMENT '玩家id',
  `amount` decimal(6,2) DEFAULT '0.00' COMMENT '提现金额(分)',
  `created_at` datetime DEFAULT NULL COMMENT '提现时间',
  `partner_trade_no` varchar(64) COLLATE utf8_bin DEFAULT '' COMMENT '商户订单号',
  `payment_no` varchar(64) COLLATE utf8_bin DEFAULT '' COMMENT '微信付款单号',
  `payment_time` varchar(32) COLLATE utf8_bin DEFAULT '' COMMENT '企业付款成功时间',
  PRIMARY KEY (`id`),
  KEY `player_date` (`player_id`,`created_at`)
) ENGINE=InnoDB AUTO_INCREMENT=51411 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `player`
--

DROP TABLE IF EXISTS `player`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player` (
  `id` bigint(20) unsigned NOT NULL,
  `open_id` varchar(128) COLLATE utf8_unicode_ci NOT NULL COMMENT '微信openId(账号)',
  `type` varchar(45) COLLATE utf8_unicode_ci DEFAULT '' COMMENT '来源',
  `name` varchar(45) COLLATE utf8_unicode_ci DEFAULT '' COMMENT '自定义昵称',
  `head` varchar(128) COLLATE utf8_unicode_ci DEFAULT '' COMMENT '自定义头像',
  `nick_name` varchar(45) COLLATE utf8_unicode_ci DEFAULT '' COMMENT '微信昵称',
  `avatar_url` varchar(128) COLLATE utf8_unicode_ci DEFAULT '' COMMENT '微信头像',
  `gender` int(4) DEFAULT '1' COMMENT '性别：1男 2女(0未知)',
  `city` varchar(45) COLLATE utf8_unicode_ci DEFAULT '' COMMENT '城市',
  `province` varchar(45) COLLATE utf8_unicode_ci DEFAULT '' COMMENT '省份',
  `country` varchar(45) COLLATE utf8_unicode_ci DEFAULT '' COMMENT '国家',
  `created_at` datetime DEFAULT NULL COMMENT '注册时间',
  `login_at` datetime DEFAULT NULL COMMENT '登录时间',
  `is_first` int(11) DEFAULT '0' COMMENT '当天登陆(0:未登陆,1:登陆)',
  `login_days` int(11) DEFAULT '0' COMMENT '登陆天数',
  `level` int(11) DEFAULT '0' COMMENT '玩家等级',
  `exp` int(11) DEFAULT '0' COMMENT '经验值',
  `gold` int(11) DEFAULT '0' COMMENT '玩家金币',
  `money` int(11) DEFAULT '0' COMMENT '玩家元宝',
  `chapter_id` int(11) DEFAULT '0' COMMENT '当前章节id',
  `chapter_level` int(11) DEFAULT '1' COMMENT '当前章节难度等级',
  `is_modify_nick_name` int(11) DEFAULT '0' COMMENT '是否已经修改过名字(0:未修改,1:修改)',
  `is_new_user_pass` int(11) DEFAULT '0' COMMENT '是否通过新手引导(0:未通过,1:通过)',
  `new_user_guide` varchar(128) COLLATE utf8_unicode_ci DEFAULT '' COMMENT '新手引导步骤',
  `new_user_chapter` varchar(500) COLLATE utf8_unicode_ci DEFAULT '' COMMENT '新手通章',
  `system_open` varchar(128) COLLATE utf8_unicode_ci DEFAULT '' COMMENT '已开放系统(模块)',
  `system_open_newuser_at` datetime DEFAULT NULL COMMENT '新手活动开启时间',
  `today_online_at` datetime DEFAULT NULL COMMENT '今天在线时间',
  `today_online_second` int(11) DEFAULT '0' COMMENT '今天在线时间秒数',
  `recharge_status` int(11) DEFAULT '0' COMMENT '首充奖励(0:未充值,1:已充值,2:已领取)',
  `ad_status` int(11) DEFAULT '0' COMMENT '广告奖励(0:未满足,1:未(可)领取,2:已领取)',
  `ad_status_at` datetime DEFAULT NULL COMMENT '广告状态领取时间(关羽)',
  `normal_chapter_num` int(11) DEFAULT '0' COMMENT '普通关进副本次数',
  `elite_chapter_num` int(11) DEFAULT '0' COMMENT '精英关进副本次数',
  `kill_monsters_num` int(11) DEFAULT '0' COMMENT '总杀怪数量',
  `monsters_packet_num` int(11) DEFAULT '0' COMMENT '怪物红包数量',
  `video_num` int(11) DEFAULT '0' COMMENT '看视频次数',
  `red_packet_num` int(11) DEFAULT '0' COMMENT '红包领取次数',
  `sign_money_num` int(11) DEFAULT '0' COMMENT '签到红包领取次数',
  `super_red_packet_num` int(11) DEFAULT '0' COMMENT '超级红包(3000)(6次)',
  `super_red_packet_at` datetime DEFAULT NULL COMMENT '超级红包领取日期',
  PRIMARY KEY (`id`),
  UNIQUE KEY `open_id` (`open_id`)
) ENGINE=InnoDB AUTO_INCREMENT=11593 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `player_limit`
--

DROP TABLE IF EXISTS `player_limit`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_limit` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `open_id` varchar(64) COLLATE utf8_bin DEFAULT NULL COMMENT '禁言账号',
  `created_id` datetime DEFAULT NULL COMMENT '添加时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7226 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `red_packet`
--

DROP TABLE IF EXISTS `red_packet`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `red_packet` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `player_id` int(11) DEFAULT NULL COMMENT '玩家id',
  `money` int(11) DEFAULT NULL COMMENT '随机红包币',
  `token` varchar(64) COLLATE utf8_bin DEFAULT NULL COMMENT '签名值',
  `time` varchar(32) COLLATE utf8_bin DEFAULT NULL COMMENT '签名时间',
  `log_date` date DEFAULT NULL COMMENT '日期',
  `status` int(4) DEFAULT '0' COMMENT '0:未领取,1:领取',
  `created_at` datetime DEFAULT NULL COMMENT '记录生成时间',
  PRIMARY KEY (`id`),
  KEY `player_id` (`player_id`,`log_date`),
  KEY `token` (`token`,`money`,`status`)
) ENGINE=InnoDB AUTO_INCREMENT=11899 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `sign_day`
--

DROP TABLE IF EXISTS `sign_day`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sign_day` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `player_id` int(11) NOT NULL COMMENT '用户id',
  `days` int(11) DEFAULT '1' COMMENT '签到天数',
  `created_at` datetime DEFAULT NULL COMMENT '签到日期',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `player_id` (`player_id`)
) ENGINE=MyISAM AUTO_INCREMENT=34835 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=FIXED;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `sign_online`
--

DROP TABLE IF EXISTS `sign_online`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sign_online` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `player_id` int(11) DEFAULT NULL COMMENT '玩家id',
  `type` int(11) DEFAULT NULL COMMENT '在线类型(分钟)',
  `second` int(11) DEFAULT '0' COMMENT '在线时间(签到时总在线时长)',
  `created_at` datetime DEFAULT NULL COMMENT '记录时间',
  PRIMARY KEY (`id`),
  KEY `player_id` (`player_id`,`created_at`)
) ENGINE=InnoDB AUTO_INCREMENT=1993 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `task`
--

DROP TABLE IF EXISTS `task`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `task` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `player_id` int(11) DEFAULT NULL COMMENT '玩家表id',
  `type` int(11) DEFAULT NULL COMMENT '任务类型',
  `number` int(11) DEFAULT '0' COMMENT '次数',
  `status` int(11) DEFAULT '0' COMMENT '领取状态(0:未领,1:已领)',
  `created_at` datetime DEFAULT NULL COMMENT '添加时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间(领取)',
  PRIMARY KEY (`id`),
  KEY `player_id` (`player_id`,`created_at`),
  KEY `player_id_2` (`player_id`,`type`,`created_at`)
) ENGINE=InnoDB AUTO_INCREMENT=116382 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `task_list`
--

DROP TABLE IF EXISTS `task_list`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `task_list` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `player_id` int(11) NOT NULL COMMENT '用户id',
  `type` smallint(3) DEFAULT '0' COMMENT '1:领取签到红包\r\n2:闯关三次\r\n3:下载app三次\r\n4:评分三星以上5次\r\n5:观看视频5次',
  `num` smallint(3) DEFAULT '0' COMMENT '任务次数',
  `money` int(11) DEFAULT '0' COMMENT '红包币',
  `name` varchar(20) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '任务名称',
  `status` smallint(3) DEFAULT '1' COMMENT '任务状态1未完成 2 已完成',
  `get_state` smallint(3) DEFAULT '1' COMMENT '领取奖励状态 1:未领取 2:已领取',
  `created_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '任务时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=MyISAM AUTO_INCREMENT=9880 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-07-08 18:07:20
