ALTER TABLE `trivia_detail` DROP CONSTRAINT  td_user ;

ALTER TABLE `trivia_detail` DROP CONSTRAINT td_trivia;

ALTER TABLE `trivia_detail` DROP CONSTRAINT td_status;

DROP TABLE IF EXISTS `trivia_detail`;
