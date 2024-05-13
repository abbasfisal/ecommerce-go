-- +migrate Up
CREATE TABLE `addresses`
(
    id          INT PRIMARY KEY AUTO_INCREMENT,
    user_id     INT         NOT NULL,
    floor       SMALLINT,
    number      VARCHAR(20) NOT NULL,
    phase       varchar(255),
    block       varchar(255),
    description VARCHAR(255),
    created_at  DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at  DATETIME,
    deleted_at  DATETIME,

    FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
);

-- +migrate Down
Drop Table `addresses`;