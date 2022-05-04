create table if not exists t_chat(
    cid     SERIAL          PRIMARY KEY,
    uid     int             not null,
    uname   varchar(21)     not null,
    room_id int             not null,
    send_time int           not null,
    message  blob,
    index utime(room_id,send_time)
);