CREATE TABLE IF NOT EXISTS `menus` (
  `id` int(5) UNSIGNED PRIMARY KEY AUTO_INCREMENT,
  `user_id` int(10) UNSIGNED NOT NULL,
  `target` varchar(255) NOT NULL,
  `comment` varchar(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS `workouts` (
  `id` int(5) UNSIGNED PRIMARY KEY AUTO_INCREMENT,
  `menu_id` int(10) UNSIGNED NOT NULL,
  `name` varchar(255) NOT NULL,
  `detail` varchar(1027) NOT NULL
);
