CREATE TABLE IF NOT EXISTS `likes` (
	`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	`send_user_id` int UNSIGNED NOT NULL,
	`receive_user_id` int UNSIGNED NOT NULL,
	`created_at` int UNSIGNED NOT NULL,
	FOREIGN KEY (send_user_id) REFERENCES users(id),
	FOREIGN KEY (receive_user_id) REFERENCES users(id),
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
