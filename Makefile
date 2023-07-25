
run:
	go run cmd/main.go

swag-init:
	swag init -g api/api.go -o api/docs

migration-up:
	migrate -path ./migration/postgres -database 'postgres://abdulbosit:946236953@localhost:5432/exam?sslmode=disable' up

migration-down:
	migrate -path ./migration/postgres -database 'postgres://abdulbosit:946236953@localhost:5432/exam?sslmode=disable' down



