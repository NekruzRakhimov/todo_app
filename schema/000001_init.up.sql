CREATE TABLE users
(
    id            serial primary key not null unique,
    name          varchar(255)       not null,
    username      varchar(255)       not null,
    password_hash varchar(255)       not null
);

CREATE TABLE todo_items
(
    id          serial primary key                          not null unique,
    user_id     int references users (id) on delete cascade not null,
    title       varchar(255)                                not null,
    description varchar(255),
    done        boolean                                     not null default false,
    is_removed  boolean                                     not null default false
);