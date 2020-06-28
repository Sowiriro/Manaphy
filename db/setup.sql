
CREATE IF NOT EXISTS `users`(
    id            integer        NOT NULL AUTO_INCREMENT,
    name          varchar(255)      NOT NULL,
) ENGINE=InnoDB;

CREATE IF NOT EXISTS `movies` (
    id             integer      NOT NULL AUTO_INCREMENT PRIMARY KEY,
    title            varchar(255) NOT NULL,
    description      text         NOT NULL,
    created_at       timestamp    not null,
    updated_at       current_timestamp not null,
) ENGINE=InnoDB;

DROP TABLE IF EXISTS `users`;
DROP TABLE IF EXISTS `movies`;
