/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50720
Source Host           : localhost:3306
Source Database       : firmoo

Target Server Type    : MYSQL
Target Server Version : 50720
File Encoding         : 65001

Date: 2018-04-04 18:05:22
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for firmoo_member
-- ----------------------------
DROP TABLE IF EXISTS `firmoo_member`;
CREATE TABLE `firmoo_member` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `real_name` varchar(32) NOT NULL DEFAULT '',
  `user_name` varchar(24) NOT NULL DEFAULT '',
  `user_pwd` varchar(255) NOT NULL DEFAULT '',
  `is_super` tinyint(1) NOT NULL DEFAULT '0',
  `status` int(11) NOT NULL DEFAULT '0',
  `mobile` varchar(16) NOT NULL DEFAULT '',
  `email` varchar(256) NOT NULL DEFAULT '',
  `avatar` varchar(256) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of firmoo_member
-- ----------------------------
INSERT INTO `firmoo_member` VALUES ('1', 'lihaitao', 'admin', 'e10adc3949ba59abbe56e057f20f883e', '0', '1', '18612348765', '123123@qq.com', '/static/upload/1.jpg');

-- ----------------------------
-- Table structure for firmoo_member_token
-- ----------------------------
DROP TABLE IF EXISTS `firmoo_member_token`;
CREATE TABLE `firmoo_member_token` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `member_id` int(11) NOT NULL COMMENT '会员id',
  `token` varchar(50) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `member_id` (`member_id`)
) ENGINE=InnoDB AUTO_INCREMENT=78 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of firmoo_member_token
-- ----------------------------
