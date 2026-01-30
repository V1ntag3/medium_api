# ⚙️ Publify API — Backend do Clone do Medium

A **Publify API** é o backend da aplicação Publify, responsável pelo gerenciamento de usuários, artigos e autenticação.

Este projeto é uma API REST inspirada na plataforma [Medium](https://medium.com), permitindo que usuários publiquem e interajam com conteúdos.

---

## 🚀 Documentação da API

A documentação está disponível no Postman:

🔗 https://marcos-team.postman.co/workspace/4927b179-fda5-47d3-b94c-14c71518e289/documentation/17463653-e3d2ac4f-a213-4357-8ddc-a0bb9dc8035f

---

## 💻 Frontend

Repositório do frontend:

🔗 https://github.com/V1ntag3/medium_clone_front

---

## 🛠️ Tecnologias Utilizadas

### Backend
- Golang
- Fiber
- GORM
- SQLite3
- JWT (Autenticação)

---

## ✨ Funcionalidades

### 🔐 Autenticação
- Login
- Cadastro
- Logout

### 👤 Usuários
- Atualizar dados do usuário
- Excluir usuário
- Seguir usuário
- Deixar de seguir usuário
- Listar seguindo
- Listar seguidores
- Upload de imagem de perfil e banner

### 📝 Artigos
- Listar artigos
- Criar artigos
- Excluir artigos

---

## 📦 Instalação

### Pré-requisitos
- Golang instalado

### Passos

Clone o repositório e execute:

```bash
go run main.go
```
A API será iniciada por padrão em:

```bash
http://localhost:8080
```

(Caso tenha configurado outra porta, verifique no projeto.)

---

## 🔐 Configuração de Segurança
Antes de utilizar a aplicação em produção, altere a chave secreta utilizada para geração de tokens JWT.

Arquivo:
```bash
controllers/authControllers.go
```

Exemplo:

```bash
const SecretKey = "newSecretKey"
```
➡️ Substitua por uma chave segura.

---
 
## 🎯 Lições Aprendidas
Durante o desenvolvimento deste projeto, foram praticados:
- Criação de APIs REST
- Autenticação com JWT
- Integração com banco de dados via GORM
- Organização em camadas (controllers, models, routes)
- Boas práticas em Golang
- Desenvolvimento com Fiber

## 📂 Estrutura de Pastas (Exemplo)

```bash
.
├── controllers
│   └── authControllers.go
├── models
├── routes
├── database
└── main.go
```

🤝 Contribuição
Contribuições são bem-vindas!

Abra uma issue ou envie um pull request.
