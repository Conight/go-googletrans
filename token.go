package translator

import (
	"fmt"
	"io"
	"math"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var ReTkk = regexp.MustCompile(`tkk:'(.+?)'`)

type tokenAcquirer struct {
	tkk    string
	host   string
	client *http.Client
}

func Token(host string, client *http.Client) *tokenAcquirer {
	host = strings.ToLower(host)
	if !strings.HasPrefix(host, "http") {
		host = "https://" + host
	}
	return &tokenAcquirer{
		tkk:    "0",
		host:   host,
		client: client,
	}
}

func (a *tokenAcquirer) do(text string) (string, error) {
	err := a.update()
	if err != nil {
		return "", err
	}
	tk := a.acquire(text)
	return tk, nil
}

func (a *tokenAcquirer) update() error {
	now := int(math.Floor(float64(time.Now().UnixNano()) / 1000000.00 / 3600000.00))

	tkk, _ := strconv.Atoi(strings.Split(a.tkk, ".")[0])
	if a.tkk != "" && tkk == now {
		return nil
	}

	resp, err := a.client.Get(a.host)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	rawTkk := ReTkk.FindStringSubmatch(string(body))
	if len(rawTkk) > 0 {
		a.tkk = rawTkk[1]
		return nil
	}
	return nil
}

func (a *tokenAcquirer) acquire(text string) string {
	var textSlice []int
	for _, value := range text {
		val := int(value)
		if value < 0x10000 {
			textSlice = append(textSlice, val)
		} else {
			textSlice = append(textSlice, int(math.Floor(float64(val-0x10000)/0x400+0xD800)))
			textSlice = append(textSlice, int(math.Floor(float64((val-0x10000)%0x400)+0xDC00)))
		}
	}
	x := ""
	if a.tkk != "0" {
		x = a.tkk
	}
	d := strings.Split(x, ".")
	b := 0
	if len(d) > 1 {
		b, _ = strconv.Atoi(d[0])
	}

	var e []int
	g := 0
	size := len(textSlice)
	for g < size {
		l := textSlice[g]
		// just append if l is less than 128(ascii: DEL)
		if l < 128 {
			e = append(e, l)
		} else {
			// append calculated value if l is less than 2048
			if l < 2048 {
				e = append(e, l>>6|192)
			} else {
				if (l&64512) == 55296 && g+1 < size && textSlice[g+1]&64512 == 56320 {
					g += 1
					l = 65536 + ((l & 1023) << 10) + (textSlice[g] & 1023)
					e = append(e, l>>18|240)
					e = append(e, l>>12&63|128)
				} else {
					e = append(e, l>>12|224)
				}
				e = append(e, l>>6&63|128)
			}
			e = append(e, l&63|128)
		}
		g += 1
	}
	temp := b
	for _, value := range e {
		temp += value
		temp = xr(temp, "+-a^+6")
	}
	temp = xr(temp, "+-3^+b+-f")
	if len(d) > 1 {
		t, _ := strconv.Atoi(d[1])
		temp ^= t
	} else {
		temp ^= 0
	}
	if temp < 0 {
		temp = (temp & 2147483647) + 2147483648
	}
	temp %= 1000000

	return fmt.Sprintf("%d.%d", temp, temp^b)
}

func rShift(val, n int) int {
	return (val % 0x100000000) >> n
}

func xr(a int, b string) int {
	lenB := len(b)
	c := 0
	for c < lenB-2 {
		stringD := string(b[c+2])
		intD := int(b[c+2])
		if stringD >= "a" {
			intD = intD - 87
		} else {
			temp, _ := strconv.Atoi(stringD)
			intD = temp
		}
		if string(b[c+1]) == "+" {
			intD = rShift(a, intD)
		} else {
			intD = a << intD
		}
		if string(b[c]) == "+" {
			a = (a + intD) & 4294967295
		} else {
			a = a ^ intD
		}
		c += 3
	}
	return a
}
