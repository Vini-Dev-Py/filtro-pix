# 🧾 filtro-pix

**filtro-pix** é uma ferramenta de linha de comando (CLI) escrita em Go para processar arquivos de log de pagamentos Pix.

Permite:
- Filtrar linhas contendo `"Pagamento Pix"`
- Extrair pares `TXID : EndToEndId` de logs JSON
- Verificar a versão da ferramenta

---

## ⚙️ Requisitos

- Go 1.20 ou superior
- Git Bash, terminal Linux/macOS ou WSL
- (opcional) `goreleaser` para releases automatizados
- (opcional) `make` para facilitar desenvolvimento local

---

## 🚀 Instalação

```bash
git clone https://github.com/seu-usuario/filtro-pix.git
cd filtro-pix
go build -o filtro-pix
```

---

## 💻 Modo de uso

### ✅ Executável compilado (CLI)

```bash
./filtro-pix filter -i entrada.log -o pagamentos_pix.log
./filtro-pix extract -i pagamentos_pix.log -o txids.txt
./filtro-pix version
```

---

### 🛠️ Ambiente de desenvolvimento (usando Makefile)

Você pode usar `make` para facilitar comandos durante o desenvolvimento:

```bash
make build            # Compila o binário
make run              # Executa com main.go
make clean            # Remove binário
make filter           # Usa entrada.log e gera pagamentos_pix.log
make extract          # Usa pagamentos_pix.log e gera txids.txt
make version          # Mostra a versão atual
make snapshot         # Gera release local com GoReleaser
make release          # Publica release real no GitHub
```

Com parâmetros personalizados:

```bash
make filter INPUT=meu.log OUTPUT=saida.log
make extract OUTPUT=saida.log EXTRACT_OUTPUT=txid_resultado.txt
```

---

## 📦 Comandos disponíveis

### 🔎 `filter`
Filtra todas as linhas que contenham a palavra `Pagamento Pix`.

```bash
filtro-pix filter -i entrada.log -o pagamentos_pix.log
```

---

### 🧬 `extract`
Extrai pares `TXID : EndToEndId` de logs filtrados que contêm JSON Pix.

```bash
filtro-pix extract -i pagamentos_pix.log
filtro-pix extract -i pagamentos_pix.log -o txids.txt
```

---

### 🏷 `version`
Exibe a versão atual da ferramenta:

```bash
filtro-pix version
```

---

## 📁 Estrutura do Projeto

```
.
├── cmd/
│   ├── extract_test.go
│   ├── extract.go
│   ├── filter_test.go
│   ├── filter.go
│   ├── root.go
│   ├── version.go
├── dist/
│   ├── filtro-pix_windows_amd64/
│   ├── filtro-pix_linux_amd64/
│   ├── filtro-pix_darwin_arm64/
│   └── filtro-pix_darwin_amd64/
├── main.go
├── go.mod
├── go.sum
├── .env
├── Makefile
├── .goreleaser.yaml
└── README.md
```

---

## 👨‍💻 Autor

Desenvolvido por [Vinicius Guilherme Batista](https://github.com/Vini-Dev-Py)

---

## 📄 Licença

Este projeto está licenciado sob a licença MIT. Veja o arquivo `LICENSE` para mais detalhes.
