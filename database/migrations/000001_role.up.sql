CREATE TABLE IF NOT EXISTS  `role` (
  `id` bigint PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `name` varchar(10) NOT NULL
);

INSERT INTO `role` (`id`, `name`) VALUES (1, 'ADMIN');
INSERT INTO `role` (`id`, `name`) VALUES (2, 'USER');
