CREATE TABLE IF NOT EXISTS  `trivia_status` (
  `id` bigint PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `name` varchar(10) NOT NULL
);

INSERT INTO `trivia_status` (`id`, `name`) VALUES (1, 'OK');
INSERT INTO `trivia_status` (`id`, `name`) VALUES (2, 'REVIEW');
INSERT INTO `trivia_status` (`id`, `name`) VALUES (3, 'REJECTED');
