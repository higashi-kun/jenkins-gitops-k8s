# syntax=docker/dockerfile:1
FROM golang:1.26 AS builder

WORKDIR /app

# 依存解決用ファイルを先にコピー → ソース変更時のキャッシュを効率化
COPY go.mod ./
# go.sum がある場合はここに COPY する(今回は依存ゼロなので不要)

COPY . .

# 静的バイナリを生成。CGO_ENABLED=0 + -ldflags "-s -w" でサイズ削減
# -s: シンボルテーブル削除 / -w: DWARF デバッグ情報削除
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o go-app

# -------- runtime --------
FROM gcr.io/distroless/static:nonroot

# nonroot ユーザ(UID 65532)で実行するため、バイナリを /home/nonroot に配置
COPY --from=builder --chown=nonroot:nonroot /app/go-app /home/nonroot/go-app

USER nonroot:nonroot

EXPOSE 8080

ENTRYPOINT ["/home/nonroot/go-app"]
