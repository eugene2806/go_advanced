#IN_MIGRATE_DIRECTORY = /opt/homebrew/Cellar/golang-migrate/4.17.1/bin/migrate
#OUT_MIGRATE_DIRECTORY = /Users/eugene/Desktop/go/go_advanced/bookmarks/bin/migrate
LOCAL_BIN:=$(CURDIR)/bin

GOBIN=$(LOCAL_BIN) go install

get-deps:
		#cp ${IN_MIGRATE_DIRECTORY} ${OUT_MIGRATE_DIRECTORY}
		GOBIN=$(LOCAL_BIN) go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.17.1

migrate-create:
		migrate create -ext sql -dir migrations $(name)

migrate-run:
		migrate -source file://migrations -database "postgres://localhost:5432/book_marks?sslmode=disable&user=eugene&password=postgres" up

migrate-drop:
		migrate -source file://migrations -database "postgres://localhost:5432/book_marks?sslmode=disable&user=eugene&password=postgres" down