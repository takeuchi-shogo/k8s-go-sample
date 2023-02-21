CREATE TABLE IF NOT EXISTS `users` (
	`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	-- `age` INT(10) UNSIGNED NOT NULL,
	`uuid` varchar(255) NOT NULL,
	`account_id` int UNSIGNED NOT NULL,
	`display_name` varchar(255) NOT NULL,
	`screen_name` varchar(255) NOT NULL,
	`gender` varchar(255) NOT NULL,
	`location` varchar(255) NULL,
	`is_authorize_email` tinyint UNSIGNED NOT NULL,
	`is_verified_age` tinyint UNSIGNED NOT NULL,
	`created_at` int UNSIGNED NOT NULL,
	`updated_at` int UNSIGNED NOT NULL,
	`deleted_at` int UNSIGNED NULL,
	FOREIGN KEY (account_id) REFERENCES accounts(id),
	UNIQUE (uuid),
	UNIQUE (screen_name),
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
