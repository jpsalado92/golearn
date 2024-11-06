This comes from https://github.com/matt4biz/go-class-profile



Run the project TODO with form


## Run the profiler on the TODO project

Run the server with `go run .` from `go-class-profile-trunk/cmd/todo/main.go`.

Open the metrics page [http://localhost:8080/metrics](http://localhost:8080/metrics).

To see pprof, you'll need to open your browser to [http://localhost:8080/debug/pprof](http://localhost:8080/debug/pprof)


## Run the profiler on the sort algorithms project

Build the project with `go build .` from `go-class-profile-trunk/cmd/sort/main.go`.

http://localhost:8080/insert
http://localhost:8080/qsort
http://localhost:8080/qsortm
http://localhost:8080/qsort3
http://localhost:8080/qsorti
http://localhost:8080/qsortf


http://localhost:8080/debug/pprof
http://localhost:8080/debug/pprof/profile

## More info on pprof

https://go.dev/blog/pprof