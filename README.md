# peloton

A client library for the undocumented Peloton API

[![Build Status](https://ci.j5y.xyz/api/badges/jchenry/peloton/status.svg)](https://ci.j5y.xyz/jchenry/peloton)

## Install

```
go get github.com/jchenry/peloton
```

## Usage

```
	j, _ := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	c := peloton.Client{
		HTTPClient: http.Client{
			Jar: j,
		},
	}
	peloton.Authenticate(c, "peloton_user", "password")
    	rides := peloton.GetRides(c, peloton.FilterSpec{
		Category: peloton.Cycling,
		Page:     0,
		Limit:    10000,
	})
	fmt.Printf("%#v", rides)
```

## Contributing

PRs accepted.

## License

MIT © Colin Henry
