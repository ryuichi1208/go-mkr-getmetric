# go-mkr-getmetric

Mackerel からメトリクス一覧を取得するシンプルな Go のツールです。

## 機能

- Mackerel のホスト一覧の取得
- 特定ホストのメトリクス一覧の取得

## 使い方

### API キーの設定

```bash
export MACKEREL_APIKEY="your-api-key"
```

または、コマンドライン引数で指定:

```bash
./go-mkr-getmetric -apikey "your-api-key"
```

### ホスト一覧の取得

```bash
./go-mkr-getmetric
```

### 特定ホストのメトリクス一覧の取得

```bash
./go-mkr-getmetric -host "host-id"
```

## Docker

### ビルド

```bash
docker build -t go-mkr-getmetric .
```

### 実行

```bash
docker run -e MACKEREL_APIKEY="your-api-key" go-mkr-getmetric
```

特定ホストのメトリクスを取得:

```bash
docker run -e MACKEREL_APIKEY="your-api-key" go-mkr-getmetric -host "host-id"
```

## GitHub Actions

このリポジトリには GitHub Actions の設定が含まれており、以下の処理が自動的に実行されます:

- テストの実行
- Docker イメージのビルドとプッシュ（main ブランチへの push のみ）

Docker イメージをプッシュするには GitHub Secrets に以下を設定する必要があります:

- `DOCKER_USERNAME`: Docker Hub のユーザー名
- `DOCKER_PASSWORD`: Docker Hub のパスワードまたはアクセストークン
