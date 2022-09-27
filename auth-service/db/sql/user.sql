/*
 *  Copyright (c), Team Oneiro and/or its affiliates. All rights reserved.
 *  TO PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 */
create table user
(
    id       BIGINT auto_increment,
    name     varchar(100) not null,
    phone    varchar(15)  not null,
    email    varchar(50)  not null,
    password varchar(50)  not null,
    created  datetime     not null,
    updated  datetime     null,
    constraint id
        primary key (id)
);

create unique index email
    on user (email);

create unique index phone
    on user (phone);

alter table user
    add avatar varchar(50) null;

alter table user
    add rating int default 0 not null;

alter table user
    add address varchar(255) null;

alter table user
    add ip_address varchar(30) null;

alter table user
    add device_id varchar(30) null;

alter table user
    add last_active_time datetime null;
