# log-line-filter

reading log file line by line in go with `bufio.Scanner`, filtering by keyword or level, and counting each level in a map.

## usage
```bash
go run . -file server.log -level ERROR -keyword timeout
```
