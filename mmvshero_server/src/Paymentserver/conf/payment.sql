-- phpMyAdmin SQL Dump
-- version 4.6.4
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: 2016-10-24 09:16:28
-- 服务器版本： 5.7.15
-- PHP Version: 5.6.24

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";

CREATE DATABASE payment DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;

use payment;


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `payment`
--

-- --------------------------------------------------------

--
-- 表的结构 `goods`
--

CREATE TABLE `goods` (
  `id` int(11) NOT NULL,
  `name` varchar(50) NOT NULL COMMENT '商品名称',
  `description` varchar(100) NOT NULL COMMENT '商品描述',
  `price` double NOT NULL COMMENT '价格',
  `goods_url` varchar(200) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- 转存表中的数据 `goods`
--

INSERT INTO `goods` (`id`, `name`, `description`, `price`, `goods_url`) VALUES
(1, '300魔钻会员', '连续30天每天可领120魔钻', 25, 'http://www.qq.coom'),
(2, '6480魔钻', '另赠6480魔钻（首次购买赠送）', 648, 'http://www.qq.com'),
(3, '3280魔钻会员', '另赠3280魔钻（首次购买赠送）', 328, 'http://www.qq.com'),
(4, '1980魔钻', '另赠1980魔钻（首次购买赠送）', 198, 'http://www.qq.com'),
(5, '980魔钻', '另赠980魔钻（首次购买赠送）', 98, 'http://www.qq.com'),
(6, '300魔钻', '另赠300魔钻（首次购买赠送）', 30, 'http://www.qq.com'),
(7, '60魔钻', '另赠60魔钻（首次购买赠送）', 6, 'http://www.qq.com');

-- --------------------------------------------------------

--
-- 表的结构 `orders`
--

CREATE TABLE `orders` (
  `id` int(11) NOT NULL,
  `order_id` varchar(50) NOT NULL COMMENT '订单编号',
  `player_id` varchar(50) NOT NULL COMMENT '角色编号',
  `goods_id` int(11) NOT NULL COMMENT '商品编号',
  `state` int(11) NOT NULL COMMENT '订单状态（0表示初始化，1表示支付成功，2表示通知游戏服务器成功）',
  `pay_channel` varchar(50) NOT NULL COMMENT '支付渠道',
  `pay_order_id` varchar(50) DEFAULT NULL COMMENT '支付订单编号',
  `total_price` double NOT NULL COMMENT '总价',
  `zone_id` varchar(20) NOT NULL COMMENT '服务器编号',
  `token` varchar(1024) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- 转存表中的数据 `orders`
--

INSERT INTO `orders` (`id`, `order_id`, `player_id`, `goods_id`, `state`, `pay_channel`, `pay_order_id`, `total_price`, `zone_id`, `token`) VALUES
(60, '20161020143608733646', '78668', 4, 0, 'midas', '5913B1EC110FF5F27E8C833E476E057F26373', 25, '1', ''),
(61, '2016102014383725863', '78668', 5, 0, 'midas', '621F7AE6257656C039FDB00F98D9593226373', 25, '1', ''),
(62, '20161020144302810145', '78668', 6, 0, 'midas', 'C8C06A31E438551CBBD475D98B30782D26373', 25, '1', ''),
(63, '20161020144447743467', '78668', 3, 0, 'midas', 'C9EE08B95D096F4524B187798E11E1BB26373', 25, '1', ''),
(64, '20161020144447799438', '78668', 2, 0, 'midas', '1C6FA4533C94AA1BA3E3C517768FD54226373', 25, '1', ''),
(65, '20161020144542221062', '78668', 5, 0, 'midas', '39DBEEF15D9722187594CDDF254F180C26373', 25, '1', ''),
(66, '20161020145902325404', '78581', 5, 0, 'midas', '5C32A4267BBB2DED3A013E0D2302289026373', 25, '1', ''),
(67, '20161020150237502235', '78581', 5, 0, 'midas', '7471762F4A8C170BA0E4EDC397223FB726373', 25, '1', ''),
(68, '2016102015162950522', '78581', 5, 0, 'midas', 'A69F440373752E47012391C250BB32FC26373', 25, '1', ''),
(69, '20161020152645265863', '78581', 1, 1, 'midas', 'D08A16F5373681C05F4B33EA5815459626373', 25, '1', '');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `goods`
--
ALTER TABLE `goods`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `orders`
--
ALTER TABLE `orders`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `id` (`id`);

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `goods`
--
ALTER TABLE `goods`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;
--
-- 使用表AUTO_INCREMENT `orders`
--
ALTER TABLE `orders`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=70;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
