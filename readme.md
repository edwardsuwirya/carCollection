# Car Collection 
## Golang Backend Server

> To run application with compile daemon
```
 CompileDaemon -build "-o carCollection github.com/edwardsuwirya/carCollection/main" -command="./carCollection --config /Users/edwardsuwirya/Downloads/config.json s"
```
> To run test and see coverage
```
go test -coverprofile cp.out github.com/edwardsuwirya/carCollection/... && go tool cover -func cp.out
```
