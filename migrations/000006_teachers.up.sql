create table if not exists public.teachers
(
    "id"         uuid default gen_random_uuid() primary key,
    "user_id"    uuid unique references public.users ("id")    not null,
    "created_at" timestamp,
    "updated_at" timestamp
);

create table if not exists public.teachers (
    teacher_id uuid references public.teachers(id),
    subject_id uuid references public.subjects(id),
    primary key (teacher_id, subject_id)
);
