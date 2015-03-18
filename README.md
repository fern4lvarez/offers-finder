offers-finder
=============

`offers-finder` is a server of Berlin-based offers. Each offer contains an
id, and latitude and longitude coordinates, so they can be located in the map.

`offers-finder` is a HTTP application written entirely in Go, using its standard
library only, as a showcase of its strength and flexibility.

Getting Started
---------------

### Install

~~~
go get github.com/fern4lvarez/offers-finder
~~~

### Run the tests

~~~
go test -v -cover .
~~~

### Run the server

~~~
# Note: It runs on port 3000 so ensure its available

# Optional if you want to set a custom password
export OFFERS_FINDER_PASSWORD=secret

offers-finder
~~~

Specification
-------------

* `GET` request to `/` returns `200`

	```
	⇒ curl -i http://127.0.0.1:3000/          
	HTTP/1.1 200 OK
	Access-Control-Allow-Headers: Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization
	Access-Control-Allow-Origin: *
	Content-Type: application/json
	Date: Wed, 18 Mar 2015 16:38:59 GMT
	Content-Length: 16

	{"status":"OK"}
	```

* Authenticated `POST` request to `/v1/token` returns `200` and random and unique
  40 character token:

	```
	⇒ curl -i -X POST -u "user:secret" http://127.0.0.1:3000/v1/token     
	HTTP/1.1 200 OK
	Access-Control-Allow-Headers: Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization
	Access-Control-Allow-Origin: *
	Content-Type: application/json
	Date: Wed, 18 Mar 2015 16:41:22 GMT
	Content-Length: 53

	{"token":"NFaQ0IJUcmn75QUXfem2rkaCZkOG8MXqm0cIgFNA"}
	```

* Authenticated `POST` request with the token (as a form parameter) to `/v1/offers`
  returns `200` and randomly and random list of Berlin offers in the format:

    ```js
    [
      {
        "id": 0,
        "lat": 52.5452407270436,
        "long": 13.412785613836808
      }
    ]
    ```

	~~~
	⇒ curl -i -X POST -u "user:secret" http://127.0.0.1:3000/v1/offers\?token\=NFaQ0IJUcmn75QUXfem2rkaCZkOG8MXqm0cIgFNA
	HTTP/1.1 200 OK
	Access-Control-Allow-Headers: Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization
	Access-Control-Allow-Origin: *
	Content-Type: application/json
	Date: Wed, 18 Mar 2015 16:45:36 GMT
	Content-Length: 357

	[{"id":0,"lat":52.512498900043845,"long":13.303739745914536},{"id":1,"lat":52.51202029761004,"long":13.350967639139323},{"id":2,"lat":52.52210593055171,"long":13.44626918440136},{"id":3,"lat":52.54331296047837,"long":13.46356050111962},{"id":4,"lat":52.509743468901824,"long":13.430912335328042},{"id":5,"lat":52.524105958211635,"long":13.295507420281938}]
	~~~

* Unauthenticated `POST` request to `/v1/token` and /v1/offers returns `401`:

	```
	⇒ curl -i -X POST -u "wrong:wrong" http://127.0.0.1:3000/v1/token 
	HTTP/1.1 401 Unauthorized
	Access-Control-Allow-Headers: Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization
	Access-Control-Allow-Origin: *
	Content-Type: application/json
	Www-Authenticate: Basic realm="Authorization Required"
	Date: Wed, 18 Mar 2015 16:50:48 GMT
	Content-Length: 0
	```

* Any other kind of request returns `404`:

	```
	⇒ curl -i  http://127.0.0.1:3000/bhjgbj                                  
	HTTP/1.1 404 Not Found
	Access-Control-Allow-Headers: Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization
	Access-Control-Allow-Origin: *
	Content-Type: text/plain; charset=utf-8
	Date: Wed, 18 Mar 2015 16:54:30 GMT
	Content-Length: 19

	404 page not found
```

Development
-----------

If you are working on `offers-finder` development, this repo provides
with a `run-server` script that compiles the project, run tests, execute
[`gometalinter`](https://github.com/alecthomas/gometalinter) and starts
the server.

~~~
./run-server
~~~

License
-------

MIT
