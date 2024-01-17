# Exemplo de configuração de gGRP gateway


Neste projeto vai ser possivel usar a comunicação via gRPC e Http

**Não é necessário fazer isso já existe no projeto** 

Arquivos de configuração necessário para usar 

- annotations.proto
- http.proto

Os dois vão está disponivel nesse link [google api]("https://github.com/googleapis/googleapis/tree/master/google/api")


Vai ser preciso instalar algumas dependecias

```bash
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

já existe um comando criado no Makefile que já instala todos esses pacotes

```bash
make init
```

Para conseguir converter os arquivos `.proto`. É necessário instalar o [protobuf]("https://grpc.io/docs/protoc-installation"). Basta seguir como instalar na sua plataforma.

## Testando projeto

Basta executar o arquivo `main.go`

```bash
go run cmd/main.go
```

Para testar o serviço para executar esse comando 

```bash
curl --location --request POST 'http://localhost:5001/v1/message' \
--header 'Content-Type: application/json' \
--data-raw '{"name": "everyone"}'
```