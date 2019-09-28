create table events
(
	id serial not null
		constraint events_pk
			primary key,
	name varchar(250) not null,
	event_type int not null,
	start_time timestamp not null,
	end_time timestamp not null,
	create_time timestamp not null,
	update_time timestamp not null
);