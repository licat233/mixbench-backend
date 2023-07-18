USE mixbench_chat;

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for message
-- ----------------------------
DROP TABLE IF EXISTS `message`;
CREATE TABLE `message` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '表主键',
  `uuid` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'UUID',
  `content` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '消息内容',
  `content_type` tinyint NOT NULL DEFAULT '0' COMMENT '0文本，1图片，2表情，3视频, 4错误信息',
  `visitor_id` int NOT NULL DEFAULT '0' COMMENT '游客id',
  `kefu_id` int NOT NULL DEFAULT '0' COMMENT '客服id',
  `project_id` int NOT NULL DEFAULT '0' COMMENT '项目id',
  `msg_type` tinyint NOT NULL DEFAULT '0' COMMENT '0游客消息，1客服消息，2公告消息，3系统消息',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态:0合法，1违规',
  `is_read` tinyint NOT NULL DEFAULT '0' COMMENT '是否已读:0未读，1已读',
  `is_send` tinyint NOT NULL DEFAULT '0' COMMENT '是否已发送:0未发送，1已发送',
  `send_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '发送时间',
  `expire_time` datetime NOT NULL DEFAULT ((now() + interval 30 day)) COMMENT '缓存过期时间, 默认为30天',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_deleted` tinyint NOT NULL DEFAULT '0' COMMENT '是否已删除',
  PRIMARY KEY (`id`),
  KEY `uuid` (`uuid`),
  KEY `visitor_id` (`visitor_id`),
  KEY `project_id` (`project_id`),
  KEY `kefu_id` (`kefu_id`),
  KEY `msg_type` (`msg_type`),
  KEY `expire_time` (`expire_time`),
  KEY `status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='消息';

-- ----------------------------
-- Table structure for message_config
-- ----------------------------
DROP TABLE IF EXISTS `message_config`;
CREATE TABLE `message_config` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `project_id` int NOT NULL DEFAULT '0' COMMENT '项目id',
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '名字',
  `msgkey` varchar(255) NOT NULL DEFAULT '' COMMENT '消息密匙',
  `content` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '消息内容',
  PRIMARY KEY (`id`),
  UNIQUE KEY `msgkey` (`msgkey`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='消息配置';

-- ----------------------------
-- Records of message_config
-- ----------------------------
BEGIN;
INSERT INTO `message_config` VALUES (1, 1, '歡迎信息', 'welcome_message', '你好！我是Anna老師，請問有什麼可以幫您？ \\n皺紋問題請回 1 \\n色斑問題請回 2 \\n眼袋問題請回 3 \\n其它肌膚問題請回 4');
INSERT INTO `message_config` VALUES (2, 1, '留言信息', 'leave_message', '你好！Anna暫時不在線，有問題可以添加Anna的line: <a href=\"line://ti/p/Jh0NreIhfp\" style=\"color:green\" target=\"_blank\">anby091</a>');
COMMIT;

-- ----------------------------
-- Table structure for visitor
-- ----------------------------
DROP TABLE IF EXISTS `visitor`;
CREATE TABLE `visitor` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '表主键',
  `uuid` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'UUID',
  `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '游客名',
  `ip` varchar(50) NOT NULL DEFAULT '' COMMENT 'IP',
  `kefu_id` int NOT NULL DEFAULT '0' COMMENT '客服id，0表示没有客服接待',
  `project_id` int NOT NULL DEFAULT '0' COMMENT '项目id',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态:0离线，1在线，2拉黑',
  `referer` varchar(500) NOT NULL DEFAULT '' COMMENT '来源，从哪个网址发起的聊天',
  `region` varchar(100) NOT NULL DEFAULT '' COMMENT '地区',
  `city` varchar(100) NOT NULL DEFAULT '' COMMENT '城市',
  `platform` varchar(100) NOT NULL DEFAULT '' COMMENT '设备平台:移动设备,电脑,平板,Xbox',
  `os` varchar(100) NOT NULL DEFAULT '' COMMENT '系统: andriod,ios,macOS,windows',
  `browser` varchar(100) NOT NULL DEFAULT '' COMMENT '浏览器: Chrome,Safari,InternalExplorer,Firefox,Edge',
  `language` varchar(100) NOT NULL DEFAULT '' COMMENT '语言',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_deleted` tinyint NOT NULL DEFAULT '0' COMMENT '是否已删除',
  PRIMARY KEY (`id`),
  KEY `uuid` (`uuid`),
  KEY `status` (`status`),
  KEY `kefu_id` (`kefu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='游客';

SET FOREIGN_KEY_CHECKS = 1;
