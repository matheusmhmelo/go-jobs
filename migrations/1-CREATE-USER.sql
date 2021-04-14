CREATE TABLE users (
    id SERIAL primary key,
    name varchar(255) not null,
    email varchar(255) unique not null,
    password varchar(255) not null,
    phone varchar(50),
    state varchar(50),
    city varchar(50),
    address varchar(50)
);