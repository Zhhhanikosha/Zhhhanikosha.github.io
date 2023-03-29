CREATE TABLE `favorites` (
  `id` int NOT NULL AUTO_INCREMENT,
  `ep_id` int NOT NULL,
  PRIMARY KEY (`id`),
  KEY `episode_idx` (`ep_id`),
  CONSTRAINT `episode` FOREIGN KEY (`ep_id`) REFERENCES `episodes` (`id`)
) 
