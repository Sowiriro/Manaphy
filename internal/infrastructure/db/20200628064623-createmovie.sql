-- +migrate Up

CREATE TABLE IF NOT EXISTS `users`(
    `id`            integer        NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `name`         varchar(255)      NOT NULL,
    `email`         varchar(255)     NOT NULL,
    `password`      varchar(255)     NOT NULL,
    `created_at`    timestamp       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`    timestamp       NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB;


CREATE TABLE IF NOT EXISTS `movies` (
    `id`             integer      NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `title`            varchar(255) NOT NULL,
    `description`      text         NOT NULL,
    `created_at`       timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`       timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB;


-- +migrate Down
DROP TABLE IF EXISTS `users`;
DROP TABLE IF EXISTS `movies`;
