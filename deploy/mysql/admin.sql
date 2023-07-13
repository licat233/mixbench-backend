/*
 Navicat MySQL Data Transfer

 Source Server         : localmysql
 Source Server Type    : MySQL
 Source Server Version : 80033
 Source Host           : localhost:3306
 Source Schema         : mixbench_admin

 Target Server Type    : MySQL
 Target Server Version : 80033
 File Encoding         : 65001

 Date: 12/07/2023 20:46:53
*/

USE mixbench_admin;
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for adminer
-- ----------------------------
DROP TABLE IF EXISTS `adminer`;
CREATE TABLE `adminer` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '表主键',
  `uuid` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'UUID',
  `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '管理员名称',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '头像',
  `access` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '身份',
  `passport` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '账号',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '密码',
  `email` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '邮箱',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '账号状态，是否可用',
  `is_super_admin` tinyint NOT NULL DEFAULT '0' COMMENT '是否为超级管理员',
  `project_id` int NOT NULL DEFAULT '0' COMMENT '项目id',
  `login_count` int NOT NULL DEFAULT '0' COMMENT '登录次数',
  `last_login` datetime NOT NULL DEFAULT '2005-01-02 15:04:05' COMMENT '最后一次登录时间',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_deleted` tinyint NOT NULL DEFAULT '0' COMMENT '是否已删除',
  PRIMARY KEY (`id`),
  KEY `uuid` (`uuid`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='管理员';

-- ----------------------------
-- Records of adminer
-- ----------------------------
BEGIN;
INSERT INTO `adminer` VALUES (1, 'b4a1d67c-97cc-463f-a366-65b5caa5e3aa', '超管', '//mixbench.net/static/avatar/avatar_1_1663575200688027000.jpg', 'superAdmin', 'licat', '$2a$10$9ZGoeKcUplsVsqUsXNNDBeF4KI35yz4w4BK4H8gX0O3Dm9mc64cAW', '1697470538@qq.com', 1, 1, 0, 18, '2022-10-14 17:04:03', '2022-07-29 09:47:08', '2022-10-14 17:04:02', 0);
COMMIT;

-- ----------------------------
-- Table structure for campaign
-- ----------------------------
DROP TABLE IF EXISTS `campaign`;
CREATE TABLE `campaign` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '表主键',
  `uuid` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'UUID',
  `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '活动名称',
  `project_id` int NOT NULL DEFAULT '0' COMMENT '项目id',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '活动状态，是否在举办，默认关闭，需手动开启',
  `pattern` tinyint NOT NULL DEFAULT '0' COMMENT '方案: 0会员制，1公开制，2会员需授权',
  `max_count` int NOT NULL DEFAULT '1' COMMENT '最大抽奖次数：-1表示无限制，0表示不可抽奖',
  `max_prizes` int NOT NULL DEFAULT '1' COMMENT '最大抽奖次数：-1表示无限制，0表示不可中奖',
  `start_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '活动开始时间',
  `end_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '活动结束时间',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_deleted` tinyint NOT NULL DEFAULT '0' COMMENT '是否已删除',
  PRIMARY KEY (`id`),
  KEY `uuid` (`uuid`),
  KEY `project_id` (`project_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='活动';

-- ----------------------------
-- Table structure for jwt_blacklist
-- ----------------------------
DROP TABLE IF EXISTS `jwt_blacklist`;
CREATE TABLE `jwt_blacklist` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '表主键',
  `token` blob NOT NULL COMMENT 'jwt-token',
  `platform` varchar(254) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '绑定平台',
  `ip` varchar(254) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '0.0.0.0' COMMENT '绑定ip',
  `expire_at` datetime NOT NULL DEFAULT '2005-01-02 15:04:05' COMMENT '过期时间',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='jwt黑名单';

-- ----------------------------
-- Table structure for project
-- ----------------------------
DROP TABLE IF EXISTS `project`;
CREATE TABLE `project` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '表主键',
  `uuid` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'UUID',
  `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '项目名称',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '项目状态: 0已下线，1已上线',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建日期',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_deleted` tinyint NOT NULL DEFAULT '0' COMMENT '是否已删除',
  PRIMARY KEY (`id`),
  KEY `uuid` (`uuid`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='项目';

-- ----------------------------
-- Records of project
-- ----------------------------
BEGIN;
INSERT INTO `project` VALUES (1, '14a1d67c-97cc-463f-a366-65b5caa5e3aa', '植题祛斑', 1, '2022-07-07 17:37:52', '2022-08-29 10:01:45', 0);
INSERT INTO `project` VALUES (2, '24a1d67c-97cc-463f-a366-65b5caa5e3aa', '李中医祛斑', 0, '2022-08-29 10:08:57', '2022-08-29 10:08:57', 0);
COMMIT;

-- ----------------------------
-- Table structure for web
-- ----------------------------
DROP TABLE IF EXISTS `web`;
CREATE TABLE `web` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '表主键',
  `uuid` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'UUID',
  `project_id` int NOT NULL DEFAULT '0' COMMENT '项目id',
  `url` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '连接',
  `host` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '域名',
  `description` varchar(254) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '备注',
  `code` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '代码',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建日期',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_deleted` tinyint NOT NULL DEFAULT '0' COMMENT '是否已删除',
  PRIMARY KEY (`id`),
  KEY `uuid` (`uuid`),
  KEY `project_id` (`project_id`),
  KEY `url` (`url`),
  KEY `host` (`host`),
  UNIQUE (`url`),
  UNIQUE (`host`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='网站';

SET FOREIGN_KEY_CHECKS = 1;
