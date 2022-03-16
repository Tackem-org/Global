# Tackem Global
Global Features for Tackem
## TODO
- concurrency RWMUTEX
- TESTING
  -config section needs a mock
  - <http://www.inanzzz.com/index.php/post/w9qr/unit-testing-golang-grpc-client-and-server-application-with-bufconn-package>
  - <https://stackoverflow.com/questions/42102496/testing-a-grpc-service>
  - <https://songrgg.github.io/programming/grpc-go-client-performance-test/>
  - <https://blog.bullgare.com/2019/05/golang-testing-http-and-grpc-servers/>
  - <https://medium.com/testvagrant/automating-grpc-api-tests-cfee49d10384>

- while adding all tests make sure:
  - any access to a service checks if its running and active and does the right action for where it is.
  - concurrency give as many places locks, make the objets all lockable with rwmutex to read and write lock.
  - any other lean up movement of code to make systems easier to read or find. move as much to its own places as possible

## Channels

## Config

## Health Check

## Helpers

## Logging

## Structs

## System Errors

## System

## Uses
- <google.golang.org/grpc>
- <github.com/xhit/go-str2duration/v2>
- <github.com/viney-shih/go-lock>
