# go-proxy

## About

HTTP server for proxying HTTP-requests to 3rd-party services.

## Installation

```shell
git clone git@github.com:zuzuleinen/go-proxy.git
go run main.go
```

## Sample request

POST http://localhost:8000

```json
{
  "method": "GET",
  "url": "http://google.ro",
  "headers": {
  	"Authentication": "Basic bG9naW46cGFzc3dvcmQ=&quot;,"
  }
}
```

## Sample response

```json
{
	"id": "4120bbe3-cb2e-405a-b43f-568243d4a966",
	"status": "200 OK",
	"headers": {
		"Cache-Control": "private, max-age=0",
		"Content-Type": "text/html; charset=ISO-8859-2",
		"Date": "Sat, 24 Apr 2021 08:04:44 GMT",
		"Expires": "-1",
		"P3p": "CP=\"This is not a P3P policy! See g.co/p3phelp for more info.\"",
		"Server": "gws",
		"Set-Cookie": "NID=214=RdGpUBztBTgzJhDqJp5WzDj2d0KpMD_ezz3pNQW2bXF4Nxo0u_0IZTyAfnvGyaKTSTJEqp49aU4dTipgSrfCl4OKZlkYGCY8Z09HobNWUWNccBmE5RKkyWm8YGsMUH_jEMd8TUFGmSN7y63aVap5ZwAtULtdYjL6PUrKwiuT0XA; expires=Sun, 24-Oct-2021 08:04:44 GMT; path=/; domain=.google.ro; HttpOnly",
		"X-Frame-Options": "SAMEORIGIN",
		"X-Xss-Protection": "0"
	},
	"length": -1
}
```

## Storing the request-response

Each request to server and response from 3rd party sevice is logged in a in-memory database:

```golang
// To save
	db := DB()
	db.Save(id, reqDump, respDump)

// To check current rows of db

	for k, v := range db.Rows() {
		fmt.Println("For id", k)
		fmt.Println("Request", string(v.Request))
		fmt.Println("Request", string(v.Response))
	}

```
		

