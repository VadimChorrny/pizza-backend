package migrations

var v1 = `
create table if not exists users (
        id text not null,
        email text not null,
        password text not null,
        first_name text not null,
        last_name text not null,
        role text not null,
        created_at timestamp not null,
                                 
    primary key (id)
);

create index if not exists idx_users_email_first_name on users (email, first_name);

`
