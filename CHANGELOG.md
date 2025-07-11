# 📦 Changelog

Todas as mudanças importantes neste projeto serão documentadas aqui.

O formato é baseado em [Keep a Changelog](https://keepachangelog.com/pt-BR/1.0.0/), e este projeto segue a [SemVer](https://semver.org/lang/pt-BR/) para versionamento.

---

## [v1.0.0] - 2025-07-10

### ✨ Adicionado
- Comando `filter` para extrair linhas contendo "Pagamento Pix"
- Comando `extract` para extrair pares `TXID : EndToEndId`
- Comando `version` com versão embutida via flag de build
- Suporte a `.env` para token do GitHub
- `Makefile` com comandos facilitados e variáveis customizadas
- Pipeline do GitHub Actions com testes e release automático
- GoReleaser com builds multiplataforma (Windows/Linux/macOS)

---

## [Unreleased]

### 🔧 Em desenvolvimento
- Suporte a logs binários
- Comando `stats` para gerar resumo estatístico de transações Pix
