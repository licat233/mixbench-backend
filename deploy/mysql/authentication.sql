USE mixbench_authentication;

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for ipblack
-- ----------------------------
DROP TABLE IF EXISTS `ipblack`;
CREATE TABLE `ipblack` (
  `id` int NOT NULL AUTO_INCREMENT,
  `ip` varchar(100) NOT NULL DEFAULT '' COMMENT 'IP',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_deleted` tinyint NOT NULL DEFAULT '0' COMMENT '是否已删除',
  PRIMARY KEY (`id`),
  UNIQUE KEY `ip` (`ip`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='IP黑名单';

-- ----------------------------
-- Table structure for jwt_blacklist
-- ----------------------------
DROP TABLE IF EXISTS `jwt_blacklist`;

CREATE TABLE `jwt_blacklist` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '表主键',
  `token` BLOB NOT NULL COMMENT 'jwt-token',
  `platform` varchar(254) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '绑定平台',
  `ip` varchar(254) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '0.0.0.0' COMMENT '绑定ip',
  `expire_at` datetime NOT NULL DEFAULT '2005-01-02 15:04:05' COMMENT '过期时间',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = 'jwt黑名单';

SET FOREIGN_KEY_CHECKS = 1;