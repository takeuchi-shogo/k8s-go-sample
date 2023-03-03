CREATE TABLE IF NOT EXISTS `blocks` (
	`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	`blocking_id` int UNSIGNED NOT NULL,
	`blocked_id` int UNSIGNED NOT NULL,
	`created_at` int UNSIGNED NOT NULL,
	FOREIGN KEY (blocking_id) REFERENCES users(id),
	FOREIGN KEY (blocked_id) REFERENCES users(id),
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
