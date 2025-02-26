create database todolist;
drop database todolist;
use todolist;

CREATE TABLE `users` (
  `id` integer PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `name` varchar(255),
  `email` varchar(255),
  `password` varchar(255),
  `created_at` timestamp,
  `updated_at` timestamp
);

CREATE TABLE `checklists` (
    `id` integer PRIMARY key not NULL AUTO_INCREMENT,
    `title` varchar(255),
    `user_id` integer,
    `created_at` timestamp,
  	`updated_at` timestamp
);

ALTER TABLE `checklists` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);