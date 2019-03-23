-- phpMyAdmin SQL Dump
-- version 4.5.2
-- http://www.phpmyadmin.net
--
-- Host: localhost
-- Generation Time: Mar 23, 2019 at 07:30 
-- Server version: 10.1.13-MariaDB
-- PHP Version: 5.6.20

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `crr`
--

-- --------------------------------------------------------

--
-- Table structure for table `message`
--

CREATE TABLE `message` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `message` text CHARACTER SET utf8 COLLATE utf8_polish_ci NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `message`
--

INSERT INTO `message` (`id`, `user_id`, `message`, `created_at`) VALUES
(1, 2, 'Siema!', '2019-03-23 10:46:30'),
(2, 3, 'Hejo!', '2019-03-23 10:46:30'),
(3, 2, 'Co tam?', '2019-03-23 18:29:48'),
(4, 2, 'Halo', '2019-03-23 18:29:52'),
(5, 2, 'Odbjuuur?', '2019-03-23 18:29:56'),
(6, 2, 'Ejjj', '2019-03-23 18:29:59');

-- --------------------------------------------------------

--
-- Table structure for table `user`
--

CREATE TABLE `user` (
  `id` int(11) NOT NULL,
  `name` varchar(50) CHARACTER SET utf8 COLLATE utf8_polish_ci NOT NULL,
  `email` varchar(128) NOT NULL,
  `password` varchar(256) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `last_visit` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `user`
--

INSERT INTO `user` (`id`, `name`, `email`, `password`, `created_at`, `last_visit`) VALUES
(1, 'Tomasz Dzielnik', 'jakismail@gmail.com', '$2a$10$03i9sOrRV94fyOMtu72mIuJ1uC.j62NKnfoKIBuuv3/1VyOvOW8KW', '2019-03-22 16:40:53', '2019-03-22 16:40:53'),
(2, 'Tomasz Dzielnik', 'jakiaaasmail@gmail.com', '$2a$10$d4Y8SSozliKIGST8hI1n5.G6aur0FXOzZj4WBlxEEvS5OHKw78KGK', '2019-03-21 15:07:00', '2019-03-21 15:07:00'),
(3, 'Tomasz Dzielnik', 'randacc@gmail.com', '$2a$10$7.CcbNlrfqAy8dkbk1eSvuE9s5O42vRvt.wxRROoaKV8N6DkIEYgO', '2019-03-21 16:12:24', '2019-03-21 16:12:24'),
(4, 'Tomasz Dzielnik', 'raasdasdasdndacc@gmail.com', '$2a$10$9dOfCBbyvpeSZWaQubEgvu69b4k.zF84KCn0dteeE9VYczRRxnpJy', '2019-03-23 18:12:13', '2019-03-23 18:12:13');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `message`
--
ALTER TABLE `message`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_id` (`user_id`);

--
-- Indexes for table `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `email` (`email`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `message`
--
ALTER TABLE `message`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=16;
--
-- AUTO_INCREMENT for table `user`
--
ALTER TABLE `user`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;
--
-- Constraints for dumped tables
--

--
-- Constraints for table `message`
--
ALTER TABLE `message`
  ADD CONSTRAINT `author` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
