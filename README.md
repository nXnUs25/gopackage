
# Creating Go package

The idea here is to test the package releases to set master branch always as the *latest* release
if this works we will not have to specify v1 or v2 in `go.mod` requirments.

So to start we will be creating a package called _`gopackage`_ and then few other packages which are going to use it.

```shell
├── callgopackage-latest
├── callgopackage-v1
├── callgopackage-v1.1
├── callgopackage-v2
├── gopackage
│   ├── README.md
│   └── gopackage.go
```

***gopackage*** will be a simple package displaying Version, where as v2 release will change a dependency with v1
and one new version will be released we will merge it into the master (aka main) branch

so to start developiong a new package:

```go
package gopackage

import "fmt"

const VERSION = "1.0.0"
const ver = VERSION

func Version() string {
	return ver
}

func ShowVersion() {
	fmt.Printf("Current version: %s\n", ver)
}
```

so now we can can initialize new package and test the package usage:

```shell
❯ cd gopackage
❯ go mod init github.com/nXnUs25/gopackage/v1
go: creating new go.mod: module github.com/nXnUs25/gopackage/v1
❯ git checkout -b v1
Switched to a new branch 'v1'
❯ git add .
❯ git commit -m "initialize version 1.0.0"
[v1 eb7bccd] initialize version 1.0.0
 3 files changed, 62 insertions(+), 1 deletion(-)
 rewrite README.md (100%)
 create mode 100644 go.mod
 create mode 100644 gopackage.go
❯ git tag v1.0.0
❯ git push --tags origin v1
Warning: Permanently added 'github.com' (ED25519) to the list of known hosts.
Enumerating objects: 7, done.
Counting objects: 100% (7/7), done.
Delta compression using up to 8 threads
Compressing objects: 100% (4/4), done.
Writing objects: 100% (5/5), 1.07 KiB | 1.07 MiB/s, done.
Total 5 (delta 0), reused 0 (delta 0)
To github.com:nXnUs25/gopackage.git
 * [new branch]      v1 -> v1
 * [new tag]         v1.0.0 -> v1.0.0
❯ git --no-pager branch -a
  main
* v1
  remotes/origin/HEAD -> origin/main
  remotes/origin/main
  remotes/origin/v1
```

Now lets use it in the following packages

```
├── callgopackage-latest
├── callgopackage-v1
```

for this we will just create `go mod` so we have a file called `go.mod`
