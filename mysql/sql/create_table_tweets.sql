CREATE TABLE IF NOT EXISTS `tweets` (
	`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	`user_id` int UNSIGNED NOT NULL,
	`text` text NOT NULL,
	`created_at` int UNSIGNED NOT NULL,
	`updated_at` int UNSIGNED NOT NULL,
	`deleted_at` int UNSIGNED NULL,
	FOREIGN KEY (user_id) REFERENCES users(id),
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
