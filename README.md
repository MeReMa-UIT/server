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

- `/api`: provides api for clients
- `/services`: handles app's core logics
- `/repo`: handles communication with db
- `/models`: describes DTOs, errors, ...

## Devs

### Required envs
- `GMAIL_USERNAME` (ex: haha@gmail.com, ...) 
- `GMAIL_PASSWORD` (google app password, 12 char long)

### Setup
- Run `docker compose up -d`
- If db isn't updated with the lastest schema, run `make migrate-up` (need to install `dbmate`). If there is an error when running `make migrate-up`, run `make migrate-down` until db refreshs to init state, then run `make migrate-up` to update 
- Connect to db via vscode sql tools (already setup in .vscode, password is "pg") or vimdadbob blah blah
- Run `go get -u github.com/swaggo/swag` (if `swag` hasn't been installed yet)
- Run `go mod tidy`
- Run `make run`
- Check `/test` for testing cmd
- swagger is also available, check it for frontend dev

### To-Do
- [x] Add mutex lock for repo layer
- [x] Ultimate rework on error handling
- [ ] Use native postgresql mutex lock rather than go mutex lock 

### Test Accounts
- Admin 1: 
  - **ID**: 123123123123 (23520199@gm.uit.edu.vn)
  - **Password**: 123456
- Admin 2: 
  - **ID**: 123412341234 (23521734@gm.uit.edu.vn)
  - **Password**: 123456
- Receptionist 1:
  - **ID**: 000000001111
  - **Password**: 123456
- Doctor 1:
  - **ID**: 000000001112
  - **Password**: 123456
- Patient 1:
  - **ID**: 000000001113
  - **Password**: 123456