# Tackem Global
Global Features for Tackem
## TODO
- concurrency RWMUTEX
- TESTING
  - To test the GPRC Server Section
  - Then The Config Section
  - Then figure out how to test last few files (run.go and server.go)
  - Once it's all done look at Github Actions to do unit testing.
  - Then move back to master and check the proto changes havn't really broken it
  - Start on unit testing of master
  - Move to User After And Test That.

- Github Actions
  - <https://github.com/mvdan/github-actions-golang>
  - <https://www.docker.com/blog/docker-golang/>
- Deeper Linter
  - <https://github.com/mgechev/revive>
## Improvements Where possible
  - any access to a service checks if its running and active and does the right action for where it is.
  - concurrency give as many places locks, make the objets all lockable with rwmutex to read and write lock.
  - any other lean up movement of code to make systems easier to read or find. move as much to its own places as possible

## Channels

## Config

## Health Check

## Helpers

## Logging

## ProtoBufs (pb and *.proto)

## Structs

## System Errors

## System

## Uses
- <google.golang.org/grpc>
- <github.com/xhit/go-str2duration/v2>
- <github.com/viney-shih/go-lock>
