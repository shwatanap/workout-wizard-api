CREATE TABLE IF NOT EXISTS `menus` (
  `id` int(5) UNSIGNED PRIMARY KEY AUTO_INCREMENT,
  `user_id` int(10) UNSIGNED NOT NULL,
  `target` varchar(255) NOT NULL,
  `comment` varchar(255) NOT NULL
);
