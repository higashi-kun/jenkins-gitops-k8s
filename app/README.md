# app

Phase 0 の Go アプリ。`/healthz` を返す HTTP サーバ(ポート 8080)。

## 概要

`/healthz` で `{"status":"ok"}` を JSON で返す最小 HTTP サーバ。
Jenkins / Kubernetes 等のヘルスチェック用途を想定。

## 前提条件

- Go 1.26 以上
- Podman 5.x(コンテナビルド・実行に必要)

## ローカル開発

### テスト

```sh
go test ./...
```

### サーバ起動

```sh
go run main.go
```

別ターミナルで:

```sh
curl -i localhost:8080/healthz
```

期待出力: `HTTP/1.1 200 OK` + `Content-Type: application/json` + `{"status":"ok"}`

## コンテナビルド・実行

```sh
podman build -t go-app .
podman run --rm -p 8080:8080 go-app
```

別ターミナルで:

```sh
curl -i localhost:8080/healthz
```
