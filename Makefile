compile_local:
	go build -o local.server main.go

compile_prod:
	env GOOS=linux GOARCH=amd64 go build -o linux.server main.go
