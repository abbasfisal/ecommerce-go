-- +migrate Up
CREATE TABLE `users`
(
    id            INT PRIMARY KEY AUTO_INCREMENT,
    first_name    VARCHAR(191) NOT NULL,
    last_name     VARCHAR(191) NOT NULL,
    phone_number  VARCHAR(191) NOT NULL UNIQUE,
    national_code varchar(11)  NOT NULL UNIQUE,
    password      VARCHAR(255) NOT NULL,
    type          ENUM('client' , 'admin') DEFAULT 'client',
    created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at    DATETIME,
    deleted_at    DATETIME
);

-- +migrate Down
Drop Table users;