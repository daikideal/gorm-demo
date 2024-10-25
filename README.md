# gorm-demo

## 環境構築

### 前提

- docker (Engine バージョン 19.03.0 以上) がインストールされていること

### postgres を起動

```shell
docker compose up -d
```

## 動作確認

```shell
go test -v ./...
```
