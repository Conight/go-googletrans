# Googletrans

[![Sourcegraph](https://sourcegraph.com/github.com/Conight/go-googletrans/-/badge.svg)](https://sourcegraph.com/github.com/Conight/go-googletrans?badge)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/Conight/go-googletrans/blob/master/LICENSE)

This is Golang version of [py-googletrans](https://github.com/ssut/py-googletrans).

Googletrans is a **free** and **unlimited** Golang library that implemented Google Translate API.
This uses the [Google Translate Ajax API from Chrome extensions](https://chrome.google.com/webstore/detail/google-translate/aapbdbdomjkkjkaonfhkkikfgjllcleb) to make calls to such methods as detect and translate.

## Download from Github
```sh
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
	t := translate.New()
	result, err := t.Translate("你好，世界！", "auto", "en")
	if err != nil {
		panic(err)
	}
	fmt.Println(result.Text)
}
```

### Using proxy
```go
c := translate.Config{
    Proxy: "http://PROXY_HOST:PROXY_PORT",
}
t := translate.New(c)
```

### Using custom service urls or user agent
```go
c := translate.Config{
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
