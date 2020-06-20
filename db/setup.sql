create table users {
    id          serial primary key,
    name           varchar(255),
    email          varchar(255) not null unique,
    password       varchar(255) not null,
    created_at     timestamp not null
};

create table sessions {
    id              serial primary key,
    email           varchar(255),
    user_id         inter references users(id),
    created_at      timestamp not null
};

create table movies {
    id              serial primary key,
    title           varchar(255),
    description     text,
    created_at      timestamp,
    updated_at      current_timestamp
};

create table likes {
    id              serial primary key,
    user_id         serial not null,
    movie_id        serial not null,
};

drop table likes;
drop table movies;
drop table sessions;
drop table users;