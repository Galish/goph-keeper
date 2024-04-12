# Goph Keeper

## Clone the project
```
$ git clone https://github.com/Galish/goph-keeper
$ cd goph-keeper
```

## Usage

Launch Postgres DB using official docker alpine image:
```
$ docker compose up
```

Run server application:

```
$ ./server.sh
```

Run client application:

```
$ ./client.sh
```

## Tests

Run unit tests:

```
$ go test ./... -tags=unit -v
```

Run behavior (integration) tests:

```
$ go test ./... -tags=integration -v
```
