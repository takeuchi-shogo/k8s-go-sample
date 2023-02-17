CREATE TABLE IF NOT EXISTS `accounts` (
	`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	`phone_number` int UNSIGNED NOT NULL,
	`email` varchar(255) NOT NULL,
	`password` varchar(255) NOT NULL,
	`login_status` varchar(255) NOT NULL,
	`access_level` varchar(255) NULL,
	`created_at` int UNSIGNED NOT NULL,
	`updated_at` int UNSIGNED NOT NULL,
	`deleted_at` int UNSIGNED NULL,
	UNIQUE (email),
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
