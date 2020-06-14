fiber-pongo2-pkger:
---

Template-project with `fiber` and `pongo2` as template + `pkger` to package static assets. 

### Usage?
- Clone it
- Update the `internal/config.go` file with your structured configuration
- Register your routes inside `pkg/app/registerroutes.go` 
- Ready to use

### Packaging:
This template-project uses `pkger` to package html-templates (`pongo2`) and static files like js, css etc. 

### Swagger specs
You can generate and serve swagger (OAP-2) specs and serve it on `swagger` end point.

### Build and run:
```
make build_linux_only && ./build/fiber-pongo2-pkger
```
Test it by navigating to
- http://localhost:3030
- http://localhost:3030/healthz


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