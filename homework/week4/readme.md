# Go 工程化实践

## /cmd
可执行代码
代码不应该放太多，尤其不应该放业务代码
业务代码放pkg 或者 internal


## /pkg
想共享给外部使用的代码,  外部程序可以使用的库代码

## /internal
不想共享给外部使用的代码

/internal/app/myapp 放一些app的业务代码
/internal/pkg  多个app共享的代码