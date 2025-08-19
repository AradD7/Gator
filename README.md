# Gator

## About

Gator is guided project from boot.dev. Though instructions were given, 100% of the code and this readme document is written by me! So if it sucks blame me not boot.dev people

Gator is a RSS feed aggregator written in GO and Postgres. See [usage](#usage) below.

## Prerequisites

Make sure you have Postgres, Go and a database migration tool (I recommend goose) installed. Check out the links below for installation guide:

Go:         [Download and install](https://go.dev/doc/install)

Postgres:   [Download](https://www.postgresql.org/download/)

Goose:
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

If the installation went well, you can use gator by running commands formatted like
```shell
Gator [command] [arguments]
```

Here is a list of available commands with their arguments:
| Command | Arguments | Description |
| ------- | --------- | ----------- |
| register | *username* | adds the given *username* to the users table |
| login | *username* | sets the current user to *username* |
| reset | None | deletes all the records in users table |
| users | None | logs all the users and marks the current user |
| feeds | None | logs all the current feeds in the database |
| addfeed | *name* *url* | adds a feed to the database and follows it |
| follow | *url* | current user follows the feed with the given url |
| unfollow | *url* | current user unfollows the feed with the given url |
| following | None | logs all the feeds that the current user is following |
| agg | *time_between_requests* | agg will fetch feeds and save them as posts! It will start with the feeds that haven't been fetched and then feeds that were fetched oldest. It will do so in intervals of *time_between_requests* which is format like "1s" or "1m".
| browse | [optional]*num_of_post* | logs the *num_of_post* most recent posts from the feeds that the user follows. By default it will return 2 |
