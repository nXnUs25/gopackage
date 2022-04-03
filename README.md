
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

```
