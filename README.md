# About
コマンドラインで数独を解くプログラムです。

いけるところまで数独の規則に則った解き方を行い、途中から深さ優先探索に切り替わります。

# インストール
```
git clone https://github.com/ken109/go-nump.git
cd go-nump
```

# Usage
```
go run main.go [問題のテキストファイル] [描画間隔(ms)]
```

# Example
question2.txtは[世界一難しいと言われている問題](https://www.sentohsharyoga.com/ja/puzzle/blog/entry/sudoku_most_difficult#:~:text=%E4%B8%96%E7%95%8C%E3%81%A7%E6%9C%80%E3%82%82%E9%9B%A3%E3%81%97%E3%81%84%E6%95%B0%E7%8B%AC%E3%81%A8%E3%81%AF,%E3%82%92%E8%A6%81%E3%81%97%E3%81%9F%E3%82%88%E3%81%86%E3%81%A7%E3%81%99%E3%80%82)です。
```
go run main.go question.txt 80
go run main.go question2.txt 0
```
