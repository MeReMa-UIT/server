# MEREMA SERVER

## Ref
- [techshool/simplebank](https://github.com/techschool/simplebank)
- [moby/moby](https://github.com/moby/moby)
- [golang-standards/project-layout](https://github.com/golang-standards/project-layout)
- [sikozonpc/GopherSocial](https://github.com/sikozonpc/GopherSocial)

## Tools
- `docker` and `docker-compose`
- [`dbmate`](https://github.com/amacneil/dbmate)

## Main structure

- `/api`: provides api for frontend
- `/services`: handles core logic
- `/repo`: handles communication with db
- `/models`: describes DTOs

## Devs
### Setup
- Run `docker compose up -d`
- If db isn't update with the lastest schema, run `make migrate-up` (need to install `dbmate`)
- Connect to db via sql tools (already setup in .vscode, password is "pg")
- Check `/test` 
- Run `go run .` and test
