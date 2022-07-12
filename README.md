# practical-go-programming

実用Go言語の写経リポジトリ

```sh
$ go env -w GO111MODULE=on # GOPATH/src配下以外で開発できるようにするおまじない
$ go mod init github.com/katsukiniwa/practical-go-programming # プロジェクト初期化
$ go get github.com/foo/bar@latest # 最新バージョンのライブラリを追加
$ go mod tidy # go.sumが作成される
```

## 参考記事

[GOPATH に(可能な限り)依存しない Go 開発環境(Go 1.15 版)](https://zenn.dev/tennashi/articles/3b87a8d924bc9c43573e)
