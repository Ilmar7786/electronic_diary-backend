create table if not exists public.groups
(
    "id"                   uuid default gen_random_uuid() primary key,
    "title"                varchar(50)                            not null,
    "user_id"              uuid references public.users ("id")    not null,
    "classroom_teacher_id" uuid references public.teachers ("id") not null,
    "created_at"           timestamp,
    "updated_at"           timestamp
)