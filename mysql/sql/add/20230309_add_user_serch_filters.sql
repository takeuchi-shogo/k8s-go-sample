ALTER TABLE `user_search_filters`
	ADD `height_id` int UNSIGNED NOT NULL AFTER `has_introduction`,
	ADD `body_type_id` int UNSIGNED NOT NULL AFTER `height_id`,
	ADD `blood_type_id` int UNSIGNED NOT NULL AFTER `body_type_id`,
	ADD `residence_country_id` int UNSIGNED NOT NULL AFTER `blood_type_id`,
	ADD `residence_state_id` int UNSIGNED NOT NULL AFTER `residence_country_id`,
	ADD `hometown_country_id` int UNSIGNED NOT NULL AFTER `residence_state_id`,
	ADD `hometown_state_id` int UNSIGNED NOT NULL AFTER `hometown_country_id`,
	ADD `occupation_id` int UNSIGNED NOT NULL AFTER `hometown_state_id`,
	ADD `education_id` int UNSIGNED NOT NULL AFTER `occupation_id`,
	ADD `annual_income_id` int UNSIGNED NOT NULL AFTER `education_id`,
	ADD `smoking_id` int UNSIGNED NOT NULL AFTER `annual_income_id`,
	ADD `drinking_id` int UNSIGNED NOT NULL AFTER `smoking_id`;

# 更新1回だけ
-- UPDATE `ticket_items` SET `usage_timezone`= "Asia/Tokyo" WHERE 1
