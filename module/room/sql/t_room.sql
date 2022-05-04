create table if not exists t_room(
    rid     SERIAL          PRIMARY KEY,
    rname   varchar(21)     not null,
);