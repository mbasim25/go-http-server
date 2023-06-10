migrate:
		migrate -source file://migrations \
						-database postgres://postgres:password@localhost/postgres?sslmode=disable up


