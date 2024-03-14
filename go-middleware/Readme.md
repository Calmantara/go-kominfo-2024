CMD => tempat main.go
PKG => ada tempat package yang membantu kita ngoding
INTERNAL => layer layer api

mockery --dir internal/repository --name UserQuery --output internal/repository/mocks

run all tests
go test ./...
go test -cover ./...

CHALLANGE:

- buatlah unit test setidaknya 1 function
- di layer Handler dan Repository

ref:

1. mock gorm: https://www.codingexplorations.com/blog/testing-gorm-with-sqlmock
2. mock handler: https://gin-gonic.com/docs/testing/
