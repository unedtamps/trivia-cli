CREATE TABLE IF NOT EXISTS  `trivia` (
  `id` bigint PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `title` varchar(512) NOT NULL,
  `question` varchar(1024) NOT NULL,
  `answer` varchar(512) NOT NULL,
  `dificulity_id` bigint NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
ALTER TABLE `trivia` ADD CONSTRAINT trivia_dificulity_id_foreign FOREIGN KEY (`dificulity_id`) REFERENCES `dificulity` (`id`);
