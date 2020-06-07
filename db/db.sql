CREATE TABLE `ollert`.`users` (
    `account_id` INT NOT NULL AUTO_INCREMENT,
    `account_name` VARCHAR(256) NULL,
    `created_at` TIMESTAMP NULL,
    PRIMARY KEY (`account_id`),
    UNIQUE INDEX `account_name_UNIQUE` (`account_name` ASC));





CREATE TABLE `ollert`.`users` (
  `user_id` INT NOT NULL AUTO_INCREMENT,
  `email_id` VARCHAR(256) NULL,
  `user_name` VARCHAR(256) NULL,
  `account_id` INT NOT NULL,
  `is_admin` TINYINT NULL,
  `created_at` TIMESTAMP NULL,
  `updated_at` TIMESTAMP NULL,
  PRIMARY KEY (`user_id`),
  UNIQUE INDEX `email_id_UNIQUE` (`email_id` ASC));

