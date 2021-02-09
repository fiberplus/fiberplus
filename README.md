# fiber plus (works in progress)
Simple command line interface to generate boilerplates codes with repository pattern based structure in go fiber framework . 
TODO
load env deom one config -- laravel like
define structure 
documenting 



# Installation
```bash
go get -u github.com/fiberplus/fiberplus
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

##Note1 

Start with capitel letters when naming model, controller or repository

##Note2 -How to Overide defualt folder structure(not recommend)
default directory structure is below

```
modelpath: "models"
pkgpath:  "pkg"
controllerpath: "controllers"
```

to override value
you need to create `.fiberplus.yaml` in your root directory of your application and add the overriden values. (This feature is experimental)

```
modelpath: "model"
pkgpath:  "packges"
controllerpath: "routers"
```




