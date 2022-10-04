# shuryak-blog

## Default migration

```bash
docker run -v $(pwd)/migrations:/migrations --network host migrate/migrate -path=migrations/ -database "postgres://user:password@localhost:5432/postgres?sslmode=disable" up
```
