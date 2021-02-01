# fiber plus
Simple command line interface to generate boilerplates codes with repository pattern based structure in go fiber framework . 

# Installation
```bash
go get -u github.com/sacsand/fiberplus
```

# Commands
### create::repository User

Upgrade Fiber cli if a newer version is available

```
fiberplus create::repository User
```
this create ,

models/User.go
pkg/user/service.go and pkg/user/repository.go

### create::model User

This creat a single model

```
fiberplus create::model User
```
this create ,

models/User.go

### create::controller

This creat a controller

```
fiberplus create::Controller UserController
```
this create ,

controllers/UserControllers.go

