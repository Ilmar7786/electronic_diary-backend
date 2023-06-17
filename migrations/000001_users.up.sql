create type role_enum as enum (
    'teacher',
    'parent',
    'student'
    );

CREATE TABLE users
(
    id           UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    surname      VARCHAR(40)  NOT NULL,
    name         VARCHAR(40)  NOT NULL,
    patronymic   VARCHAR(40)  NOT NULL,
    address      VARCHAR(255) NOT NULL,
    phone        VARCHAR(20)  NOT NULL,
    email        VARCHAR(50)  NOT NULL UNIQUE,
    password     TEXT         NOT NULL,
    role         role_enum  NOT NULL,
    is_active    BOOLEAN          DEFAULT false,
    is_superuser BOOLEAN      NOT NULL,
    created_at   TIMESTAMPTZ,
    updated_at   TIMESTAMPTZ
);