CREATE TABLE IF NOT EXISTS `trivia_detail` (
  `user_id` bigint NOT NULL,
  `trivia_id` bigint NOT NULL,
  `trivia_status_id` bigint NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`trivia_id`, `user_id`)
);

ALTER TABLE `trivia_detail` ADD CONSTRAINT td_user FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);

ALTER TABLE `trivia_detail` ADD CONSTRAINT td_trivia FOREIGN KEY (`trivia_id`) REFERENCES `trivia` (`id`);

ALTER TABLE `trivia_detail` ADD CONSTRAINT td_status FOREIGN KEY (`trivia_status_id`) REFERENCES `trivia_status` (`id`);
