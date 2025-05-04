# MEREMA SERVER

## Ref
- [techshool/simplebank](https://github.com/techschool/simplebank)
- [moby/moby](https://github.com/moby/moby)
- [golang-standards/project-layout](https://github.com/golang-standards/project-layout)
- [sikozonpc/GopherSocial](https://github.com/sikozonpc/GopherSocial)

## Tools
- `docker` and `docker-compose`
- [`dbmate`](https://github.com/amacneil/dbmate)
- [`swaggo/swag`](https://github.com/swaggo/swag)

## Main structure

- `/api`: provides api for frontend
- `/services`: handles core logic
- `/repo`: handles communication with db
- `/models`: describes DTOs

## Devs

### Required envs
- `GMAIL_USERNAME` (ex: haha@gmail.com, ...) 
- `GMAIL_PASSWORD` (google app password, 12 char long)

### Setup
- Run `docker compose up -d`
- If db isn't updated with the lastest schema, run `make migrate-up` (need to install `dbmate`)
- Connect to db via sql tools (already setup in .vscode, password is "pg")
- Run `go get -u github.com/swaggo/swag` (if `swag` hasn't been installed yet)
- Run `go mod tidy`
- Run `make run`
- Check `/test` for testing cmd
- swagger is also available, check it for frontend  

### To-Do
- [x] Add mutex lock for repo layer