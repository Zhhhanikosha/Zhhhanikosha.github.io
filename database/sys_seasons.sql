CREATE TABLE `seasons` (
  `id` int NOT NULL AUTO_INCREMENT,
  `season_num` int NOT NULL,
  `ep_quantity` int NOT NULL,
  `episode` int NOT NULL,
  PRIMARY KEY (`id`),
  KEY `episode_idx` (`episode`),
  CONSTRAINT `episodes` FOREIGN KEY (`episode`) REFERENCES `episodes` (`id`)
) 
