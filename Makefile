run: #run without hot reload
	go run cmd/main.go

serve: #run with hot reload
	cd cmd/;  air

db_migrate: #create db migration file dynamically
	cd migrations/;  goose create ${filename} sql

db_up: #run db migration file
	cd migrations/;  goose postgres postgresql://postgres:root@localhost:5432/clean-arc-db up


db_down: #run db migration file
	cd migrations/;  goose postgres postgresql://postgres:root@localhost:5432/clean-arc-db down