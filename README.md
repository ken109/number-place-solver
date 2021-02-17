# About
コマンドラインで数独を解くプログラムです。

いけるところまで数独の規則に則った解き方を行い、途中から深さ優先探索に切り替わります。

# インストール
```
git clone https://github.com/ken109/go-nump.git
cd go-nump
go build
```

# Usage
```
./go-nump [問題のテキストファイル] [描画間隔(ms)]
```

# Example
```
./go-nump question.txt 80
./go-nump question2.txt 0
```
