# fungo

Just some fun stuff with golang

## Running Tests

To run all tests:

```bash
go test ./...
```

To run tests with verbose output:

```bash
go test -v ./...
```

To run benchmarks:

```bash
go test -bench=. ./tinkering
```

### Token Package Integration Tests

The token package includes integration tests that require a PostgreSQL database. To run these tests:

1. Start the database:
```bash
cd token
docker-compose up -d
```

2. Remove the skip statement in `token/token_test.go` (line 10: `t.Skip()`)

3. Run the tests:
```bash
go test ./token -v
```

4. Stop the database when done:
```bash
cd token
docker-compose down
```

## Running

