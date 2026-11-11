-- Change your database name here if needed:
CREATE DATABASE my_database_name;

-- ============================================================
-- gosql_learn schema
-- MariaDB 10.5+ / InnoDB / utf8mb4
-- ============================================================

SET NAMES utf8mb4;

-- Drop child first to satisfy fk_posts_user
DROP TABLE IF EXISTS `posts`;
DROP TABLE IF EXISTS `users`;

-- ------------------------------------------------------------
-- users
-- ------------------------------------------------------------
CREATE TABLE `users` (
    `id`         INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `name`       VARCHAR(255)     NOT NULL,
    `email`      VARCHAR(255)     NOT NULL,
    `password`   VARCHAR(255)     NOT NULL,
    `created_at` TIMESTAMP        NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_users_email` (`email`)
) ENGINE=InnoDB
  DEFAULT CHARSET=utf8mb4
  COLLATE=utf8mb4_unicode_ci;

-- ------------------------------------------------------------
-- posts
-- ------------------------------------------------------------
CREATE TABLE `posts` (
    `id`         INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id`    INT(10) UNSIGNED NOT NULL,
    `content`    TEXT             NOT NULL,
    `created_at` TIMESTAMP        NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `idx_posts_user_id` (`user_id`),
    CONSTRAINT `fk_posts_user`
        FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
        ON DELETE CASCADE
        ON UPDATE CASCADE
) ENGINE=InnoDB
  DEFAULT CHARSET=utf8mb4
  COLLATE=utf8mb4_unicode_ci;

-- password = 'password'
INSERT INTO `users` (`name`, `email`, `password`) VALUES
    ('Alice Chen',      'alice.chen@example.com',     '$2a$12$dgPKJrTABJdfF0CPTw7zk.CWhMwgAHDfG4qjf00UmtVTVLdr8UCHW'),
    ('Bob Martinez',    'bob.martinez@example.com',   '$2a$12$dgPKJrTABJdfF0CPTw7zk.CWhMwgAHDfG4qjf00UmtVTVLdr8UCHW'),
    ('Carla Rossi',     'carla.rossi@example.com',    '$2a$12$dgPKJrTABJdfF0CPTw7zk.CWhMwgAHDfG4qjf00UmtVTVLdr8UCHW'),
    ('Diego Silva',     'diego.silva@example.com',    '$2a$12$dgPKJrTABJdfF0CPTw7zk.CWhMwgAHDfG4qjf00UmtVTVLdr8UCHW'),
    ('Emma Novak',      'emma.novak@example.com',     '$2a$12$dgPKJrTABJdfF0CPTw7zk.CWhMwgAHDfG4qjf00UmtVTVLdr8UCHW'),
    ('Farid Haddad',    'farid.haddad@example.com',   '$2a$12$dgPKJrTABJdfF0CPTw7zk.CWhMwgAHDfG4qjf00UmtVTVLdr8UCHW'),
    ('Gina Tanaka',     'gina.tanaka@example.com',    '$2a$12$dgPKJrTABJdfF0CPTw7zk.CWhMwgAHDfG4qjf00UmtVTVLdr8UCHW'),
    ('Hugo Muller',     'hugo.muller@example.com',    '$2a$12$dgPKJrTABJdfF0CPTw7zk.CWhMwgAHDfG4qjf00UmtVTVLdr8UCHW'),
    ('Ines Garcia',     'ines.garcia@example.com',    '$2a$12$dgPKJrTABJdfF0CPTw7zk.CWhMwgAHDfG4qjf00UmtVTVLdr8UCHW'),
    ('Jonas Kim',       'jonas.kim@example.com',      '$2a$12$dgPKJrTABJdfF0CPTw7zk.CWhMwgAHDfG4qjf00UmtVTVLdr8UCHW');
