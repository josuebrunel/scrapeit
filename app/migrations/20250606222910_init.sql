-- +goose Up
-- +goose StatementBegin
create table if not exists api_keys (
    api_key text primary key,
    user_id uuid unique not null references auth.users(id) on delete cascade,
    created_at timestamptz default now()
);

create index if not exists idx_api_keys_user_id on api_keys(user_id);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
drop index if exists idx_api_keys_user_id;

drop table if exists api_keys;

-- +goose StatementEnd