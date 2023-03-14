# learn-go
golang learn record

> go 初始化模块

`go mod init xxx/xxx`

> go 模块简化

`go mod tidy`

> go 模块引用替换

`go mod edit -replace learn-go/pkg/greetings=../../greeting`

> test 显示详情

`go test -v`

> 查找打包后 exe 位置

`go list -f  '{{.Target}}'`

> window  设置 bin  路径

`go env -w GOBIN=F:\learn-go\bin`
`go install`
