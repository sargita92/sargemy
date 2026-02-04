# Marketplace de Cursos em Vídeo (Backend)

Backend de um **marketplace de cursos em vídeo**, desenvolvido em **Go**, seguindo **arquitetura de microserviços orientados a domínio**, **Clean Architecture / Hexagonal Architecture** e preparado para **produção**, **escala**, **segurança** e **rastreabilidade financeira**.

Este repositório simula um **produto real de mercado**, não um projeto didático.

---

## 🎯 Objetivo do Projeto

Construir uma base sólida para uma plataforma onde:

* Criadores publicam cursos em vídeo
* Usuários compram cursos pela plataforma
* A plataforma intermedia pagamentos
* Há controle de acesso para cursos públicos e privados
* Todas as transações financeiras são auditáveis

O foco é **engenharia de software de alto nível**, com decisões arquiteturais que resistam a crescimento, refatorações e code reviews sênior.

---

## 🧱 Arquitetura

### Visão Geral

* **Microserviços por domínio (Bounded Context)**
* **Monorepo**, preparado para extração futura em múltiplos repositórios
* **API Gateway** como ponto único de entrada
* **Frontend simples como microserviço em Go (BFF)**
* **Banco de dados único em DEV, com múltiplos databases isolados**
* **Domínio isolado de infraestrutura e frameworks**

```
Browser → Frontend (Go) → API Gateway → Microserviços
```

### Princípios Arquiteturais

* Clean Architecture / Hexagonal Architecture
* Separação explícita entre:

    * Domain
    * Application (Use Cases)
    * Interfaces (HTTP)
    * Infrastructure
* Inversão de dependência
* Código orientado a interfaces
* Regras de negócio puras e testáveis

---

## 📦 Organização do Repositório

Cada pasta de primeiro nível representa um **microserviço completo e independente**:

```
/auth
/courses
/order
/payment
/deploy
```

### Estrutura Interna de um Serviço

```
<service>
├── cmd/api            # Entry point
├── internal
│   ├── domain         # Regras de negócio puras
│   ├── application    # Casos de uso
│   ├── interfaces     # HTTP / handlers
│   └── infra          # Banco, gateways, adapters
├── migrations         # Migrations versionadas
├── Dockerfile
└── go.mod
```

Cada serviço possui:

* Banco de dados próprio
* `go.mod` próprio
* Dockerfile próprio
* Deploy independente

---

## 🔐 Serviços

### Auth

Responsável por:

* Autenticação
* Gestão de usuários
* Emissão e validação de JWT

---

### Courses

Responsável por:

* Cadastro e exibição de cursos
* Cursos públicos e privados
* Convites e controle de acesso

---

### Order

Responsável por:

* Checkout
* Criação de ordens de compra
* Orquestração do fluxo de compra

---

### Payment (Domínio Crítico)

Responsável por:

* Cobranças
* Taxas da plataforma
* Repasse a criadores
* Ledger financeiro auditável

Características:

* Idempotência
* Transações explícitas
* Histórico imutável

---

## 🌐 API Gateway

O sistema utiliza um **API Gateway** para:

* Roteamento
* Autenticação inicial
* Rate limiting
* Observabilidade
* TLS termination

Nenhuma regra de negócio reside no gateway.

---

## 📊 Observabilidade

Desde o início o projeto considera:

* Logs estruturados
* Tracing distribuído
* Métricas de latência e erro
* Endpoints `/health` e `/ready`

---

## 🧪 Testes

* Testes unitários focados em regras de negócio
* Testes de integração para banco e pagamentos
* Uso consciente de mocks
* Cenários reais de marketplace

---

## 🐳 Ambiente de Desenvolvimento

* Docker
* Docker Compose
* Configuração via variáveis de ambiente
* Migrations automáticas

```
docker-compose up
```

---

## 🚀 Status do Projeto

🛠 **Em desenvolvimento ativo**

O projeto evolui de forma incremental, respeitando bases sólidas antes de novas features.

---

## 📄 Licença

Este projeto está licenciado sob a **Apache License 2.0**.

Você é livre para usar, modificar e distribuir este software, inclusive para fins comerciais, desde que respeite os termos da licença.

Veja o arquivo `LICENSE` para mais detalhes.

---

## 🧠 Filosofia

> Toda decisão arquitetural aqui deve sobreviver a um code review de um time sênior.

Simplicidade **sem fragilidade**.
Evolução **sem refatorações traumáticas**.
