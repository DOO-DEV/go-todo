CREATE TABLE IF NOT EXISTS `user`(
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `username` VARCHAR(15) NOT NULL UNIQUE,
    `password` VARCHAR(30) NOT NULL,
    `email` VARCHAR(50) NOT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS `todo`(
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `creator_id` INT NOT NULL,
    `category_id` INT NOT NULL,
    `title` VARCHAR(100) NOT NULL,
    `description` TEXT NOT NULL,
    `priority` ENUM('HIGH', 'MEDIUM', 'LOW') NOT NULL DEFAULT 'LOW',
    `complete` TINYINT(1) NOT NULL DEFAULT '0',
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS `category`(
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `creator_id` INT NOT NULL,
    `Name` VARCHAR(20) NOT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

ALTER TABLE `todo` ADD CONSTRAINT fk_todo_user_id FOREIGN KEY (`creator_id`) REFERENCES `user`(`id`);
ALTER TABLE `todo` ADD CONSTRAINT fk_todo_category_id FOREIGN KEY (`category_id`) REFERENCES `category`(`id`);

ALTER TABLE `category` ADD CONSTRAINT fk_category_user_id FOREIGN KEY (`creator_id`) REFERENCES `user`(`id`);