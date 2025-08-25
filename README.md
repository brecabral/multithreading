# Desafio CEP Multithreading

Este projeto tem como objetivo exercitar **concorrência em Go (Golang)** utilizando **multithreading** e consumo de **APIs externas**.  
O sistema consulta simultaneamente dois provedores de CEP:
- **BrasilAPI** (brasilapi.com.br/api/cep/v1/ + CEP)
- **ViaCEP** (viacep.com.br/ws/ + CEP + /json/)

A aplicação retorna apenas o **primeiro resultado válido** e descarta o mais lento.  
Caso nenhuma resposta seja recebida em até **1 segundo**, é retornado um erro de **timeout**.

---

## Como executar

1. Clone o repositório
```sh
git clone https://github.com/brecabral/multithreading.git
cd multithreading
```
2. Instale as dependências e rode a aplicação
```sh
go mod tidy
go run ./cmd/server
```
O servidor subirá em: http://localhost:8000

## Exemplo de uso
Busque por um CEP com
```sh
curl http://localhost:8000/01153000
```
Possível retorno
```json
{
  "cep": "01153000",
  "logradouro": "Rua das Palmeiras",
  "bairro": "Barra Funda",
  "cidade": "São Paulo",
  "estado": "SP",
  "provider": "viacep"
}
```
Ou erro (timeout ou CEP inválido)
```json
{
  "error": "timeout"
}
```
## Documentação Swagger

Após rodar o servidor, acesse:
http://localhost:8000/docs/index.html

## Critérios atendidos

- [x] Multithreading para chamadas concorrentes
- [x] Retornar apenas a API mais rápida
- [x] Timeout de 1 segundo
- [x] Exibição dos dados no terminal
