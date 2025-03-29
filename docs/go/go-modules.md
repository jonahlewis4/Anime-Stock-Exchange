<h1>GO MODULES</h1>

Go modules are a dependency management system for go.

Before go modules, you needed to develop your go code inside a certain directory that matched the `GOPATH` variable. 

Go modules will handle external dependencies for you.

<h2>Initialize all modules to start with the project's repository name on github</h2>
```bash
go mod init github.com/jonahlewis4/Anime-Stock-Exchange/scraper
```

be sure to create `go.sum` file (this happens automatically if there are actual dependencies, but if initially starting out and trying to use a Dockerfile that copies the go.sum file, there will be issues)
```bash
touch go.sum
```

<h3>Run go mod tidy to update and verify all dependencies</h3>
```bash
go mod tidy
```
