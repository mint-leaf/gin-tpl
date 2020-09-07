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

- `models` models represents the most basic abstract structures, also binding with database and cache stored in there, functions related with orm and cache are also in this directory. any io error throws by any function in models must be written to log system. packages in models can not import object from any other package except those with the same level, on the other word, any other package should be able to import models' packages without worry about cycle import problem. object in models package could contain `trans` method used to trans from elements in `params` package.

- `param` holds params from outside, usually from frontend, usually combined with some objects from models package. common structures like `pager` which used to deal with pagination request also can be found in it.

- `router` routes defined in this package. usually routes will be registered in `init` function. and with function `WrapHandler`, you can write handlers in a common way. usually binding params and trans them to object in models here.

- `service` functions hold your logic in this package, usually combine your code in model, pass messages to MQ and other.

- `db.config.yaml` config file.

- `main.go` app entry, route register here

## Usage

- clone:

        git clone github.com/mint-leaf/gin-tpl

- replace:

  - in your project root dir and replace all `github.com/mint-leaf/gin-tpl` in import block with `YOUR_APP_NAME`

- go mod:

        go mod init {{YOUR_APP_NAME}}

- if GOPROXY:

  - use [goproxy.io](https://goproxy.io/) or [goproxy.cn](https://goproxy.cn/)

## Contribute

- PRs welcome
