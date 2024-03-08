Begin;

CREATE TABLE IF NOT EXISTS `token` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `user_id` INT NOT NULL,
    `jti` VARCHAR(255) NOT NULL,
    `refresh_token` VARCHAR(255) NOT NULL,
    `access_token_expires_at` TIMESTAMP NOT NULL,
    `refresh_token_expires_at` TIMESTAMP NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

ALTER TABLE `token` ADD CONSTRAINT fk_user_id FOREIGN KEY (`user_id`) REFERENCES `user`(`id`);

commit;