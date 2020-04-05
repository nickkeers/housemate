create table house_members
(
	id serial
		constraint users_pk
			primary key,
	name text not null,
	email text not null,
	birthday date not null
);

insert into house_members (name, email, birthday) values ('Nick', 'nick@housemate.com', '01-13-1994');