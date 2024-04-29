-- +migrate Up
CREATE TABLE categories
(
    id         INT PRIMARY KEY AUTO_INCREMENT,
    title      VARCHAR(255) NOT NULL,
    slug       VARCHAR(255) NOT NULL UNIQUE,
    image      VARCHAR(255),
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME,
    deleted_at DATETIME,

);

-- +migrate Down
Drop Table categories;