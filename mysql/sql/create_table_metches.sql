CREATE TABLE IF NOT EXISTS `matches` (
	`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	`male_user_id` int UNSIGNED NOT NULL,
	`female_user_id` int UNSIGNED NOT NULL,
	`created_at` int UNSIGNED NOT NULL,
	FOREIGN KEY (male_user_id) REFERENCES users(id),
	FOREIGN KEY (female_user_id) REFERENCES users(id),
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
