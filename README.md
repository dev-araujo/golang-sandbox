# Golang Sandbox

Sandbox para estudos e experimentos com Go.

## Projetos

### go-job
Sistema de agendamento de tarefas desenvolvido em Go.

- **Repositório original:** [dev-araujo/go-job](https://github.com/dev-araujo/go-job) (arquivado)
- **Localização:** [/go-job](./go-job)
- **Tecnologias:** Go, Gin, GORM, SQLite
- **Descrição:** API REST para gerenciamento de vagas de emprego com operações CRUD

---

## Estrutura

```
golang-sandbox/
├── README.md (este arquivo)
└── go-job/
    ├── config/
    ├── handler/
    ├── router/
    ├── schemas/
    ├── main.go
    └── go.mod
```

## Como usar

### Clonar o repositório

```bash
git clone https://github.com/dev-araujo/golang-sandbox.git
cd golang-sandbox
```

### Executar um projeto específico

Cada subpasta contém um projeto independente com seu próprio README e instruções.

#### Exemplo: go-job

```bash
cd go-job
go mod download
go run main.go
```

A API estará disponível em `http://localhost:8080`.

Consulte o README dentro de cada pasta para instruções detalhadas.

---

## Adicionar novos projetos

Para adicionar um novo projeto de estudo em Go:

1. Crie uma nova pasta na raiz do repositório
2. Inicialize o módulo Go: `go mod init`
3. Adicione o código do projeto
4. Inclua um README.md com instruções
5. Atualize este README principal com a descrição do novo projeto

---

## Tecnologias

- **Go** - Linguagem de programação
- **Gin** - Framework web
- **GORM** - ORM para Go
- **SQLite** - Banco de dados

---

**Nota:** Este repositório consolida múltiplos projetos de estudo em Go para facilitar organização e manutenção.
