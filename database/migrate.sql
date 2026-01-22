-- GCX Database Migration
-- Complete schema from Clever Cloud database
-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Sep 17, 2025 at 11:16 AM
-- Server version: 10.4.32-MariaDB
-- PHP Version: 8.2.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Schema cleaned for Clever Cloud import

-- --------------------------------------------------------

--
-- Table structure for table `blog_categories`
--

CREATE TABLE `blog_categories` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `name` longtext NOT NULL,
  `slug` longtext NOT NULL,
  `description` longtext DEFAULT NULL,
  `color` longtext DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `blog_posts`
--

CREATE TABLE `blog_posts` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `title` longtext NOT NULL,
  `slug` longtext NOT NULL,
  `content` longtext DEFAULT NULL,
  `excerpt` text DEFAULT NULL,
  `featured_image` longtext DEFAULT NULL,
  `tags` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL CHECK (json_valid(`tags`)),
  `status` varchar(191) DEFAULT 'draft',
  `author_id` bigint(20) UNSIGNED DEFAULT NULL,
  `published_at` datetime(3) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `blog_posts`
--

INSERT INTO `blog_posts` (`id`, `title`, `slug`, `content`, `excerpt`, `featured_image`, `tags`, `status`, `author_id`, `published_at`, `created_at`, `updated_at`) VALUES
(2, 'testing again', 'testing-again', '<p>kjhskjdhvv sjhdvfhvlshidv sliudvgsialuvsudi SJDGVYFSUYGCSILUD uksdfygysgalsd IDGFILUDHD ildsuhsoudfhun</p><p></p><p>skdjvsljdvsldij sDILUFGSIDLUKLDVSJS SDILFUGSILDUGDSV ASILDGVUSIOUDAV a8foishfhdsf soduhsujdh.hgh</p>', 'testig so i n=know im okat', NULL, NULL, 'published', 1, '2025-08-18 15:18:53.248', '2025-08-18 13:29:05.544', '2025-08-18 15:18:53.313'),
(3, 'New post ', 'new-post-', '<p>trying to test the upload power so let see</p><p></p><p></p><img src=\"http://localhost:8080/uploads/images/1755531021_66666666.png\" alt=\"Uploaded image\"><p></p><p>💡 Tip: You can resize this image by right-clicking and selecting image options</p><p></p><p><br></p><p></p>', 'trying to test the upload power so let see   \n      \n         \n        \n        💡 Tip: You can resize this image by right-clicking and selecting image options', 'http://localhost:8080/uploads/images/1755530991_iiiiiiii.png', NULL, 'published', 1, '2025-08-18 15:33:14.206', '2025-08-18 15:31:25.486', '2025-08-18 15:33:14.210');

-- --------------------------------------------------------

--
-- Table structure for table `board_members`
--

CREATE TABLE `board_members` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `name` longtext NOT NULL,
  `position` longtext NOT NULL,
  `image` longtext DEFAULT NULL,
  `description` text NOT NULL,
  `linked_in` longtext DEFAULT NULL,
  `facebook` longtext DEFAULT NULL,
  `instagram` longtext DEFAULT NULL,
  `order_index` bigint(20) DEFAULT 0,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `board_members`
--

INSERT INTO `board_members` (`id`, `name`, `position`, `image`, `description`, `linked_in`, `facebook`, `instagram`, `order_index`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'micheal eshun', 'sectary', '/uploads/images/1756992930444273700.jpg', 'testing ', '', '', '', 1, '2025-09-04 13:35:42.124', '2025-09-04 13:35:42.124', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `brokers`
--

CREATE TABLE `brokers` (
  `id` int(11) NOT NULL,
  `name` longtext NOT NULL,
  `company` varchar(255) DEFAULT NULL,
  `license_number` varchar(100) DEFAULT NULL,
  `phone_no` varchar(50) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `address` text DEFAULT NULL,
  `specialization` varchar(255) DEFAULT NULL,
  `experience_years` int(11) DEFAULT 0,
  `status` enum('Active','Inactive','Suspended') DEFAULT 'Active',
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `brokers`
--

INSERT INTO `brokers` (`id`, `name`, `company`, `license_number`, `phone_no`, `email`, `address`, `specialization`, `experience_years`, `status`, `created_at`, `updated_at`) VALUES
(1, 'John Mensah', 'AgriTrade Ltd', 'BRK001', '0244 123 456', 'john.mensah@agritrade.com', 'Accra, Ghana', 'Agricultural Commodities', 5, 'Active', '2025-09-05 12:59:17.000', '2025-09-05 12:59:17.000'),
(2, 'Sarah Osei', 'Commodity Brokers Ghana', 'BRK002', '0244 234 567', 'sarah.osei@cbg.com', 'Kumasi, Ghana', 'Grain Trading', 8, 'Active', '2025-09-05 12:59:17.000', '2025-09-05 12:59:17.000'),
(3, 'Michael Asante', 'West Africa Trading Co', 'BRK003', '0244 345 678', 'michael.asante@watc.com', 'Tema, Ghana', 'Export/Import', 12, 'Active', '2025-09-05 12:59:17.000', '2025-09-05 12:59:17.000'),
(4, 'Grace Adjei', 'FarmLink Brokers', 'BRK004', '0244 456 789', 'grace.adjei@farmlink.com', 'Tamale, Ghana', 'Livestock Trading', 6, 'Active', '2025-09-05 12:59:17.000', '2025-09-05 12:59:17.000'),
(5, 'David Nkrumah', 'Ghana Commodity Brokers', 'BRK005', '0244 567 890', 'david.nkrumah@gcb.com', 'Cape Coast, Ghana', 'Oil Seeds', 10, 'Active', '2025-09-05 12:59:17.000', '2025-09-05 12:59:17.000');

-- --------------------------------------------------------

--
-- Table structure for table `careers`
--

CREATE TABLE `careers` (
  `id` int(11) NOT NULL,
  `title` longtext NOT NULL,
  `description` text DEFAULT NULL,
  `category` enum('Job Openings','Internship','Job Functional Areas') NOT NULL,
  `department` varchar(255) DEFAULT NULL,
  `location` varchar(255) DEFAULT NULL,
  `employment_type` enum('Full-time','Part-time','Contract','Internship') NOT NULL,
  `experience_level` enum('Entry Level','Mid Level','Senior Level','Executive') NOT NULL,
  `requirements` text DEFAULT NULL,
  `responsibilities` text DEFAULT NULL,
  `benefits` text DEFAULT NULL,
  `salary_range` varchar(100) DEFAULT NULL,
  `application_deadline` date DEFAULT NULL,
  `start_date` date DEFAULT NULL,
  `status` enum('Open','Closed','On Hold') DEFAULT 'Open',
  `application_count` int(11) DEFAULT 0,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `careers`
--

INSERT INTO `careers` (`id`, `title`, `description`, `category`, `department`, `location`, `employment_type`, `experience_level`, `requirements`, `responsibilities`, `benefits`, `salary_range`, `application_deadline`, `start_date`, `status`, `application_count`, `created_at`, `updated_at`) VALUES
(1, 'Senior Market Analyst', 'Lead market analysis and research for commodity trading', 'Job Openings', 'Market Research', 'Accra, Ghana', 'Full-time', 'Senior Level', 'Masters in Economics, 5+ years experience, Strong analytical skills', 'Analyze market trends, Prepare reports, Lead research projects', 'Health insurance, Pension, Professional development', 'GHC 8,000 - GHC 12,000', '2025-02-15', '2025-03-01', 'Open', 0, '2025-09-05 13:46:13.000', '2025-09-05 13:46:13.000'),
(2, 'IT Support Intern', 'Support IT operations and system maintenance', 'Internship', 'Information Technology', 'Accra, Ghana', 'Internship', 'Entry Level', 'Computer Science degree, Basic programming knowledge', 'Support users, Maintain systems, Learn new technologies', 'Stipend, Mentorship, Certificate', 'GHC 1,500 - GHC 2,000', '2025-01-31', '2025-02-15', 'Open', 0, '2025-09-05 13:46:13.000', '2025-09-05 13:46:13.000'),
(3, 'Trading Operations', 'Manage day-to-day trading operations and member services', 'Job Functional Areas', 'Trading Operations', 'Accra, Ghana', 'Full-time', 'Mid Level', 'Finance background, Trading experience preferred', 'Execute trades, Manage member accounts, Monitor markets', 'Competitive salary, Performance bonus, Career growth', 'GHC 5,000 - GHC 8,000', '2025-02-28', '2025-03-15', 'Open', 0, '2025-09-05 13:46:13.000', '2025-09-05 13:46:13.000'),
(4, 'Marketing Coordinator', 'Coordinate marketing activities and member engagement', 'Job Openings', 'Marketing', 'Accra, Ghana', 'Full-time', 'Mid Level', 'Marketing degree, 3+ years experience, Creative thinking', 'Develop campaigns, Manage social media, Organize events', 'Health insurance, Flexible hours, Team building', 'GHC 4,000 - GHC 6,000', '2025-02-20', '2025-03-01', 'Open', 0, '2025-09-05 13:46:13.000', '2025-09-05 13:46:13.000'),
(5, 'Finance Intern', 'Assist with financial reporting and analysis', 'Internship', 'Finance', 'Accra, Ghana', 'Internship', 'Entry Level', 'Accounting/Finance degree, Excel proficiency', 'Prepare reports, Assist with audits, Data analysis', 'Stipend, Learning opportunities, Reference letter', 'GHC 1,200 - GHC 1,800', '2025-01-25', '2025-02-01', 'Open', 0, '2025-09-05 13:46:13.000', '2025-09-05 13:46:13.000');

-- --------------------------------------------------------

--
-- Table structure for table `commodities`
--

CREATE TABLE `commodities` (
  `id` int(11) NOT NULL,
  `name` longtext NOT NULL,
  `code` longtext NOT NULL,
  `description` text DEFAULT NULL,
  `specifications` text DEFAULT NULL,
  `trading_hours` varchar(255) DEFAULT NULL,
  `contract_size` varchar(100) DEFAULT NULL,
  `price_unit` varchar(50) DEFAULT NULL,
  `minimum_price` decimal(10,2) DEFAULT NULL,
  `maximum_price` decimal(10,2) DEFAULT NULL,
  `current_price` decimal(10,2) DEFAULT NULL,
  `price_change` decimal(10,2) DEFAULT NULL,
  `price_change_percent` decimal(5,2) DEFAULT NULL,
  `trading_volume` bigint(20) DEFAULT 0,
  `market_status` enum('Open','Closed','Suspended') DEFAULT 'Open',
  `image_path` varchar(500) DEFAULT NULL,
  `category` varchar(100) DEFAULT NULL,
  `origin_country` varchar(100) DEFAULT NULL,
  `harvest_season` varchar(100) DEFAULT NULL,
  `storage_requirements` text DEFAULT NULL,
  `quality_standards` text DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `commodities`
--

INSERT INTO `commodities` (`id`, `name`, `code`, `description`, `specifications`, `trading_hours`, `contract_size`, `price_unit`, `minimum_price`, `maximum_price`, `current_price`, `price_change`, `price_change_percent`, `trading_volume`, `market_status`, `image_path`, `category`, `origin_country`, `harvest_season`, `storage_requirements`, `quality_standards`, `created_at`, `updated_at`) VALUES
(1, 'Maize', 'GAPWM2', 'White maize trading information and specifications', 'Moisture content: 12-14%, Foreign matter: 2% max, Broken kernels: 5% max', 'Monday-Friday 9:00-17:00', '50kg bags', 'GHC per bag', 1800.00, 2000.00, 1880.00, 30.00, 1.62, 150000, 'Open', '/commodities/maize.jpg', 'Grains', 'Ghana', 'June-September', 'Cool, dry storage, 12-14% moisture', 'Ghana Standards Authority Grade A', '2025-09-05 13:47:55.000', '2025-09-05 13:47:55.000'),
(2, 'Soya Bean', 'GEJWM2', 'Soya bean market data and trading details', 'Protein content: 35% min, Oil content: 18% min, Moisture: 12% max', 'Monday-Friday 9:00-17:00', '50kg bags', 'GHC per bag', 3800.00, 4200.00, 4030.00, 50.00, 1.26, 75000, 'Open', '/commodities/soya-bean.jpg', 'Oilseeds', 'Ghana', 'October-December', 'Dry storage, 12% moisture max', 'Ghana Standards Authority Grade A', '2025-09-05 13:47:55.000', '2025-09-05 13:47:55.000'),
(3, 'Sorghum', 'GSAWM2', 'Sorghum commodity information and trading specs', 'Moisture: 12% max, Foreign matter: 2% max, Broken grains: 5% max', 'Monday-Friday 9:00-17:00', '50kg bags', 'GHC per bag', 4600.00, 4900.00, 4745.00, 25.00, 0.53, 45000, 'Open', '/commodities/sorghum.jpg', 'Grains', 'Ghana', 'September-December', 'Dry storage, 12% moisture', 'Ghana Standards Authority Grade A', '2025-09-05 13:47:55.000', '2025-09-05 13:47:55.000'),
(4, 'Sesame', 'GKUWM2', 'Sesame trading specifications and market data', 'Oil content: 45% min, Moisture: 8% max, Foreign matter: 1% max', 'Monday-Friday 9:00-17:00', '50kg bags', 'GHC per bag', 4500.00, 4800.00, 4645.00, 45.00, 0.98, 25000, 'Open', '/commodities/sesame.jpg', 'Oilseeds', 'Ghana', 'November-January', 'Cool, dry storage, 8% moisture', 'Ghana Standards Authority Grade A', '2025-09-05 13:47:55.000', '2025-09-05 13:47:55.000'),
(5, 'Rice', 'GRIWM2', 'Rice commodity market data and trading information', 'Moisture: 14% max, Broken grains: 5% max, Foreign matter: 1% max', 'Monday-Friday 9:00-17:00', '50kg bags', 'GHC per bag', 2200.00, 2500.00, 2350.00, -20.00, -0.84, 85000, 'Open', '/commodities/rice.jpg', 'Grains', 'Ghana', 'October-December', 'Dry storage, 14% moisture max', 'Ghana Standards Authority Grade A', '2025-09-05 13:47:55.000', '2025-09-05 13:47:55.000');

-- --------------------------------------------------------

--
-- Table structure for table `commodity_info`
--

CREATE TABLE `commodity_info` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `name` longtext NOT NULL,
  `code` longtext NOT NULL,
  `description` text DEFAULT NULL,
  `specifications` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL CHECK (json_valid(`specifications`)),
  `trading_info` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL CHECK (json_valid(`trading_info`)),
  `is_active` tinyint(1) DEFAULT 1,
  `sort_order` bigint(20) DEFAULT 0,
  `image` longtext DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `languages`
--

CREATE TABLE `languages` (
  `code` varchar(191) NOT NULL,
  `name` longtext DEFAULT NULL,
  `native_name` longtext DEFAULT NULL,
  `flag` longtext DEFAULT NULL,
  `is_active` tinyint(1) DEFAULT 1,
  `is_rtl` tinyint(1) DEFAULT 0,
  `sort_order` bigint(20) DEFAULT 0,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `market_analytics`
--

CREATE TABLE `market_analytics` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `commodity` longtext NOT NULL,
  `date` datetime(3) NOT NULL,
  `total_volume` double DEFAULT NULL,
  `average_price` double DEFAULT NULL,
  `price_change` double DEFAULT NULL,
  `price_change_percent` double DEFAULT NULL,
  `high_price` double DEFAULT NULL,
  `low_price` double DEFAULT NULL,
  `open_price` double DEFAULT NULL,
  `close_price` double DEFAULT NULL,
  `transaction_count` bigint(20) DEFAULT NULL,
  `market_sentiment` longtext DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `market_data`
--

CREATE TABLE `market_data` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `commodity` longtext NOT NULL,
  `price` double NOT NULL,
  `currency` varchar(191) DEFAULT 'GHS',
  `unit` varchar(191) DEFAULT 'metric_ton',
  `change` double DEFAULT NULL,
  `change_percent` double DEFAULT NULL,
  `volume` double DEFAULT NULL,
  `high` double DEFAULT NULL,
  `low` double DEFAULT NULL,
  `open` double DEFAULT NULL,
  `close` double DEFAULT NULL,
  `market_date` datetime(3) DEFAULT NULL,
  `source` longtext DEFAULT NULL,
  `metadata` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL CHECK (json_valid(`metadata`)),
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `media_files`
--

CREATE TABLE `media_files` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `original_name` longtext NOT NULL,
  `filename` varchar(191) NOT NULL,
  `url` longtext NOT NULL,
  `thumbnail_url` longtext DEFAULT NULL,
  `mime_type` longtext NOT NULL,
  `size` bigint(20) NOT NULL,
  `uploaded_by` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `media_files`
--

INSERT INTO `media_files` (`id`, `original_name`, `filename`, `url`, `thumbnail_url`, `mime_type`, `size`, `uploaded_by`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'deep-blue-shirt-a-high-quality-e-commerce-fashion-photoshoot-model-wearing-a-deep-blue-premium-dress-shirt-wrinkle-free-detailed-stitching-realistic-textures-in-a-clean-neutral-studio-setting-with-soft-professional.png', '1755526687_qqqqqqqq.png', '/uploads/images/1755526687_qqqqqqqq.png', NULL, 'image/png', 2202297, 1, '2025-08-18 14:18:07.519', '2025-08-18 14:18:07.519', NULL),
(2, 'deep-blue-shirt-a-high-quality-e-commerce-fashion-photoshoot-model-wearing-a-deep-blue-premium-dress-shirt-wrinkle-free-detailed-stitching-realistic-textures-in-a-clean-neutral-studio-setting-with-soft-professional.png', '1755527015_66666666.png', '/uploads/images/1755527015_66666666.png', NULL, 'image/png', 2202297, 1, '2025-08-18 14:23:35.963', '2025-08-18 14:23:35.963', NULL),
(3, 'black-shirt-a-high-quality-e-commerce-fashion-photoshoot-model-wearing-a-sleek-black-premium-dress-shirt-wrinkle-free-detailed-stitching-realistic-textures-in-a-clean-neutral-studio-setting-with-soft-professional-l.png', '1755529452_yyyyyyyy.png', '/uploads/images/1755529452_yyyyyyyy.png', NULL, 'image/png', 2637222, 1, '2025-08-18 15:04:12.061', '2025-08-18 15:04:12.061', NULL),
(4, 'black-shirt-a-high-quality-e-commerce-fashion-photoshoot-model-wearing-a-sleek-black-premium-dress-shirt-wrinkle-free-detailed-stitching-realistic-textures-in-a-clean-neutral-studio-setting-with-soft-professional-l.png', '1755530567_yyyyyyyy.png', '/uploads/images/1755530567_yyyyyyyy.png', NULL, 'image/png', 2637222, 1, '2025-08-18 15:22:47.749', '2025-08-18 15:22:47.749', NULL),
(5, 's-single-picture-will-do-too-please.png', '1755530991_iiiiiiii.png', 'http://localhost:8080/uploads/images/1755530991_iiiiiiii.png', NULL, 'image/png', 2541169, 1, '2025-08-18 15:29:51.936', '2025-08-18 15:29:51.936', NULL),
(6, '20250815_1212_Crisp White Elegance_simple_compose_01k2pt74nvfvx91f8ts7s1qz8z.png', '1755531021_66666666.png', 'http://localhost:8080/uploads/images/1755531021_66666666.png', NULL, 'image/png', 2281687, 1, '2025-08-18 15:30:21.321', '2025-08-18 15:30:21.321', NULL),
(7, 'deep-blue-shirt-a-high-quality-e-commerce-fashion-photoshoot-model-wearing-a-deep-blue-premium-dress-shirt-wrinkle-free-detailed-stitching-realistic-textures-in-a-clean-neutral-studio-setting-with-soft-professional.png', '1755531513_yyyyyyyy.png', 'http://localhost:8080/uploads/images/1755531513_yyyyyyyy.png', NULL, 'image/png', 2202297, 1, '2025-08-18 15:38:33.673', '2025-08-18 15:38:33.673', NULL),
(8, 's-single-picture-will-do-too-please.png', '1755531961_aaaaaaaa.png', 'http://localhost:8080/uploads/images/1755531961_aaaaaaaa.png', NULL, 'image/png', 2541169, 1, '2025-08-18 15:46:01.573', '2025-08-18 15:46:01.573', NULL),
(9, 's-single-picture-will-do-too-please.png', '1755532044_iiiiiiii.png', 'http://localhost:8080/uploads/images/1755532044_iiiiiiii.png', NULL, 'image/png', 2541169, 1, '2025-08-18 15:47:24.173', '2025-08-18 15:47:24.173', NULL),
(10, 's-single-picture-will-do-too-please.png', '1755532344_mmmmmmmm.png', 'http://localhost:8080/uploads/images/1755532344_mmmmmmmm.png', NULL, 'image/png', 2541169, 1, '2025-08-18 15:52:24.372', '2025-08-18 15:52:24.372', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `menus`
--

CREATE TABLE `menus` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `name` longtext NOT NULL,
  `location` longtext DEFAULT NULL,
  `is_active` tinyint(1) DEFAULT 1,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `menu_items`
--

CREATE TABLE `menu_items` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `menu_id` bigint(20) UNSIGNED DEFAULT NULL,
  `label` longtext NOT NULL,
  `url` longtext DEFAULT NULL,
  `target` varchar(191) DEFAULT '_self',
  `icon_class` longtext DEFAULT NULL,
  `parent_id` bigint(20) UNSIGNED DEFAULT NULL,
  `sort_order` bigint(20) DEFAULT 0,
  `is_active` tinyint(1) DEFAULT 1,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `pages`
--

CREATE TABLE `pages` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `title` longtext NOT NULL,
  `slug` longtext NOT NULL,
  `content` longtext DEFAULT NULL,
  `excerpt` text DEFAULT NULL,
  `template` varchar(191) DEFAULT 'default',
  `status` varchar(191) DEFAULT 'draft',
  `featured_image` longtext DEFAULT NULL,
  `meta_title` longtext DEFAULT NULL,
  `meta_description` longtext DEFAULT NULL,
  `meta_keywords` longtext DEFAULT NULL,
  `parent_id` bigint(20) UNSIGNED DEFAULT NULL,
  `sort_order` bigint(20) DEFAULT 0,
  `author_id` bigint(20) UNSIGNED DEFAULT NULL,
  `published_at` datetime(3) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `pages`
--

INSERT INTO `pages` (`id`, `title`, `slug`, `content`, `excerpt`, `template`, `status`, `featured_image`, `meta_title`, `meta_description`, `meta_keywords`, `parent_id`, `sort_order`, `author_id`, `published_at`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'About GCX', 'about', '{\"about_title\":\"About GCX\",\"about_description\":\"The Ghana Commodity Exchange is a private company limited by shares, structured as a Public-Private Partnership, with the Government of Ghana currently the sole shareholder. The Exchange aims to establish linkages between agricultural and commodity producers and buyers, secure competitive prices for their products, assure the market quantity and quality, and settle trade promptly.\",\"ceo_name\":\"Mr. Robert Dowuona Owooi\",\"ceo_title\":\"Acting Chief Executive Officers\",\"ceo_image\":\"/Mr. Robert Dowuona Owoo.jpeg\",\"ceo_intro\":\"Ghana Commodity Exchange\'s Management team is led by Mr. Robert Dowuona Owooi, the Acting Chief Executive Officer.\",\"key_goal_title\":\"Our Key Goal\",\"key_goal_description\":\"To link Ghanaian smallholder farmers to agricultural and financial markets in Ghana and across the West Africa Region to ensure Ghana farmers secure competitive prices for their commodities, as well as supply good quality commodities which meet the nutritional needs of the Ghanaian people.\",\"mission_title\":\"Mission..\",\"mission_subtitle\":\"Ghana Commodity Exchange (GCX) is a leading commodity exchange platform dedicated to transforming agricultural trading in Ghana and West Africa.\",\"our_mission_title\":\"Our Mission\",\"our_mission_text1\":\"We provide innovative solutions for farmers, traders, and stakeholders, ensuring transparency, efficiency, and growth in the agricultural sector.\",\"our_mission_text2\":\"Through our regulated marketplace, we connect agricultural producers with buyers, enabling fair pricing and secure transactions.\",\"our_mission_text3\":\"Our platform ensures quality assurance, timely delivery, and reliable settlement for all market participants.\",\"values_title\":\"Our Core Values...\",\"transparency_title\":\"Transparency\",\"transparency_text\":\"Open and fair trading practices with clear pricing mechanisms\",\"efficiency_title\":\"Efficiency\",\"efficiency_text\":\"Streamlined processes for optimal market performance\",\"statistics_title\":\"Agricultural Trading\",\"farmers_count\":\"50K+\",\"farmers_label\":\"Farmers Connected\",\"mission_image\":\"/Service/testing.jpg\",\"mission_image_alt\":\"GCX Mission\",\"statistics_image\":\"/crop.jpg\",\"statistics_image_alt\":\"Agricultural Statistics\",\"vision_purpose_title\":\"Our Vision, Purpose & Objectives\",\"vision_purpose_subtitle\":\"Driving Ghana\'s economic transformation through innovative commodity trading solutions\",\"vision_title\":\"Our Vision\",\"vision_description\":\"To transform Ghana\'s economy by creating prosperity for all in commodity value chains and become a regional and global trading hub for all commodities.\",\"purpose_title\":\"Our Purpose\",\"purpose_description\":\"To connect markets, connect people & provide opportunities through innovation; using the most interactive and appropriate technology that meets the needs of our stakeholders and honouring earth\'s resources through sustainable practices across our business.\",\"objectives_title\":\"Our Objectives...\",\"objective_1\":\"Provide liquid, efficient markets for instruments such as auctions, spot, forward, futures and options....\",\"objective_2\":\"Support access to finance through efficient linkages along the value chain\",\"objective_3\":\"Build capacity of members to use GCX infrastructure\",\"objective_4\":\"Establish fair and transparent price discovery mechanisms\",\"objective_5\":\"Ensure market integrity through rigorous surveillance systems\",\"objective_6\":\"Develop standards, contracts, and dispute mechanisms to reduce risk\",\"objective_7\":\"Implement a reliable, efficient, and guaranteed clearing & settlement system...\",\"leadership_title\":\"Leadership Team\",\"leadership_subtitle\":\"Experienced leaders driving Ghana Commodity Exchange\'s strategic vision and operational Excellence.\",\"ceo_bio\":\"Experienced leader with extensive background in commodity trading and financial markets.\",\"deputy_ceo_name\":\"Ms. Ophelia Martekuo Atoklo\",\"deputy_ceo_title\":\"Acting Deputy Chief Executive Officer\",\"deputy_ceo_image\":\"/Ms. Ophelia Martekuo Atoklo \'Deputy Chief Executive Officer\'.jpg\",\"deputy_ceo_bio\":\"Strategic leader with deep expertise in operational excellence and stakeholder management.\",\"board_title\":\"Board of Directors\",\"board_subtitle\":\"Governance and strategic oversight by experienced industry leaders\",\"chairman_name\":\"Mr. Kofi S. Yamoah\",\"chairman_title\":\"Board Chairman\",\"chairman_image\":\"/Board of directors/Mr. Kofi S. Yamoah \'Board Chairman\'.png\",\"chairman_bio\":\"Experienced leader providing strategic direction and governance oversight.\",\"director1_name\":\"Mr. Kwame Daaku\",\"director1_title\":\"Non-Executive Director\",\"director1_image\":\"/Board of directors/Mr. Kwame Daaku \'Non-excutive director\'.jpg\",\"director1_bio\":\"Independent director bringing valuable industry expertise and governance experience.\",\"director2_name\":\"Mr. Stephen Antwi-Asimeng\",\"director2_title\":\"Non-Executive Director\",\"director2_image\":\"/Board of directors/Mr. Stephen Antwi-Asimeng \' non-excutive director\'.png\",\"director2_bio\":\"Independent director with extensive experience in financial markets and corporate governance.\",\"secretary_name\":\"Mrs. Wendy Malm\",\"secretary_title\":\"Board Secretary\",\"secretary_image\":\"/Board of directors/Mrs. Wendy Malm Board Secretary.png\",\"secretary_bio\":\"Experienced professional ensuring effective board governance and compliance.\",\"functional_heads_title\":\"Functional Heads....\",\"functional_heads_subtitle\":\" functional leaders ensure operational excellence across all departments.\",\"auditor_name\":\"Mr. Opoku Debrah\",\"auditor_title\":\"Internal Auditor\",\"auditor_image\":\"/Functional Heads/Mr. Opoku Debrah.jpg\",\"auditor_bio\":\"Ensuring compliance and risk management across all organizational processes.\",\"auditor_description\":\"Responsible for internal audit functions, risk assessment, and ensuring compliance with regulatory requirements and internal policies.\",\"special_project_name\":\"Mr. Richard Ankrah\",\"special_project_title\":\"Special Project\",\"special_project_image\":\"/Functional Heads/Mr. Richard Ankrah.jpg\",\"special_project_bio\":\"Leading strategic initiatives and special projects to drive organizational growth.\",\"special_project_description\":\"Oversees special projects and strategic initiatives that drive innovation and organizational development.\",\"risk_membership_name\":\"Mr. Vitus Ninfaakang\",\"risk_membership_title\":\"Risk, Membership and Partnership\",\"risk_membership_image\":\"/Functional Heads/Mr. Vitus Ninfaakang.jpg\",\"risk_membership_bio\":\"Managing risk assessment and building strategic partnerships with key stakeholders.\",\"risk_membership_description\":\"Manages risk assessment, membership services, and strategic partnerships to ensure market stability and growth.\",\"corporate_services_name\":\"Mrs. Jemimah Naa Adjeley Oppong-Gyamfi\",\"corporate_services_title\":\"Corporate Services\",\"corporate_services_image\":\"/Functional Heads/Mrs. Jemimah Naa Adjeley Oppong-Gyamfi.jpg\",\"corporate_services_bio\":\"Overseeing corporate governance and administrative excellence.\",\"corporate_services_description\":\"Oversees corporate governance, administrative functions, and ensures compliance with corporate policies and procedures.\",\"operations_name\":\"Mrs. Wendy Malmm\",\"operations_title\":\"Operations\",\"operations_image\":\"/Functional Heads/Mrs. Wendy Malm.jpg\",\"operations_bio\":\"Ensuring smooth operational processes and service delivery excellence.\",\"operations_description\":\"Manages day-to-day operations, ensures efficient service delivery, and maintains operational excellence across all departments.\",\"value_chain_name\":\"Mr. Godfred Kofi Nyamekye\",\"value_chain_title\":\"Value Chain and Product Development\",\"value_chain_image\":\"/Functional Heads/Mr. Godfred Kofi Nyamekye.jpg\",\"value_chain_bio\":\"Mr. Albert Tagoe is a Chartered Management Accountant and is currently the Head of Finance and Investment at the Ghana Commodity Exchange and has over a decade of experience in Financial Management. Albert was awarded in March 2023 as part of the Top 30 Public Sector Finance Leaders in Ghana. He holds an MBA in Finance with Distinction from Coventry University (UK) and holds a Chartered Professional Membership with the Chartered Institute of Management Accountants (CIMA-UK), and the Institute of Chartered Accountants Ghana (ICAG-GH).\\n\\nHe holds a First-Class Honours degree in Bachelor of Business Administration (Accounting Option) from Valley View University and a certification by the Ghana Stock Exchange for Investment and Securities. Albert joined the Exchange for its implementation in January 2018 and has supported in setting up a robust and dynamic finance and control unit for the business with his expertise in Financial Management and Reporting, Financial Analysis, Budgeting and Financial projections, Treasury Management, Investment, Internal and External Audit, and Financial controls and Taxation. As part of his portfolio, he currently leads the teams’ engagement with Financial Institutions to extend financing to smallholder farmers in Ghana under its’ Warehouse Receipt and Aggregation Financing programme where the commodities serve as collateral for the financial institutions.\\n\\nAlbert worked for (7) years with Stanbic Bank Ghana in Credit Risk Management and Banking operations and received a Beyond Excellence Award in 2017 for service excellence. He also received an award as the Top Performing Student in Ghana in the CIMA Strategic Case Study Exam and the Top Performing Student in Ghana in the CIMA Performance Strategy Paper in 2014.\",\"value_chain_description\":\"Develops new products, optimizes value chains, and drives innovation to enhance market offerings and efficiency.\",\"warehouse_quality_name\":\"Mr. Gabriel Aryeetey\",\"warehouse_quality_title\":\"Warehouse & Quality\",\"warehouse_quality_image\":\"/Functional Heads/Mr. Gabriel Aryeetey.jpg\",\"warehouse_quality_bio\":\"Maintaining quality standards and efficient warehouse operations.\",\"warehouse_quality_description\":\"Manages warehouse operations, quality control, and ensures proper storage and handling of commodities.\",\"it_systems_name\":\"Dr. Harold Okai-Tettey\",\"it_systems_title\":\"Information Technology and Information System\",\"it_systems_image\":\"/Functional Heads/Dr. Harold Okai-Tettey.jpg\",\"it_systems_bio\":\"Leading digital transformation and technology infrastructure development.\",\"it_systems_description\":\"Leads digital transformation, manages IT infrastructure, and ensures secure and efficient technology systems.\",\"finance_investments_name\":\"Mr. Albert Nii Ayi Tagoe\",\"finance_investments_title\":\"Finance and Investments...\",\"finance_investments_image\":\"/Functional Heads/Mr. Albert Nii Ayi Tagoe.jpg\",\"finance_investments_bio\":\"Managing financial strategy and investment opportunities for sustainable growth.\",\"finance_investments_description\":\"Manages financial planning, investment strategies, and ensures sound financial management and growth opportunities.\"}', 'Learn about Ghana Commodity Exchange, our mission, vision, and leadership team.', 'default', 'published', NULL, 'About GCX - Ghana Commodity Exchange', 'Learn about Ghana Commodity Exchange, our mission, vision, and leadership team.', 'about, gcx, ghana commodity exchange, mission, vision, leadership', NULL, 0, 1, '2025-09-04 12:28:26.693', '2025-09-04 10:54:21.518', '2025-09-05 10:13:27.866', NULL),
(2, 'About GCX', 'about-1', '{\"hero_title\":\"What is GCb\",\"hero_subtitle\":\"A regulated market that links buyers and sellers of commodities to trade by Rules, while we assure the market quantity and quality, timely delivery and settlement.\",\"hero_image\":\"/trading dashboard.jpg\",\"about_title\":\"About GCX\",\"about_description\":\"The Ghana Commodity Exchange is a private company limited by shares, structured as a Public-Private Partnership, with the Government of Ghana currently the sole shareholder.\",\"ceo_name\":\"Mr. Robert Dowuona Owoo\",\"ceo_title\":\"Acting Chief Executive Officer\",\"ceo_image\":\"/Mr. Robert Dowuona Owoo.jpeg\",\"key_goal_title\":\"Our Key Goal\",\"key_goal_description\":\"To link Ghanaian smallholder farmers to agricultural and financial markets in Ghana and across the West Africa Region to ensure Ghana farmers secure competitive prices for their commodities.\"}', 'Learn about Ghana Commodity Exchange, our mission, vision, and leadership team.', 'default', 'published', NULL, 'About GCX - Ghana Commodity Exchange', 'Learn about Ghana Commodity Exchange, our mission, vision, and leadership team.', 'about, gcx, ghana commodity exchange, mission, vision, leadership', NULL, 0, 1, '2025-09-04 10:55:45.881', '2025-09-04 10:55:45.881', '2025-09-04 10:55:45.881', NULL),
(3, 'About GCX', 'about-2', '{\"hero_title\":\"What is GCB\",\"hero_subtitle\":\"A regulated market that links buyers and sellers of commodities to trade by Rules, while we assure the market quantity and quality, timely delivery and settlement.\",\"hero_image\":\"/trading dashboard.jpg\",\"about_title\":\"About GCX\",\"about_description\":\"The Ghana Commodity Exchange is a private company limited by shares, structured as a Public-Private Partnership, with the Government of Ghana currently the sole shareholder.\",\"ceo_name\":\"Mr. Robert Dowuona Owoo\",\"ceo_title\":\"Acting Chief Executive Officer\",\"ceo_image\":\"/Mr. Robert Dowuona Owoo.jpeg\",\"key_goal_title\":\"Our Key Goal\",\"key_goal_description\":\"To link Ghanaian smallholder farmers to agricultural and financial markets in Ghana and across the West Africa Region to ensure Ghana farmers secure competitive prices for their commodities.\"}', 'Learn about Ghana Commodity Exchange, our mission, vision, and leadership team.', 'default', 'published', NULL, 'About GCX - Ghana Commodity Exchange', 'Learn about Ghana Commodity Exchange, our mission, vision, and leadership team.', 'about, gcx, ghana commodity exchange, mission, vision, leadership', NULL, 0, 1, '2025-09-04 10:56:35.247', '2025-09-04 10:56:35.248', '2025-09-04 10:56:35.248', NULL),
(4, 'About GCX', 'about-3', '{\"hero_title\":\"What is GCB\",\"hero_subtitle\":\"A regulated market that links buyers and sellers of commodities to trade by Rules, while we assure the market quantity and quality, timely delivery and settlement.\",\"hero_image\":\"/trading dashboard.jpg\",\"about_title\":\"About GCX\",\"about_description\":\"The Ghana Commodity Exchange is a private company limited by shares, structured as a Public-Private Partnership, with the Government of Ghana currently the sole shareholder.\",\"ceo_name\":\"Mr. Robert Dowuona Owoo\",\"ceo_title\":\"Acting Chief Executive Officer\",\"ceo_image\":\"/Mr. Robert Dowuona Owoo.jpeg\",\"key_goal_title\":\"Our Key Goal\",\"key_goal_description\":\"To link Ghanaian smallholder farmers to agricultural and financial markets in Ghana and across the West Africa Region to ensure Ghana farmers secure competitive prices for their commodities.\"}', 'Learn about Ghana Commodity Exchange, our mission, vision, and leadership team.', 'default', 'published', NULL, 'About GCX - Ghana Commodity Exchange', 'Learn about Ghana Commodity Exchange, our mission, vision, and leadership team.', 'about, gcx, ghana commodity exchange, mission, vision, leadership', NULL, 0, 1, '2025-09-04 11:09:29.761', '2025-09-04 11:09:29.762', '2025-09-04 11:09:29.762', NULL),
(5, 'About GCX', 'about-4', '{\"hero_title\":\"What is GCX\",\"hero_subtitle\":\"A negulated market that links buyers and sellers of commodities to trade by Rules, while we assure the market quantity and quality, timely delivery and settlement.\",\"hero_image\":\"/trading dashboard.jpg\",\"about_title\":\"About GCX\",\"about_description\":\"The Ghana Commodity Exchange is a private company limited by shares, structured as a Public-Private Partnership, with the Government of Ghana currently the sole shareholder.\",\"ceo_name\":\"Mr. Robert Dowuona Owoo\",\"ceo_title\":\"Acting Chief Executive Officer\",\"ceo_image\":\"/Mr. Robert Dowuona Owoo.jpeg\",\"key_goal_title\":\"Our Key Goal\",\"key_goal_description\":\"To link Ghanaian smallholder farmers to agricultural and financial markets in Ghana and across the West Africa Region to ensure Ghana farmers secure competitive prices for their commodities.\"}', 'Learn about Ghana Commodity Exchange, our mission, vision, and leadership team.', 'default', 'published', NULL, 'About GCX - Ghana Commodity Exchange', 'Learn about Ghana Commodity Exchange, our mission, vision, and leadership team.', 'about, gcx, ghana commodity exchange, mission, vision, leadership', NULL, 0, 1, '2025-09-04 11:12:44.976', '2025-09-04 11:12:44.977', '2025-09-04 11:12:44.977', NULL),
(6, 'About GCX', 'about-5', '{\"hero_title\":\"What is GCX\",\"hero_subtitle\":\"A regulated market that links buyers and sellers of commodities to trade by Rules, while we assure the market quantity and quality, timely delivery and settlement.\",\"hero_image\":\"/trading dashboard.jpg\",\"about_title\":\"About GCX\",\"about_description\":\"The Ghana Commodity Exchange is a private company limited by shares, structured as a Public-Private Partnership, with the Government of Ghana currently the sole shareholder.\",\"ceo_name\":\"Mr. Robert Dowuona Owoo\",\"ceo_title\":\"Acting Chief Executive Officer\",\"ceo_image\":\"/Mr. Robert Dowuona Owoo.jpeg\",\"key_goal_title\":\"Our Key Goal\",\"key_goal_description\":\"To link Ghanaian smallholder farmers to agricultural and financial markets in Ghana and across the West Africa Region to ensure Ghana farmers secure competitive prices for their commodities.\"}', 'Learn about Ghana Commodity Exchange, our mission, vision, and leadership team.', 'default', 'published', NULL, 'About GCX - Ghana Commodity Exchange', 'Learn about Ghana Commodity Exchange, our mission, vision, and leadership team.', 'about, gcx, ghana commodity exchange, mission, vision, leadership', NULL, 0, 1, '2025-09-04 11:27:36.940', '2025-09-04 11:27:36.940', '2025-09-04 11:27:36.940', NULL),
(7, 'About GCX', 'about-6', '{\"hero_title\":\"What is GCBB\",\"hero_subtitle\":\"A regulated market that links buyers and sellers of commodities to trade by Rules, while we assure the market quantity and quality, timely delivery and settlement.\",\"hero_image\":\"/trading dashboard.jpg\",\"about_title\":\"About GCX\",\"about_description\":\"The Ghana Commodity Exchange is a private company limited by shares, structured as a Public-Private Partnership, with the Government of Ghana currently the sole shareholder.\",\"ceo_name\":\"Mr. Robert Dowuona Owoo\",\"ceo_title\":\"Acting Chief Executive Officer\",\"ceo_image\":\"/Mr. Robert Dowuona Owoo.jpeg\",\"key_goal_title\":\"Our Key Goal\",\"key_goal_description\":\"To link Ghanaian smallholder farmers to agricultural and financial markets in Ghana and across the West Africa Region to ensure Ghana farmers secure competitive prices for their commodities.\"}', 'Learn about Ghana Commodity Exchange, our mission, vision, and leadership team.', 'default', 'published', NULL, 'About GCX - Ghana Commodity Exchange', 'Learn about Ghana Commodity Exchange, our mission, vision, and leadership team.', 'about, gcx, ghana commodity exchange, mission, vision, leadership', NULL, 0, 1, '2025-09-04 11:27:57.275', '2025-09-04 11:27:57.275', '2025-09-04 11:27:57.275', NULL),
(8, 'About GCX', 'about-7', '{\"hero_title\":\"What is GCBB\",\"hero_subtitle\":\"A negulated market that links buyers and sellers of commodities to trade by Rules, while we assure the market quantity and quality, timely delivery and settlement.\",\"hero_image\":\"/trading dashboard.jpg\",\"about_title\":\"About GCX\",\"about_description\":\"The Ghana Commodity Exchange is a private company limited by shares, structured as a Public-Private Partnership, with the Government of Ghana currently the sole shareholder.\",\"ceo_name\":\"Mr. Robert Dowuona Owoo\",\"ceo_title\":\"Acting Chief Executive Officer\",\"ceo_image\":\"/Mr. Robert Dowuona Owoo.jpeg\",\"key_goal_title\":\"Our Key Goal\",\"key_goal_description\":\"To link Ghanaian smallholder farmers to agricultural and financial markets in Ghana and across the West Africa Region to ensure Ghana farmers secure competitive prices for their commodities.\"}', 'Learn about Ghana Commodity Exchange, our mission, vision, and leadership team.', 'default', 'published', NULL, 'About GCX - Ghana Commodity Exchange', 'Learn about Ghana Commodity Exchange, our mission, vision, and leadership team.', 'about, gcx, ghana commodity exchange, mission, vision, leadership', NULL, 0, 1, '2025-09-04 11:28:42.371', '2025-09-04 11:28:42.372', '2025-09-04 11:28:42.372', NULL),
(9, 'About GCX', 'about-8', '{\"hero_title\":\"What is GCBB\",\"hero_subtitle\":\"A begulated market that links buyers and sellers of commodities to trade by Rules, while we assure the market quantity and quality, timely delivery and settlement.\",\"hero_image\":\"/trading dashboard.jpg\",\"about_title\":\"About GCX\",\"about_description\":\"The Ghana Commodity Exchange is a private company limited by shares, structured as a Public-Private Partnership, with the Government of Ghana currently the sole shareholder.\",\"ceo_name\":\"Mr. Robert Dowuona Owoo\",\"ceo_title\":\"Acting Chief Executive Officer\",\"ceo_image\":\"/Mr. Robert Dowuona Owoo.jpeg\",\"key_goal_title\":\"Our Key Goal\",\"key_goal_description\":\"To link Ghanaian smallholder farmers to agricultural and financial markets in Ghana and across the West Africa Region to ensure Ghana farmers secure competitive prices for their commodities.\"}', 'Learn about Ghana Commodity Exchange, our mission, vision, and leadership team.', 'default', 'published', NULL, 'About GCX - Ghana Commodity Exchange', 'Learn about Ghana Commodity Exchange, our mission, vision, and leadership team.', 'about, gcx, ghana commodity exchange, mission, vision, leadership', NULL, 0, 1, '2025-09-04 11:28:58.188', '2025-09-04 11:28:58.189', '2025-09-04 11:28:58.189', NULL),
(10, 'About GCX', 'about-9', '{\"hero_title\":\"what is gcx\",\"hero_subtitle\":\"A negulated market that links buyers and sellers of commodities to trade by Rules, while we assure the market quantity and quality, timely delivery and settlement.\"}', 'Learn about Ghana Commodity Exchange, our mission, vision, and leadership team.', 'default', 'published', NULL, 'About GCX - Ghana Commodity Exchange', 'Learn about Ghana Commodity Exchange, our mission, vision, and leadership team.', 'about, gcx, ghana commodity exchange, mission, vision, leadership', NULL, 0, 1, '2025-09-04 11:33:51.526', '2025-09-04 11:33:51.527', '2025-09-04 11:33:51.527', NULL),
(11, 'Our Services', 'services', '{\"services_title\":\"Our Core Services\",\"services_subtitle\":\"Ghana Commodity Exchange provides comprehensive trading solutions for agricultural commodities\",\"service_trading_title\":\"Regulated Trading\",\"service_trading_description\":\"National and regional market linking buyers and sellers under established rules with transparent pricing and secure settlement.\",\"service_price_title\":\"Price Discovery\",\"service_price_description\":\"Transparent price discovery mechanism ensuring competitive pricing for commodities through real-time market data.\",\"service_quality_title\":\"Quality Assurance\",\"service_quality_description\":\"Assured market quantity and quality with timely delivery and settlement through rigorous standards.\"}', 'Discover the comprehensive trading solutions and services offered by GCX.', 'default', 'published', NULL, 'Services - Ghana Commodity Exchange', 'Discover the comprehensive trading solutions and services offered by GCX.', 'services, trading, commodity exchange, gcx', NULL, 0, 1, '2025-09-04 11:45:49.071', '2025-09-04 11:45:49.071', '2025-09-04 16:36:59.795', NULL),
(12, 'Membership', 'membership', '{\"hero_title\":\"Join GCX\",\"hero_subtitle\":\"Become a member and access our trading platform\"}', 'Join GCX as a member and access our trading platform and services.', 'default', 'published', NULL, 'Membership - Ghana Commodity Exchange', 'Join GCX as a member and access our trading platform and services.', 'membership, join, gcx, trading platform', NULL, 0, 1, '2025-09-04 11:46:03.786', '2025-09-04 11:46:03.786', '2025-09-04 11:46:03.786', NULL),
(13, 'About GCX', 'about-10', '{\"hero_title\":\"Test Title 1756984859094\",\"hero_subtitle\":\"A negulated market that links buyers and sellers of commodities to trade by Rules, while we assure the market quantity and quality, timely delivery and settlement.\"}', 'Learn about Ghana Commodity Exchange, our mission, vision, and leadership team.', 'default', 'published', NULL, 'About GCX - Ghana Commodity Exchange', 'Learn about Ghana Commodity Exchange, our mission, vision, and leadership team.', 'about, gcx, ghana commodity exchange, mission, vision, leadership', NULL, 0, 1, '2025-09-04 12:33:17.943', '2025-09-04 12:33:17.944', '2025-09-04 12:33:17.944', NULL),
(14, 'About GCX', 'about-11', '{\"hero_title\":\"1756984859094\",\"hero_subtitle\":\"A negulated market that links buyers and sellers of commodities to trade by Rules, while we assure the market quantity and quality, timely delivery and settlement.\"}', 'Learn about Ghana Commodity Exchange, our mission, vision, and leadership team.', 'default', 'published', NULL, 'About GCX - Ghana Commodity Exchange', 'Learn about Ghana Commodity Exchange, our mission, vision, and leadership team.', 'about, gcx, ghana commodity exchange, mission, vision, leadership', NULL, 0, 1, '2025-09-04 12:33:34.103', '2025-09-04 12:33:34.104', '2025-09-04 12:33:34.104', NULL),
(15, 'Homepage', 'homepage', '{\"hero_title\":\"Ghana Commodity Exchange\",\"hero_subtitle\":\"Connecting Markets, Connecting People, Providing Opportunities\",\"hero_description\":\"A regulated market that links buyers and sellers of commodities to trade by Rules, while we assure the market quantity and quality, timely delivery and settlement.\",\"hero_cta_primary\":\"Explore Platform\",\"hero_cta_secondary\":\"Learn More\"}', 'Content for the homepage sections', 'default', 'published', NULL, 'Homepage - GCX', 'Content for the homepage sections', 'homepage', NULL, 0, 1, '2025-09-04 16:36:59.741', '2025-09-04 16:36:59.742', '2025-09-04 16:36:59.742', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `partners`
--

CREATE TABLE `partners` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `name` longtext NOT NULL,
  `description` longtext DEFAULT NULL,
  `category` longtext NOT NULL,
  `logo` longtext DEFAULT NULL,
  `website` longtext DEFAULT NULL,
  `email` longtext DEFAULT NULL,
  `phone` longtext DEFAULT NULL,
  `address` longtext DEFAULT NULL,
  `status` varchar(191) DEFAULT 'active',
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `partners`
--

INSERT INTO `partners` (`id`, `name`, `description`, `category`, `logo`, `website`, `email`, `phone`, `address`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'test', '', 'financial', '', '', '', '', '', 'active', '2025-09-10 14:57:15.940', '2025-09-10 14:57:15.940', NULL),
(2, 'Ghana Export-Import Bank', 'Ghana\'s premier export-import bank providing financial services to support international trade and economic development.', 'financial', '/Partners/ghana-exim-bank.jpg', 'https://www.ghanaeximbank.com', 'info@ghanaeximbank.com', '+233 302 500 000', 'Accra, Ghana', 'active', '2025-09-10 15:02:58.000', '2025-09-10 15:02:58.000', NULL),
(3, 'Ghana Export Promotion Authority', 'Promoting Ghana\'s exports and facilitating trade relationships with international partners.', 'financial', '/Partners/ghana-export-promotion-authority-gepa.png', 'https://www.gepa.gov.gh', 'info@gepa.gov.gh', '+233 302 500 000', 'Accra, Ghana', 'active', '2025-09-10 15:02:58.000', '2025-09-10 15:02:58.000', NULL),
(4, 'Ghana Standards Authority', 'Ensuring quality standards and certification for Ghana\'s products and services.', 'financial', '/Partners/ghana-standard-authority-gsa.png', 'https://www.gsa.gov.gh', 'info@gsa.gov.gh', '+233 302 500 000', 'Accra, Ghana', 'active', '2025-09-10 15:02:58.000', '2025-09-10 15:02:58.000', NULL),
(5, 'Standard Chartered', 'International banking services supporting Ghana\'s economic growth and development.', 'financial', '/Partners/standard-chartered.jpg', 'https://www.sc.com/gh', 'ghana@sc.com', '+233 302 500 000', 'Accra, Ghana', 'active', '2025-09-10 15:02:58.000', '2025-09-10 15:02:58.000', NULL),
(6, 'Fidelity Bank', 'Leading Ghanaian bank providing comprehensive financial services.', 'financial', '/Partners/fidelity-bank.png', 'https://www.fidelitybank.com.gh', 'info@fidelitybank.com.gh', '+233 302 500 000', 'Accra, Ghana', 'active', '2025-09-10 15:02:58.000', '2025-09-10 15:02:58.000', NULL),
(7, 'Ecobank', 'Pan-African bank supporting economic development across the continent.', 'financial', '/Partners/ecobank.png', 'https://www.ecobank.com/gh', 'ghana@ecobank.com', '+233 302 500 000', 'Accra, Ghana', 'active', '2025-09-10 15:02:58.000', '2025-09-10 15:02:58.000', NULL),
(8, 'Ghana Grains Council', 'Promoting the development of Ghana\'s grain industry and market access.', 'financial', '/Partners/ghana-grains-council-ggc.png', 'https://www.ghanagrainscouncil.org', 'info@ghanagrainscouncil.org', '+233 302 500 000', 'Accra, Ghana', 'active', '2025-09-10 15:02:58.000', '2025-09-10 15:02:58.000', NULL),
(9, 'IPMC', 'Information technology and business solutions provider.', 'financial', '/Partners/ipmc.jpg', 'https://www.ipmc.com.gh', 'info@ipmc.com.gh', '+233 302 500 000', 'Accra, Ghana', 'active', '2025-09-10 15:02:58.000', '2025-09-10 15:02:58.000', NULL),
(10, 'Africa Cashew Alliance', 'Promoting sustainable cashew production and trade across Africa.', 'financial', '/Partners/africa-cashew-alliance.png', 'https://www.africacashewalliance.org', 'info@africacashewalliance.org', '+233 302 500 000', 'Accra, Ghana', 'active', '2025-09-10 15:02:58.000', '2025-09-10 15:02:58.000', NULL),
(11, 'CIAG', 'Supporting agricultural development and food security initiatives.', 'financial', '/Partners/6-ciag.jpg', 'https://www.ciag.org.gh', 'info@ciag.org.gh', '+233 302 500 000', 'Accra, Ghana', 'active', '2025-09-10 15:02:58.000', '2025-09-10 15:02:58.000', NULL),
(12, 'USAID', 'United States Agency for International Development supporting development programs in Ghana.', 'donor', '/Donors/usaid.png', 'https://www.usaid.gov/ghana', 'ghana@usaid.gov', '+233 302 500 000', 'Accra, Ghana', 'active', '2025-09-10 15:02:58.000', '2025-09-10 15:02:58.000', NULL),
(13, 'AGRA', 'Alliance for a Green Revolution in Africa promoting sustainable agricultural development.', 'donor', '/Donors/agra.png', 'https://www.agra.org', 'info@agra.org', '+233 302 500 000', 'Accra, Ghana', 'active', '2025-09-10 15:02:58.000', '2025-09-10 15:02:58.000', NULL),
(14, 'IFC', 'International Finance Corporation supporting private sector development in Ghana.', 'donor', '/Donors/ifc.png', 'https://www.ifc.org/ghana', 'ghana@ifc.org', '+233 302 500 000', 'Accra, Ghana', 'active', '2025-09-10 15:02:58.000', '2025-09-10 15:02:58.000', NULL),
(15, 'GIZ', 'German development agency supporting sustainable development initiatives.', 'donor', '/Donors/giz-logo.gif', 'https://www.giz.de/en/worldwide/ghana.html', 'info@giz.de', '+233 302 500 000', 'Accra, Ghana', 'active', '2025-09-10 15:02:58.000', '2025-09-10 15:02:58.000', NULL),
(16, 'SNV', 'Netherlands Development Organisation supporting inclusive development.', 'donor', '/Donors/snv.png', 'https://www.snv.org/country/ghana', 'ghana@snv.org', '+233 302 500 000', 'Accra, Ghana', 'active', '2025-09-10 15:02:58.000', '2025-09-10 15:02:58.000', NULL),
(17, 'UNDP', 'United Nations Development Programme supporting sustainable development goals.', 'donor', '/Donors/undp_100.png', 'https://www.gh.undp.org', 'ghana@undp.org', '+233 302 500 000', 'Accra, Ghana', 'active', '2025-09-10 15:02:58.000', '2025-09-10 15:02:58.000', NULL),
(18, 'UNIDO', 'United Nations Industrial Development Organization promoting industrial development.', 'donor', '/Donors/unido.png', 'https://www.unido.org/ghana', 'ghana@unido.org', '+233 302 500 000', 'Accra, Ghana', 'active', '2025-09-10 15:02:58.000', '2025-09-10 15:02:58.000', NULL),
(19, 'IFAD', 'International Fund for Agricultural Development supporting rural development.', 'donor', '/Donors/ifad-a-edit.png', 'https://www.ifad.org/en/ghana', 'ghana@ifad.org', '+233 302 500 000', 'Accra, Ghana', 'active', '2025-09-10 15:02:58.000', '2025-09-10 15:02:58.000', NULL),
(20, 'World Food Programme', 'United Nations agency fighting hunger and promoting food security.', 'donor', '/Donors/world-food-programme-wfp.jpg', 'https://www.wfp.org/countries/ghana', 'ghana@wfp.org', '+233 302 500 000', 'Accra, Ghana', 'active', '2025-09-10 15:02:58.000', '2025-09-10 15:02:58.000', NULL),
(21, 'Youth in Agriculture Programme', 'Government initiative supporting youth participation in agricultural development.', 'government', '/government/youth-in-agriculture-programme-yiap.png', 'https://www.moagd.gov.gh', 'yiap@moagd.gov.gh', '+233 302 500 000', 'Accra, Ghana', 'active', '2025-09-10 15:02:58.000', '2025-09-10 15:02:58.000', NULL),
(22, 'Ghana Incentive-Based Risk-Sharing System for Agricultural Lending', 'Government program facilitating agricultural financing and risk management.', 'government', '/government/ghana-incentive-based-risk-sharing-system-for-agricultural-lending-girsal.png', 'https://www.girsal.com', 'info@girsal.com', '+233 302 500 000', 'Accra, Ghana', 'active', '2025-09-10 15:02:58.000', '2025-09-10 15:02:58.000', NULL),
(23, 'Ghana Investment Promotion Centre', 'Government agency promoting investment opportunities in Ghana.', 'government', '/government/ghana-investment-promotion-centregipc.png', 'https://www.gipc.gov.gh', 'info@gipc.gov.gh', '+233 302 500 000', 'Accra, Ghana', 'active', '2025-09-10 15:02:58.000', '2025-09-10 15:02:58.000', NULL),
(24, 'Ministry of Finance and Economic Planning', 'Government ministry overseeing economic policy and financial planning.', 'government', '/government/ministry-of-finance-and-economic-planning-100.png', 'https://www.mofep.gov.gh', 'info@mofep.gov.gh', '+233 302 500 000', 'Accra, Ghana', 'active', '2025-09-10 15:02:58.000', '2025-09-10 15:02:58.000', NULL),
(25, 'Ministry of Trade and Industry', 'Government ministry promoting trade and industrial development.', 'government', '/government/ministry-of-trade-and-industry-moti-100.png', 'https://www.moti.gov.gh', 'info@moti.gov.gh', '+233 302 500 000', 'Accra, Ghana', 'active', '2025-09-10 15:02:58.000', '2025-09-10 15:02:58.000', NULL),
(26, 'Ghana Made', 'Promoting locally made products and supporting Ghanaian manufacturers.', 'ngo', '/NGO/ghana-made.png', 'https://www.ghanamade.org', 'info@ghanamade.org', '+233 302 500 000', 'Accra, Ghana', 'active', '2025-09-10 15:02:58.000', '2025-09-10 15:02:58.000', NULL),
(27, 'PEF', 'Private Enterprise Federation supporting private sector development.', 'ngo', '/NGO/pef.png', 'https://www.pef.org.gh', 'info@pef.org.gh', '+233 302 500 000', 'Accra, Ghana', 'active', '2025-09-10 15:02:58.000', '2025-09-10 15:02:58.000', NULL),
(28, 'Aflasafe', 'Promoting aflatoxin control and food safety in agricultural products.', 'ngo', '/NGO/aflasafe2.png', 'https://www.aflasafe.com', 'info@aflasafe.com', '+233 302 500 000', 'Accra, Ghana', 'active', '2025-09-10 15:02:58.000', '2025-09-10 15:02:58.000', NULL),
(29, 'Socodevi', 'Supporting cooperative development and rural entrepreneurship.', 'ngo', '/NGO/socodevi.jpeg', 'https://www.socodevi.org', 'ghana@socodevi.org', '+233 302 500 000', 'Accra, Ghana', 'active', '2025-09-10 15:02:58.000', '2025-09-10 15:02:58.000', NULL),
(30, 'Ghana National Chamber of Commerce and Industry', 'Representing the interests of Ghana\'s business community.', 'ngo', '/NGO/ghana-national-chamber-of-commerce-and-industry.jpg', 'https://www.ghanachamber.org', 'info@ghanachamber.org', '+233 302 500 000', 'Accra, Ghana', 'active', '2025-09-10 15:02:58.000', '2025-09-10 15:02:58.000', NULL),
(31, 'Ghana Rice Inter-Professional Body', 'Promoting rice production and processing in Ghana.', 'ngo', '/NGO/ghana-rice-inter-professional-body-grib.png', 'https://www.grib.org.gh', 'info@grib.org.gh', '+233 302 500 000', 'Accra, Ghana', 'active', '2025-09-10 15:02:58.000', '2025-09-10 15:02:58.000', NULL),
(32, 'Peasant Farmers Association of Ghana', 'Representing smallholder farmers and promoting their interests.', 'ngo', '/NGO/peasant-farmers-association-of-ghana-pfag.jpg', 'https://www.pfag.org.gh', 'info@pfag.org.gh', '+233 302 500 000', 'Accra, Ghana', 'active', '2025-09-10 15:02:58.000', '2025-09-10 15:02:58.000', NULL),
(33, 'The John A. Kufuor Foundation', 'Promoting good governance and development initiatives.', 'ngo', '/NGO/the-john-a-kufuor-foundation.jpg', 'https://www.kufuorfoundation.org', 'info@kufuorfoundation.org', '+233 302 500 000', 'Accra, Ghana', 'active', '2025-09-10 15:02:58.000', '2025-09-10 15:02:58.000', NULL),
(34, 'IITA', 'International Institute of Tropical Agriculture supporting agricultural research.', 'ngo', '/NGO/iita.png', 'https://www.iita.org', 'ghana@iita.org', '+233 302 500 000', 'Accra, Ghana', 'active', '2025-09-10 15:02:58.000', '2025-09-10 15:02:58.000', NULL),
(35, 'DMT Collateral', 'Providing collateral management services for agricultural commodities.', 'private', '/Private/dmt-collateral.png', 'https://www.dmtcollateral.com', 'info@dmtcollateral.com', '+233 302 500 000', 'Accra, Ghana', 'active', '2025-09-10 15:02:58.000', '2025-09-10 15:02:58.000', NULL),
(36, 'Intervalle Geneve', 'International trading company supporting agricultural commodity trade.', 'private', '/Private/intervalle-geneve.png', 'https://www.intervalle-geneve.com', 'info@intervalle-geneve.com', '+233 302 500 000', 'Accra, Ghana', 'active', '2025-09-10 15:02:58.000', '2025-09-10 15:02:58.000', NULL),
(37, 'Wienco', 'Agricultural input supplier and commodity trading company.', 'private', '/Private/wienco.jpg', 'https://www.wienco.com', 'info@wienco.com', '+233 302 500 000', 'Accra, Ghana', 'active', '2025-09-10 15:02:58.000', '2025-09-10 15:02:58.000', NULL),
(38, 'TATF', 'Technology and agricultural training foundation supporting farmers.', 'private', '/Private/tatf.png', 'https://www.tatf.org.gh', 'info@tatf.org.gh', '+233 302 500 000', 'Accra, Ghana', 'active', '2025-09-10 15:02:58.000', '2025-09-10 15:02:58.000', NULL),
(39, 'DESS Inc', 'Development and environmental services supporting sustainable agriculture.', 'private', '/Private/dess-inc.png', 'https://www.dessinc.com', 'info@dessinc.com', '+233 302 500 000', 'Accra, Ghana', 'active', '2025-09-10 15:02:58.000', '2025-09-10 15:02:58.000', NULL),
(40, 'Ghana Procurement Authority', 'Government agency managing public procurement processes.', 'tender', '/government/ghana-investment-promotion-centregipc.png', 'https://www.ppa.gov.gh', 'info@ppa.gov.gh', '+233 302 500 000', 'Accra, Ghana', 'active', '2025-09-10 15:02:58.000', '2025-09-10 15:02:58.000', NULL),
(41, 'Public Procurement Regulatory Authority', 'Regulating and monitoring public procurement in Ghana.', 'tender', '/government/ministry-of-finance-and-economic-planning-100.png', 'https://www.ppra.gov.gh', 'info@ppra.gov.gh', '+233 302 500 000', 'Accra, Ghana', 'active', '2025-09-10 15:02:58.000', '2025-09-10 15:02:58.000', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `price_alerts`
--

CREATE TABLE `price_alerts` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `user_id` bigint(20) UNSIGNED DEFAULT NULL,
  `commodity` longtext DEFAULT NULL,
  `target_price` double DEFAULT NULL,
  `condition` longtext DEFAULT NULL,
  `is_active` tinyint(1) DEFAULT 1,
  `triggered_at` datetime(3) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `publications`
--

CREATE TABLE `publications` (
  `id` int(11) NOT NULL,
  `title` longtext NOT NULL,
  `description` text DEFAULT NULL,
  `category` enum('Research Papers','Annual Reports','Policy Documents') NOT NULL,
  `file_path` varchar(500) DEFAULT NULL,
  `file_name` varchar(255) DEFAULT NULL,
  `file_size` bigint(20) DEFAULT NULL,
  `file_type` varchar(100) DEFAULT NULL,
  `publication_date` date DEFAULT NULL,
  `author` varchar(255) DEFAULT NULL,
  `tags` text DEFAULT NULL,
  `status` enum('Published','Draft','Archived') DEFAULT 'Published',
  `download_count` int(11) DEFAULT 0,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `publications`
--

INSERT INTO `publications` (`id`, `title`, `description`, `category`, `file_path`, `file_name`, `file_size`, `file_type`, `publication_date`, `author`, `tags`, `status`, `download_count`, `created_at`, `updated_at`) VALUES
(1, 'GCX Market Analysis 2024', 'Comprehensive analysis of commodity markets and trading patterns for 2024', 'Research Papers', '/publications/research/gcx-market-analysis-2024.pdf', 'gcx-market-analysis-2024.pdf', 2048576, 'application/pdf', '2024-12-01', 'GCX Research Team', 'market analysis, commodities, trading', 'Published', 0, '2025-09-05 13:45:50.000', '2025-09-05 13:45:50.000'),
(2, 'Annual Report 2023', 'GCX Annual Performance Report for the year 2023', 'Annual Reports', '/publications/annual/gcx-annual-report-2023.pdf', 'gcx-annual-report-2023.pdf', 5120000, 'application/pdf', '2024-03-15', 'GCX Management', 'annual report, performance, financial', 'Published', 0, '2025-09-05 13:45:50.000', '2025-09-05 13:45:50.000'),
(3, 'Trading Rules and Regulations', 'Updated trading rules and regulations for GCX members', 'Policy Documents', '/publications/policy/trading-rules-2024.pdf', 'trading-rules-2024.pdf', 1536000, 'application/pdf', '2024-01-01', 'GCX Legal Team', 'trading rules, regulations, policy', 'Published', 0, '2025-09-05 13:45:50.000', '2025-09-05 13:45:50.000'),
(4, 'Maize Market Outlook 2024', 'Research paper on maize market trends and future outlook', 'Research Papers', '/publications/research/maize-outlook-2024.pdf', 'maize-outlook-2024.pdf', 1024000, 'application/pdf', '2024-06-15', 'Dr. Sarah Mensah', 'maize, market outlook, research', 'Published', 0, '2025-09-05 13:45:50.000', '2025-09-05 13:45:50.000'),
(5, 'GCX Compliance Guidelines', 'Guidelines for member compliance and regulatory requirements', 'Policy Documents', '/publications/policy/compliance-guidelines.pdf', 'compliance-guidelines.pdf', 768000, 'application/pdf', '2024-02-01', 'GCX Compliance Team', 'compliance, guidelines, regulatory', 'Published', 0, '2025-09-05 13:45:50.000', '2025-09-05 13:45:50.000');

-- --------------------------------------------------------

--
-- Table structure for table `settings`
--

CREATE TABLE `settings` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `key` varchar(100) NOT NULL,
  `value` text DEFAULT NULL,
  `type` varchar(50) DEFAULT 'string',
  `group` varchar(50) DEFAULT NULL,
  `label` varchar(200) DEFAULT NULL,
  `description` text DEFAULT NULL,
  `is_public` tinyint(1) DEFAULT 0,
  `sort_order` bigint(20) DEFAULT 0,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `settings`
--

INSERT INTO `settings` (`id`, `key`, `value`, `type`, `group`, `label`, `description`, `is_public`, `sort_order`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'site_name', 'Ghana Commodity Exchange', 'text', 'general', 'Site Name', 'The name of your website', 1, 1, '2025-09-04 09:42:10.066', '2025-09-04 09:42:10.066', NULL),
(2, 'site_tagline', 'Connecting Markets, Connecting People, Providing Opportunities', 'text', 'general', 'Site Tagline', 'Your site tagline or motto', 1, 2, '2025-09-04 09:42:10.090', '2025-09-04 09:42:10.090', NULL),
(3, 'site_logo', '/logo_black.png', 'image', 'general', 'Site Logo', 'Your site logo', 1, 3, '2025-09-04 09:42:10.098', '2025-09-04 09:42:10.098', NULL),
(4, 'contact_email', 'info@gcx.com.gh', 'email', 'contact', 'Contact Email', 'Main contact email address', 1, 1, '2025-09-04 09:42:10.103', '2025-09-04 09:42:10.103', NULL),
(5, 'contact_phone', '+233 302 123 456', 'text', 'contact', 'Contact Phone', 'Main contact phone number', 1, 2, '2025-09-04 09:42:10.106', '2025-09-04 09:42:10.106', NULL),
(6, 'hero', '{\"title\":\"hana Commodity Exchange\",\"subtitle\":\"Connecting Markets, Connecting People, Providing Opportunities\",\"cta_primary_text\":\"Start Trading\",\"cta_primary_url\":\"/trading\",\"cta_secondary_text\":\"View Platform\",\"cta_secondary_url\":\"/platform\",\"background_image_1\":\"/trading dashboard.jpg\",\"background_image_2\":\"/crop.jpg\",\"background_image_3\":\"/trading.jpg\"}', 'json', 'homepage', 'Homepage hero', 'Homepage section data for hero', 1, 0, '2025-09-04 10:05:44.152', '2025-09-04 10:05:44.152', NULL),
(7, 'about_content', '{\"about_title\":\"abouting Gcx\"}', 'json', 'page_content', 'About Page Content', 'Content for about page', 1, 0, '2025-09-04 10:33:48.930', '2025-09-04 10:33:48.930', '2025-09-04 11:34:46.515');

-- --------------------------------------------------------

--
-- Table structure for table `subscription_features`
--

CREATE TABLE `subscription_features` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `plan_id` bigint(20) UNSIGNED NOT NULL,
  `name` longtext NOT NULL,
  `description` text DEFAULT NULL,
  `is_enabled` tinyint(1) DEFAULT 1,
  `limit` bigint(20) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `subscription_plans`
--

CREATE TABLE `subscription_plans` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `name` longtext NOT NULL,
  `description` text DEFAULT NULL,
  `price` double NOT NULL,
  `currency` varchar(191) DEFAULT 'GHS',
  `duration` bigint(20) NOT NULL,
  `features` text DEFAULT NULL,
  `max_users` bigint(20) DEFAULT 1,
  `is_active` tinyint(1) DEFAULT 1,
  `sort_order` bigint(20) DEFAULT 0,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `team_members`
--

CREATE TABLE `team_members` (
  `id` int(11) NOT NULL,
  `name` longtext NOT NULL,
  `title` longtext NOT NULL,
  `description` text DEFAULT NULL,
  `image` longtext DEFAULT NULL,
  `type` varchar(191) NOT NULL,
  `order_index` int(11) DEFAULT 0,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `linkedin_url` varchar(500) DEFAULT NULL,
  `twitter_url` varchar(500) DEFAULT NULL,
  `facebook_url` varchar(500) DEFAULT NULL,
  `instagram_url` varchar(500) DEFAULT NULL,
  `linked_in_url` varchar(500) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `team_members`
--

INSERT INTO `team_members` (`id`, `name`, `title`, `description`, `image`, `type`, `order_index`, `created_at`, `updated_at`, `linkedin_url`, `twitter_url`, `facebook_url`, `instagram_url`, `linked_in_url`) VALUES
(32, 'Mr. Kwame Daaku', 'Non-Executive Director', 'Independent director bringing valuable industry expertise and governance experience.', '/Board of directors/Mr. Kwame Daaku \'Non-excutive director\'.jpg', 'board', 2, '2025-09-05 11:58:42.000', '2025-09-05 11:58:42.000', NULL, NULL, NULL, NULL, NULL),
(33, 'Mr. Stephen Antwi-Asimeng', 'Non-Executive Director', 'Independent director with extensive experience in financial markets and corporate governance.', '/Board of directors/Mr. Stephen Antwi-Asimeng \' non-excutive director\'.png', 'board', 3, '2025-09-05 11:58:42.000', '2025-09-05 11:58:42.000', NULL, NULL, NULL, NULL, NULL),
(34, 'Mrs. Wendy Malm', 'Board Secretary', 'Experienced professional ensuring effective board governance and compliance.', '/Board of directors/Mrs. Wendy Malm Board Secretary.png', 'board', 4, '2025-09-05 11:58:42.000', '2025-09-05 11:58:42.000', NULL, NULL, NULL, NULL, NULL),
(35, 'Mr. Robert Dowuona', 'Acting Chief Executive Officer', 'Experienced leader with extensive background in commodity trading and financial markets.', '/Mr. Robert Dowuona Owoo.jpeg', 'executive', 1, '2025-09-05 11:58:42.000', '2025-09-05 12:15:11.000', NULL, NULL, NULL, NULL, NULL),
(36, 'Ms. Ophelia Martekuo Atoklo', 'Acting Deputy Chief Executive Officer', 'Strategic leader with deep expertise in operational excellence and stakeholder management.', '/Ms. Ophelia Martekuo Atoklo \'Deputy Chief Executive Officer\'.jpg', 'executive', 2, '2025-09-05 11:58:42.000', '2025-09-05 11:58:42.000', NULL, NULL, NULL, NULL, NULL),
(37, 'Mr. Opoku Debrah', 'Internal Auditor', 'Ensuring compliance and risk management across all organizational processes.', '/Functional Heads/Mr. Opoku Debrah (Internal Auditor).jpg', 'functional', 1, '2025-09-05 11:58:42.000', '2025-09-05 11:58:42.000', NULL, NULL, NULL, NULL, NULL),
(38, 'Mr. Richard Ankrah', 'Special Project', 'Leading strategic initiatives and special projects to drive organizational growth.', '/Functional Heads/Mr. Richard Ankrah ( Special Project).jpg', 'functional', 2, '2025-09-05 11:58:42.000', '2025-09-05 11:58:42.000', NULL, NULL, NULL, NULL, NULL),
(39, 'Mr. Vitus Ninfaakang', 'Risk, Membership and Partnership', 'Managing risk assessment and building strategic partnerships with key stakeholders.', '/Functional Heads/Mr. Vitus Ninfaakang (Risk, Membership and Partnership).jpg', 'functional', 3, '2025-09-05 11:58:42.000', '2025-09-05 11:58:42.000', NULL, NULL, NULL, NULL, NULL),
(40, 'Mrs. Jemimah Naa Adjeley Oppong-Gyamfi', 'Corporate Services', 'Overseeing corporate governance and administrative excellence.', '/Functional Heads/Mrs. Jemimah Naa Adjeley Oppong-Gyamfi ( Corporate Services).jpg', 'functional', 4, '2025-09-05 11:58:42.000', '2025-09-05 11:58:42.000', NULL, NULL, NULL, NULL, NULL),
(41, 'Mrs. Wendy Malm', 'Operations', 'Ensuring smooth operational processes and service delivery excellence.', '/Functional Heads/Mrs. Wendy Malm (Operations).jpg', 'functional', 5, '2025-09-05 11:58:42.000', '2025-09-05 11:58:42.000', NULL, NULL, NULL, NULL, NULL),
(42, 'Mr. Godfred Kofi Nyamekye', 'Value Chain and Product Development', 'Driving innovation in product development and value chain optimization.', '/Functional Heads/Mr. Godfred Kofi Nyamekye (Value Chain and Product Development).jpg', 'functional', 6, '2025-09-05 11:58:42.000', '2025-09-05 11:58:42.000', NULL, NULL, NULL, NULL, NULL),
(43, 'Mr. Gabriel Aryeetey', 'Warehouse & Quality', 'Maintaining quality standards and efficient warehouse operations.', '/Functional Heads/Mr. Gabriel Aryeetey (Warehouse & Quality).jpg', 'functional', 7, '2025-09-05 11:58:42.000', '2025-09-05 11:58:42.000', NULL, NULL, NULL, NULL, NULL),
(44, 'Dr. Harold Okai-Tettey', 'Information Technology and Information System', 'Leading digital transformation and technology infrastructure development.', '/Functional Heads/Dr. Harold Okai-Tettey ( Information Technology and Information System).jpg', 'functional', 8, '2025-09-05 11:58:42.000', '2025-09-05 11:58:42.000', NULL, NULL, NULL, NULL, NULL),
(45, 'Mr. Albert Nii Ayi Tagoe', 'Finance and Investments', 'Managing financial strategy and investment opportunities for sustainable growth.', '/Functional Heads/Mr. Albert Nii Ayi Tagoe (Finance and Investments).jpg', 'functional', 9, '2025-09-05 11:58:42.000', '2025-09-05 11:58:42.000', NULL, NULL, NULL, NULL, NULL),
(46, 'test', 'testing ', 'testing a feat', '', 'functional', 10, '2025-09-05 12:04:24.000', '2025-09-05 12:04:24.000', NULL, NULL, NULL, NULL, NULL);

-- --------------------------------------------------------

--
-- Table structure for table `traders`
--

CREATE TABLE `traders` (
  `id` int(11) NOT NULL,
  `name` longtext NOT NULL,
  `industry` varchar(255) DEFAULT NULL,
  `member_type` enum('Associates','Full Members','Brokers','Warehouse Operators') NOT NULL DEFAULT 'Associates',
  `phone_no` varchar(50) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `address` text DEFAULT NULL,
  `registration_date` date DEFAULT NULL,
  `status` enum('Active','Inactive','Suspended') DEFAULT 'Active',
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `traders`
--

INSERT INTO `traders` (`id`, `name`, `industry`, `member_type`, `phone_no`, `email`, `address`, `registration_date`, `status`, `created_at`, `updated_at`) VALUES
(1, 'Eric Kwabena Agyei', '-', 'Associates', '0244 692 089/ 0245 071 061', 'eric.agyei@example.com', 'Accra, Ghana', '2024-01-15', 'Active', '2025-09-05 12:58:46.000', '2025-09-05 12:58:46.000'),
(2, 'Joseph Awudu Malik', '-', 'Associates', '0244 869 387', 'joseph.malik@example.com', 'Kumasi, Ghana', '2024-01-20', 'Active', '2025-09-05 12:58:46.000', '2025-09-05 12:58:46.000'),
(3, 'Kofi Adusei Koduah', '-', 'Associates', '0540 122295', 'kofi.koduah@example.com', 'Tamale, Ghana', '2024-02-01', 'Active', '2025-09-05 12:58:46.000', '2025-09-05 12:58:46.000'),
(4, 'Kwabena Duah Agyemang', '-', 'Associates', '0201 689497', 'kwabena.agyemang@example.com', 'Cape Coast, Ghana', '2024-02-10', 'Active', '2025-09-05 12:58:46.000', '2025-09-05 12:58:46.000'),
(5, 'Maame Adjoa Thompson', '-', 'Associates', '0302 200748/ 0266 802388', 'maame.thompson@example.com', 'Tema, Ghana', '2024-02-15', 'Active', '2025-09-05 12:58:46.000', '2025-09-05 12:58:46.000'),
(6, 'Monica . .', '-', 'Associates', '0269 382146', 'monica@example.com', 'Takoradi, Ghana', '2024-02-20', 'Active', '2025-09-05 12:58:46.000', '2025-09-05 12:58:46.000'),
(7, 'Praise Awisi Bogobley', '-', 'Associates', '0244 590 621', 'praise.bogobley@example.com', 'Ho, Ghana', '2024-03-01', 'Active', '2025-09-05 12:58:46.000', '2025-09-05 12:58:46.000'),
(8, 'Roland Apindem Ajiabadek', '-', 'Associates', '020 535 1223', 'roland.ajiabadek@example.com', 'Koforidua, Ghana', '2024-03-05', 'Active', '2025-09-05 12:58:46.000', '2025-09-05 12:58:46.000');

-- --------------------------------------------------------

--
-- Table structure for table `trading_sessions`
--

CREATE TABLE `trading_sessions` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `date` datetime(3) DEFAULT NULL,
  `open_time` datetime(3) DEFAULT NULL,
  `close_time` datetime(3) DEFAULT NULL,
  `is_open` tinyint(1) DEFAULT 0,
  `status` longtext DEFAULT NULL,
  `volume` double DEFAULT NULL,
  `transactions` bigint(20) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `translations`
--

CREATE TABLE `translations` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `resource_type` longtext DEFAULT NULL,
  `resource_id` bigint(20) UNSIGNED DEFAULT NULL,
  `language_code` longtext DEFAULT NULL,
  `field_name` longtext DEFAULT NULL,
  `content` longtext DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `name` longtext NOT NULL,
  `email` longtext NOT NULL,
  `password` longtext NOT NULL,
  `role` varchar(191) DEFAULT 'user',
  `avatar` longtext DEFAULT NULL,
  `bio` longtext DEFAULT NULL,
  `is_active` tinyint(1) DEFAULT 1,
  `last_login` datetime(3) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `company` longtext DEFAULT NULL,
  `phone` longtext DEFAULT NULL,
  `country` longtext DEFAULT NULL,
  `time_zone` longtext DEFAULT NULL,
  `preferences` text DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `name`, `email`, `password`, `role`, `avatar`, `bio`, `is_active`, `last_login`, `created_at`, `updated_at`, `company`, `phone`, `country`, `time_zone`, `preferences`) VALUES
(1, 'GCX Admin', 'admin@gcx.com', '$2a$10$X3YMVQ4cILJlWyb0wQPV/O3VVMKtSXixGYJMTTpdwewP29X/3srOy', 'admin', NULL, NULL, 1, '2025-09-16 13:29:23.695', '2025-08-18 09:31:40.346', '2025-09-16 13:29:23.695', NULL, NULL, NULL, NULL, ''),
(2, 'test', 'test@gmail.com', '$2a$10$0FUxgqZss12LXM6ZSYwNquWQv5gbWSX8vCzVA8eJPI9VK.3GrxQ9e', 'user', NULL, NULL, 1, '2025-09-12 10:12:12.889', '2025-08-21 16:08:08.409', '2025-09-12 10:12:12.889', NULL, NULL, NULL, NULL, '');

-- --------------------------------------------------------

--
-- Table structure for table `user_data_access`
--

CREATE TABLE `user_data_access` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `user_id` bigint(20) UNSIGNED NOT NULL,
  `data_type` longtext NOT NULL,
  `access_level` longtext NOT NULL,
  `commodities` text DEFAULT NULL,
  `max_requests` bigint(20) DEFAULT NULL,
  `request_count` bigint(20) DEFAULT 0,
  `last_reset` datetime(3) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `user_subscriptions`
--

CREATE TABLE `user_subscriptions` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `user_id` bigint(20) UNSIGNED NOT NULL,
  `plan_id` bigint(20) UNSIGNED NOT NULL,
  `status` varchar(191) DEFAULT 'active',
  `start_date` datetime(3) NOT NULL,
  `end_date` datetime(3) NOT NULL,
  `auto_renew` tinyint(1) DEFAULT 1,
  `payment_method` longtext DEFAULT NULL,
  `payment_reference` longtext DEFAULT NULL,
  `amount_paid` double DEFAULT NULL,
  `currency` varchar(191) DEFAULT 'GHS',
  `last_billing_date` datetime(3) DEFAULT NULL,
  `next_billing_date` datetime(3) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `blog_categories`
--
ALTER TABLE `blog_categories`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `idx_blog_categories_slug` (`slug`(191));

--
-- Indexes for table `blog_posts`
--
ALTER TABLE `blog_posts`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `idx_blog_posts_slug` (`slug`(191)),
  ADD KEY `fk_blog_posts_author` (`author_id`);

--
-- Indexes for table `board_members`
--
ALTER TABLE `board_members`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_board_members_deleted_at` (`deleted_at`);

--
-- Indexes for table `brokers`
--
ALTER TABLE `brokers`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `careers`
--
ALTER TABLE `careers`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `commodities`
--
ALTER TABLE `commodities`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `idx_commodities_code` (`code`(191));

--
-- Indexes for table `commodity_info`
--
ALTER TABLE `commodity_info`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `idx_commodity_info_code` (`code`(191));

--
-- Indexes for table `languages`
--
ALTER TABLE `languages`
  ADD PRIMARY KEY (`code`),
  ADD KEY `idx_languages_deleted_at` (`deleted_at`);

--
-- Indexes for table `market_analytics`
--
ALTER TABLE `market_analytics`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `market_data`
--
ALTER TABLE `market_data`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `media_files`
--
ALTER TABLE `media_files`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `uni_media_files_filename` (`filename`),
  ADD KEY `idx_media_files_deleted_at` (`deleted_at`),
  ADD KEY `fk_media_files_user` (`uploaded_by`);

--
-- Indexes for table `menus`
--
ALTER TABLE `menus`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_menus_deleted_at` (`deleted_at`);

--
-- Indexes for table `menu_items`
--
ALTER TABLE `menu_items`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_menu_items_deleted_at` (`deleted_at`),
  ADD KEY `fk_menu_items_children` (`parent_id`),
  ADD KEY `fk_menus_items` (`menu_id`);

--
-- Indexes for table `pages`
--
ALTER TABLE `pages`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `idx_pages_slug` (`slug`(191)),
  ADD KEY `idx_pages_deleted_at` (`deleted_at`),
  ADD KEY `fk_pages_children` (`parent_id`),
  ADD KEY `fk_pages_author` (`author_id`);

--
-- Indexes for table `partners`
--
ALTER TABLE `partners`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_partners_deleted_at` (`deleted_at`);

--
-- Indexes for table `price_alerts`
--
ALTER TABLE `price_alerts`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_price_alerts_user` (`user_id`);

--
-- Indexes for table `publications`
--
ALTER TABLE `publications`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `settings`
--
ALTER TABLE `settings`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `idx_settings_key` (`key`),
  ADD KEY `idx_settings_deleted_at` (`deleted_at`),
  ADD KEY `idx_settings_group` (`group`);

--
-- Indexes for table `subscription_features`
--
ALTER TABLE `subscription_features`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_subscription_features_plan` (`plan_id`);

--
-- Indexes for table `subscription_plans`
--
ALTER TABLE `subscription_plans`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `team_members`
--
ALTER TABLE `team_members`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_type` (`type`),
  ADD KEY `idx_order` (`type`,`order_index`),
  ADD KEY `idx_team_members_type` (`type`);

--
-- Indexes for table `traders`
--
ALTER TABLE `traders`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `trading_sessions`
--
ALTER TABLE `trading_sessions`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `idx_trading_sessions_date` (`date`);

--
-- Indexes for table `translations`
--
ALTER TABLE `translations`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_translations_deleted_at` (`deleted_at`),
  ADD KEY `fk_menu_items_translations` (`resource_id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `idx_users_email` (`email`(191));

--
-- Indexes for table `user_data_access`
--
ALTER TABLE `user_data_access`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_user_data_access_user` (`user_id`);

--
-- Indexes for table `user_subscriptions`
--
ALTER TABLE `user_subscriptions`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_user_subscriptions_user` (`user_id`),
  ADD KEY `fk_user_subscriptions_plan` (`plan_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `blog_categories`
--
ALTER TABLE `blog_categories`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `blog_posts`
--
ALTER TABLE `blog_posts`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `board_members`
--
ALTER TABLE `board_members`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `brokers`
--
ALTER TABLE `brokers`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT for table `careers`
--
ALTER TABLE `careers`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT for table `commodities`
--
ALTER TABLE `commodities`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT for table `commodity_info`
--
ALTER TABLE `commodity_info`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `market_analytics`
--
ALTER TABLE `market_analytics`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `market_data`
--
ALTER TABLE `market_data`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `media_files`
--
ALTER TABLE `media_files`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;

--
-- AUTO_INCREMENT for table `menus`
--
ALTER TABLE `menus`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `menu_items`
--
ALTER TABLE `menu_items`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `pages`
--
ALTER TABLE `pages`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=16;

--
-- AUTO_INCREMENT for table `partners`
--
ALTER TABLE `partners`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=42;

--
-- AUTO_INCREMENT for table `price_alerts`
--
ALTER TABLE `price_alerts`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `publications`
--
ALTER TABLE `publications`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT for table `settings`
--
ALTER TABLE `settings`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;

--
-- AUTO_INCREMENT for table `subscription_features`
--
ALTER TABLE `subscription_features`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `subscription_plans`
--
ALTER TABLE `subscription_plans`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `team_members`
--
ALTER TABLE `team_members`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=47;

--
-- AUTO_INCREMENT for table `traders`
--
ALTER TABLE `traders`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;

--
-- AUTO_INCREMENT for table `trading_sessions`
--
ALTER TABLE `trading_sessions`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `translations`
--
ALTER TABLE `translations`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `user_data_access`
--
ALTER TABLE `user_data_access`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `user_subscriptions`
--
ALTER TABLE `user_subscriptions`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `blog_posts`
--
ALTER TABLE `blog_posts`
  ADD CONSTRAINT `fk_blog_posts_author` FOREIGN KEY (`author_id`) REFERENCES `users` (`id`),
  ADD CONSTRAINT `fk_users_blog_posts` FOREIGN KEY (`author_id`) REFERENCES `users` (`id`);

--
-- Constraints for table `media_files`
--
ALTER TABLE `media_files`
  ADD CONSTRAINT `fk_media_files_user` FOREIGN KEY (`uploaded_by`) REFERENCES `users` (`id`);

--
-- Constraints for table `menu_items`
--
ALTER TABLE `menu_items`
  ADD CONSTRAINT `fk_menu_items_children` FOREIGN KEY (`parent_id`) REFERENCES `menu_items` (`id`),
  ADD CONSTRAINT `fk_menus_items` FOREIGN KEY (`menu_id`) REFERENCES `menus` (`id`) ON DELETE CASCADE;

--
-- Constraints for table `pages`
--
ALTER TABLE `pages`
  ADD CONSTRAINT `fk_pages_author` FOREIGN KEY (`author_id`) REFERENCES `users` (`id`),
  ADD CONSTRAINT `fk_pages_children` FOREIGN KEY (`parent_id`) REFERENCES `pages` (`id`);

--
-- Constraints for table `price_alerts`
--
ALTER TABLE `price_alerts`
  ADD CONSTRAINT `fk_price_alerts_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

--
-- Constraints for table `subscription_features`
--
ALTER TABLE `subscription_features`
  ADD CONSTRAINT `fk_subscription_features_plan` FOREIGN KEY (`plan_id`) REFERENCES `subscription_plans` (`id`);

--
-- Constraints for table `translations`
--
ALTER TABLE `translations`
  ADD CONSTRAINT `fk_menu_items_translations` FOREIGN KEY (`resource_id`) REFERENCES `menu_items` (`id`) ON DELETE CASCADE,
  ADD CONSTRAINT `fk_pages_translations` FOREIGN KEY (`resource_id`) REFERENCES `pages` (`id`) ON DELETE CASCADE,
  ADD CONSTRAINT `fk_settings_translations` FOREIGN KEY (`resource_id`) REFERENCES `settings` (`id`) ON DELETE CASCADE;

--
-- Constraints for table `user_data_access`
--
ALTER TABLE `user_data_access`
  ADD CONSTRAINT `fk_user_data_access_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

--
-- Constraints for table `user_subscriptions`
--
ALTER TABLE `user_subscriptions`
  ADD CONSTRAINT `fk_user_subscriptions_plan` FOREIGN KEY (`plan_id`) REFERENCES `subscription_plans` (`id`),
  ADD CONSTRAINT `fk_user_subscriptions_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
