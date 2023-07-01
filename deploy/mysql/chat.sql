USE mixbench_chat;

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

DROP TABLE IF EXISTS `visitor`;
CREATE TABLE `visitor` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '表主键',
  `uuid` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'UUID',
  `username` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '游客名',
  `ip` varchar(50) NOT NULL DEFAULT '' COMMENT 'IP',
  `kefu_id` int NOT NULL DEFAULT '-1' COMMENT '客服id，-1表示没有客服接待',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态:0离线，1在线，2拉黑',
  `referer` varchar(500) NOT NULL DEFAULT '' COMMENT '来源，从哪个网址发起的聊天',
  `city` varchar(100) NOT NULL DEFAULT '' COMMENT '城市或地区',
  `platform` varchar(100) NOT NULL DEFAULT '' COMMENT '设备',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_deleted` tinyint NOT NULL DEFAULT '0' COMMENT '是否已删除',
  PRIMARY KEY (`id`),
  KEY `uuid` (`uuid`),
  KEY `username` (`username`),
  KEY `status` (`status`),
  KEY `kefu_id` (`kefu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='游客';

DROP TABLE IF EXISTS `kefu`;
CREATE TABLE `kefu` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '表主键',
  `uuid` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'UUID',
  `passport` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '客服账号',
  `username` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '客服名',
  `password` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '密码',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态:0离线，1在线，2冻结',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_deleted` tinyint NOT NULL DEFAULT '0' COMMENT '是否已删除',
  PRIMARY KEY (`id`),
  KEY `uuid` (`uuid`),
  KEY `passport` (`passport`),
  KEY `username` (`username`),
  KEY `status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='客服';

DROP TABLE IF EXISTS `chat_record`;
CREATE TABLE `chat_record` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '表主键',
  `uuid` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'UUID',
  `content` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '消息内容',
  `visitor_id` int NOT NULL DEFAULT '0' COMMENT '游客id',
  `kefu_id` int NOT NULL DEFAULT '0' COMMENT '客服id',
  `msg_type` tinyint NOT NULL DEFAULT '0' COMMENT '0游客消息，1客服消息，2公告消息，3系统消息',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态:0合法，1违规',
  `is_read` tinyint NOT NULL DEFAULT '0' COMMENT '状态:0未读，1已读',
  `is_send` tinyint NOT NULL DEFAULT '0' COMMENT '状态:0未发送，1已发送',
  `send_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '发送时间',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_deleted` tinyint NOT NULL DEFAULT '0' COMMENT '是否已删除',
  PRIMARY KEY (`id`),
  KEY `uuid` (`uuid`),
  KEY `visitor_id` (`visitor_id`),
  KEY `kefu_id` (`kefu_id`),
  KEY `msg_type` (`msg_type`),
  KEY `status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='聊天记录';

DROP TABLE IF EXISTS `message`;
CREATE TABLE `message` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '表主键',
  `uuid` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'UUID',
  `content` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '消息内容',
  `visitor_id` int NOT NULL DEFAULT '0' COMMENT '游客id',
  `kefu_id` int NOT NULL DEFAULT '0' COMMENT '客服id',
  `msg_type` tinyint NOT NULL DEFAULT '0' COMMENT '0游客消息，1客服消息，2公告消息，3系统消息',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态:0合法，1违规',
  `is_read` tinyint NOT NULL DEFAULT '0' COMMENT '状态:0未读，1已读',
  `is_send` tinyint NOT NULL DEFAULT '0' COMMENT '状态:0未发送，1已发送',
  `send_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '发送时间',
  `expire_time` datetime NOT NULL DEFAULT (CURRENT_TIMESTAMP + INTERVAL 3 DAY) COMMENT '过期时间, 默认为3天',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_deleted` tinyint NOT NULL DEFAULT '0' COMMENT '是否已删除',
  PRIMARY KEY (`id`),
  KEY `uuid` (`uuid`),
  KEY `visitor_id` (`visitor_id`),
  KEY `kefu_id` (`kefu_id`),
  KEY `msg_type` (`msg_type`),
  KEY `expire_time` (`expire_time`),
  KEY `status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='消息, 主要用于缓存未发送成功的消息';

DROP TABLE IF EXISTS `ipblack`;
CREATE TABLE `ipblack` (
  `id` int NOT NULL AUTO_INCREMENT,
  `ip` varchar(100) NOT NULL DEFAULT '' COMMENT 'IP',
  `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `kefu_id` varchar(100) NOT NULL DEFAULT '' COMMENT '客服id',
  PRIMARY KEY (`id`),
  UNIQUE KEY `ip` (`ip`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='IP黑名单';

DROP TABLE IF EXISTS `message_config`;
CREATE TABLE `message_config` (
  `id` int unsigned NOT NULL AUTO_INCREMENT  COMMENT '主键',
  `name` varchar(255) NOT NULL DEFAULT ''  COMMENT '名字',
  `key` varchar(255) NOT NULL DEFAULT '' COMMENT '密匙',
  `content` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '消息内容',
  PRIMARY KEY (`id`),
  UNIQUE KEY `key` (`key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='消息配置';
