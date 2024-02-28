ALTER TABLE `trivia_choice` DROP CONSTRAINT tc_trivia;
DROP INDEX `trivia_choice_trivia_id_foreign` ON `trivia_choice`;
DROP TABLE IF EXISTS `trivia_choice`;
