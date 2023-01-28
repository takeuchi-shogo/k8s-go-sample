CREATE TABLE IF NOT EXISTS `users` (
	`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	`name` varchar(50) NOT NULL,
	-- `screen_name` varchar(50) NOT NULL,
	`age` INT(10) UNSIGNED NOT NULL,
	`email` varchar(30),
	`password` varchar(255) NOT NULL, 
	`created_at` int UNSIGNED NOT NULL,
	`updated_at` int UNSIGNED NOT NULL,
	`deleted_at` int UNSIGNED NULL,
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
