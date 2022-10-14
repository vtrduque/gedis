# gedis
Go implementation of Redis for **study purposes** =)

Implementation list:

- [x] TCP server
- [x] PING
- [ ] ECHO
- [ ] SET and GET
- [ ] Integration test/e2e with redis client
- [ ] Unit tests

# Run it
```sh
$ make run
$ echo "aeHO" | netcat localhost 6379
```

# Run tests
```sh
$ make test-all
```

or

```sh
$ make test-integration
$ make test-unit
```
