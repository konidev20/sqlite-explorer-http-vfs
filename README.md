# sqlite-explorer - CLI tool to query SQL over HTTP
Run SELECT queries against SQLite databases over HTTP using a VFS driver. This repository also contains examples to use Pre-signed URLs to access SQLite databases stored in S3

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