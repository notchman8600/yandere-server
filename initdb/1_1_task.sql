create table tasks
(
    task_id text UNIQUE,
    name text not null default '',
    description text not null default '',
    is_done boolean default false,
    user_id text not null default '',
    status integer not null default 0,
    deadline timestamp with time zone,
    created_at timestamp with time zone default current_timestamp,
    updated_at timestamp with time zone default current_timestamp
);
