
# Street Project

People can buy/sell house/building

[links](https://docs.google.com/document/d/1ATjUVawcWBM_axQBJU2Xdcu2WkF4sz9-NSGCSDzBHJ8/edit)

## Architecture

nothing fancy, just normal monolith but with multiple database connection

```
browser --> svelte+monolith --> tarantool/clickhouse
mobile  --> monolith --> tarantool/clickhouse
```

## Project Structure

- conf: shared configuration
- main: inject all dependencies
- svelte: web frontend

MVC-like structure

- presentation -> serialization and transport
- domain -> business logic, DTO
- model -> persistence/3rd party endpoints, DAO

## CQRS

- rq = request/query (OLTP)
- wc = write/command (OLTP)
- sa = statistics/analytics (OLAP)

## Start dev mode

```shell
# start golang backend server
air web

# manually
go run main.go web

# start frontend auto build 
cd svelte
npm run watch
```

## Generate ORM

```shell
# input: 
# - model/m*.go

./gen-orm.sh
# or
cd model
go test -bench=BenchmarkGenerateOrm
# then go generate each file

# output:
# - model/m*/rq*/*.go  # -- read/query models
# - model/m*/wc*/*.go  # -- write/command mutation models
# - model/m*/sa*/*.go  # -- statistic/analytics models
```

## Generate Views

```shell
# input: 
# - domain/*.go
# - model/m*/*/*.go
# - svelte/*.svelte

go get -bench=BenchmarkGenerateViews

# output:
# - presentation/actions.GEN.go     # -- all possible commands
# - presentation/api_routes.GEN.go  # -- automatic API routes
# - presentation/web_viwe.GEN.go    # -- all template that can be used in web_static.go
# - presentation/cmd_run.GEN.go     # -- all CLI commands
# - svelte/jsApi.GEN.js             # -- all API client SDK 
```

## Manual test

```shell
# using command line
go run main.go cli guest/register '{"email":"a@b.c"}'

# using curl
go run main.go web

curl -X POST -d '{"email":"test@a.com"}' localhost:1234/guest/register
```

## Note

```shell
# docker spawning failed (because test terminated improperly), run this:
alias dockill='docker kill $(docker ps -q); docker container prune -f; docker network prune -f'
```