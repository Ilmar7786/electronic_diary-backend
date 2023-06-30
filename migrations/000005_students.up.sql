create table if not exists public.students
(
    "id"                  uuid default gen_random_uuid() primary key,
    "residential_address" text                                  not null,
    "user_id"             uuid references public.users ("id")   not null,
    "parent_id"           uuid references public.parents ("id") not null,
    "created_at"          timestamp,
    "updated_at"          timestamp
)