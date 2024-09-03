IN_MIGRATE_DIRECTORY = /opt/homebrew/Cellar/golang-migrate/4.17.1/bin/migrate
OUT_MIGRATE_DIRECTORY = /Users/eugene/Desktop/go/go_advanced/bookmarks/bin/migrate

get-deps:
		cp ${IN_MIGRATE_DIRECTORY} ${OUT_MIGRATE_DIRECTORY}

migrate-create:
		migrate create -ext sql -dir migrations BookMarksMigration

migrate-run:
		migrate -source file://migrations -database "postgres://localhost:5432/book_marks?sslmode=disable&user=eugene&password=postgres" up

migrate-drop:
		migrate -source file://migrations -database "postgres://localhost:5432/book_marks?sslmode=disable&user=eugene&password=postgres" down