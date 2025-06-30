# Hello-golang

> A simple example showing how to create and publish a public Go module.
>
> 📖 Original blog post (中文): [Publish Golang Module](https://blog.tommylin.tech/post/golang/publish-golang-module/)

## Description

This repository provides a step-by-step guide to publishing a Go module.

## Publishing Steps

### 1. Create a new GitHub repository

<img src="./picture/create_github_repository.png" width="70%"/>

### 2. Configure repository settings

Choose the MIT license.

<img src="./picture/setting_repository_confi.png" width="70%"/>

### 3. Clone project to local

```shell
$ cd ~/workspace

$ git clone https://github.com/TommyLin81/hello-golang.git

$ cd hello-golang
```

### 4. Initialize Go module

```shell
$ go mod init github.com/TommyLin81/hello-golang
go: creating new go.mod: module github.com/TommyLin81/hello-golang

$ touch hellogolang.go

$ touch hellogolang_test.go
```

In `~/workspace/hello-golang/hellogolang.go`

```golang
package hellogolang

func Hello() string {
    return "hello golang !"
}
```

In `~/workspace/hello-golang/hellogolang_test.go`

```golang
package hellogolang

import "testing"

func TestHello(t *testing.T) {
    result := Hello()
    if result != "hello golang !" {
    	t.Fatal("Hello function return Fail")
    }
}
```

Run unit tests

```shell
$ go test
PASS
ok      github.com/TommyLin81/hello-golang      0.319s
```

### 5. Commit and tag a version

```shell
$ git add .

$ git commit -m 'feat: add Hello function'

$ git tag v0.1.0

$ git push

$ git push --tags
```

### 6. Publish module

```shell
$ GOPROXY=proxy.golang.org go list -m github.com/TommyLin81/hello-golang@v0.1.0
github.com/TommyLin81/hello-golang v0.1.0
```

After a short while, your published module will be available at: https://pkg.go.dev/github.com/TommyLin81/hello-golang

## References

- [Golang Document - Publishing a module](https://go.dev/doc/modules/publishing)