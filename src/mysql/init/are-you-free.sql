create table if not exists users(
    id int(10)  AUTO_INCREMENT NOT NULL primary key,
    name CHAR(20) NOT NULL,
    password CHAR(20) NOT NULL,
    is_admin BOOLEAN NOT NULL DEFAULT 0,
    created_at DATETIME DEFAULT current_timestamp,
    updated_at DATETIME DEFAULT current_timestamp
)
