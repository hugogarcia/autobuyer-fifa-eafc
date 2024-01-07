# Sobre
Um simples auto buyer de EA FC que criei para uso pessoal. Mas fique à vontade para utilizá-lo

Não tem interface e sua configuração é manual. A única parte automatizada é a compra de cartas.

> Este auto buyer simula o EA FC Web App, portanto, se você estiver jogando no console, não poderá executar este bot.

# Stack
- Go 1.19

# AVISO!!!!!!!!!!!!!!!!!
PODE ACONTECER DA EA BLOQUEAR SUA CONTA NA WEB PARA FAZER TRADE

TENHA EM MENTE ESSE RISCO

# O que o bot faz?
- Lance nas cartas
- Observa as cartas que já sofreram lance, e faz o lance novamente caso necessário

# Configurar Planilha (CSV)
1. Encontre no [Futbin](https://www.futbin.com/) o jogador que quer buscar
2. Pegue o código do `ID` dele, vai estar do lado esquerdo com as informações do jogador
3. Adicione na planilha o ID, Nome e Valor Máximo do lance para aquela carta. As informações devem estar separadas por ponto e virgula(;)

## Example
```
243812;Rodrygo;5000
206517;Jack Grealish;3200
186345;Kieran Trippier;3100
220834;Marco Asensio;800
266933;Arianna Caruso;800
227119;Steph Catley;800
```

# Configurar variáveis de ambiente
Primeiramente, abra o arquivo `.env` que estará na pasta. Vou mostrar abaixo a configuração de cada variável:

### `GAME_SKU`
  - Informe o código de qual a plataforma você usa para jogar, as opções são:
    - FFA24XBO - Xbox One
    - FFA24XSX - Xbox X
    - FFA24PS5 - PS5
    - FFA24PS4 - PS4
    - FFA24PCC - PC
### `USER_ID`
  - Acesse o site do EA FC Web App e faça o login.
  - Quando estiver na tela inicial pressione F12.
  - Vá na aba `Console`.
  - Digite o comando `localStorage.getItem('_eadp.identity.pidId')`.
  - Copie esse código e cole no USER_ID do arquivo.
### `TOKEN`
  - Acesse o site do EA FC Web App e faça o login.
  - Quando estiver na tela inicial pressione F12.
  - Vá na aba `Console`.
  - Digite o comando `localStorage.getItem('_eadp.identity.access_token')`.
  - Copie esse texto e cole no TOKEN do arquivo.

# Como executar
Após ter executado o passo de [configurar variáveis de ambiente](#configurar-variáveis-de-ambiente), é possível executar o bot.

## Executável (Windows Apenas)
Na pasta terá um executável que iniciará o bot.

## Terminal
- Instalar Go 1.19
- Executar o comando `go run main.go`

# Meu bot está fechando sozinho
## Causa 1: 
  Se ao executar o bot a tela já fechar em seguida é porque está com problema na leitura do CSV, verifique se está configurado certo.

## Causa 2:
  Se o bot fechar depois que estiver executando e fazendo lances, é porque a EA bloqueou temporariamente o acesso. Isso é normal, a EA bloqueia quando há muitas tentativas de trade pela Web. Saia da conta e refaça o login no site, e refaça o processo de configurar a variável [TOKEN](#token).
## Causa 3: 
  Se o bot fechar depois que estiver executando e fazendo os lances, tente logar novamente no EA FC Web e verifique se está conseguindo logar. Pode ser que a EA bloqueie sua conta no Web definitivamente, nesse caso o bot nunca mais funcionará 😢 O bloqueio será só na versão Web, você poderá fazer trade normalmente por dentro do jogo.

# Doação
Se gostou do bot e puder me ajudar:

https://www.buymeacoffee.com/hugogarcia

