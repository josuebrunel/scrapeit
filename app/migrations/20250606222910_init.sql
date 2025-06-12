-- +goose Up
-- +goose StatementBegin
-- API KEY
create table if not exists api_keys (
    api_key text primary key,
    user_id uuid unique not null references auth.users(id) on delete cascade,
    created_at timestamptz default now()
);

create index if not exists idx_api_keys_user_id on api_keys(user_id);

-- ETHER ACCOUNT
create table if not exists ether_accounts (
    user_id uuid primary key references auth.users(id) on delete cascade,
    address text not null,
    created_at timestamptz default now(),
    updated_at timestamptz default now()
);

create index if not exists idx_ether_accounts_user_id on ether_accounts(user_id);

create index if not exists idx_ether_accounts_address on ether_accounts(address);

-- TASKS
create table if not exists tasks (
    id uuid primary key default gen_random_uuid(),
    user_id uuid not null references auth.users(id) on delete cascade,
    created_at timestamptz default now(),
    updated_at timestamptz default now(),
    assignee TEXT,
    contract TEXT,
    input TEXT,
    output TEXT,
    request_data JSONB,
    reward_wei NUMERIC,
    status TEXT,
    submitter TEXT,
    transaction TEXT
);

create index if not exists idx_tasks_user_id on tasks(user_id);

create index if not exists idx_tasks_status on tasks(status);

create index if not exists idx_tasks_assignee on tasks(assignee);

create index if not exists idx_tasks_submitter on tasks(submitter);

create index if not exists idx_tasks_input on tasks(input);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
-- API KEY
drop index if exists idx_api_keys_user_id;

drop table if exists api_keys;

-- ETHEREUM ACCOUNT
drop index if exists idx_ether_accounts_user_id;

drop index if exists idx_ether_accounts_address;

drop table if exists ether_accounts;

-- TASKS
drop index if exists idx_tasks_user_id;

drop index if exists idx_tasks_status;

drop index if exists idx_tasks_assignee;

drop index if exists idx_tasks_submitter;

drop index if exists idx_tasks_input;

drop table if exists tasks;

-- +goose StatementEnd