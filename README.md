fiber-pongo2-pkger:
---

Base project with `fiber` and `pongo2` as template + `pkger` to package static assets. 

### Usage?
- Clone it
- Update the `internal/config.go` file with your structured configuration
- Register your routes inside `pkg/app/registerroutes.go` 
- Ready to use

### Using Makefile
```
$ make

help                 Show available options with this Makefile
test                 Run all the tests
clean                Clean the application
vendor               Go get vendor
release              Create a release build.
bench                Benchmark the code.
prof                 Run the profiler.
prof_svg             Run the profiler and generate image.
generate_swagger     Generate swagger definitions from the comments
package              Package the html, css, js files etc
compress             Run upx on the generated binary in `build` directory
build_linux_only     Helper target to quickly build for linux without creating tar
```