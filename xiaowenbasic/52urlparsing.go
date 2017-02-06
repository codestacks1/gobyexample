// 52urlparsing
/*
	output:
			URL Parsing main
			postgres
			user:pass
			user
			pass
			host.com:5432
			host.com
			5432
			/path
			f
			k=v
			map[k:[v]]
			v
*/

package xiaowenbasic

import (
	"fmt"
	"net"
	"net/url"
)

func URLParsingMain() {
	fmt.Println("URL Parsing main")

	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	pl := fmt.Println

	pl(u.Scheme)
	pl(u.User)
	pl(u.User.Username())
	p, _ := u.User.Password()
	pl(p)

	pl(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	pl(host)
	pl(port)

	pl(u.Path)
	pl(u.Fragment)

	pl(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])

}
