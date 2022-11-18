#!/bin/sh

CMD_MYSQL="mysql -u${MYSQL_USER} -p${MYSQL_PASSWORD} ${MYSQL_DATABASE}"
$CMD_MYSQL -e "create table schedule (
    id int(10)  AUTO_INCREMENT NOT NULL primary key,
    year INTEGER NOT NULL,
    month INTEGER NOT NULL,
    day INTEGER NOT NULL,
    starthour INTEGER NOT NULL,
    startminute INTEGER NOT NULL,
    endhour INTEGER NOT NULL,
    endminute INTEGER NOT NULL
    );"
$CMD_MYSQL -e  "insert into schedule values (1, 2022, 11, 18, 18, 00, 24, 00);"
$CMD_MYSQL -e  "insert into schedule values (2, 2022, 11, 20, 10, 00, 12, 00);"
