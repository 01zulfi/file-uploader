package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	databaseUrl = "postgres://postgres:postgres@localhost:5432/file_uploader"

	createUserTableQuery = `
	create table if not exists users (
		id integer primary key generated always as identity,
		username varchar ( 255 ) unique not null,
		password varchar ( 255 ) not null
	)
`
	createDefaultUserQuery = `
	insert into users ( username, password ) values ( 'default', 'default' ) 
	on conflict ( username ) do nothing;
`

	createSessionsTableQuery = `
	create table if not exists sessions (
		id integer primary key generated always as identity,
		token varchar ( 255 ) unique not null,
		user_id integer references users( id ) on delete cascade
	);	
`

	createFilesMetadataTableQuery = `
	create table if not exists files (
		id integer primary key generated always as identity,
		filename varchar ( 255 ) not null,
		filepath text unique not null,
		owner integer references users( id ) on delete cascade
	);
`
)

var pool *pgxpool.Pool
var ctx = context.Background()

func Init() error {
	newPool, err := pgxpool.New(ctx, databaseUrl)
	if err != nil {
		fmt.Println(err)
		fmt.Println("error while connecting to database")
		return err
	}
	pool = newPool

	_, err = pool.Exec(ctx, createUserTableQuery)
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = pool.Exec(ctx, createDefaultUserQuery)
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = pool.Exec(ctx, createSessionsTableQuery)
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = pool.Exec(ctx, createFilesMetadataTableQuery)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return err
}

func Get() *pgxpool.Pool {
	return pool
}
