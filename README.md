<h1 align="center">
  <img src="https://raw.githubusercontent.com/wiki/greymd/ojichat/img/ojichat_logo.png" height="178" width="485" />
   <h4 align="center">Ojisan Nanchatte (ojichat) Generator</h2>
</h1>

<p align="center">
  <a href="https://github.com/greymd/ojichat/releases/latest"><img src="https://img.shields.io/github/release/greymd/ojichat.svg" alt="Latest version" /></a>
  <a href="https://travis-ci.org/greymd/ojichat"><img src="https://travis-ci.org/greymd/ojichat.svg?branch=master" alt="Build Status" /></a>
  <a href="LICENSE" alt="MIT License"><img src="http://img.shields.io/badge/license-MIT-blue.svg?style=flat" /></a>
  <a href="https://codecov.io/gh/greymd/ojichat"><img src="https://codecov.io/gh/greymd/ojichat/branch/master/graph/badge.svg" alt="codecov" /></a>
</p>

## なんだこれは

おじさんがLINEやメールで送ってきそうな文を生成するコマンド。

## 開発環境

```bash
$ go version
go version go1.12 linux/amd64
```

## インストール

インストールには Go の開発環境が必要です。

```bash
go install github.com/greymd/ojichat@latest
```

## 使い方

```bash
$ ojichat -h
Usage:
  ojichat [options] [<name>]

Options:
  -h, --help      ヘルプを表示.
  -V, --version   バージョンを表示.
  -e <number>     絵文字/顔文字の最大連続数 [default: 4].
  -p <level>      句読点挿入頻度レベル [min:0, max:3] [default: 0].
```

そのまま実行すると文言が出力される。
文章は参考文献[1]で提唱される感情表現の順番で、いくつかのテンプレートの組み合わせにより自動生成がされる。

```bash
$ ojichat
ヤッホー😍😃れいこちゃん、元気かな⁉😜⁉️🤔オレは、近所に新しく できたラーメン屋さん🍜に行ってきたよ。味はまぁまぁだったかナ💕
```

文言には特定の人名が含まれることもあるが、第一引数で指定可能。

```bash
$ ojichat 山田
山田ちゃん、オハヨウ〜(^з<)😚（笑）山田ちゃんも今日も2時までお仕事かナ❓寒いけど、頑張ってね(＃￣З￣)🙂💤
```

`-p` オプションの数字を大きくする(最大3)することで文章に句読点が含まれやすくなる。
おじさんの文章には句読点が多い傾向が見られるため[1][2]、より実際の状況を模したユースケースに対応できる。


```bash
$ ojichat -p 3 オレとオマエと大五郎
オレと、オマエと、大五郎ﾁｬﾝ、オッハー❗(^_^)🎵オレと、オマエと 、大五郎ﾁｬﾝにとって、素敵な、1日に、なります、ようニ😘
```

`-e` オプションの数字を大きくすることで、絵文字/顔文字がより連続で含まれやすくなる。
一部のおじさんの文章にはそれらが多用される傾向があるためである。
また、引数を0とすることで真面目なおじさんにもなる。
より柔軟に実際の状況を模したユースケースに対応できる。

```bash
$ ojichat -e 10
おはよー、！チュッ😚😘😘😃☀ 😆❗😚😆(^з<)

$ ojichat -e 0
ヤッホー。はなみちゃん、元気かな。はなみちゃんにとって素敵な1日になりますようニ。
```

また、適宜、文節の終わりが最大2文字までカタカナとなる活用がされる。
これにより実際の状況を模したユースケースに(ry

```bash
$ ojichat
...ご要望とかはあるのかな❗💕😚😘😜❓

$ ojichat
...ご要望とかはあるのカナ❗🎵😆💕❓😜
```

## Dockerコンテナ版
おじさんで環境を汚したくない、Goの実行環境を持っていないなどの状況でもお手軽におじさんになるために、Dockerコンテナでもojichatを用意してある ( [greymd/ojichat](https://hub.docker.com/r/greymd/ojichat) )。

### 使い方

- `docker run --rm -i greymd/ojichat:latest` はオプション等を含めて全て `ojichat` と同じ動きをする。

```
$ docker run --rm -i greymd/ojichat:latest
ヤッホー(^з<)🎵（笑）キララチャン、元気かな😜⁉️土曜日は仕事〜❗❓キララチャン😚😃♥ 💗元気、ないのかなァ(^▽^;)💦大丈夫⁉😜⁉️✋❓❓
```

- `ojichat 坂東まりも` と同じ動きをする
```
$ docker run --rm -i greymd/ojichat:latest 坂東まりも
坂東まりもちゃん、久しぶり(^з<)(^з<)そういえば、昨日は例のラーメン屋さん🍜に行ってきたよ。結構いい雰囲気だったから、オススメだよ😚😚😍
```

## 関連ツール等 (Wiki)
https://github.com/greymd/ojichat/wiki/Related-services-and-tools

## 参考文献

[1]【SNSあるある】「おじさん」がLINEやメールで送ってきそうな文が話題に！【ソーシャルハラスメント？】 | こぐま速報
https://kogusoku.com/archives/2939

[2] 女子高生「おじさんLINEごっこ」の実例に学ぶキモがられる態度とは | ニュース3面鏡 | ダイヤモンド・オンライン
https://diamond.jp/articles/-/143111?page=2

[3] 女子同士がオジサンになりきってオジサンとオジサンがキャッキャする謎の「オジサンLINEごっこ」が流行の兆し - Togetter
https://togetter.com/li/1111905

## ライセンス

| 種類 | ライセンス |
| -- | -- |
| ソースコード | [MIT](./LISENCE) |
| ロゴ | <a rel="license" href="http://creativecommons.org/licenses/by-nc/4.0/"><img alt="クリエイティブ・コモンズ・ライセンス" style="border-width:0" src="https://i.creativecommons.org/l/by-nc/4.0/88x31.png" /></a>|

二次利用について詳しくは [Wiki > Copy and Redistribute](https://github.com/greymd/ojichat/wiki/Copy-and-Redistribute) 参照してください。
