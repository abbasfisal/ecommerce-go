-- +migrate Up
CREATE TABLE `image_product`
(
    id         INT PRIMARY KEY AUTO_INCREMENT,
    product_id INT          NOT NULL,
    path       VARCHAR(255) NOT NULL,

    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME,
    deleted_at DATETIME,

    FOREIGN KEY (`product_id`) REFERENCES `products` (`id`)
);

-- +migrate Down
Drop Table `image_product`;