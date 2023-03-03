CREATE TABLE IF NOT EXISTS `messages` (
	`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	`match_id` int UNSIGNED NOT NULL,
	`sender_id` int UNSIGNED NOT NULL,
	`receiver_id` int UNSIGNED NOT NULL,
	`message` text NOT NULL,
	`created_at` int UNSIGNED NOT NULL,
	`updated_at` int UNSIGNED NOT NULL,
	`deleted_at` int NULL,
	FOREIGN KEY (match_id) REFERENCES matches(id),
	FOREIGN KEY (sender_id) REFERENCES users(id),
	FOREIGN KEY (receiver_id) REFERENCES users(id),
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
