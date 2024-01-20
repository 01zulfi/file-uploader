create table if not exists users (
    id integer primary key generated always as identity,
    username varchar ( 255 ) unique not null,
    password varchar ( 255 ) not null
);

insert into users ( username, password ) values ( 'default', 'default' ) 
on conflict ( username ) do nothing;

create table if not exists sessions (
    id integer primary key generated always as identity,
    token varchar ( 255 ) unique not null,
    user_id integer references users( id ) on delete cascade
);	

create table if not exists folders (
    id integer primary key generated always as identity,
    name varchar ( 255 ) not null,
    owner integer references users( id ) on delete cascade
);

create table if not exists folders_users (
    id integer primary key generated always as identity,
    folder_id integer references folders( id ) on delete cascade,
    user_id integer references users( id ) on delete cascade,
    unique ( folder_id, user_id )
);

create table if not exists files (
    id integer primary key generated always as identity,
    filename varchar ( 255 ) not null,
    unique_filename text not null,
    content bytea not null,
    owner integer references users( id ) on delete cascade,
    folder integer references folders( id ) on delete cascade
);