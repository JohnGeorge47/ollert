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


CREATE TABLE `ollert`.`passwords` (
    `user_email` VARCHAR(256) NOT NULL,
    `password` VARCHAR(256) NOT NULL,
    `user_id` INT(11) NOT NULL,
    `updated_at` TIMESTAMP NULL,
    PRIMARY KEY (`user_email`),
    UNIQUE INDEX `user_email_UNIQUE` (`user_email` ASC),
    UNIQUE INDEX `user_id_UNIQUE` (`user_id` ASC));


ALTER TABLE `ollert`.`passwords`
    ADD CONSTRAINT `fk_passwords_user_id`
        FOREIGN KEY (`user_id`)
            REFERENCES `ollert`.`users` (`user_id`)
            ON DELETE CASCADE
            ON UPDATE CASCADE;
