# Automatos

Conversão de uma linguagem livre de contexto para a forma de Chomsky

## Representações
1. O vazio é representado por: `0`
2. O alfabeto é dado pelo range: `[aA-zZ]`

## Input

Para criar um `automato`

```golang
newAutomato := models.NewAutomato([]string{})
```

A função `models.NewAutomato` recebe uma lista de string, onde cada item dessa lista seria a representação de uma expressão, e.g:

```golang
expressionList := []string{
  "A -> B",
  "B -> b",
}

newAutomato := models.NewAutomato(expressionList)
```
## Resolvendo
Ja para rodar a transformação para forma de chomsky, basta passar o automato como argumento para a função `solver.Solve`, que retorna um novo `automato` na forma de Chomsky, eg:

```golang
chomskyAutomato := solver.Solve(oldAutomato, false)
```

## Mostrando os resultados
Basta chamar o `.Explain` de um `automato`
eg:
```golang
chomskyAutomato.Explain()
```

## Como rodar

Rodar diretamente pelo terminal:

```bash
$ go run main.go
```
ou

Compilar e executar:

```bash
$ go build
$ ./automatos
```

Obs: O comando build ira compilar para o SO em que o comando foi executado, caso queira compilar para outro SO:

```bash
$ export GOOS={ windows | linux  | ... }
$ go build
$ ./automatos
```
Mais informações sobre compile do Go [aqui](https://golang.org/doc/install/source#environment)

### Requerimentos
[GoLang](https://golang.org/doc/install)
```bash
#Testado em:
  $ go version
  -> go version go1.8.3 linux/amd64
```