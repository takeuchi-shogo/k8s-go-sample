CREATE TABLE IF NOT EXISTS `user_search_filters` (
	`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	`user_id`  int UNSIGNED NOT NULL,
	`gender` varchar(255) NULL,
	`location` varchar(255) NULL,
	`created_at` int UNSIGNED NOT NULL,
	`updated_at` int UNSIGNED NOT NULL,
	FOREIGN KEY (user_id) REFERENCES users(id),
	UNIQUE(`user_id`),
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
