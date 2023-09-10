START TRANSACTION;

DROP INDEX `idx_user_id` ON `users`;

DROP TABLE `users`;

COMMIT;