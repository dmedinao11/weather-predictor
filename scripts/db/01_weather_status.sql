CREATE TABLE `weather_database`.`weather_status`
(
    `id`     integer PRIMARY KEY AUTO_INCREMENT,
    `status` varchar(255)
);

INSERT INTO `weather_database`.`weather_status` (`status`)
VALUES ('optimal_weather'),
       ('rainy'),
       ('drought'),
       ('normal');