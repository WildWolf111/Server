CREATE TABLE IF NOT EXISTS "users" (
    "id" bigserial   not null PRIMARY KEY,
    "login" varchar not null unique,
    password varchar not null
);


CREATE TABLE articles(
    id bigserial not null PRIMARY KEY,
    title varchar not null unique,
    author varchar not null,
    content varchar not null
);