# Gator

## Prerequisites

Make sure you have Postgres, Go and a database migration tool (I recommend goose) installed. Check out the links below for installation guide:

Go:         [Download and install](https://go.dev/doc/install)

Postgres:   [Download](https://www.postgresql.org/download/)

Goose:      Run
```shell
go install github.com/pressly/goose/v3/cmd/goose@latest
```

## Setup

First clone the repo:
```shell
git clone github.com/AradD7/Gator
```

Then start the postgres server in the background. For Linux:
```shell
sudo service postgresql start
```

Connect to the server. You can use any client but for psql just run:
```shell
sudo -u postgres psql
```
You should see a new prompt that looks like:
```shell
postgres=#
```

And finally create the database once you are connected to the server:
```shell
CREATE DATABASE gator;
```

Once the database is set up, run the migrations in sql/schema. For goose just run:
```shell
goose postgres <connection_string> up
```
Replace connection string with your own connection string. For example `postgres://postgres:@localhost:5432/gator`

Now, it's time to install Gator! Simply run the following at the root of the repo:
```shell
go install .
```

Now you're all set! You can delete the cloned repo.

## Usage
