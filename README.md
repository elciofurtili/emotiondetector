# Emotion Detector

Projeto simples em Go utilizando o padrão MVC para detectar a emoção predominante em um texto enviado pelo usuário, utilizando a API OpenAI (GPT-3.5-turbo) para processamento e TailwindCSS para estilização da interface web.

## Estrutura do Projeto

```
emotion-detector/
├── controllers/
│   └── emotion.go
├── models/
│   └── emotion.go
├── views/
│   └── index.gohtml
├── main.go
├── go.mod
├── go.sum
```

## Tecnologias Utilizadas

- Go (Golang)
- OpenAI API (gpt-3.5-turbo)
- TailwindCSS
- HTML Templates

## Como Usar

### 1. Configurar a variável de ambiente para a API Key do OpenAI

```bash
export OPENAI_API_KEY=seu_token_aqui
```

### 2. Instalar a dependência do OpenAI

```bash
go get github.com/sashabaranov/go-openai@latest
```

### 3. Rodar o projeto

```bash
go run main.go
```

### 4. Abrir no navegador

Acesse:

```
http://localhost:8080
```

Digite um texto no campo e envie para detectar a emoção predominante.

## Estrutura MVC

- **controllers/emotion.go**: contém a lógica do controlador para manipular requisições, interagir com a API OpenAI e renderizar a view.
- **views/index.gohtml**: template HTML com TailwindCSS para a interface.
- **main.go**: configura o servidor HTTP e as rotas.

## Detalhes da API OpenAI

- Utiliza o modelo `gpt-3.5-turbo`
- O prompt é enviado em português para retornar a emoção no mesmo idioma
- O resultado é exibido diretamente na página

## Estilização

- TailwindCSS via CDN
- Layout simples, responsivo e clean

## Contato

Para dúvidas ou melhorias, abra uma issue no repositório ou envie um contato direto.