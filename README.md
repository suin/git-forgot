# :::非推奨:::

このプロジェクトは非推奨です。https://github.com/suin/git-remind を代わりに使ってください。

# git forgot

git add したが commit を忘れていた。コミット漏れがあって同僚に迷惑をかけた。複数のリポジトリを同時に触っていたら、片方のリポジトリをpushするのを忘れてしまった。gitを使ったことがあれば、このような失敗をしたことがあるのではないだろうか？

忘れることは人間である以上しかたない。であれば、思い出させてくれる何かがあればいい。例えば、通知センターでコミット漏れを教えてくれるだけでも十分だ。

![Screen Shot 2014-08-25 at 11.12.22.png](https://qiita-image-store.s3.amazonaws.com/0/889/cfd35194-7924-4e52-a20f-59f587843ed8.png "Screen Shot 2014-08-25 at 11.12.22.png")

そこで、今回作ったツールが[git-forgot]。このツールは、`git forgot`と打つとpushやcommitがされてないリポジトリを教えてくれるものだ。

![2__Shell.png](https://qiita-image-store.s3.amazonaws.com/0/889/16b42c12-af04-8867-2310-acc5da83446e.png "2__Shell.png")

この記事では、git-forgtについて紹介したい。

## インストール方法

インストールはMacであればHomebrewでできる。

```
brew tap suin/suin
brew install git-forgot
```

LinuxやWindowsは[ダウンロードコーナー]にバイナリを配布しているのでダウンロードしてほしい。

## セットアップ

git-forgotでチェックしたいディレクトリを環境変数 `GIT_FORGOT_DIR` に設定しておく必要がある。複数またはアスタリスクで指定できる。.gitがあるディレクトリのみチェックされるので、関係ないディレクトリが含まれていても構わない。

```
export GIT_FORGOT_DIR="$HOME/projects/* $HOME/github/*"
```

## 使い方

セットアップが完了したら、あとはコマンドを叩くだけでチェックできるようになる。コミット漏れやプッシュ忘れがあるリポジトリがあれば、そのパスが出る。問題なければ何も出ない。

```console
$ git-forgot
```

なお、gitのサブコマンドとしても使える。

```console
$ git forgot
```

リポーター(`--reporter`, `-r`)を指定すれば、通知センターに通知を送ることもできる。通知には[terminal-notifier]が必要なので、インストールしておいて欲しい。

```console
$ git forgot --reporter terminal-notifier
```

また、`git forgot -r terminal-notifier -t iTerm` のようにコンソールを`-t`で渡すと、通知センターの通知をクリックしたときに、該当プロジェクトを開けるようになる([m0aさん提供])。

crontabにしこんでおけば、定期的にコミット漏れなどをチェックし、通知センターで通知を受け取ることができる。

```sh:crontab
PATH=/usr/local/bin:/usr/bin:/bin

*/20 * * * * git-forgot -r terminal-notifier $HOME/projects/* $HOME/github/* > /dev/null 2>&1
```

pecoと組み合わせても便利。

```
alias gf='arr=(`git-forgot|peco`) && cd $arr[1]'
```

その他のオプションについては `-h` でヘルプを見れる。

```console
$ git forgot -h
```



[git-forgot]: https://github.com/suin/git-forgot
[ダウンロードコーナー]: https://drone.io/github.com/suin/git-forgot/files
[terminal-notifier]: https://github.com/alloy/terminal-notifier
[m0aさん提供]: https://github.com/suin/git-forgot/pull/1
