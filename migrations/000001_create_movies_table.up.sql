--- Using golang-migrate/migrate from github

--- Generate a pair of migration files using the "migrate create"
--- migrate create -seq -ext=.sql -dir=./migrations create_movies_table

--- To apply migrations
--- migrate -path=./migrations -database=$GREENLIGHT_DB_DSN up

--- To check which migration version database is currently on
--- migrate -path=./migrations -database=$GREENLIGHT_DB_DSN version

--- if "error: Dirty database version 6. Fix and force version." happens
--- Do this procedure
--- migrate -path ./migrations -database $GREENLIGHT_DB_DSN version (Get version)
--- migrate -path ./migrations -database $GREENLIGHT_DB_DSN force 5 (Go to the prev version)
--- migrate -path ./migrations -database $GREENLIGHT_DB_DSN version (Check DB in prev version)
--- migrate -path ./migrations -database $GREENLIGHT_DB_DSN up (Now apply new changes again)

CREATE TABLE IF NOT EXISTS movies (
  id bigserial PRIMARY KEY,
  created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
  title text NOT NULL,
  year integer NOT NULL,
  runtime integer NOT NULL,
  genres text[] NOT NULL,
  version integer NOT NULL DEFAULT 1
);