-- +migrate Up
CREATE TABLE products
(
    id             INT PRIMARY KEY AUTO_INCREMENT,
    category_id    INT          NOT NULL,
    title          VARCHAR(255) NOT NULL,
    slug           VARCHAR(255) NOT NULL UNIQUE,
    sku            VARCHAR(255) NOT NULL UNIQUE,
    status         BOOLEAN  DEFAULT TRUE,
    quantity       INT UNSIGNED DEFAULT 0,
    original_price INT UNSIGNED DEFAULT 0 ,
    sale_price     INT UNSIGNED DEFAULT 0,
    description    TEXT,

    created_at     DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at     DATETIME,
    deleted_at     DATETIME,

    FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`)
);

-- +migrate Down
Drop Table products;