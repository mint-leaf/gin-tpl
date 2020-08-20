## Gin template

`gin-tpl` is backend template for [gin](https://github.com/gin-gonic/gin) web application

## Structure

    |--config/
    |   |--db/db.go
    |   |--route
    |       |  router.go
    |       |  middleware.go
    |--lib/
    |   |  errors.go
    |--models/
    |--param/
    |--router/
    |--service/
    |--utils/
    |  main.go
    |  db.config.yaml

- `config` init your application and load config. `gin-tpl` use yaml files as its config files, also you can use any other type file such as json file and gin-tpl provide function in `utils/file_help.go` called `LoadDataFromFile` to help load config from file easily. examples for load database config and route config can be found in `config/db/db.go` and `config/route/route.go`

- `lib` errors definitions and some wraps for packages such as cdn, oss, sms and so on

- `models` models represents the most basic abstract structures, also binding with database and cache store in there, functions related with orm and cache model also in this directory. any io error throws by any function in models must be written in log system. package in models can not import object from other package in application except packages with the same level, on the other word, any other package should be able to import models package without worry about cycle import problem. Object in models package can contain `trans` method used to trans from elements in `params` package.

- `param` holds params from outside, usually from frontend, built from some objects from models package. common structures like `pager` which used to deal with pagination request also can be found in it.

- `router` routes defined in this package. usually routes will be registered in `init` function. And with function `WrapHandler`, you can write handlers in a common way. usually binding params and trans them to model in models here.

- `service` functions hold your logic in this package, usually combine your code in model, pass messages to MQ and other.

- `db.config.yaml` config file.

- `main.go` app entry, route register here

## Usage

- clone:

        git clone github.com/mint-leaf/gin-tpl

- replace:

  replace all `github.com/mint-leaf/gin-tpl` in import block with your app name

- write your code:

  write your code in your app and do not forget register routes in main.go

## Contribute

- PRs welcome
