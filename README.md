# docker

## env
[docker\-composeから環境変数をDockerfileに渡す方法 \- Qiita](https://qiita.com/Targityen/items/2717511ca9f12c1c667f)
dockerにわたす環境変数は
.env => docker-compose.yml内で${}
.env => docker-compose.ymlでargsに渡す => Dockerfile内でARGにキー名を渡して${}

## マルチステージビルド
[マルチステージビルドの利用 \| Docker ドキュメント](https://matsuand.github.io/docs.docker.jp.onthefly/develop/develop-images/multistage-build/)

# 参考

## 開発環境

- [VSCode に Go の開発環境を整える \- iriya\-ufo's blog](https://iriya-ufo.net/blog/2019/12/08/go-env-in-vscode/)
- [tools/user\.md at master · golang/tools](https://github.com/golang/tools/blob/master/gopls/doc/user.md)
- [VSCode で Go の Language Server である gopls を有効にする \- Qiita](https://qiita.com/ryysud/items/1cf66ee4363aec22394a)
- [Error installing tools when using Go version managers \(asdf\-vm\) · Issue \#3087 · microsoft/vscode\-go](https://github.com/microsoft/vscode-go/issues/3087)
  - [all: cleanup use of \`go\.goroot\` · Issue \#146 · golang/vscode\-go](https://github.com/golang/vscode-go/issues/146)
- [【golang 環境構築】Modules で package 管理をする \- Qiita](https://qiita.com/fox777/items/a8cb025df5439902b6c4)
- [GOPATH モードからモジュール対応モードへ移行せよ \- Qiita](https://qiita.com/spiegel-im-spiegel/items/5cb1587cb55d6f6a34d7)

