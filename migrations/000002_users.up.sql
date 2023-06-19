create table students
(
    "id"         UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    "user_id"    UUID UNIQUE references users ("id") NOT NULL,
    "created_at" TIMESTAMP,
    "updated_at" TIMESTAMP
)