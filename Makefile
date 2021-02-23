setup:
	go get -tags 'postgres' -u github.com/golang-migrate/migrate/cmd/migrate
  
migrate-up:
	migrate -path migrations -database "postgresql://@localhost/gormtesting" up

migrate-down:
	migrate -path migrations -database "postgresql://@localhost/gormtesting" down
