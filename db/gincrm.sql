/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 80032
 Source Host           : localhost:3307
 Source Schema         : gincrm

 Target Server Type    : MySQL
 Target Server Version : 80032
 File Encoding         : 65001

 Date: 18/10/2024 14:32:15
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for contracts
-- ----------------------------
DROP TABLE IF EXISTS `contracts`;
CREATE TABLE `contracts`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `amount` double NULL DEFAULT NULL,
  `begin_time` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `over_time` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `remarks` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `cid` bigint NULL DEFAULT NULL,
  `product_list` json NULL,
  `status` bigint NULL DEFAULT NULL,
  `creator` bigint NULL DEFAULT NULL,
  `created` bigint NULL DEFAULT NULL,
  `updated` bigint NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of contracts
-- ----------------------------
INSERT INTO `contracts` VALUES (4, 'JetBrainss全家桶采购', 1230, '2024-10-17', '2024-10-31', 'JetBrainss全家桶采购3年license', 1, '[{\"id\": 2, \"name\": \"华为云服务\", \"type\": 1, \"unit\": \"个\", \"count\": 10, \"price\": 123, \"total\": 1230}]', 1, 1, 1729133185, 0);
INSERT INTO `contracts` VALUES (6, '数据分析项目', 3330, '2024-10-18', '2024-10-31', '数据分析项目', 1, '[{\"id\": 3, \"name\": \"电商推荐系统\", \"type\": 1, \"unit\": \"个\", \"count\": 10, \"price\": 333, \"total\": 3330}]', 1, 1, 1729223049, 0);

-- ----------------------------
-- Table structure for customers
-- ----------------------------
DROP TABLE IF EXISTS `customers`;
CREATE TABLE `customers`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `source` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `phone` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `email` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `industry` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `level` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `remarks` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `region` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `address` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `status` bigint NULL DEFAULT NULL,
  `creator` bigint NULL DEFAULT NULL,
  `created` bigint NULL DEFAULT NULL,
  `updated` bigint NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of customers
-- ----------------------------
INSERT INTO `customers` VALUES (1, '杨达成', '促销', '18456789111', '3285844058@qq.com', '互联网', '重点客户', '重点客户，客户是基础服务公司', '广东省,深圳市,南山区', '深圳湾', 1, 1, 1718851683, 1729129835);
INSERT INTO `customers` VALUES (2, '华为', '搜索引擎', '15512345623', 'hw@qq.com', '互联网', '重点客户', '234', '广东省,东莞市', '塘朗', 1, 1, 1718855208, 0);
INSERT INTO `customers` VALUES (4, '张三', '广告', '', '', '房地产', '普通客户', '', '', '', 1, 1, 1729223147, 0);
INSERT INTO `customers` VALUES (5, '李四', '线上注册', '', '', '物流运输', '', '', '', '', 1, 1, 1729223174, 0);

-- ----------------------------
-- Table structure for mail_configs
-- ----------------------------
DROP TABLE IF EXISTS `mail_configs`;
CREATE TABLE `mail_configs`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `stmp` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `port` bigint NULL DEFAULT NULL,
  `auth_code` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `email` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `status` bigint NULL DEFAULT NULL,
  `creator` bigint NULL DEFAULT NULL,
  `created` bigint NULL DEFAULT NULL,
  `updated` bigint NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of mail_configs
-- ----------------------------
INSERT INTO `mail_configs` VALUES (2, 'smtp.qq.com', 465, '666', '3285844058@qq.com', 1, 1, 1720581412, 1728982936);

-- ----------------------------
-- Table structure for notices
-- ----------------------------
DROP TABLE IF EXISTS `notices`;
CREATE TABLE `notices`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `status` bigint NULL DEFAULT NULL,
  `creator` bigint NULL DEFAULT NULL,
  `created` bigint NULL DEFAULT NULL,
  `updated` bigint NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 41 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of notices
-- ----------------------------
INSERT INTO `notices` VALUES (1, '你登录了账号', 2, 1, 1718263242, 0);
INSERT INTO `notices` VALUES (2, '你登录了账号', 2, 1, 1718605642, 0);
INSERT INTO `notices` VALUES (3, '你登录了账号', 2, 1, 1718609652, 0);
INSERT INTO `notices` VALUES (4, '你登录了账号', 2, 1, 1718617126, 0);
INSERT INTO `notices` VALUES (5, '你登录了账号', 2, 1, 1718680424, 0);
INSERT INTO `notices` VALUES (6, '你登录了账号', 2, 1, 1718766071, 0);
INSERT INTO `notices` VALUES (7, '你登录了账号', 2, 1, 1718849901, 0);
INSERT INTO `notices` VALUES (8, '你登录了账号', 2, 1, 1718959310, 0);
INSERT INTO `notices` VALUES (9, '你登录了账号', 2, 1, 1719458476, 0);
INSERT INTO `notices` VALUES (10, '你登录了账号', 2, 1, 1720518072, 0);
INSERT INTO `notices` VALUES (11, '你登录了账号', 2, 1, 1720518079, 0);
INSERT INTO `notices` VALUES (12, '你登录了账号', 2, 1, 1720518479, 0);
INSERT INTO `notices` VALUES (13, '你登录了账号', 2, 1, 1720518483, 0);
INSERT INTO `notices` VALUES (14, '你登录了账号', 2, 1, 1720577033, 0);
INSERT INTO `notices` VALUES (15, '你登录了账号', 2, 1, 1720680809, 0);
INSERT INTO `notices` VALUES (16, '你登录了账号', 2, 1, 1720770833, 0);
INSERT INTO `notices` VALUES (17, '你登录了账号', 2, 1, 1723540135, 0);
INSERT INTO `notices` VALUES (18, '你登录了账号', 2, 1, 1723603628, 0);
INSERT INTO `notices` VALUES (19, '你登录了账号', 2, 1, 1728550658, 0);
INSERT INTO `notices` VALUES (20, '你登录了账号', 2, 1, 1728611064, 0);
INSERT INTO `notices` VALUES (21, '你登录了账号', 2, 1, 1728611481, 0);
INSERT INTO `notices` VALUES (22, '你登录了账号', 2, 1, 1728611487, 0);
INSERT INTO `notices` VALUES (23, '你登录了账号', 2, 1, 1728611542, 0);
INSERT INTO `notices` VALUES (24, '你登录了账号', 2, 1, 1728611721, 0);
INSERT INTO `notices` VALUES (25, '你登录了账号', 2, 1, 1728629574, 0);
INSERT INTO `notices` VALUES (26, '你登录了账号', 2, 1, 1728704946, 0);
INSERT INTO `notices` VALUES (27, '你登录了账号', 2, 1, 1728720530, 0);
INSERT INTO `notices` VALUES (28, '你登录了账号', 2, 1, 1728724305, 0);
INSERT INTO `notices` VALUES (29, '你登录了账号', 2, 1, 1728871611, 0);
INSERT INTO `notices` VALUES (30, '你登录了账号', 2, 1, 1728890170, 0);
INSERT INTO `notices` VALUES (31, '你登录了账号', 2, 1, 1728974958, 0);
INSERT INTO `notices` VALUES (32, '你登录了账号', 2, 1, 1728975950, 0);
INSERT INTO `notices` VALUES (33, '你登录了账号', 2, 1, 1728975986, 0);
INSERT INTO `notices` VALUES (34, '你登录了账号', 2, 1, 1728976328, 0);
INSERT INTO `notices` VALUES (35, '你登录了账号', 2, 1, 1728977068, 0);
INSERT INTO `notices` VALUES (36, '你登录了账号', 2, 1, 1728977767, 0);
INSERT INTO `notices` VALUES (37, '你订阅了专业版', 2, 1, 1728978141, 0);
INSERT INTO `notices` VALUES (38, '你登录了账号', 2, 1, 1728981718, 0);
INSERT INTO `notices` VALUES (39, '你订阅了专业版', 2, 1, 1728981821, 0);
INSERT INTO `notices` VALUES (40, '你登录了账号', 2, 1, 1729048454, 0);
INSERT INTO `notices` VALUES (41, '你登录了账号', 2, 1, 1729221493, 0);

-- ----------------------------
-- Table structure for products
-- ----------------------------
DROP TABLE IF EXISTS `products`;
CREATE TABLE `products`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `type` bigint NULL DEFAULT NULL,
  `unit` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `code` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `price` double NULL DEFAULT NULL,
  `description` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `status` bigint NULL DEFAULT NULL,
  `creator` bigint NULL DEFAULT NULL,
  `created` bigint NULL DEFAULT NULL,
  `updated` bigint NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of products
-- ----------------------------
INSERT INTO `products` VALUES (1, 'NVIDIA A100 Tensor Core GPU', 1, '个', '10001', 111, '可针对 AI、数据分析和 HPC 应用场景，在不同规模下实现出色的加速，有效助力更高性能的弹性数据中心。', 1, 1, 1723603912, 1729130043);
INSERT INTO `products` VALUES (2, '华为云服务', 1, '个', '10002', 123, '华为云弹性服务器', 1, 1, 1723604275, 1729130124);
INSERT INTO `products` VALUES (3, '电商推荐系统', 1, '个', '10003', 333, '电商推荐系统，精准匹配人群，投送广告', 2, 1, 1723618677, 1729130200);

-- ----------------------------
-- Table structure for subscribes
-- ----------------------------
DROP TABLE IF EXISTS `subscribes`;
CREATE TABLE `subscribes`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `uid` bigint NULL DEFAULT NULL,
  `version` bigint NULL DEFAULT NULL,
  `expired` bigint NULL DEFAULT NULL,
  `created` bigint NULL DEFAULT NULL,
  `updated` bigint NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of subscribes
-- ----------------------------
INSERT INTO `subscribes` VALUES (1, 1, 2, 1787024176, 1718768176, 1728981814);

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `email` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `password` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `status` bigint NULL DEFAULT NULL,
  `created` bigint NULL DEFAULT NULL,
  `updated` bigint NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, '', '3285844058@qq.com', '$2a$10$Mk.wh0oGxj5JdNOn3LvySOyOOTf0Oi/EdB6IMTwS03VW0vBYzCcW.', 1, 1718261941, 0);

SET FOREIGN_KEY_CHECKS = 1;
