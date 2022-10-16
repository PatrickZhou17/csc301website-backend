DROP TABLE IF EXISTS `product_info`;
CREATE TABLE `product_info` (
     `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'add ID',
     `name` varchar(50) NOT NULL COMMENT 'name',
     `price` float NOT NULL COMMENT 'price',
     `total` int DEFAULT 0 COMMENT 'amount',
     `created_at` datetime DEFAULT NULL COMMENT 'create time',
     `updated_at` datetime DEFAULT NULL COMMENT 'update time',
     `deleted_at` datetime DEFAULT NULL COMMENT 'delete time',
     PRIMARY KEY (`id`),
     UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='list of items';
