CREATE TABLE IF NOT EXISTS `reports` (
	`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	`reporter_id` int UNSIGNED NOT NULL,
	`reported_id` int UNSIGNED NOT NULL,
	`reason` text NULL,
	`created_at` int UNSIGNED NOT NULL,
	FOREIGN KEY (male_user_id) REFERENCES users(id),
	FOREIGN KEY (female_user_id) REFERENCES users(id),
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
