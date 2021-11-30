begin;

create table users(
    id BIGSERIAL primary key,
    account varchar(64) not null,
    name varchar(64),
    password varchar(32) not null,
    created_at timestamp without time zone not null,
    updated_at timestamp without time zone not null
);

CREATE UNIQUE INDEX "users_uk_account" ON users(account);
CREATE INDEX "users_idx_name" ON users(name);


create table projects(
    id BIGSERIAL primary key,
    user_id bigint not null,
    name varchar(64) not null,
    description varchar(1024),
    status int,
    created_at timestamp without time zone not null,
    updated_at timestamp without time zone not null
);

CREATE INDEX "projects_uidx_un" ON projects(user_id, name);

commit;