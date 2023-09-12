## Getting Started 🎃

```sh
docker-compose up -d
cp ./config.example.yaml ./config.yaml
```

Make sure to update config.yaml with corresponding env from `docker-compose.yml`.

## Debug

```sh
make debug
```

## Watch mode

```sh
make watch
```

## Migration

To create a migration file, run
```sh
make gen-migration name=<migration_file_name>
```

To run migration
```sh
make migrate
```

We're using [golang-migrate](github.com/golang-migrate/migrate) underlying, please go check their documentation.
