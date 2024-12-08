-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema mydb
-- -----------------------------------------------------
-- -----------------------------------------------------
-- Schema marketplace
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema marketplace
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `marketplace` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci ;
USE `marketplace` ;

-- -----------------------------------------------------
-- Table `marketplace`.`users`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `marketplace`.`users` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(255) NOT NULL,
  `email` VARCHAR(255) NOT NULL,
  `password` VARCHAR(255) NOT NULL,
  `role` ENUM('buyer', 'seller', 'admin') NOT NULL DEFAULT 'buyer',
  `created_at` TIMESTAMP NULL DEFAULT NULL,
  `updated_at` TIMESTAMP NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `users_email_unique` (`email` ASC))
ENGINE = InnoDB
AUTO_INCREMENT = 8
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_unicode_ci;


-- -----------------------------------------------------
-- Table `marketplace`.`carts`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `marketplace`.`carts` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `buyer_id` BIGINT UNSIGNED NOT NULL,
  `created_at` TIMESTAMP NULL DEFAULT NULL,
  `updated_at` TIMESTAMP NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  INDEX `carts_buyer_id_foreign` (`buyer_id` ASC),
  CONSTRAINT `carts_buyer_id_foreign`
    FOREIGN KEY (`buyer_id`)
    REFERENCES `marketplace`.`users` (`id`)
    ON DELETE CASCADE)
ENGINE = InnoDB
AUTO_INCREMENT = 2
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_unicode_ci;


-- -----------------------------------------------------
-- Table `marketplace`.`categories`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `marketplace`.`categories` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(255) NOT NULL,
  `description` TEXT NULL DEFAULT NULL,
  `created_at` TIMESTAMP NULL DEFAULT NULL,
  `updated_at` TIMESTAMP NULL DEFAULT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB
AUTO_INCREMENT = 201
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_unicode_ci;


-- -----------------------------------------------------
-- Table `marketplace`.`products`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `marketplace`.`products` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(255) NOT NULL,
  `description` TEXT NULL DEFAULT NULL,
  `image` VARCHAR(255) NULL DEFAULT NULL,
  `price` DECIMAL(10,2) NOT NULL,
  `stock` INT NOT NULL,
  `seller_id` BIGINT UNSIGNED NOT NULL,
  `category_id` BIGINT UNSIGNED NOT NULL,
  `created_at` TIMESTAMP NULL DEFAULT NULL,
  `updated_at` TIMESTAMP NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  INDEX `products_seller_id_foreign` (`seller_id` ASC),
  INDEX `products_category_id_foreign` (`category_id` ASC),
  CONSTRAINT `products_category_id_foreign`
    FOREIGN KEY (`category_id`)
    REFERENCES `marketplace`.`categories` (`id`)
    ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `products_seller_id_foreign`
    FOREIGN KEY (`seller_id`)
    REFERENCES `marketplace`.`users` (`id`)
    ON DELETE CASCADE ON UPDATE CASCADE)
ENGINE = InnoDB
AUTO_INCREMENT = 377
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_unicode_ci;


-- -----------------------------------------------------
-- Table `marketplace`.`cart_items`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `marketplace`.`cart_items` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `cart_id` BIGINT UNSIGNED NOT NULL,
  `product_id` BIGINT UNSIGNED NOT NULL,
  `quantity` INT NOT NULL DEFAULT '1',
  `created_at` TIMESTAMP NULL DEFAULT NULL,
  `updated_at` TIMESTAMP NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  INDEX `cart_items_cart_id_foreign` (`cart_id` ASC),
  INDEX `cart_items_product_id_foreign` (`product_id` ASC),
  CONSTRAINT `cart_items_cart_id_foreign`
    FOREIGN KEY (`cart_id`)
    REFERENCES `marketplace`.`carts` (`id`)
    ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `cart_items_product_id_foreign`
    FOREIGN KEY (`product_id`)
    REFERENCES `marketplace`.`products` (`id`)
    ON DELETE CASCADE ON UPDATE CASCADE)
ENGINE = InnoDB
AUTO_INCREMENT = 4
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_unicode_ci;


-- -----------------------------------------------------
-- Table `marketplace`.`discussions`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `marketplace`.`discussions` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `product_id` BIGINT UNSIGNED NOT NULL,
  `user_id` BIGINT UNSIGNED NOT NULL,
  `parent_id` BIGINT UNSIGNED DEFAULT NULL,
  `content` TEXT NOT NULL,
  `created_at` TIMESTAMP NULL DEFAULT NULL,
  `updated_at` TIMESTAMP NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  INDEX `discussions_product_id_foreign` (`product_id` ASC),
  INDEX `discussions_user_id_foreign` (`user_id` ASC),
  INDEX `discussions_parent_id_foreign` (`parent_id` ASC),
  CONSTRAINT `discussions_product_id_foreign`
    FOREIGN KEY (`product_id`)
    REFERENCES `marketplace`.`products` (`id`)
    ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `discussions_user_id_foreign`
    FOREIGN KEY (`user_id`)
    REFERENCES `marketplace`.`users` (`id`)
    ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `discussions_parent_id_foreign`
    FOREIGN KEY (`parent_id`)
    REFERENCES `marketplace`.`discussions` (`id`)
    ON DELETE CASCADE ON UPDATE CASCADE)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_unicode_ci;

-- Table structures for core tables continue...

-- -----------------------------------------------------
-- Table `marketplace`.`orders`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `marketplace`.`orders` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `buyer_id` BIGINT UNSIGNED NOT NULL,
  `status` VARCHAR(255) NOT NULL,
  `total` DECIMAL(8,2) NOT NULL,
  `created_at` TIMESTAMP NULL DEFAULT NULL,
  `updated_at` TIMESTAMP NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  INDEX `orders_buyer_id_foreign` (`buyer_id` ASC),
  CONSTRAINT `orders_buyer_id_foreign`
    FOREIGN KEY (`buyer_id`)
    REFERENCES `marketplace`.`users` (`id`)
    ON DELETE CASCADE)
ENGINE = InnoDB
AUTO_INCREMENT = 2
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_unicode_ci;

-- -----------------------------------------------------
-- Table `marketplace`.`order_items`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `marketplace`.`order_items` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `order_id` BIGINT UNSIGNED NOT NULL,
  `product_id` BIGINT UNSIGNED NOT NULL,
  `quantity` INT NOT NULL,
  `price` DECIMAL(8,2) NOT NULL,
  `created_at` TIMESTAMP NULL DEFAULT NULL,
  `updated_at` TIMESTAMP NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  INDEX `order_items_order_id_foreign` (`order_id` ASC),
  INDEX `order_items_product_id_foreign` (`product_id` ASC),
  CONSTRAINT `order_items_order_id_foreign`
    FOREIGN KEY (`order_id`)
    REFERENCES `marketplace`.`orders` (`id`)
    ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `order_items_product_id_foreign`
    FOREIGN KEY (`product_id`)
    REFERENCES `marketplace`.`products` (`id`)
    ON DELETE CASCADE ON UPDATE CASCADE)
ENGINE = InnoDB
AUTO_INCREMENT = 2
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_unicode_ci;

-- -----------------------------------------------------
-- Table `marketplace`.`replies`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `marketplace`.`replies` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `discussion_id` BIGINT UNSIGNED NOT NULL,
  `user_id` BIGINT UNSIGNED NOT NULL,
  `content` TEXT NOT NULL,
  `created_at` TIMESTAMP NULL DEFAULT NULL,
  `updated_at` TIMESTAMP NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  INDEX `replies_discussion_id_foreign` (`discussion_id` ASC),
  INDEX `replies_user_id_foreign` (`user_id` ASC),
  CONSTRAINT `replies_discussion_id_foreign`
    FOREIGN KEY (`discussion_id`)
    REFERENCES `marketplace`.`discussions` (`id`)
    ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `replies_user_id_foreign`
    FOREIGN KEY (`user_id`)
    REFERENCES `marketplace`.`users` (`id`)
    ON DELETE CASCADE ON UPDATE CASCADE)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_unicode_ci;

-- -----------------------------------------------------
-- Table `marketplace`.`reviews`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `marketplace`.`reviews` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `product_id` BIGINT UNSIGNED NOT NULL,
  `user_id` BIGINT UNSIGNED NOT NULL,
  `rating` INT NOT NULL,
  `comment` TEXT NULL DEFAULT NULL,
  `created_at` TIMESTAMP NULL DEFAULT NULL,
  `updated_at` TIMESTAMP NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  INDEX `reviews_product_id_foreign` (`product_id` ASC),
  INDEX `reviews_user_id_foreign` (`user_id` ASC),
  CONSTRAINT `reviews_product_id_foreign`
    FOREIGN KEY (`product_id`)
    REFERENCES `marketplace`.`products` (`id`)
    ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `reviews_user_id_foreign`
    FOREIGN KEY (`user_id`)
    REFERENCES `marketplace`.`users` (`id`)
    ON DELETE CASCADE ON UPDATE CASCADE)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_unicode_ci;

-- -----------------------------------------------------
-- Table `marketplace`.`transactions`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `marketplace`.`transactions` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` BIGINT UNSIGNED NOT NULL,
  `product_id` BIGINT UNSIGNED NOT NULL,
  `quantity` INT NOT NULL,
  `total_price` DECIMAL(10,2) NOT NULL,
  `status` VARCHAR(255) NOT NULL DEFAULT 'pending',
  `created_at` TIMESTAMP NULL DEFAULT NULL,
  `updated_at` TIMESTAMP NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  INDEX `transactions_user_id_foreign` (`user_id` ASC),
  INDEX `transactions_product_id_foreign` (`product_id` ASC),
  CONSTRAINT `transactions_product_id_foreign`
    FOREIGN KEY (`product_id`)
    REFERENCES `marketplace`.`products` (`id`)
    ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `transactions_user_id_foreign`
    FOREIGN KEY (`user_id`)
    REFERENCES `marketplace`.`users` (`id`)
    ON DELETE CASCADE ON UPDATE CASCADE)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_unicode_ci;

-- -----------------------------------------------------
-- Table `marketplace`.`wishlists`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `marketplace`.`wishlists` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` BIGINT UNSIGNED NOT NULL,
  `product_id` BIGINT UNSIGNED NOT NULL,
  `created_at` TIMESTAMP NULL DEFAULT NULL,
  `updated_at` TIMESTAMP NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  INDEX `wishlists_user_id_foreign` (`user_id` ASC),
  INDEX `wishlists_product_id_foreign` (`product_id` ASC),
  CONSTRAINT `wishlists_product_id_foreign`
    FOREIGN KEY (`product_id`)
    REFERENCES `marketplace`.`products` (`id`)
    ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `wishlists_user_id_foreign`
    FOREIGN KEY (`user_id`)
    REFERENCES `marketplace`.`users` (`id`)
    ON DELETE CASCADE ON UPDATE CASCADE)
ENGINE = InnoDB
AUTO_INCREMENT = 3
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_unicode_ci;

-- Insert initial data
INSERT INTO `marketplace`.`users` (`id`, `name`, `email`, `password`, `role`, `created_at`, `updated_at`) VALUES
(1, 'John Doe', 'john@example.com', 'password', 'buyer', NOW(), NOW());

INSERT INTO `marketplace`.`categories` (`id`, `name`, `description`, `created_at`, `updated_at`) VALUES
(1, 'Electronics', 'Electronic items', NOW(), NOW());

INSERT INTO `marketplace`.`products` (`id`, `name`, `description`, `image`, `price`, `stock`, `seller_id`, `category_id`, `created_at`, `updated_at`) VALUES
(1, 'Laptop', 'A powerful laptop', 'laptop.png', 1000.00, 10, 1, 1, NOW(), NOW());

INSERT INTO `marketplace`.`carts` (`id`, `buyer_id`, `created_at`, `updated_at`) VALUES
(1, 1, NOW(), NOW());

INSERT INTO `marketplace`.`orders` (`id`, `buyer_id`, `status`, `total`, `created_at`, `updated_at`) VALUES
(1, 1, 'pending', 100.00, NOW(), NOW());

INSERT INTO `marketplace`.`order_items` (`id`, `order_id`, `product_id`, `quantity`, `price`, `created_at`, `updated_at`) VALUES
(1, 1, 1, 1, 100.00, NOW(), NOW());

INSERT INTO `marketplace`.`cart_items` (`id`, `cart_id`, `product_id`, `quantity`, `created_at`, `updated_at`) VALUES
(1, 1, 1, 2, NOW(), NOW());

INSERT INTO `marketplace`.`discussions` (`id`, `product_id`, `user_id`, `content`, `created_at`, `updated_at`) VALUES
(1, 1, 1, 'Is this product available?', NOW(), NOW());

INSERT INTO `marketplace`.`replies` (`id`, `discussion_id`, `user_id`, `content`, `created_at`, `updated_at`) VALUES
(1, 1, 1, 'Yes, it is available.', NOW(), NOW());

INSERT INTO `marketplace`.`reviews` (`id`, `product_id`, `user_id`, `rating`, `comment`, `created_at`, `updated_at`) VALUES
(1, 1, 1, 5, 'Great product!', NOW(), NOW());

INSERT INTO `marketplace`.`transactions` (`id`, `user_id`, `product_id`, `quantity`, `total_price`, `status`, `created_at`, `updated_at`) VALUES
(1, 1, 1, 1, 100.00, 'pending', NOW(), NOW());

INSERT INTO `marketplace`.`wishlists` (`id`, `user_id`, `product_id`, `created_at`, `updated_at`) VALUES
(1, 1, 1, NOW(), NOW());


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;