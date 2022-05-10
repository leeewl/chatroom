create table if not exists t_room(
    rid             SERIAL          PRIMARY KEY,
    name            varchar(21)     not null,
    create_time     int             not null,
    create_uid      int             not null
);

insert into t_room (name, create_time, create_uid) values ('鼹鼠亭',1652150816,1);
insert into t_room (name, create_time, create_uid) values ('红茶馆',1652150817,1);