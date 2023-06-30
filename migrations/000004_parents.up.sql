create table if not exists public.parents
(
    "id"         uuid default gen_random_uuid() primary key,
    "guardian"   varchar(20),
    "user_id"    uuid unique references public.users ("id")   not null,
    "created_at" timestamp,
    "updated_at" timestamp
)