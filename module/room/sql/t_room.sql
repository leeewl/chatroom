create table if not exists t_room(
    rid             SERIAL          PRIMARY KEY,
    name            varchar(21)     not null,
    create_time     int             not null,
    create_uid      int             not null
);