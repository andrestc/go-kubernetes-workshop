# go-kubernetes-workshop
Aprenda a customizar o Kubernetes utilizando Go.

## Agenda

### Conceitos básicos do Kubernetes (1h)

#### Apresentação 

1. containers
2. pods
3. rs
4. deployments
5. services
6. endpoint
7. API do Kubernetes

#### Exercicio 

Apresentar o ambiente dos exercicios (katacoda ou minikube)
Utilizar o kubectl para escutar as alterações de um tipo de objeto e alterar outro objeto na API do Kubernetes.

### Client-go

#### Apresentação

1. Mostrar a estrutura dos pacotes do client, como vendorizar etc...
2. Exercicio basico de consulta/alteracao na API utilizando o client-go
3. Testes unitários usando o client-go

#### Exercicio

Utilizar o client-go para dar watch nas alterações de um tipo de object e alterar outro objeto (mesmo exercício que do modulo anterior).

Escrever testes para o Watch utilizando reaction.

### Controllers

#### Apresentação

0. Qual o problema do exemplo anterior? Dica: muitas chamadas a API e perda de eventos.
1. O que são os controllers?
2. Informers
3. Como criar um controller simples em Go

#### Exercicio: 

Criar um controller utilizando o boilerplate do https://github.com/kubernetes/sample-controller que faz as mesmas coisas que os ultimos 2 exercicios.

### Custom Resource Definitions

#### Apresentação

1. O que são CRDs?
2. Como criar um CRD

#### Exercicio

Como criar o CR, CRD e toda a parte de geração de código.
Utilizar o kubectl apply para inserir o CR no cluster.
Controller que escuta as alterações nos CRs e faz alguma ação.

### Operators

#### Apresentação

0. Problema da abordagem de controller. Dica: muito boilerplate
1. O que são operators?
2. Operator Framework (https://github.com/operator-framework)
0. Comandos do operator-sdk

#### Exercicio

Evoluir o exercicio de controllers utilizando o operator-sdk

