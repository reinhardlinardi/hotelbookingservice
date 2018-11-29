-- Adminer 4.6.3 MySQL dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

USE `hotel_booking`;

SET NAMES utf8mb4;

CREATE TABLE `agent` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `address` varchar(100) NOT NULL,
  `expire_date` date NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

INSERT INTO `agent` (`id`, `name`, `address`, `expire_date`) VALUES
(1,	'Traveloka',	'Jakarta',	'2020-01-01'),
(2,	'Airy',	'Jakarta',	'2021-01-01');

CREATE TABLE `customer` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `identity` varchar(100) NOT NULL,
  `email` varchar(100) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

INSERT INTO `customer` (`id`, `name`, `identity`, `email`) VALUES
(1,	'Cisco',	'123412341234',	'cisco@coki.com'),
(2,	'Roland',	'112233445566',	'roland@avatar.com');

CREATE TABLE `employee` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `address` varchar(100) NOT NULL,
  `job_desc` varchar(100) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

INSERT INTO `employee` (`id`, `name`, `address`, `job_desc`) VALUES
(1,	'Dono',	'Jakarta',	'Membersihkan semua kamar di lantai 2');

CREATE TABLE `invoice` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `room_id` int(11) NOT NULL,
  `customer_id` int(11) NOT NULL,
  `in_date` date NOT NULL,
  `out_date` date NOT NULL,
  `price` int(11) NOT NULL,
  `paid` tinyint(4) NOT NULL,
  `cancelled` tinyint(4) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `room_id` (`room_id`),
  KEY `customer_id` (`customer_id`),
  CONSTRAINT `invoice_ibfk_1` FOREIGN KEY (`room_id`) REFERENCES `room` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `invoice_ibfk_2` FOREIGN KEY (`customer_id`) REFERENCES `customer` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


CREATE TABLE `room` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `type` varchar(100) NOT NULL,
  `tv` tinyint(4) NOT NULL,
  `ac` tinyint(4) NOT NULL,
  `internet` tinyint(4) NOT NULL,
  `water` tinyint(4) NOT NULL,
  `refrigerator` tinyint(4) NOT NULL,
  `deposit_box` tinyint(4) NOT NULL,
  `wardrobe` tinyint(4) NOT NULL,
  `window` tinyint(4) NOT NULL,
  `balcony` tinyint(4) NOT NULL,
  `price` int(11) NOT NULL,
  `description` varchar(200) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

INSERT INTO `room` (`id`, `type`, `tv`, `ac`, `internet`, `water`, `refrigerator`, `deposit_box`, `wardrobe`, `window`, `balcony`, `price`, `description`) VALUES
(1,	'Single',	1,	1,	1,	1,	1,	1,	1,	0,	0,	200000,	'Kamar untuk satu orang dengan single bed yang luas dan nyaman serta fasilitas yang lengkap.'),
(2,	'Double',	1,	1,	1,	1,	1,	1,	1,	1,	0,	400000,	'Kamar untuk dua orang dengan double bed yang luas dan nyaman, dengan fasilitas lengkap dan pemandangan langsung ke kolam renang.'),
(3,	'Family',	1,	1,	1,	1,	1,	1,	1,	1,	1,	800000,	'Kamar untuk satu keluarga yang luas dan nyaman, dengan fasilitas lengkap, sliding window, dan private balcony.'),
(10,	'Single',	1,	1,	1,	1,	1,	1,	1,	0,	0,	200000,	'Kamar single bed yang luas dengan fasilitas yang lengkap.'),
(11,	'Single',	1,	1,	1,	1,	1,	1,	1,	0,	0,	200000,	'Kamar single bed yang luas dengan fasilitas yang lengkap.'),
(12,	'Double',	1,	1,	1,	1,	1,	1,	1,	1,	0,	400000,	'Kamar untuk dua orang dengan double bed yang luas dan nyaman, dengan fasilitas lengkap dan pemandangan langsung ke kolam renang.'),
(13,	'Family',	1,	1,	1,	1,	1,	1,	1,	1,	1,	800000,	'Kamar untuk satu keluarga yang luas dan nyaman, dengan fasilitas lengkap, sliding window, dan private balcony.'),
(14,	'Double',	1,	1,	1,	1,	1,	1,	1,	1,	0,	400000,	'Kamar untuk dua orang dengan double bed yang luas dan nyaman, dengan fasilitas lengkap dan pemandangan langsung ke kolam renang.'),
(15,	'Single',	1,	1,	1,	1,	1,	1,	1,	0,	0,	200000,	'Kamar single bed yang luas dengan fasilitas yang lengkap.');

-- 2018-11-30 10:39:14
