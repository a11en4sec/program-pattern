## go generate
格式:
```
//go:generate command arg1 arg2
```
这样在同一个目录下执行命令go generate就会自动运行命令command arg1 arg2.

示例:
```go
//go:generate ./gen.sh ./template/container.tmp.go gen uint32 container
func fn(){}
```