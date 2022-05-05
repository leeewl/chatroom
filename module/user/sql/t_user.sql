create table if not exists t_user(
    uid             SERIAL          PRIMARY KEY,
    uname           varchar(21)     not null    UNIQUE,
    passwd          varchar(21)     not null,
    create_time     int             not null,
    ban_chat_time   int             not null,
    ban_time        int             not null
);