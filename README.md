# sqlite-explorer - CLI tool to query SQL over HTTP

# SELECT Query
```
./sqlite-explorer -url <url-to-sqlite-database> -query <SELECT-SQL-query-to-be-executed>
```
## Example
```
./sqlitehttpcli -url 'https://www.sanford.io/demo.db' -query 'select * from csv'
```

# References
[https://github.com/psanford/sqlite3vfshttp](sqlitehttpcli/sqlitehttpcli.go)