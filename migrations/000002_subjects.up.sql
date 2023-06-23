create table public.subjects
(
    "id"         uuid default gen_random_uuid() primary key,
    "title"      varchar(40) not null unique,
    "created_at" timestamp,
    "updated_at" timestamp
);