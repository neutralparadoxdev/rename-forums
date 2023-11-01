# Some Notes

For the time being, I will need scripts to start up goforum dependencies

### Postgres

> podman run -p 5432:5432 --name some-postgres -e POSTGRES_PASSWORD=mysecretpassword -d docker.io/library/postgres

### Rabbitmq 

> podman run -p 15672:15672 -d --hostname my-rabbit --name some-rabbit docker.io/library/rabbitmq:3 


## Refactoring Auth on webapi endpoints

The Idea is to have helper functions that require auth for additional functionality, but is not require for basic
behavior.

### fn OptionalAuth(authfunc, nonauthfunc) => handle(ctx) error

### fn OptionalAuth(authfunc) => handler(ctx) error 