/*
 Navicat MySQL Data Transfer

 Source Server         : looklook
 Source Server Type    : MySQL
 Source Server Version : 80028
 Source Host           : 127.0.0.1:33069
 Source Schema         : looklook_cms

 Target Server Type    : MySQL
 Target Server Version : 80028
 File Encoding         : 65001

 Date: 31/01/2026 10:00:00
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for article
-- ----------------------------
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '文章id',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `delete_time` datetime DEFAULT NULL COMMENT '删除时间',
  `del_state` tinyint NOT NULL DEFAULT '0' COMMENT '删除状态 0:未删除 1:已删除',
  `version` bigint NOT NULL DEFAULT '0' COMMENT '版本号',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '文章标题',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '文章内容',
  `publish_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '文章发布时间',
  `category` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '文章分类',
  `like_count` bigint NOT NULL DEFAULT '0' COMMENT '文章点赞数量',
  `author_id` bigint NOT NULL DEFAULT '0' COMMENT '作者id',
  `cover_image` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '封面图片',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态 0:草稿 1:已发布 2:已下架',
  PRIMARY KEY (`id`),
  KEY `idx_category` (`category`),
  KEY `idx_author_id` (`author_id`),
  KEY `idx_publish_time` (`publish_time`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='文章表';

SET FOREIGN_KEY_CHECKS = 1;
