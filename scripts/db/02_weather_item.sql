CREATE TABLE `weather_database`.`weather_item`
(
    `id`             integer PRIMARY KEY AUTO_INCREMENT,
    `day_id`         integer,
    `weather_status` integer,
    `period_id`      integer,
    `perimeter`      double
);

ALTER TABLE `weather_database`.`weather_item` ADD FOREIGN KEY (`weather_status`) REFERENCES `weather_database`.`weather_status` (`id`);