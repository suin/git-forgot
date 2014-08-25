# git forgot


* git commitを忘れてた…
* git pushを忘れてた…

というがっかりな事態を回避するためのツールです。


## インストール

Homebrewでインストールできます。

```
brew tap suin/suin
brew install git-forgot
```

(LinuxやWindowsのバイナリは[ダウンロード](https://drone.io/github.com/suin/git-forgot/files)から落とせます。)


## セットアップ

チェックしたいリポジトリのパスを環境変数 `GIT_FORGOT_DIR` にて設定します。.bashrcなどに書いておきます。

```
export GIT_FORGOT_DIR="$HOME/projects/* $HOME/github/*"
```

## 使い方


```console
$ git-forgot
~/project/backend needs to push and commit.
~/github/my-app needs to commit.
~/github/git-forgot needs to commit.
```

gitのコマンドの一部としても動作します。

```
git forgot
```

cronで定期的にチェックし、20分おきに通知センターで通知を受けることもできます。通知には[alloy/terminal-notifier](https://github.com/alloy/terminal-notifier)をインストールしておいてください。

```
PATH=/usr/local/bin:/usr/bin:/bin

*/20 * * * * /Users/suin/bin/git-forgot -r terminal-notifier $HOME/projects/* $HOME/github/* > /dev/null 2>&1
```