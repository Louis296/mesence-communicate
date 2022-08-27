create table if not exists friend_relation
(
    id               int auto_increment
    primary key,
    create_time      datetime          not null,
    update_time      datetime          not null,
    status           tinyint default 0 not null,
    user_id          int               not null,
    user_phone       varchar(255)      not null,
    friend_id        int               not null,
    friend_phone     varchar(255)      null,
    friend_note_name varchar(255)      null
    )
    comment '好友关系';

create index user_id__index
    on friend_relation (user_id);

create index user_phone__index
    on friend_relation (user_phone);

create table if not exists friend_request
(
    id             int auto_increment
    primary key,
    create_time    datetime          not null,
    update_time    datetime          not null,
    status         tinyint default 0 not null,
    sender         varchar(255)      not null,
    candidate      varchar(255)      not null,
    request_status int               not null,
    start_time     datetime          not null,
    content        text              null
    )
    comment '好友请求';

create index candidate__index
    on friend_request (candidate);

create index sender__index
    on friend_request (sender);

create table if not exists message
(
    id          int auto_increment
    primary key,
    create_time datetime          not null,
    update_time datetime          not null,
    status      tinyint default 0 not null,
    content     text              null,
    `from`      varchar(255)      not null,
    `to`        varchar(255)      not null,
    send_time   datetime          not null
    )
    comment '消息记录';

create index from__index
    on message (`from`);

create index to__index
    on message (`to`);

create table if not exists user
(
    id          int auto_increment
    primary key,
    create_time datetime          not null,
    update_time datetime          not null,
    status      tinyint default 0 not null,
    phone       varchar(255)      not null,
    name        varchar(255)      not null,
    password    varchar(255)      not null,
    avatar      varchar(255)      not null,
    location    varchar(255)      not null,
    constraint user_phone_uindex
    unique (phone)
    )
    comment '用户';

