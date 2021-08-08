CREATE TABLE `b_user`
(
    `id`           INTEGER PRIMARY KEY AUTOINCREMENT,
    `username`     VARCHAR(64) NOT NULL,
    `password`     VARCHAR(256) NULL,
    `created_time` DATETIME NULL DEFAULT current_timestamp
);