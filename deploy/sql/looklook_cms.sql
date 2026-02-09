-- ----------------------------
-- Chat2DB export data , export time: 2026-02-06 14:50:29
-- ----------------------------
SET
FOREIGN_KEY_CHECKS=0;
-- ----------------------------
-- Table structure for table article
-- ----------------------------
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article`
(
    `id`           bigint                                                        NOT NULL AUTO_INCREMENT COMMENT '文章id',
    `create_time`  datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`  datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `delete_time`  datetime                                                               DEFAULT NULL COMMENT '删除时间',
    `del_state`    tinyint                                                       NOT NULL DEFAULT '0' COMMENT '删除状态 0:未删除 1:已删除',
    `version`      bigint                                                        NOT NULL DEFAULT '0' COMMENT '版本号',
    `title`        varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '文章标题',
    `content`      text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '文章内容',
    `publish_time` datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '文章发布时间',
    `category`     varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci  NOT NULL DEFAULT '' COMMENT '文章分类',
    `like_count`   bigint                                                        NOT NULL DEFAULT '0' COMMENT '文章点赞数量',
    `author_id`    bigint                                                        NOT NULL DEFAULT '0' COMMENT '作者id',
    `cover_image`  varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '封面图片',
    `status`       tinyint                                                       NOT NULL DEFAULT '0' COMMENT '状态 0:草稿 1:已发布 2:已下架',
    PRIMARY KEY (`id`),
    KEY            `idx_category` (`category`),
    KEY            `idx_author_id` (`author_id`),
    KEY            `idx_publish_time` (`publish_time`),
    KEY            `idx_status` (`status`)
) ENGINE=InnoDB AUTO_INCREMENT=118 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='文章表';

-- ----------------------------
-- Table structure for table t_pri_clearing_data_general_yunnan
-- ----------------------------
DROP TABLE IF EXISTS `t_pri_clearing_data_general_yunnan`;
CREATE TABLE `t_pri_clearing_data_general_yunnan`
(
    `id`                       int NOT NULL AUTO_INCREMENT COMMENT '主键',
    `province_id`              int                                                           DEFAULT NULL COMMENT '省份id',
    `company_id`               varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '公司id',
    `company_name`             varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci  DEFAULT NULL COMMENT '公司名称',
    `unit_id`                  varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci  DEFAULT NULL COMMENT '单元或者机组单元id',
    `unit_name`                varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '单元或者机组单元名称',
    `target_date`              date                                                          DEFAULT NULL COMMENT '标的日期',
    `timeperiod`               varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci  DEFAULT NULL COMMENT '时段',
    `dayahead_clearing_energy` decimal(10, 3)                                                DEFAULT NULL COMMENT '日前出清电量（MWh）',
    `dayahead_clearing_price`  decimal(10, 3)                                                DEFAULT NULL COMMENT '日前出清价格（元/MWh）',
    `realtime_clearing_energy` decimal(10, 3)                                                DEFAULT NULL COMMENT '实时出清电量（MWh）',
    `realtime_clearing_price`  decimal(10, 3)                                                DEFAULT NULL COMMENT '实时出清价格（元/MWh）',
    `create_time`              datetime                                                      DEFAULT (now()) COMMENT '创建时间',
    `update_time`              datetime                                                      DEFAULT (now()) ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `dayahead_bid_power`       decimal(10, 3)                                                DEFAULT NULL COMMENT '日前中标出力(MW)',
    `realtime_bid_power`       decimal(10, 3)                                                DEFAULT NULL COMMENT '实时中标出力(MW)',
    `actual_energy`            decimal(10, 3)                                                DEFAULT NULL COMMENT '实际上网电量（MWh）',
    `actual_power`             decimal(10, 3)                                                DEFAULT NULL COMMENT '实际发电出力（MW）',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_uniq` (`province_id`,`company_id`,`target_date`,`timeperiod`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=7396 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='私有-出清电量电价-通用表';

SET
FOREIGN_KEY_CHECKS=1;
