SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

SET NAMES utf8mb4;

DROP TABLE IF EXISTS `table_truncate1`;
CREATE TABLE `table_truncate1` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

INSERT INTO `table_truncate1` (`id`) VALUES (1), (2), (3);


DROP TABLE IF EXISTS `table_truncate2`;
CREATE TABLE `table_truncate2` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `delete_me` tinyint(4) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

INSERT INTO `table_truncate2` (`id`, `delete_me`) VALUES
(1,	1),
(2,	1),
(3,	0);

DROP TABLE IF EXISTS `table_update`;
CREATE TABLE `table_update` (
  `id` int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `col_string` varchar(255) NULL,
  `col_bool` int NULL,
  `col_int` int NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

INSERT INTO `table_update` (`id`, `col_string`, `col_bool`, `col_int`) VALUES
(1,	'foo',	1,	1),
(2,	'bar',	0,	2),
(3,	NULL,	NULL,	NULL);
