-- +migrate Up
CREATE TABLE sessions
(
    id         INT PRIMARY KEY AUTO_INCREMENT,
    session_id VARCHAR(255) NOT NULL UNIQUE,
    user_id    INT          NOT NULL,
    expired_at DATETIME,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME,
    deleted_at DATETIME,
    FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
);

-- +migrate Down
Drop Table sessions;