# loog
View json logs in a human-readable format, for terminal lovers

## Usage
Pipe any stream of logs into loog in order to see it formatted
```shell
tail -f /var/log/gibrish.json | loog
```

## Roadmap
- [ ] Allow printing as a vertical table
- [ ] Allow printing as a horizontal table
- [ ] Allow printing a log separator
- [ ] Use cobra to define the CLI
- [ ] Compile for multiple platforms
- [ ] Test piping logs into loog
- [ ] Refactor code and add tests
- [ ] Add examples to documentation
- [ ] Allow marking specific fields using a flag
