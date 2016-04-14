/*
CREATE DATABASE `game1`;
CREATE USER 'dev'@'localhost' IDENTIFIED BY 'password';
GRANT ALL PRIVILEGES ON *.* TO `dev`@`localhost`;
FLUSH PRIVILEGES;



REVOKE ALL PRIVILEGES ON *.* FROM `dev`@`localhost`;
DROP DATABASE IF EXISTS `game1`;
DROP USER 'dev'@'localhost';
FLUSH PRIVILEGES;


SET PASSWORD FOR 'dev'@'localhost' = PASSWORD('dev');
*/


use game1;

CREATE TABLE IF NOT EXISTS `users` (
   `uid` VARCHAR(40) PRIMARY KEY NOT NULL,
   `user_name` VARCHAR(30) NOT NULL,
   `password` VARCHAR(44) NOT NULL,
   `created_at` CHAR(64) NOT NULL,
   `updated_at` CHAR(64) NOT NULL)