create type role_enum as enum ('teacher', 'parent', 'student');
create extension if not exists "pgcrypto";

create table users
(
    "id"            UUID    DEFAULT gen_random_uuid() PRIMARY KEY,
    "surname"       VARCHAR(40)           NOT NULL,
    "name"          VARCHAR(40)           NOT NULL,
    "patronymic"    VARCHAR(40)           NOT NULL,
    "address"       VARCHAR(255)          NOT NULL,
    "phone"         VARCHAR(30)           NOT NULL,
    "email"         VARCHAR(50)           NOT NULL UNIQUE,
    "password"      TEXT                  NOT NULL,
    "role"          role_enum             NOT NULL,
    "is_active"     BOOLEAN DEFAULT FALSE NOT NULL,
    "is_super_user" BOOLEAN DEFAULT FALSE NOT NULL,
    "created_at"    TIMESTAMP,
    "updated_at"    TIMESTAMP
)