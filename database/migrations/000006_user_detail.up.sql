CREATE TABLE IF NOT EXISTS  `user_detail` (
  `id` bigint  PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL,
  `trivia_round` bigint NOT NULL DEFAULT 0,
  `right_answer` bigint NOT NULL DEFAULT 0,
  `wrong_answer` bigint NOT NULL DEFAULT 0,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE INDEX `user_detail_index_0` ON `user_detail` (`user_id`);
ALTER TABLE `user_detail` ADD CONSTRAINT ud_user FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);
