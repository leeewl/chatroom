create table if not exists t_user(
    uid             int             not null,
    uname           varchar(21)     not null    UNIQUE,
    create_time     int             not null,
    ban_chat_time   int             not null,
    ban_time        int             not null,
    primary key(uid)
);