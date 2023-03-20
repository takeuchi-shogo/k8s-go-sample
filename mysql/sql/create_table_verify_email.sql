CREATE TABLE IF NOT EXISTS `verify_emails` (
	`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	-- `age` INT(10) UNSIGNED NOT NULL,
	`email` varchar(255) NOT NULL,
	`token` varchar(255) NOT NULL,
	`pin_code` varchar(255) NOT NULL,
	`expire_at` int UNSIGNED NOT NULL,
	`created_at` int UNSIGNED NOT NULL,
	UNIQUE(`token`),
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
