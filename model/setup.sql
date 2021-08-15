CREATE DATABASE IF NOT EXISTS `fast-gin` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;

USE `fast-gin`;

SET FOREIGN_KEY_CHECKS=0;

DROP TABLE IF EXISTS `blog_article`;
CREATE TABLE `blog_article` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `tag_id` int(10) unsigned DEFAULT '0',
  `title` varchar(100) DEFAULT '',
  `desc` varchar(255) DEFAULT '',
  `content` text,
  `created_on` int(10) unsigned DEFAULT '0',
  `created_by` varchar(100) DEFAULT '',
  `modified_on` int(10) unsigned DEFAULT '0',
  `modified_by` varchar(255) DEFAULT '',
  `deleted_on` int(10) unsigned DEFAULT '0',
  `state` tinyint(3) unsigned DEFAULT '1',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `blog_tag`;
CREATE TABLE `blog_tag` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT '',
  `created_on` int(10) unsigned DEFAULT '0',
  `created_by` varchar(100) DEFAULT '',
  `modified_on` int(10) unsigned DEFAULT '0',
  `modified_by` varchar(100) DEFAULT '',
  `deleted_on` int(10) unsigned DEFAULT '0',
  `state` tinyint(3) unsigned DEFAULT '1',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `user` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(100) NOT NULL DEFAULT '' COMMENT '用户名',
    `state` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '状态',
    `ctime` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `mtime` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';