USE mixbench_admin;

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for adminer
-- ----------------------------
DROP TABLE IF EXISTS `adminer`;
CREATE TABLE `adminer` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '表主键',
  `uuid` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '唯一识别码',
  `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '管理员名称',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '头像',
  `passport` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT 'admin' COMMENT '账号',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT 'admin' COMMENT '密码',
  `email` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '邮箱',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '账号状态，是否可用',
  `is_super_admin` tinyint NOT NULL DEFAULT '0' COMMENT '是否为超级管理员',
  `login_count` int NOT NULL DEFAULT '0' COMMENT '登录次数',
  `last_login` datetime NOT NULL DEFAULT '2005-01-02 15:04:05' COMMENT '最后一次登录时间',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_deleted` tinyint NOT NULL DEFAULT '0' COMMENT '是否已被删除',
  `delete_at` datetime NOT NULL DEFAULT '2005-01-02 15:04:05' COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `uuid` (`uuid`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='管理员';

-- ----------------------------
-- Records of adminer
-- ----------------------------
BEGIN;
INSERT INTO `adminer` VALUES (1, '3a37bef3-65da-4d1f-9acf-f13b308ad213', '超管', '//mixbench.net/static/avatar/avatar_1_1663575200688027000.jpg', 'licat', '', '1697470538@qq.com', 1, 1, 18, '2022-10-14 17:04:03', '2022-07-29 09:47:08', '2022-10-14 17:04:02', 0, '2005-01-02 15:04:05');
COMMIT;

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
-- DROP TABLE IF EXISTS `casbin_rule`;
-- CREATE TABLE `casbin_rule` (
--   `p_type` varchar(32) NOT NULL DEFAULT '',
--   `v0` varchar(255) NOT NULL DEFAULT '',
--   `v1` varchar(255) NOT NULL DEFAULT '',
--   `v2` varchar(255) NOT NULL DEFAULT '',
--   `v3` varchar(255) NOT NULL DEFAULT '',
--   `v4` varchar(255) NOT NULL DEFAULT '',
--   `v5` varchar(255) NOT NULL DEFAULT '',
--   KEY `idx_casbin_rule` (`p_type`,`v0`,`v1`)
-- ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for jwt_blacklist
-- ----------------------------
DROP TABLE IF EXISTS `jwt_blacklist`;
CREATE TABLE `jwt_blacklist` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '表主键',
  `uuid` varchar(254) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '用户或者管理员的唯一识别码',
  `token` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'jwt-token',
  `platform` varchar(254) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '绑定平台',
  `ip` varchar(254) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '0.0.0.0' COMMENT '绑定ip',
  `expire_at` datetime NOT NULL DEFAULT '2005-01-02 15:04:05' COMMENT '过期时间',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `uuid` (`uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='jwt黑名单';

SET FOREIGN_KEY_CHECKS = 1;

