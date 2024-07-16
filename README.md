**The Basics of Messenger Handler** é um projeto desenvolvido para o teste técnico da empresa [Smarters](https://smarte.rs/). Este projeto serve como um simples manipulador que envia e recebe mensagens do Messenger, incluindo recursos de log para rastrear todas as mensagens enviadas e recebidas. Além disso, o projeto inclui outro servidor TCP rodando em modo concorrente que emula um webhook do Messenger, realizando requisições para os endpoints do sistema principal.

#### Considerações gerais do projeto
* Este foi o meu primeiro projeto e praticamente o meu primeiro contato com a linguagem Go. Portanto, houve um atraso para aprimorar a web API com mais funcionalidades essenciais (testes automatizados, organização estrutural do projeto para evitar problemas futuros com ciclos de importação, adição de regras/técnicas de segurança, sistemas de banco de dados/cache e monitoramento) devido aos estudos necessários para entender a semântica e as técnicas específicas da linguagem Go.
* Um dos requisitos do teste é a detecção do envio de mensagens através do Messenger. Nesse sentido, é obrigatório implementar um webhook que enviará dados para um endpoint na plataforma da Meta. Como o endpoint precisa ser registrado com o protocolo HTTPS, seria necessário realizar o deploy do sistema para continuar com a funcionalidade. Porém, criei um servidor TCP HTTP para emular o sistema de recebimento de mensagens e um webhook que "captará" as mensagens através de eventos e enviará para a rota de recebimento no TCP principal.
* O emulador TCP do webhook do Messenger foi desenvolvido às pressas porque o Facebook ainda não havia aprovado a criação da minha conta para a utilização da Meta API. Sem um registro ativo, estou sujeito a uma espera considerável pela aprovação. Portanto, optei por criar o emulador.

#### Como rodar o projeto
Antes de tudo, é necessário ter o Go instalado na sua máquina.

```
git clone git@github.com:pyurips/the-basics-of-messenger-handler.git  
```

```
go mod download
```

```
go build -o main
```
Ou, caso esteja utilizando Windows:
```
go build -o main.exe
```

#### Informações sobre as variáveis de ambiente
No projeto, estão configuradas duas variáveis de ambiente: `ACCESS_TOKEN` e `EMULATOR`. Para rodar o projeto com o emulador do servidor TCP, a primeira deve ser configurada como `12345690`. Caso prefira não usar o emulador, o token de acesso deve ser obtido através da Meta API. A segunda variável determina se o sistema será executado com o emulador, podendo ter apenas dois valores: `true` ou `false`.

#### Sobre o registro de logs
Para cada mensagem enviada (para o emulador ou não) ou respostas de requisições quando o usuário é especificado no corpo,  será criada um diretório `logs` contendo os arquivos `.log`. A nomeação dos arquivos é de acordo com o ID do usuário do Messenger.
