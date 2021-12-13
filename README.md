# taskctf21-public
12/12(日)00:30 〜 12/12(日)23:59で開催したtaskctf21の公開用リポジトリです。

開催レポート: [誕生日CTFにてコードゴルフを裏開催した話](https://qiita.com/task4233/items/3460e60a65aa49de661c)

## ディレクトリ構成
ジャンル>問題名という順に並んでいます。
各問題には下記の構成でファイルが配置されています。

```bash
├── build               # 問題のビルドに必要なファイルが格納されています
├── dist(dist.zip)      # 配布されたファイルが格納されています(オプショナル)
├── solver              # 問題を解くためのスクリプトが格納されています(オプショナル)
├── src                 # 問題のソースコードが格納されています
├── docker-compose.yml
├── Dockerfile
├── writeup.md          # 問題の想定解法が書かれています
└── README.md           # 問題の概要等が書かれています
```


## LICENSE
MIT
