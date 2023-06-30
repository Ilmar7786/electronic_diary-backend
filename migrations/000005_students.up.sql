create table if not exists public.students
(
    "id"                  uuid default gen_random_uuid() primary key,
    "residential_address" text,
    "user_id"             uuid unique references public.users ("id") not null,
    "created_at"          timestamp,
    "updated_at"          timestamp
)