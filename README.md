# Firebase Token Creator for CLI

Firebase Authentication を使ったサーバを実装するときにフロント作るのが面倒だったので、 CLI で取れるやつを作りました。

# How to use

```shell
$ export FTC_APIKEY=xxx
$ curl -H "Authorization:Bearer $(ftc)" localhost:8080
```

or

```shell
$ curl -H "Authorization:Bearer $(ftc --apikey xxx)" localhost:8080
```