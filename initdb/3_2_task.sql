create table tasks
(
    task_id text not null UNIQUE,
    name text,
    is_done boolean default false,
    agenda_id text,
    created_at timestamp with time zone default current_timestamp,
    updated_at timestamp with time zone default current_timestamp
);
