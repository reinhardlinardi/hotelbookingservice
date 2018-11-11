-- phpMyAdmin SQL Dump
-- version 4.5.1
-- http://www.phpmyadmin.net
--
-- Host: 127.0.0.1
-- Generation Time: Nov 11, 2018 at 04:47 PM
-- Server version: 10.1.19-MariaDB
-- PHP Version: 5.6.28

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `hotel_booking`
--

-- --------------------------------------------------------

--
-- Table structure for table `agent`
--

CREATE TABLE `agent` (
  `id` int(11) NOT NULL,
  `name` varchar(100) NOT NULL,
  `address` varchar(100) NOT NULL,
  `expire_date` date NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `customer`
--

CREATE TABLE `customer` (
  `id` int(11) NOT NULL,
  `name` varchar(100) NOT NULL,
  `identity` varchar(100) NOT NULL,
  `email` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `customer`
--

INSERT INTO `customer` (`id`, `name`, `identity`, `email`) VALUES
(1, 'Francisco', 'HUEHUEHUE', 'francisco@example.com'),
(2, 'Tasya', '12345', 'tasya@example.com'),
(3, 'reinhard', '12345', 'reinhard@example.com'),
(4, 'ReinhardCarry', '12345', 'reinhardcarry@example.com'),
(5, 'PPLS MATA PANCING', '12345', 'ppls@example.com');

-- --------------------------------------------------------

--
-- Table structure for table `employee`
--

CREATE TABLE `employee` (
  `id` int(11) NOT NULL,
  `name` varchar(100) NOT NULL,
  `address` varchar(100) NOT NULL,
  `job_desc` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `invoice`
--

CREATE TABLE `invoice` (
  `id` int(11) NOT NULL,
  `room_id` int(11) NOT NULL,
  `customer_id` int(11) NOT NULL,
  `in_date` date NOT NULL,
  `out_date` date NOT NULL,
  `price` int(11) NOT NULL,
  `paid` tinyint(4) NOT NULL,
  `cancelled` tinyint(4) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `invoice`
--

INSERT INTO `invoice` (`id`, `room_id`, `customer_id`, `in_date`, `out_date`, `price`, `paid`, `cancelled`) VALUES
(4, 1, 1, '2018-11-21', '2018-11-22', 500000, 1, 0),
(5, 2, 1, '2018-11-12', '2018-11-13', 200000, 1, 0),
(6, 3, 1, '2018-11-10', '2018-11-11', 1000000, 0, 1),
(7, 3, 4, '2018-11-14', '2018-11-15', 800000, 0, 0),
(8, 3, 4, '2018-11-14', '2018-11-15', 800000, 0, 0),
(9, 3, 4, '2018-11-14', '2018-11-15', 800000, 0, 0),
(10, 3, 5, '2018-11-14', '2018-11-15', 800000, 0, 0),
(11, 3, 5, '2018-11-14', '2018-11-15', 800000, 0, 0),
(12, 3, 5, '2018-11-14', '2018-11-15', 800000, 0, 0),
(13, 3, 5, '2018-11-14', '2018-11-15', 800000, 0, 0),
(14, 3, 1, '2018-11-20', '2018-11-21', 800000, 0, 0),
(15, 3, 1, '2018-11-20', '2018-11-21', 800000, 0, 0),
(16, 3, 1, '2018-11-20', '2018-11-21', 800000, 0, 0),
(17, 3, 1, '2018-11-20', '2018-11-21', 800000, 0, 0),
(18, 3, 3, '2018-11-25', '2018-11-26', 800000, 0, 0);

-- --------------------------------------------------------

--
-- Table structure for table `room`
--

CREATE TABLE `room` (
  `id` int(11) NOT NULL,
  `type` varchar(100) NOT NULL,
  `description` varchar(200) NOT NULL,
  `tv` tinyint(4) NOT NULL,
  `ac` tinyint(4) NOT NULL,
  `internet` tinyint(4) NOT NULL,
  `water` tinyint(4) NOT NULL,
  `refrigerator` tinyint(4) NOT NULL,
  `deposit_box` tinyint(4) NOT NULL,
  `wardrobe` tinyint(4) NOT NULL,
  `window` tinyint(4) NOT NULL,
  `balcony` tinyint(4) NOT NULL,
  `price` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `room`
--

INSERT INTO `room` (`id`, `type`, `description`, `tv`, `ac`, `internet`, `water`, `refrigerator`, `deposit_box`, `wardrobe`, `window`, `balcony`, `price`) VALUES
(1, 'Single', 'Kamar untuk satu orang dengan single bed yang luas dan nyaman serta fasilitas yang lengkap.', 1, 1, 1, 1, 1, 1, 1, 0, 0, 200000),
(2, 'Double', 'Kamar untuk dua orang dengan double bed yang luas dan nyaman, dengan fasilitas lengkap dan pemandangan langsung ke kolam renang.', 1, 1, 1, 1, 1, 1, 1, 1, 0, 400000),
(3, 'Family', 'Kamar untuk satu keluarga yang luas dan nyaman, dengan fasilitas lengkap, sliding window, dan private balcony.', 1, 1, 1, 1, 1, 1, 1, 1, 1, 800000);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `agent`
--
ALTER TABLE `agent`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `customer`
--
ALTER TABLE `customer`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `employee`
--
ALTER TABLE `employee`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `invoice`
--
ALTER TABLE `invoice`
  ADD PRIMARY KEY (`id`),
  ADD KEY `room_id` (`room_id`),
  ADD KEY `customer_id` (`customer_id`);

--
-- Indexes for table `room`
--
ALTER TABLE `room`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `agent`
--
ALTER TABLE `agent`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT for table `customer`
--
ALTER TABLE `customer`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;
--
-- AUTO_INCREMENT for table `employee`
--
ALTER TABLE `employee`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT for table `invoice`
--
ALTER TABLE `invoice`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=19;
--
-- AUTO_INCREMENT for table `room`
--
ALTER TABLE `room`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;
--
-- Constraints for dumped tables
--

--
-- Constraints for table `invoice`
--
ALTER TABLE `invoice`
  ADD CONSTRAINT `invoice_ibfk_1` FOREIGN KEY (`room_id`) REFERENCES `room` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `invoice_ibfk_2` FOREIGN KEY (`customer_id`) REFERENCES `customer` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
