CREATE TABLE IF NOT EXISTS  `user` (
  `id` bigint PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `name` varchar(512) NOT NULL,
  `user_name` varchar(512) UNIQUE NOT NULL,
  `role_id` bigint NOT NULL,
  `email` varchar(512) UNIQUE NOT NULL,
  `password` varchar(512) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

ALTER TABLE `user` ADD CONSTRAINT user_role FOREIGN KEY (`role_id`) REFERENCES `role` (`id`);
