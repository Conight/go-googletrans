<div align="center">
    <img src="https://socialify.git.ci/Conight/go-googletrans/image?description=1&font=Inter&forks=1&language=1&logo=https%3A%2F%2Fupload.wikimedia.org%2Fwikipedia%2Fcommons%2Fd%2Fd7%2FGoogle_Translate_logo.svg&name=1&owner=1&pattern=Floating%20Cogs&stargazers=1&theme=Auto" alt="go-googletrans" width="640" height="320" />
</div>

# Googletrans

[![Sourcegraph](https://sourcegraph.com/github.com/Conight/go-googletrans/-/badge.svg)](https://sourcegraph.com/github.com/Conight/go-googletrans?badge)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/Conight/go-googletrans/blob/master/LICENSE)

This is Golang version of [py-googletrans](https://github.com/ssut/py-googletrans).

Googletrans is a **free** and **unlimited** Golang library that implemented Google Translate API.
This uses the [Google Translate Ajax API from Chrome extensions](https://chrome.google.com/webstore/detail/google-translate/aapbdbdomjkkjkaonfhkkikfgjllcleb) to make calls to such methods as detect and translate.

## Download from Github
```shell script
GO111MODULE=on go get github.com/Conight/go-googletrans
```

## Quick Start Example

### Simple translate
```go
package main

import (
	"fmt"
	"github.com/Conight/go-googletrans"
)

func main() {
	t := translator.New()
	result, err := t.Translate("你好，世界！", "auto", "en")
	if err != nil {
		panic(err)
	}
	fmt.Println(result.Text)
}
```

### Using proxy
```go
c := translator.Config{
    Proxy: "http://PROXY_HOST:PROXY_PORT",
}
t := translate.New(c)
```

### Using custom service urls or user agent
```go
c := translator.Config{
    UserAgent: []string{"Custom Agent"},
    ServiceUrls: []string{"translate.google.com.hk"},
}
t := translate.New(c)
```

See [Examples](./examples) for more examples.

## Special thanks

* [py-googletrans](https://github.com/ssut/py-googletrans)

## License
This SDK is distributed under the [The MIT License](https://opensource.org/licenses/MIT), see [LICENSE](./LICENSE) for more information.
