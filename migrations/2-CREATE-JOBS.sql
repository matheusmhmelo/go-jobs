CREATE TABLE jobs (
    id SERIAL primary key,
    title varchar(50) not null,
    description varchar(255) not null,
    created date not null,
    user_id int,
    CONSTRAINT fk_user
      FOREIGN KEY(user_id)
	  REFERENCES users(id)
);

