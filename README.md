# go-template

Hello! This repository can be used as a starting point for command-line tools that are written in Go. Included is some boilerplate code for handling flags inside of a `main` function as well as a `Makefile` for installation and distribution across platforms.

This repository does not assume any specific package manager, although I recommend using Go modules.

## Makefile Tasks

**Note:** Before using the Makefile you should change the `PROJECT_NAME` variable near the top. By default it is `go-template`.

### `run`

Run the utility. It's exactly like executing `go run *.go`.

### `dist[-win|-macos|-linux]`

If no platform is specified, it will build the utility for all three supported platforms and plop them into the `dist` directory. If a platform is specified, it will only build for that platform.

### `build`

This will build the utility for your current platform and put it in the `bin` directory.

### `install`

This will run the `build` task then install the output file to `/usr/local/bin` by default. You can change this path either directly in the Makefile or via environment variable (`INSTALLPATH`).

### `uninstall`

Uninstalls the binary from `INSTALLPATH`.

## License

MIT