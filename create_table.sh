#!/bin/sh

# create schedules table
CMD_MYSQL="mysql -u${MYSQL_USER} -p${MYSQL_PASSWORD} ${MYSQL_DATABASE}"
$CMD_MYSQL -e "create table schedules (
    id int(10)  AUTO_INCREMENT NOT NULL primary key,
    year INTEGER NOT NULL,
    month INTEGER NOT NULL,
    day INTEGER NOT NULL,
    starthour INTEGER NOT NULL,
    startminute INTEGER NOT NULL,
    endhour INTEGER NOT NULL,
    endminute INTEGER NOT NULL,
    created_at DATETIME DEFAULT current_timestamp,
    updated_at DATETIME DEFAULT current_timestamp
    );"

# create users table
CMD_MYSQL="mysql -u${MYSQL_USER} -p${MYSQL_PASSWORD} ${MYSQL_DATABASE}"
$CMD_MYSQL -e "create table users (
    id int(10)  AUTO_INCREMENT NOT NULL primary key,
    name CHAR(20) NOT NULL,
    password CHAR(20) NOT NULL,
    is_admin BOOLEAN NOT NULL DEFAULT 0,
    created_at DATETIME DEFAULT current_timestamp,
    updated_at DATETIME DEFAULT current_timestamp
    );"

$CMD_MYSQL -e  "insert into schedules (year, month, day, starthour, startminute, endhour, endminute) values (2022, 11, 18, 18, 00, 24, 00);"
$CMD_MYSQL -e  "insert into schedules (year, month, day, starthour, startminute, endhour, endminute) values (2022, 11, 20, 10, 00, 12, 00);"

$CMD_MYSQL -e  "insert into users (name, password) values (hoge, hoge);"
$CMD_MYSQL -e  "insert into users (name, password) values (fuga, fuga);"
