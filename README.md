# 🌸 Flowerly - Floricultura Online

Bem-vindo ao Flowerly! Este projeto é uma aplicação web para uma floricultura online, construída com Go e o framework Fiber. No Flowerly, os clientes podem navegar por nossa seleção de flores, realizar compras e gerenciar seus pedidos de maneira fácil e prática.

## 🚀 Funcionalidades

- **Catálogo de Flores**: Explore nossa vasta coleção de flores disponíveis para venda.
- **Autenticação de Usuários**: Cadastre-se, faça login e gerencie suas informações de conta.
- **Carrinho de Compras**: Adicione flores ao carrinho e finalize suas compras de forma simples.
- **Gestão de Pedidos**: Acompanhe seus pedidos anteriores e os que estão em andamento.
- **Painel Administrativo**: Admins podem gerenciar o inventário de flores, visualizar e gerenciar pedidos.

## 🐳 Configuração com Docker

### Pré-requisitos

- Docker e Docker Compose instalados em seu sistema.

### Passos

**Clone o repositório**:
````bash
git clone https://github.com/seuusuario/flowerly.git 
cd flowerly
````

**Configure as variáveis de ambiente**:
```makefile
DB_HOST=db
DB_USER=seu_usuario_db
DB_PASSWORD=sua_senha_db
DB_NAME=flowerly
```

**Inicie a aplicação com Docker Compose**:
```bash
docker-compose up -d
```

## 📦 Dependências

- **Fiber**: Framework web minimalista e rápido para Go.
- **PostgreSQL**: Sistema de banco de dados relacional utilizado para armazenar as informações da floricultura.
