-- phpMyAdmin SQL Dump
-- version 4.8.5
-- https://www.phpmyadmin.net/
--
-- 主机： localhost
-- 生成日期： 2020-12-05 23:25:12
-- 服务器版本： 8.0.12
-- PHP 版本： 7.3.4

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- 数据库： `imdb`
--

-- --------------------------------------------------------

--
-- 表的结构 `groups`
--

CREATE TABLE `groups` (
  `id` bigint(20) NOT NULL COMMENT '主键',
  `name` varchar(50) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '名称',
  `appid` tinyint(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '应用id',
  `state` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态(1:有效，0无效，9删除)',
  `ctime` int(10) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `mtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '修改时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `groups_member`
--

CREATE TABLE `groups_member` (
  `id` bigint(20) NOT NULL COMMENT '主键',
  `groups_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '群id',
  `uid` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户id',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '类型（0:普通1群主）',
  `ctime` int(10) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `mtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '修改时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `groups_message`
--

CREATE TABLE `groups_message` (
  `id` int(11) NOT NULL,
  `msg_ser_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '服务端消息id',
  `uid` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户id',
  `readed` tinyint(4) NOT NULL DEFAULT '0' COMMENT '读取状态（0:未读,1:已读）',
  `mtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '修改时间',
  `ctime` int(10) NOT NULL DEFAULT '0' COMMENT '创建时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `member`
--

CREATE TABLE `member` (
  `id` bigint(20) NOT NULL,
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '姓名',
  `state` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态(1:有效，0无效，9删除)',
  `ctime` int(10) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `mtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- 表的结构 `message`
--

CREATE TABLE `message` (
  `id` bigint(20) UNSIGNED NOT NULL COMMENT 'ID',
  `msg_ser_id` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '服务端消息id',
  `groups_id` int(10) DEFAULT '0' COMMENT '群id',
  `from_uid` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '发送方UID',
  `to_uid` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '接收方UID',
  `content` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '内容',
  `type` tinyint(1) UNSIGNED NOT NULL DEFAULT '0' COMMENT '消息类型',
  `appname` tinyint(3) UNSIGNED NOT NULL DEFAULT '0' COMMENT '应用名称',
  `send` tinyint(1) UNSIGNED NOT NULL DEFAULT '0' COMMENT '发送状态',
  `readed` tinyint(1) UNSIGNED NOT NULL DEFAULT '0' COMMENT '读取状态',
  `ext` varchar(1024) DEFAULT '' COMMENT '扩展字段',
  `state` tinyint(1) UNSIGNED NOT NULL DEFAULT '0' COMMENT '状态',
  `admin_id` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '审核人',
  `ctime` int(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建时间',
  `mtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='消息表' ROW_FORMAT=COMPACT;

--
-- 转储表的索引
--

--
-- 表的索引 `groups`
--
ALTER TABLE `groups`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `groups_member`
--
ALTER TABLE `groups_member`
  ADD PRIMARY KEY (`id`),
  ADD KEY `uid` (`uid`);

--
-- 表的索引 `groups_message`
--
ALTER TABLE `groups_message`
  ADD PRIMARY KEY (`id`),
  ADD KEY `msg_ser_id` (`msg_ser_id`,`uid`) USING BTREE;

--
-- 表的索引 `member`
--
ALTER TABLE `member`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `message`
--
ALTER TABLE `message`
  ADD KEY `msg_ser_id` (`msg_ser_id`),
  ADD KEY `groups_id` (`groups_id`),
  ADD KEY `from_uid` (`from_uid`),
  ADD KEY `to_uid` (`to_uid`);

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `groups`
--
ALTER TABLE `groups`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键';

--
-- 使用表AUTO_INCREMENT `groups_member`
--
ALTER TABLE `groups_member`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键';

--
-- 使用表AUTO_INCREMENT `groups_message`
--
ALTER TABLE `groups_message`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `member`
--
ALTER TABLE `member`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
