USE mixbench_luckydraw;

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for activity
-- ----------------------------
DROP TABLE IF EXISTS `activity`;
CREATE TABLE `activity` (
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
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='活动';

-- ----------------------------
-- Records of activity
-- ----------------------------
BEGIN;
INSERT INTO `activity` VALUES (1, 'b411d67c-97cc-463f-a366-65b5caa5e3aa','618购物节', 1, 0, 0, 1, 1, '2022-07-07 17:38:50', '2022-07-31 17:38:52', '2022-07-07 17:38:56', '2022-08-29 11:34:43', 0);
INSERT INTO `activity` VALUES (2, 'b431d67c-97cc-463f-a366-65b5caa5e3aa','七夕促销', 1, 1, 1, 1, 1, '2022-07-07 17:39:39', '2022-08-31 17:39:41', '2022-07-07 17:39:46', '2022-07-07 17:39:46', 0);
INSERT INTO `activity` VALUES (3, 'b441d67c-97cc-463f-a366-65b5caa5e3aa','中秋活动', 1, 1, 2, 1, 1, '2022-08-30 14:56:10', '2022-08-31 14:56:14', '2022-08-30 14:59:43', '2022-09-20 17:03:04', 0);
COMMIT;

-- ----------------------------
-- Table structure for lucky
-- ----------------------------
DROP TABLE IF EXISTS `lucky`;
CREATE TABLE `lucky` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '表主键',
  `uuid` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'UUID',
  `user_id` int NOT NULL DEFAULT '0' COMMENT '中奖用户id',
  `prize_pool_id` int NOT NULL DEFAULT '0' COMMENT '所属奖品池id',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态: 0未领奖，1已领奖，2弃领',
  `luck_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '抽奖时间',
  `ip` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '0.0.0.0' COMMENT 'ip地址',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_deleted` tinyint NOT NULL DEFAULT '0' COMMENT '是否已删除',
  PRIMARY KEY (`id`),
  KEY `uuid` (`uuid`),
  KEY `user_id` (`user_id`),
  KEY `prize_pool_id` (`prize_pool_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='中奖名单';

-- ----------------------------
-- Records of lucky
-- ----------------------------
BEGIN;
INSERT INTO `lucky` VALUES (1, 'b4a1d17c-97cc-463f-a366-65b5caa5e3aa', 1, 2, 1, '2022-07-29 17:49:13', '127.0.0.1', '2022-07-07 17:49:15', '2022-09-09 15:24:25', 0);
COMMIT;

-- ----------------------------
-- Table structure for prize
-- ----------------------------
DROP TABLE IF EXISTS `prize`;
CREATE TABLE `prize` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '表主键',
  `uuid` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'UUID',
  `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '奖品名称',
  `image` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '奖品图片',
  `price` decimal(10,2) NOT NULL DEFAULT 0.00 COMMENT '奖品价格',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '是否可用',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_deleted` tinyint NOT NULL DEFAULT '0' COMMENT '是否已删除',
  PRIMARY KEY (`id`),
  KEY `uuid` (`uuid`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='奖品名单';

-- ----------------------------
-- Records of prize
-- ----------------------------
BEGIN;
INSERT INTO `prize` VALUES (1, 'b4a1d67c-97cc-463f-a366-65b5caa3e3aa', '美的蒸烤箱', 'https://img.alicdn.com/imgextra/i3/917298378/O1CN01Us4mhB2BlAymmJh8F_!!917298378.png', 325.99, 0, '2022-07-07 17:51:03', '2022-09-02 09:45:09', 0);
INSERT INTO `prize` VALUES (2, 'b4a1d67c-97cc-463f-a366-65b5caa4e3aa', 'YSL-聖羅蘭', 'https://img.alicdn.com/imgextra/i3/917298378/O1CN01WlLtDx2BlAyrbjRQs_!!917298378.png', 199.89, 1, '2022-07-07 17:51:38', '2022-07-07 17:51:38', 0);
INSERT INTO `prize` VALUES (3, 'b4a1d67c-97cc-463f-a366-65b5caase3aa', '雪肌淨膚面膜', 'https://img.alicdn.com/imgextra/i3/917298378/O1CN01Nr8x9L2BlAyrbiAQ8_!!917298378.png', 668.99, 1, '2022-07-07 17:36:33', '2022-07-07 17:36:33', 0);
INSERT INTO `prize` VALUES (4, 'b4a1d67c-97cc-463f-a366-65b5caaae3aa', '谢谢参与', '', 0.00, 1, '2022-07-07 17:44:39', '2022-09-01 17:40:00', 0);
INSERT INTO `prize` VALUES (5, 'b4a1d67c-97cc-463f-a366-65b5caafe3aa', '黑头仪', '//mixbench.net/static/img/img_1_1663557715940496000.jpg', 32.00, 1, '2022-09-19 11:22:09', '2022-09-19 11:26:23', 0);
INSERT INTO `prize` VALUES (6, 'b4a1d67c-97cc-463f-a366-65b5caaze3aa', '黑头仪', '//mixbench.net/static/img/img_1_1663557803475449000.jpg', 32.00, 1, '2022-09-19 11:23:39', '2022-09-19 11:29:44', 0);
INSERT INTO `prize` VALUES (7, 'b4a1d67c-97cc-463f-a366-65b5caaxe3aa', '黑头仪', '//mixbench.net/static/img/img_1_1663557945410676000.jpg', 32.00, 1, '2022-09-19 11:25:48', '2022-09-19 11:25:48', 0);
COMMIT;

-- ----------------------------
-- Table structure for prize_grade
-- ----------------------------
DROP TABLE IF EXISTS `prize_grade`;
CREATE TABLE `prize_grade` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '表主键',
  `uuid` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'UUID',
  `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '级别名称',
  `is_win` tinyint NOT NULL DEFAULT '0' COMMENT '是否中奖',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_deleted` tinyint NOT NULL DEFAULT '0' COMMENT '是否已删除',
  PRIMARY KEY (`id`),
  KEY `uuid` (`uuid`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='奖品级别';

-- ----------------------------
-- Records of prize_grade
-- ----------------------------
BEGIN;
INSERT INTO `prize_grade` VALUES (1, 'b4a1d671-97cc-463f-a366-65b5caa5e3aa','未中奖', 0, '2022-07-07 17:40:55', '2022-07-07 17:40:55', 0);
INSERT INTO `prize_grade` VALUES (2, 'b4a1d672-97cc-463f-a366-65b5caa5e3aa','幸运奖', 1, '2022-07-07 17:41:04', '2022-07-07 17:41:04', 0);
INSERT INTO `prize_grade` VALUES (3, 'b4a1d673-97cc-463f-a366-65b5caa5e3aa','一等奖', 1, '2022-07-07 17:41:21', '2022-07-07 17:41:21', 0);
INSERT INTO `prize_grade` VALUES (4, 'b4a1d674-97cc-463f-a366-65b5caa5e3aa','二等奖', 1, '2022-07-07 17:41:30', '2022-07-07 17:41:30', 0);
INSERT INTO `prize_grade` VALUES (5, 'b4a1d675-97cc-463f-a366-65b5caa5e3aa','三等奖', 1, '2022-07-07 17:41:36', '2022-07-07 17:41:36', 0);
INSERT INTO `prize_grade` VALUES (6, 'b4a1d676-97cc-463f-a366-65b5caa5e3aa','四等奖', 1, '2022-07-07 17:41:42', '2022-07-07 17:41:42', 0);
INSERT INTO `prize_grade` VALUES (7, 'b4a1d677-97cc-463f-a366-65b5caa5e3aa','特等奖', 1, '2022-07-07 17:41:58', '2022-07-07 17:41:58', 0);
COMMIT;

-- ----------------------------
-- Table structure for prize_pool
-- ----------------------------
DROP TABLE IF EXISTS `prize_pool`;
CREATE TABLE `prize_pool` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '表主键',
  `uuid` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'UUID',
  `activity_id` int NOT NULL DEFAULT '0' COMMENT '活动id',
  `prize_id` int NOT NULL DEFAULT '0' COMMENT '奖品id',
  `grade_id` int NOT NULL DEFAULT '1' COMMENT '奖品等级',
  `luck_value` int NOT NULL DEFAULT '0' COMMENT '幸运值',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '是否可用，0不可用，1可用',
  `total` int NOT NULL DEFAULT '1' COMMENT '奖品总数量',
  `remainder` int NOT NULL DEFAULT '1' COMMENT '奖品剩余数量',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_deleted` tinyint NOT NULL DEFAULT '0' COMMENT '是否已删除',
  PRIMARY KEY (`id`),
  KEY `uuid` (`uuid`),
  KEY `activity_id` (`activity_id`),
  KEY `prize_id` (`prize_id`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='活动奖品池，某活动中的奖品信息';

-- ----------------------------
-- Records of prize_pool
-- ----------------------------
BEGIN;
INSERT INTO `prize_pool` VALUES (1, 'aE1gM6eE5kK0eO2bI3cA1iK1bR0qD0rC0aB4', 1, 4, 1, 100, 1, 1, 1, '2022-07-07 17:44:54', '2022-09-16 11:14:53', 0);
INSERT INTO `prize_pool` VALUES (2, 'bK0bA0uD0kH0bB0oC0cA0oH3eB0aH0wI6cR4', 1, 1, 2, 100, 1, 1, 1, '2022-07-07 17:46:01', '2022-09-20 16:55:59', 0);
INSERT INTO `prize_pool` VALUES (3, 'iD3hE0gD2yD3mT0jI5jV0bB5bC2aA1qA2fC0', 1, 2, 3, 100, 1, 1, 1, '2022-07-07 17:46:26', '2022-07-07 17:46:26', 0);
INSERT INTO `prize_pool` VALUES (4, 'aW1cI0aB0cB0gA6aR4bC1mA1bD0fJ2bB5hB8', 1, 3, 4, 100, 1, 1, 1, '2022-07-07 17:46:44', '2022-07-07 17:46:44', 0);
INSERT INTO `prize_pool` VALUES (5, 'a101c466-27d8-4bc5-a7ee-e336a95eedc0', 2, 1, 3, 100, 1, 2, 1, '2022-09-15 15:58:35', '2022-09-15 18:34:49', 0);
INSERT INTO `prize_pool` VALUES (6, '373470dd-fe52-4fce-b494-f1e68614fa8b', 3, 3, 3, 100, 1, 1, 1, '2022-09-15 16:44:59', '2022-09-16 10:35:29', 0);
INSERT INTO `prize_pool` VALUES (7, '8bbefce6-58e2-49f5-9892-678b15bf7919', 3, 4, 3, 100, 1, 1, 1, '2022-09-15 16:48:41', '2022-09-16 10:00:08', 0);
INSERT INTO `prize_pool` VALUES (8, 'b71dbda2-1259-4c91-8465-1c3a642318cf', 3, 2, 2, 100, 1, 1, 1, '2022-09-15 16:53:43', '2022-09-17 17:25:45', 0);
INSERT INTO `prize_pool` VALUES (9, 'de1b6a70-3258-4642-b62a-cccf286e9cae', 3, 1, 7, 100, 1, 1, 1, '2022-09-15 16:58:17', '2022-09-15 16:58:17', 0);
INSERT INTO `prize_pool` VALUES (10, '073ceb11-8a13-4df0-a24f-85e0510da708', 2, 2, 4, 100, 1, 1, 1, '2022-09-15 16:58:32', '2022-09-16 10:03:17', 0);
INSERT INTO `prize_pool` VALUES (11, '6f59aec7-47c8-4067-98ed-2e748bfc7679', 2, 3, 5, 100, 1, 1, 1, '2022-09-15 16:58:54', '2022-09-17 17:12:02', 0);
INSERT INTO `prize_pool` VALUES (12, 'd4abe8bc-3422-40d0-9a17-b44a6674f39f', 2, 1, 3, 100, 1, 1, 1, '2022-09-15 17:53:57', '2022-09-20 16:26:53', 0);
INSERT INTO `prize_pool` VALUES (13, '98350e4d-8112-447d-a332-982a0197983c', 2, 2, 4, 100, 1, 1, 1, '2022-09-15 17:57:02', '2022-09-20 17:08:04', 0);
INSERT INTO `prize_pool` VALUES (14, '79030ce6-e043-479e-a98f-9a0c910886ba', 3, 3, 5, 100, 1, 1, 1, '2022-09-16 10:35:37', '2022-09-17 17:25:42', 0);
COMMIT;

-- ----------------------------
-- Table structure for project
-- ----------------------------
DROP TABLE IF EXISTS `project`;
CREATE TABLE `project` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '表主键',
  `uuid` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'UUID',
  `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '项目名称',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '项目状态',
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
INSERT INTO `project` VALUES (1,  '14a1d67c-97cc-463f-a366-65b5caa5e3aa','植题祛斑', 1, '2022-07-07 17:37:52', '2022-08-29 10:01:45', 0);
INSERT INTO `project` VALUES (2, '24a1d67c-97cc-463f-a366-65b5caa5e3aa', '李中医祛斑', 0, '2022-08-29 10:08:57', '2022-08-29 10:08:57', 0);
COMMIT;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '表主键',
  `uuid` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'UUID',
  `passport` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '凭证，可以是手机号或者社交账号，一般是由客户填入',
  `email` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '邮箱',
  `phone` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '手机号',
  `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `address` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '收货地址',
  `score` int NOT NULL DEFAULT '0' COMMENT '积分',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '状态:0未授权，1正常，2已冻结',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_deleted` tinyint NOT NULL DEFAULT '0' COMMENT '是否已删除',
  PRIMARY KEY (`id`),
  KEY `uuid` (`uuid`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户名单';

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `user` VALUES (1, 'b4a1d67c-97cc-463f-a366-a5b5caa5e3aa','licat', 'licat233@gmail.com', '13570923455', '李哥', '广东省广州市白云区金沙洲', 0, 1, '2022-07-07 17:50:22', '2022-09-07 17:19:38', 0);
COMMIT;

-- ----------------------------
-- Table structure for user_pool
-- ----------------------------
DROP TABLE IF EXISTS `user_pool`;
CREATE TABLE `user_pool` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '表主键',
  `uuid` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'UUID',
  `activity_id` int NOT NULL DEFAULT '0' COMMENT '活动id',
  `user_id` int NOT NULL DEFAULT '0' COMMENT '用户id',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '状态:0未授权，1正常，2已冻结',
  `score` int NOT NULL DEFAULT '0' COMMENT '用户在当前活动的积分',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_deleted` tinyint NOT NULL DEFAULT '0' COMMENT '是否已删除',
  PRIMARY KEY (`id`),
  KEY `uuid` (`uuid`),
  KEY `activity_id` (`activity_id`),
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='活动用户池，可以参与该活动抽奖的用户，及其在该活动中的积分管理';

-- ----------------------------
-- Records of user_pool
-- ----------------------------
BEGIN;
INSERT INTO `user_pool` VALUES (1,'b4a1d67c-97cc-4631-a366-65b5caa5e3aa', 1, 1, 1, 0, '2022-07-07 17:53:05', '2022-07-07 17:53:05', 0);
INSERT INTO `user_pool` VALUES (3, 'b4a1d67c-97cc-4632-a366-65b5caa5e3aa', 2, 1, 1, 0, '2022-09-17 17:30:12', '2022-09-17 17:30:18', 0);
INSERT INTO `user_pool` VALUES (4, 'b4a1d67c-97cc-4633-a366-65b5caa5e3aa', 3, 1, 1, 0, '2022-09-17 17:30:26', '2022-09-19 00:12:39', 0);
INSERT INTO `user_pool` VALUES (5, 'b4a1d67c-97cc-4634-a366-65b5caa5e3aa', 3, 1, 1, 0, '2022-09-17 17:31:12', '2022-09-19 00:16:29', 0);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
