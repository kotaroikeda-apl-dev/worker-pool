# Go の Worker Pool を使った並列処理

## 概要

このプロジェクトは、Go 言語の `goroutine` と `channel` を活用し、**ワーカープール (Worker Pool)** を実装するサンプルです。固定数のワーカー (`goroutine`) を使用して、複数のジョブを並行処理する方法を示します。

## `Worker Pool` とは？

**Worker Pool** は、一定数の `goroutine`（ワーカー）を用意し、ジョブを `channel` に送信しながら効率的に並列処理を行う設計パターンです。

## **実行方法**

```sh
go run cmd/basic/main.go  #ワーカープールの基本実装
```

## **学習ポイント**

1. **`channel`** を使うことで **`goroutine`** 間でデータをやり取りできる
2. **`sync.WaitGroup`** を活用し、ワーカーの終了を待機できる
3. **`close(jobs)`** により、チャネルの送信を終了し、ワーカーにジョブがないことを通知する
4. **ワーカープールを使うことで、CPU 負荷を分散しつつ効率的に並列処理を行える**

## 作成者

- **池田虎太郎** | [GitHub プロフィール](https://github.com/kotaroikeda-apl-dev)
