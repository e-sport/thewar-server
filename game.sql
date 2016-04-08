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

CREATE TABLE Users(
   uid INT NOT NULL,
   name VARCHAR(100) NOT NULL,
   password VARCHAR(40) NOT NULL,
   PRIMARY KEY (uid)
);