# Real-Time Chat WebSocket Avançado

Este projeto é um **chat em tempo real** desenvolvido em Go, utilizando WebSockets para comunicação bidirecional. Ele foi projetado de forma modular, com funcionalidades avançadas, incluindo múltiplas salas, mensagens privadas, comandos especiais, rate limiting, histórico de mensagens e logging assíncrono.

---

## Funcionalidades

- **Chat em Tempo Real:**  
  - Envio e recepção instantânea de mensagens entre clientes conectados à mesma sala.

- **Múltiplas Salas:**  
  - Cada usuário escolhe uma sala (padrão: `default`) para isolar as conversas.

- **Mensagens Privadas:**  
  - Comando `/pm <usuário> <mensagem>` para enviar mensagens privadas a um usuário específico na mesma sala.

- **Comandos de Sistema:**  
  - `/users` – Lista os usuários conectados na sala.  
  - `/kick <usuário>` – Expulsa (kick) um usuário da sala (disponível apenas para o usuário **admin**).

- **Histórico de Mensagens:**  
  - Armazena as últimas 50 mensagens de cada sala, acessível via endpoint REST e pelo botão "Carregar Histórico" no frontend.

- **Indicador de "Digitando":**  
  - Exibe uma notificação quando um usuário está digitando.

- **Rate Limiter:**  
  - Limita a quantidade de mensagens enviadas por cada usuário em um intervalo de tempo, prevenindo spam.

- **Logging Assíncrono:**  
  - Registra eventos importantes (como entrada/saída de usuários e erros) no arquivo `chat.log`.

- **Endpoints REST para Monitoramento:**  
  - `/history`: Retorna o histórico de mensagens da sala.  
  - `/users`: Retorna a lista de usuários conectados na sala.  
  - `/status`: Exibe informações básicas do servidor (ex.: total de clientes conectados).

---

## Instalação e Execução

### Pré-requisitos

- [Go (Golang)](https://golang.org/dl/) (versão 1.18 ou superior recomendada)
- Git (opcional, para clonar o repositório)

### Passo a Passo

1. **Clone o repositório ou baixe o código-fonte:**

   ```bash
   git clone https://github.com/seu-usuario/real-time-chat.git
   cd real-time-chat
   ```

2. **Inicialize o módulo Go (se ainda não estiver feito):**

   ```bash
   go mod init chat-websocket
   go mod tidy
   ```

3. **Instale as dependências:**

   O principal pacote de dependência é o Gorilla WebSocket. Caso não seja instalado automaticamente, execute:

   ```bash
   go get github.com/gorilla/websocket
   ```

4. **Execute o servidor:**

   Na raiz do projeto, execute:

   ```bash
   go run main.go
   ```

   Você deverá ver uma mensagem informando que o servidor foi iniciado na porta `:8080`.

5. **Acesse a aplicação:**

   Abra o navegador e acesse:

   ```
   http://localhost:8080/index.html
   ```

---

## Utilizando o Chat

- **Tela de Login:**  
  Informe seu nome e escolha uma sala (padrão: `default`). Clique em **Entrar** para acessar o chat.

- **Interface do Chat:**  
  - **Envio de Mensagens:** Digite sua mensagem e clique em **Enviar**.
  - **Carregar Histórico:** Clique em **Carregar Histórico** para recarregar as últimas 50 mensagens da sala.
  - **Indicador de "Digitando":** Enquanto um usuário estiver digitando, um indicador aparecerá.
  - **Atualização de Usuários:** A lista de usuários é atualizada automaticamente a cada 5 segundos.

- **Comandos de Chat:**  
  - **`/users`**: Lista os usuários conectados na sala.  
  - **`/pm <usuário> <mensagem>`**: Envia uma mensagem privada ao usuário especificado.  
  - **`/kick <usuário>`**: Expulsa um usuário da sala (somente se o remetente for "admin").

- **Endpoints REST:**  
  - `GET /history?room=<nome_da_sala>`: Retorna o histórico de mensagens da sala.  
  - `GET /users?room=<nome_da_sala>`: Retorna a lista de usuários conectados na sala.  
  - `GET /status`: Retorna informações básicas do servidor.

---

## Detalhes Técnicos

- **Modularização:**  
  O código está organizado em pacotes (`hub`, `models`, `handlers`, `utils`) para facilitar a manutenção e futuras expansões.

- **Comunicação em Tempo Real:**  
  Utiliza WebSockets (com Gorilla WebSocket) para permitir a comunicação bidirecional entre o servidor e os clientes.

- **Goroutines e Canais:**  
  - **Hub:** A função `hub.Run()` é executada em uma goroutine para processar e distribuir mensagens (via canal `Broadcast`) para todos os clientes conectados à mesma sala.
  - **Conexões WebSocket:** Cada nova conexão inicia seu próprio loop de leitura em uma goroutine.
  - **Rate Limiter:** Para cada cliente, uma goroutine (através de `ratelimiter.StartRateLimiter`) gerencia a reposição de tokens para limitar a taxa de envio de mensagens.
  - **Logger:** Uma goroutine é utilizada para registrar eventos assíncronos no arquivo `chat.log`.

- **Rate Limiting:**  
  Implementa controle para evitar spam, limitando quantas mensagens um usuário pode enviar em um intervalo de tempo.

- **Logging Assíncrono:**  
  Eventos importantes são registrados em `chat.log`, facilitando o monitoramento e auditoria do sistema.

---

## Possíveis Melhorias

- **Autenticação e Autorização:**  
  Implementar um sistema de login real para controlar o acesso e permissões.

- **Persistência:**  
  Utilizar um banco de dados para armazenar o histórico de mensagens e dados dos usuários.

- **Interface do Usuário:**  
  Atualizar o frontend com frameworks modernos (React, Vue.js, etc.) para uma experiência aprimorada.

- **Escalabilidade:**  
  Suporte a múltiplas instâncias do servidor e balanceamento de carga para ambientes de alta demanda.

---

## Contribuição

Contribuições são bem-vindas! Se desejar colaborar, por favor, envie _pull requests_ ou abra _issues_ para reportar bugs ou sugerir melhorias.

---

## Licença

Este projeto está licenciado sob a [MIT License](LICENSE).

---