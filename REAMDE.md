# 1.安装 protoc 以及对应语言的插件
https://www.grpc.io/docs/languages/go/quickstart/
# 2.根据proto文件生成对应语言的代码
```shell script
# 生成go代码：
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative helloworld/helloworld.proto
# 生成python代码：
python3 -m grpc_tools.protoc -I . --python_out=. --grpc_python_out=. helloworld.proto
```