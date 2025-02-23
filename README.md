# GO SOCIAL

```bash
docker exec -it postgres-db /bin/bash
```

```bash
psql postgres://admin:adminpswrd@localhost:5432/gosocial_db?sslmode=disable
```

```sql
INSERT INTO users (email, username, password, role_id) VALUES ('mail@mail.com', 'mail', 'mail', 1);
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

## Seeding

```sql
go run cmd/migrate/seed/main.go
```

1. Have line numbers in your editor
2. Name your files based on folder structure. store/posts.go = store/postStore.go and api/posts.go = api/postApi.go = helps differentiate the files