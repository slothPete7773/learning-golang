## Source code execution

**Run code**

Run only `main.go` file.

```sh
$ go run main.go
```

Run several files in the current directory.

```sh
$ go run .
```

**Compile code**

```sh
$ go build main.go
```

## Modules and packages organization

### **Basic package**

The code in modname.go declares the package with:

```
project-root-directory/
  go.mod
  modname.go
  modname_test.go
```

A Go package can be split into multiple files, all residing within the same directory, e.g.:

```
project-root-directory/
  go.mod
  modname.go
  modname_test.go
  auth.go
  auth_test.go
  hash.go
  hash_test.go
```

### **Basic Command**

Command in Go is simply the executable program (CLI tool).

The simplest program can consist of a single Go file where func main is defined. 

```
project-root-directory/
  go.mod
  auth.go
  auth_test.go
  client.go
  main.go
```

### **Multiple commands**

- Multiple programs in the same repository will typically have separate directories:
- In each directory, the program’s Go files declare package main.
- A top-level internal directory can contain shared packages used by all commands in the repository.
```
project-root-directory/
  go.mod
  internal/
    ... shared internal packages
  prog1/
    main.go
  prog2/
    main.go
```

### **Package with supporting package**

- Supporting package here is `auth/`. Inside the `internal/` directory.
- Effectively, placing supporting package is to prevent other modules from depending on the package and avoiding expose and external use.

```
project-root-directory/
  internal/
    auth/
      auth.go
      auth_test.go
    hash/
      hash.go
      hash_test.go
  go.mod
  modname.go
  modname_test.go
```

### **Multiple packages**

- each package has its own directory, and can be structured hierarchically.

```
project-root-directory/
  go.mod
  modname.go
  modname_test.go
  auth/
    auth.go
    auth_test.go
    token/
      token.go
      token_test.go
  hash/
    hash.go
  internal/
    trace/
      trace.go
```

- Sub-packages can be imported by users as follows:

```
import "github.com/someuser/modname/auth"
import "github.com/someuser/modname/auth/token"
import "github.com/someuser/modname/hash"
```

### **Package for server project**

**it’s recommended to **
- Keep the Go packages implementing the server’s logic in the internal directory.
- Since the project is likely to have many other directories with non-Go files, it’s a good idea to keep all Go commands together in a cmd directory:

```
project-root-directory/
  go.mod
  internal/
    auth/
      ...
    metrics/
      ...
    model/
      ...
  cmd/
    api-server/
      main.go
    metrics-analyzer/
      main.go
    ...
  ... the project's other directories with non-Go code
```