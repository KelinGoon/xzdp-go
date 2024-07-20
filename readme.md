# *** 小众点评

## introduce

本项目使用go语言重构黑马点评项目，方便使用go语言的同学学习黑马程序员的redis课程，欢迎一起交流学习。  

- Use the [Hertz](https://github.com/cloudwego/hertz/) framework
- Integration of pprof, cors, recovery, access_log, gzip and other extensions of Hertz.
- Generating the base code for unit tests.
- Provides basic profile functions.
- Provides the most basic MVC code hierarchy.

## Directory structure

|  catalog   | introduce  |
|  ----  | ----  |
| conf  | Configuration files |
| main.go  | Startup file |
| hertz_gen  | Hertz generated model |
| biz/handler  | Used for request processing, validation and return of response. |
| biz/service  | The actual business logic. |
| biz/dal  | Logic for operating the storage layer |
| biz/route  | Routing and middleware registration |
| biz/utils  | Wrapped some common methods |

## How to run

```shell
sh build.sh
sh output/bootstrap.sh
```