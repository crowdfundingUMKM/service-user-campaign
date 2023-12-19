-- Tabel users
CREATE TABLE `users` (
    `id` INT(11) NOT NULL AUTO_INCREMENT,
    `unix_id` CHAR(12) NOT NULL,
    `name` VARCHAR(255) NOT NULL,
    `email` VARCHAR(255),
    `phone` INT(11),
    `country` VARCHAR(255),
    `address` VARCHAR(255),
    `bio_user` TEXT,
    `fb_link` VARCHAR(255),
    `ig_link` VARCHAR(255),
    `password_hash` VARCHAR(255) DEFAULT NULL,
    `status_account` VARCHAR(10),
    `avatar_file_name` VARCHAR(255),
    `token` VARCHAR(255),
    `update_id_admin` CHAR(12),
    `update_at_admin` DATETIME,
    `created_at` DATETIME DEFAULT NULL,
    `updated_at` DATETIME DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- Tabel notif_campaigns
CREATE TABLE `notif_campaigns` (
    `id` INT(11) NOT NULL AUTO_INCREMENT,
    `user_investor` CHAR(12),
    `title` VARCHAR(255),
    `description` TEXT,
    `type_error` VARCHAR(11),
    `document` VARCHAR(255),
    `status_notif` TINYINT(1),
    `created_at` DATETIME,
    `updated_at` DATETIME,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


-- Indexes for table `users`
--
-- ALTER TABLE `users`
--   ADD PRIMARY KEY (`id`);

-- Remove token from table users
-- DELIMITER //

-- CREATE EVENT delete_expired_tokens
-- ON SCHEDULE EVERY 1 HOUR
-- DO
-- BEGIN
--     DELETE FROM users
--     WHERE token IS NOT NULL
--     AND created_at < NOW() - INTERVAL 2 DAY;
-- END //

-- DELIMITER ;

-- Backup database
-- SELECT *
-- INTO OUTFILE '/path/to/backup/users_backup.csv'
-- FIELDS TERMINATED BY ','
-- ENCLOSED BY '"'
-- LINES TERMINATED BY '\n'
-- FROM users;
