CREATE DATABASE `videdit` /*!40100 COLLATE 'utf8mb4_general_ci' */

USE `videdit`

CREATE TABLE `video_create` (
	`id` VARCHAR(191) NOT NULL COLLATE 'utf8mb4_general_ci',
	`created_at` TIMESTAMP NULL DEFAULT NULL,
	`request` TEXT NULL DEFAULT NULL COLLATE 'utf8mb4_general_ci',
	`file_path` VARCHAR(191) NULL DEFAULT NULL COLLATE 'utf8mb4_general_ci',
	PRIMARY KEY (`id`) USING BTREE
)
COLLATE='utf8mb4_general_ci'
ENGINE=InnoDB
;

CREATE TABLE `video_upload` (
	`id` VARCHAR(191) NOT NULL COLLATE 'utf8mb4_general_ci',
	`created_at` TIMESTAMP NULL DEFAULT NULL,
	`file_name` VARCHAR(191) NULL DEFAULT NULL COLLATE 'utf8mb4_general_ci',
	`file_path` VARCHAR(191) NULL DEFAULT NULL COLLATE 'utf8mb4_general_ci',
	PRIMARY KEY (`id`) USING BTREE
)
COLLATE='utf8mb4_general_ci'
ENGINE=InnoDB
;
