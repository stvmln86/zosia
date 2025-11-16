// Package sqls implements SQLite pragma and schema constants.
package sqls

// Pragma is the default always-on database pragma.
const Pragma = `
	pragma encoding = 'utf-8';
	pragma foreign_keys = on;
`

// Schema is the default first-run database schema.
const Schema = `
	create table if not exists Users (
		id   integer primary  key asc,
		addr text    not null,
		uuid text    not null default (lower(hex(randomblob(8)))),
		init integer not null default (unixepoch()),

		unique(uuid)
	);

	create table if not exists Pairs (
		id   integer primary  key asc,
		user integer not null references Users(id),
		name text    not null,
		body text    not null,
		init integer not null default (unixepoch()),

		unique(user, name)
	);

	create index if not exists PairNames on Pairs(user, name);
`
