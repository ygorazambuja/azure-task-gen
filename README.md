# Azure Task Generator

Uma ferramenta CLI que gera automaticamente tarefas no Azure DevOps baseadas em alterações de código, utilizando IA para criar descrições detalhadas e relevantes.

## 🚀 Funcionalidades

- Geração automática de tarefas baseada em alterações de código
- Integração com OpenAI para criar descrições inteligentes das alterações
- Suporte a múltiplos tipos de alterações (novos arquivos, modificações e deleções)
- Configuração personalizável para padrões de tarefas
- Geração de arquivo CSV compatível com Azure DevOps
- Interface interativa para configuração inicial
- Suporte para Windows e Linux

## 📋 Pré-requisitos

- Go 1.23.4 ou superior
- Chave de API do OpenAI
- Acesso ao Azure DevOps
- Git (para detectar alterações de código)

## 🛠️ Instalação

### Método 1: Instalação Global (Recomendado)

```bash
go install github.com/ygorazambuja/azure-task-gen@latest
```

Após a instalação, o comando `azure-task-gen` estará disponível globalmente no seu sistema.

### Método 2: Instalação Local

1. Clone o repositório:
```bash
git clone https://github.com/ygorazambuja/azure-task-gen.git
cd azure-task-gen
```

2. Instale as dependências:
```bash
go mod download
```

3. Compile o projeto:
```bash
go build
```

## ⚙️ Configuração

Na primeira execução, o programa solicitará as seguintes informações:

- Chave da API do OpenAI
- Nome padrão do responsável
- Email padrão do responsável
- Atividade padrão (ex: Development)
- Tipo de item de trabalho padrão (ex: Task)
- Unidade de Story Point padrão (ex: 4)

As configurações são salvas em:
- Windows: `%LOCALAPPDATA%\azure-task-gen\config.yaml`
- Linux: `~/.azure-task-gen/config.yaml`

## 🎯 Uso

1. Execute o programa:
```bash
azure-task-gen
```

2. Quando solicitado, insira:
   - ID da Sprint do Azure DevOps
   - ID do Caminho da Área do Azure DevOps

3. O programa irá:
   - Detectar alterações de código (novos, modificados e deletados)
   - Gerar descrições inteligentes usando IA
   - Criar um arquivo CSV com as tarefas geradas

4. Importe o arquivo CSV gerado no Azure DevOps

## 📝 Estrutura do CSV Gerado

O arquivo CSV gerado inclui os seguintes campos:
- ID
- Work Item Type
- Title
- Assigned To
- State
- Area ID
- Iteration ID
- Item Contrato
- ID SPF
- UST
- Complexidade
- Activity
- Description
- Estimate Made
- Remaining Work
- Original Estimate

## 🔧 Configuração Personalizada

Você pode personalizar as configurações padrão editando o arquivo `config.yaml` ou definindo variáveis de ambiente com o prefixo `AZURE_TASK_GEN_`.

## 🤝 Contribuindo

Contribuições são bem-vindas! Por favor, sinta-se à vontade para submeter pull requests.

## 📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

## 👥 Autor

- Ygor Azambuja

## 🙏 Agradecimentos

- OpenAI por fornecer a API de IA
- Azure DevOps por fornecer a plataforma de gerenciamento de projetos
- Todos os contribuidores e mantenedores das dependências utilizadas