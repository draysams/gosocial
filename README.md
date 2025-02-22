# GO SOCIAL

```bash
docker exec -it postgres-db /bin/bash
```

```bash
psql postgres://admin:adminpswrd@localhost:5432/gosocial_db?sslmode=disable
```

## Dirty database version

panic: Dirty database version 1. Fix and force version.

```sql
select * from schema_migrations;
update schema_migrations set dirty =false where version=1;
```

## Migrations

```bash
go run cmd/migrate/migrate.go -direction=down
```

```bash
go run cmd/migrate/migrate.go -direction=up
```
