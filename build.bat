@echo off
protoc  -I=. -I=%GOPATH%\bin\include --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./drivercommon/common.proto
protoc  -I=. -I=%GOPATH%\bin\include --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./thingmodel/thingmodel.proto
protoc  -I=. -I=%GOPATH%\bin\include --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./driverdevice/device.proto
protoc  -I=. -I=%GOPATH%\bin\include --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./gateway/gateway.proto
protoc  -I=. -I=%GOPATH%\bin\include --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./devicecallback/devicecallback.proto
protoc  -I=. -I=%GOPATH%\bin\include --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./productcallback/productcallback.proto
protoc  -I=. -I=%GOPATH%\bin\include --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./cloudinstancecallback/cloudinstancecallback.proto
protoc  -I=. -I=%GOPATH%\bin\include --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./cloudinstance/cloudinstance.proto
protoc  -I=. -I=%GOPATH%\bin\include --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./driverproduct/product.proto
protoc  -I=. -I=%GOPATH%\bin\include --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./custommqttmessage/custommqttmessage.proto
protoc  -I=. -I=%GOPATH%\bin\include --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./driverstorge/driverstorge.proto
pause