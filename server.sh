go run cmd/server/main.go -d postgres://gophkeeper:userpassword@localhost:5432/gophkeeper -i ./init.sql -l debug -s secret_key -c certs/localhost.crt -k certs/localhost.key