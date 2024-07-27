# Introdução

Este projeto implementa uma ferramenta de linha de comando simples para executar testes de estresse de urls utilizando golang.

## Execução Local

Para executar o projeto localmente, existem duas opções:

1. Utilizando golang diretamente (request go 1.22.2)
2. Utilizando o Docker (recomendado)

### Execução com go

Com o `go` devidamente instalado na sua máquina, execute o seguinte comando:

```bash
go run cmd/cli/main.go --url http://example.com --concurrency 1 --requests 1
```

Substitua os valores passados para os parâmetros de entrada com os valores desejados.

### Execução com o Docker

Com o Docker devidamente instalado no seu ambiente, execute os seguintes comandos:

```bash
docker build -t stresser:v1 .
```

O comando acima realiza o build do projeto como uma imagem Docker, utizando o nome `stresser` e a tag `v1`. Após o processo de `build`, execute o comando abaixo para executar o projeto em um container Docker:

```bash
docker run stresser:v1 --url http://example.com --concurrency 1 --requests 1
```
