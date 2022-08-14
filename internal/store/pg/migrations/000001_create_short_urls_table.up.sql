CREATE TABLE  IF NOT EXISTS short
		(id serial primary key,
		original_url varchar(4096) not null,
		short_url varchar(32) UNIQUE not null,
		user_id varchar(36) not null,
		correlation_id varchar(36) null,
		status smallint not null DEFAULT 0);