## Создание GRPC
-   Выполняем команду `protoc --proto_path=./proto --go_out=. --go-grpc_out=. ./proto/transmitter.proto`
-   Получаем `./transmitter/transmitter_grpc.pb.go` и `./transmitter/transmitter.pb.go`