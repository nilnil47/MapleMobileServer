## installation

For the generating the grpc files you will need:
* protoc compiler
* go-grpc plugin
* make

###Install For Windows:

From Powershell with **admin** 
```
# install protoc
choco install protoc --pre

# install make
choco install make 

# install go-grpc plugin
go get github.com/golang/protobuf/protoc-gen-go
```
 