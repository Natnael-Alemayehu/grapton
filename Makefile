goose-up:
	cd sql/schema && goose postgres "postgres://postgres:postgres@localhost:5432/gator" up

goose-down:
	cd sql/schema && goose postgres "postgres://postgres:postgres@localhost:5432/gator" down

goose-re: goose-down goose-up
