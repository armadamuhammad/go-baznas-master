# Base Project

[API Specs](https://gospecs.monstercode.net)

Running test:
```bash
docker-compose up test
# or
go test -v -coverprofile cover.txt ./...
```

Get Coverage:
```bash
go tool cover -html cover.txt
# or
go tool cover -func cover.txt
```