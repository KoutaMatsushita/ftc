# Firebase Token Creator for CLI

Firebase Authentication を使ったサーバを実装するときにフロント作るのが面倒だったので、 CLI で取れるやつを作りました。

# How to use

`GOOGLE_APPLICATION_CREDENTIALS` については https://firebase.google.com/docs/admin/setup?hl=ja#initialize-sdk を参照。

```shell
$ export GOOGLE_APPLICATION_CREDENTIALS=path/to/service-account.json
$ export FTC_APIKEY=xxx
$ curl -H "Authorization:Bearer $(ftc)" localhost:8080
```

or

```shell
$ export GOOGLE_APPLICATION_CREDENTIALS=path/to/service-account.json
$ curl -H "Authorization:Bearer $(ftc --apikey xxx)" localhost:8080
```

# How to install

## Homebrew

```shell
$ brew tap KoutaMatsushita/ftc
$ brew install ftc
```
