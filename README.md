# kursho

### How to test the url shotener
1. Clone the repository
2. Run the command `make run` to start the server

3. shotren a url by running the following command:
```bash
$ curl -X POST -H "Content-Type: application/json" -d '{"url": "https://www.google.com"}' http://localhost:8080/shorten
```

4. Get the original url and follow it by running the following command:
```bash
$ curl -L http://localhost:8080/1735879616919141000
```
### Test App
The project includes a kind of test for consume the API, you can run it by running the following command:
```bash
$ make client
$ ./bin/client
```
