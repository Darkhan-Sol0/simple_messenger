NAME = simple_message

EXEC = app

SOURCE = ./cmd/app/main.go

PACKAGE =\
		github.com/labstack/echo/v4\
		github.com/jackc/pgx/v5\
		github.com/jackc/pgx/v5/pgxpool\
		github.com/ilyakaznacheev/cleanenv\
		github.com/joho/godotenv\
		github.com/golang-jwt/jwt/v5\
		github.com/Darkhan-Sol0/simple_qb\
		github.com/redis/go-redis/v9\

.PHONY = all, run, build, build_run, clean, mod, get

all: run

run:
	go run $(SOURCE)

build:
	go build $(SOURCE) -o $(EXEC)

build_run:
	./$(EXEC)

clean:
	rm $(EXEC)

mod:
	go mod init $(NAME)

get:
	go get -u $(PACKAGE)

docker_up:

docker_down: