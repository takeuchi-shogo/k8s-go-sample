CREATE TABLE IF NOT EXISTS `user_profiles` (
	`id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	`user_id` int UNSIGNED NOT NULL,
	`introduction` varchar(255) NULL,
	`height_id` int UNSIGNED NOT NULL,
	`body_type_id` int UNSIGNED NOT NULL,
	`blood_type_id` int UNSIGNED NOT NULL,
	`residence_id` int UNSIGNED NOT NULL,
	`hometown_id` int UNSIGNED NOT NULL,
	`occupation_id` int UNSIGNED NOT NULL,
	`education_id` int UNSIGNED NOT NULL,
	`annual_income_id` int UNSIGNED NOT NULL,
	`smoking_id` int UNSIGNED NOT NULL,
	`drinking_id` int UNSIGNED NOT NULL,
-- 
	`siblings_id` int UNSIGNED NOT NULL,
	`language_id` int UNSIGNED NOT NULL,
	`interests_id` int UNSIGNED NOT NULL,
	`looking_for_id` int UNSIGNED NOT NULL,
-- 
	`school_name` varchar(255) NULL,
	`job_title` varchar(255) NULL,
-- 
	`marital_history_id` int UNSIGNED NOT NULL,
	`presence_of_children_id` int UNSIGNED NOT NULL,
	`intentions_towards_marriage_id` int UNSIGNED NOT NULL,
	`desire_for_children_id` int UNSIGNED NOT NULL,
	`household_chores_and_child_rearing_id` int UNSIGNED NOT NULL,
	`ideal_first_encounter_id` int UNSIGNED NOT NULL,
	`dating_expenses_id` int UNSIGNED NOT NULL,
-- 
	`personality_type_id` int UNSIGNED NOT NULL,
	`sociability_id` int UNSIGNED NOT NULL,
	`roommates_id` int UNSIGNED NOT NULL,
	`days_off_id` int UNSIGNED NOT NULL,
	`hobbies_id` int UNSIGNED NOT NULL,
-- 
	`created_at` int UNSIGNED NOT NULL ,
	`updated_at` int UNSIGNED NOT NULL,
	`deleted_at` int UNSIGNED NULL,
	FOREIGN KEY (user_id) REFERENCES users(id),
	UNIQUE(`user_id`),
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
