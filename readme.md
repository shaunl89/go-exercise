# First Go


## Getting Started (OSX)

Download Go [here](https://golang.org/dl/) or run
```brew update``` then
```brew install go```

Run ```go env``` to check if installation is successful.

### Setup

Create a single workspace to manage all Go projects.

Setting GOPATH:

Open ```~/.bashrc``` and ```~/.bash_profile``` with your text editor and add in the following:

```
export GOPATH="/Users/your-username/go-workspace-name"
export PATH=$PATH:$GOPATH/bin
```

Create ```src``` ```pkg``` ```bin``` folders in the GOPATH directory.

$GOPATH/src : Contains Go source files

$GOPATH/pkg : Contains package objects

$GOPATH/bin : Contains executable commands

For more information, look [here](https://golang.org/doc/code.html#Workspaces) for the official documentation to setup your libraries and workspace.


### How to Use

Clone this repository into the appropriate folder within $GOPATH/src.

To run the program, enter the directory ```exercise``` in terminal and type
```
go run main/main.go
```

### Packages Used
```
"fmt"
"sort"
"strings"
"bufio"
"os"
```

For more information on packages, the official documentation can be found [here](https://golang.org/pkg/).

## Improvements

* Using a regex expression to filter out common text elements such as "!", "?" and "'s" so that processed text will be only words

* Incorporate some form of testing or validation

### Documentation used
* https://golang.org/doc/install
* https://golang.org/doc/code.html
* https://tour.golang.org/list
* https://golang.org/doc/effective_go.html
* https://rominirani.com/setup-go-development-environment-with-atom-editor-a87a12366fcf
