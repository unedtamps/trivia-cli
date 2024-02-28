CREATE TABLE IF NOT EXISTS  `dificulity` (
  `id` bigint PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `name` varchar(10) NOT NULL
);

INSERT INTO `dificulity` (`id`, `name`) VALUES (1, 'EASY');
INSERT INTO `dificulity` (`id`, `name`) VALUES (2, 'MEDIUM');
INSERT INTO `dificulity` (`id`, `name`) VALUES (3, 'HARD');
