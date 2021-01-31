## Concurrency Testing

Using `ab` from the [apache-tools](http://httpd.apache.org/docs/2.4/programs/ab.html) package

```
ab -c 100 -n 1000 -T application/x-www-form-urlencoded -p ab-test.txt http://localhost:8080/
```

Where:
 - `-c` are the number of concurrent requests
 - `-n` is the total number of requests
 - `-T application/x-www-form-urlencoded` is content-type
 - `-p` points to the file `ab-test.txt` in this directory for an echo value
 - `http://localhost:8080` the url for where `echo-api` is running
