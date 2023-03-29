CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `first_name` varchar(45) DEFAULT NULL,
  `last_name` varchar(45) DEFAULT NULL,
  `favorite` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `favorite_idx` (`favorite`),
  CONSTRAINT `favorite` FOREIGN KEY (`favorite`) REFERENCES `favorites` (`id`)
)
