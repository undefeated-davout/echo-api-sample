CREATE TABLE `user` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` varchar(20) NOT NULL,
    `password` VARCHAR(80) NOT NULL,
    `role` VARCHAR(80) NOT NULL,
    `created_at` DATETIME(6) NOT NULL,
    `updated_at` DATETIME(6) NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uix_name` (`name`) USING BTREE
) Engine = InnoDB DEFAULT CHARSET = utf8mb4;

CREATE TABLE `task` (
    `id` BIGINT UNSIGNED NOT NULL,
    `user_id` BIGINT UNSIGNED NOT NULL,
    `title` VARCHAR(128) NOT NULL,
    `status` VARCHAR(20) NOT NULL,
    `created_at` DATETIME(6) NOT NULL,
    `updated_at` DATETIME(6) NOT NULL,
    PRIMARY KEY (`id`),
    CONSTRAINT `fk_user_id` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) Engine = InnoDB DEFAULT CHARSET = utf8mb4;