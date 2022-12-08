-- users table
create table if not exists users(
    id int(10)  AUTO_INCREMENT NOT NULL primary key,
    name CHAR(20) NOT NULL,
    password CHAR(20) NOT NULL,
    is_admin BOOLEAN NOT NULL DEFAULT 0,
    created_at DATETIME DEFAULT current_timestamp,
    updated_at DATETIME DEFAULT current_timestamp
);

-- schedules table
create table if not exists schedules(
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
);

insert into schedules (year, month, day, starthour, startminute, endhour, endminute) values (2022, 11, 18, 18, 00, 24, 00);
insert into schedules (year, month, day, starthour, startminute, endhour, endminute) values (2022, 11, 20, 10, 00, 12, 00);

insert into users (name, password) values ("hoge", "hoge");
insert into users (name, password) values ("fuga", "fuga");
