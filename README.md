# go-sqlite

## Start the container

```
hmmm@hmmm:/tmp$ docker run --rm -ti -v ~/GolangPlayground/go-sqlite-gpibdvhc:/app golang:1.14-buster /bin/bash
```

## Start the project (needed only once)

How did I find the fancy name? Easy:
```
RNDPART=`tr -dc a-z </dev/urandom | head -c 8`
echo "github.com/igiannoulas/go-sqlite-$RNDPART"
```

Initialize the modules:

```
root@9271f5352b13:~# cd /app/
root@9271f5352b13:/app# go mod init github.com/igiannoulas/go-sqlite-gpibdvhc
```

## Useful information

It seems that we need 2 things in order to work with SQLite:

- https://golang.org/pkg/database/sql/ : Package sql provides a generic interface around SQL (or SQL-like) databases. The sql package must be used in conjunction with a database driver. See https://golang.org/s/sqldrivers for a list of drivers.

- https://github.com/mattn/go-sqlite3: sqlite3 driver conforming to the built-in database/sql interface

In order to install the package `github.com/mattn/go-sqlite3`:

```
root@9271f5352b13:/app# go get -u github.com/mattn/go-sqlite3
```