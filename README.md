# AddonDeployer

X operationsの自作アドオンをまとめてaddonフォルダに移動するためのソフトウェア

**自分(駄場)の作成したアドオンアーカイブ**を配備するためのソフトウェアなので、他の方のアドオンに対しては使用できません。
ディレクトリ構造を自分のアドオンアーカイブと同じにすれば使用することもできますが、そこまでして使用する意味はないと思います。

## 使い方

### オプション一覧

|     オプション      |                    意味                    |
| :-----------------: | :----------------------------------------: |
| -i, --inputRootDir  | 入力ディレクトリ(アーカイブのディレクトリ) |
| -o, --outputRootDir | 出力ディレクトリ(XOPSのaddonディレクトリ)  |
|     -h, --help      |                ヘルプを表示                |
|    -v, --version    |            バージョン情報を表示            |

### 使用例

```
deploy.exe -i Archive -o addon
```

## プログラム情報

### 作者

駄場

### バージョン

1.0.0

### ライセンス

MIT

