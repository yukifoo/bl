# Backlog CLI (bl) プロジェクトコンテキスト

## プロジェクト概要
Backlog用のCLIツール「bl」を開発しています。GitHub CLI (gh)と同じコマンド体系・基本的な動作を目指しています。

## 重要な参照情報

### 公式ドキュメント
- Backlog API: https://developer.nulab.com/docs/backlog/
- GitHub CLI (参考実装): https://github.com/cli/cli

### 使用ライブラリ
- Backlog Go Client: https://github.com/kenzo0107/backlog
- 既存の類似ツール（参考）: https://github.com/shufo/backlog-cli

## アーキテクチャ

### コマンド構造
```
bl <resource> <action> [flags]
```

### 主要コマンド
- `bl auth` - 認証管理
- `bl issue` - 課題管理
- `bl pr` - プルリクエスト管理
- `bl project` - プロジェクト管理
- `bl wiki` - Wiki管理
- `bl repo` - Gitリポジトリ管理

### ディレクトリ構成
```
bl/
├── cmd/          # コマンド実装
├── pkg/          # 共通パッケージ
│   ├── api/      # APIクライアント
│   ├── config/   # 設定管理
│   ├── format/   # 出力フォーマット
│   └── interactive/ # 対話的UI
└── internal/     # 内部パッケージ
```

## 技術スタック
- 言語: Go 1.21+
- CLIフレームワーク: Cobra
- 設定管理: Viper
- テーブル表示: tablewriter
- インタラクティブUI: survey

## 開発ガイドライン

### コーディング規約
1. Go標準のコーディングスタイルに従う
2. エラーハンドリングは明示的に行う
3. テストカバレッジ80%以上を維持

### コマンド実装パターン
```go
var cmdName = &cobra.Command{
    Use:   "action",
    Short: "簡潔な説明",
    RunE: func(cmd *cobra.Command, args []string) error {
        // 実装
    },
}
```

### 出力フォーマット
- デフォルト: テーブル形式
- オプション: JSON (`--format json`), CSV (`--format csv`)

## 現在の状況
- フェーズ1: 基盤構築中
- 設計ドキュメント: plan.md に記載
- Claude Code統合完了（GitHub Actions）
- 次のタスク: プロジェクト初期化とディレクトリ構造の作成

## Claude Code統合

### 使用方法
このプロジェクトではClaude Codeが統合されており、以下の方法で利用できます：

#### @claudeメンション
- PRやissueのコメントで `@claude` とメンションすることでClaude Codeを呼び出し
- 例: `@claude この関数のテストを書いて`

#### 自動PRレビュー
- 新しいPRが作成または更新されると自動的にClaude Codeがレビューを実行
- コード品質、バグ、パフォーマンス、セキュリティの観点からフィードバック

### 権限
- Claude Codeは以下の権限を持ちます：
  - `contents: read` - ファイル読み取り
  - `pull-requests: write` - PRコメント投稿
  - `issues: write` - issueコメント投稿
  - `actions: read` - CI結果の確認

### 利用できるツール
- ファイル読み書き（Read, Write, Edit）
- コード検索（Grep, Glob）
- バッシュコマンド実行（制限あり）
- PRコメント、レビュー投稿

## 開発時の注意点

### APIキー管理
- APIキーは環境変数または設定ファイルで管理
- 設定ファイルパス: `~/.config/bl/config.yml`
- 絶対にハードコードしない

### エラーメッセージ
- ユーザーフレンドリーなメッセージを心がける
- 詳細なデバッグ情報は`--verbose`フラグで表示

### テスト
- ユニットテストは必須
- モックを使用してAPI呼び出しをテスト
- 統合テストは別途実装

## よく使うコマンド

### ビルド
```bash
make build
```

### テスト実行
```bash
make test
```

### リント
```bash
make lint
```

## トラブルシューティング

### 認証エラー
1. APIキーが正しく設定されているか確認
2. スペースキーが正しいか確認
3. `bl auth status`で認証状態を確認

### API制限
- Backlog APIのレート制限に注意
- キャッシュ機能を活用して呼び出し回数を削減

## 今後の拡張予定
1. キャッシュ機能
2. プラグインシステム
3. シェル補完機能
4. オフラインモード