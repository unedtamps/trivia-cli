CREATE TABLE IF NOT EXISTS `trivia_choice` (
  `id` bigint PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `trivia_id` bigint NOT NULL,
  `choice` varchar(255) NOT NULL
);

CREATE INDEX `trivia_choice_trivia_id_foreign` ON `trivia_choice` (`trivia_id`);
ALTER TABLE `trivia_choice` ADD CONSTRAINT tc_trivia FOREIGN KEY (`trivia_id`) REFERENCES `trivia` (`id`);
