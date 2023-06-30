create table if not exists public.auth_refresh_tokens
(
    "id"         uuid default gen_random_uuid() primary key,
    "hash"       varchar not null unique,
    "user_id"    uuid unique references public.users ("id") not null,
    "created_at" timestamp,
    "updated_at" timestamp
);