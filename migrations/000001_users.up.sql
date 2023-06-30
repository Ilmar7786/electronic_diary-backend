create type public.role_enum as enum ('teacher', 'parent', 'student');
create extension if not exists "pgcrypto";

create table if not exists public.users
(
    "id"            uuid    default gen_random_uuid() primary key,
    "surname"       varchar(40)           not null,
    "name"          varchar(40)           not null,
    "patronymic"    varchar(40)           not null,
    "address"       varchar(255)          not null,
    "phone"         varchar(30)           not null,
    "email"         varchar(50)           not null unique,
    "password"      TEXT                  not null,
    "role"          role_enum             not null,
    "is_active"     boolean default false not null,
    "is_super_user" boolean default false not null,
    "created_at"    timestamp,
    "updated_at"    timestamp
);

create table public.users_email_activate
(
    "id"         uuid default gen_random_uuid() primary key,
    "link"       varchar                             not null,
    "user_id"    uuid unique references users ("id") not null,
    "created_at" timestamp,
    "updated_at" timestamp
);