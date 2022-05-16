# purchase-service
###### How to run ?

Go Forked and clone this project and cd inside msgo.

`go run .`

###### Run test ?

testing state

`go test -v ./dto`

testing route

`go test -v ./controller`

testing services

`go test -v ./service`

test all with one liners

`go test --v ./service ./controller ./dto`

###### With Docker ?

Note: If you have mysql already running you can turn it off: `systemctl stop mysql`

`docker-compose up --build`

or

`docker-compose up`
