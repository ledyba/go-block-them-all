# 全自動Twitterブロック機

このツールは、PublicStreamを監視して、設定したキーワードが見つかったら自動で投稿者をブロックします。

ツイート本体だけでなく、ユーザーのdescriptionやURL、名前に対してもマッチを行います。

# How to use

conf/cred.goを設定してください。

```go
package conf

const ConsumerKey = ""
const ConsumerSecret = ""
const OAuthToken = ""
const OAuthSecret = ""
```

次に、NGワードを設定してください。

conf/words.go

```go
package conf

var NGCommon = []string{
	"雑学",
	"格言",
	"名言",
}

var NGText = []string{
	"#うちの子どうしてこうなった選手権",
	"#インターネット老人会",
	"nijie.info",
	"pixiv",
	"描きました",
	"入社式",
	"新社会人",
}

var NGDescription = []string{
	"公式ツイッター",
	"公式Twitter",
	"公式アカウント",
}

var NGScreenName = []string{
	"公式",
}

var NGURL = []string{
	"nijie.info",
	"pixiv",
}
```

気が済むまでこの世の嫌いなワードを並べたら、

```bash
make run
```
