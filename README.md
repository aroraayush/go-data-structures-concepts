
```
 nodemon --watch './*.go' --signal SIGTERM --exec 'go' run server.go 
```
---
<details>
<summary>Setup</summary>

#### Add GOPATH to your PATH
```
vim ~/.zshrc
export GOPATH=/Users/ayush/go
export PATH=$PATH:$GOPATH
export PATH=$PATH:$GOPATH/bin
source ~/.zshrc
sudo chmod 777 /bin/goimports
```
- Copy project from `/usr/local/go/src` to `/Users/ayush/go/src/`
### Change VS Code Settings
```
"go.gopath": "/Users/ayush/go",
"go.toolsGopath": "/Users/ayush/go",
```
- Clone project in the `User/ayush/go/src`
#### Initialize our project as a Go module
```
go mod init github.com/aroraayush/golang-data-structures
```
#### Output
`go: creating new go.mod: module github.com/aroraayush/golang-data-structures`
- A new file will appear in your directory called `go.mod`
    - `go.mod` contains information about our module for others, such as the module name and Go version it is intended for

</details>

<details>
<summary>GOPATH vs GOROOT</summary>

Installing Golang via Homebrew automatically generates two directories critical to running Go:
- `GOROOT` ( /usr/local/go ): 
    
    The Go "root" directory contains <ins> **Go's source code**</ins>. Homebrew will automatically register this path for you; there's little reason to mess around in here unless you're a Go contributor or if you're attempting to run multiple versions of Go.
- `GOPATH` ( /Users/ayush/go ): 
    
    Unlike most programming languages, Go takes an opinionated stance that <ins> **all projects** </ins>  and dependencies of the language <ins> **should exist in a single directory known as the GOPATH**</ins> . Any time we develop a Go project or install a third-party module, the actions taken ultimately happen inside this directory
    
</details>

---
<details>
<summary>Go Packages</summary>

- Go programs are made up of "packages," which mirror packaging concepts in other programming languages (think modules in Python or packages in Java). 

- Every Golang program contains a **package called**  `main`, which serves as the project's entry point.

---

</details>
<br>

To make the package available to other apps, we need to install corresponding package in to `go runtime`
```
cd <package_name>
go install
```
- It will be stored in `User/ayush/go/pkg/<os_type>/<project_name>`
- For mac, os_type is `darwin_amd64`
    - `<package_name>.a` file is generated | Compiled package, linkable
    
- package is now avaible in `main.go`

---

<details>
<summary>Go Modules [third-party libraries]</summary>

Go modules are third-party libraries installed by Go. Modules are essentially projects which have been published for general use as dependencies in your projects.

</details>

---

<details>
<summary>Go Vendors</summary>
<br>

While modules can be installed to the `/pkg/mod` directory for global use, **source projects** can contain their **own versions of these modules to avoid clashing dependency** versions between projects (**similar to Python virtual environments**). While not required, you can choose to keep module versions project-specific (we will do this in our example).
</details>

#### Vendor directory
```
brew install dep
dep init # in src/main
```
- `Gopkg.toml` like pakage.json will be generated
- dep puts dependencies into the vendor folder


---
Go is a compiled language, we need to build our project before we can run it 
```
go build
go run main.go
```
<details>
<summary>Code Formatting</summary>
    
    go fmt
</details>

---
Do `go run <filename>` in each directory

---
## nil reference
- golang uses `nil` instead of `null` or `None`

## Primitives
- #### var
    ```
    var x string
    x = "Hello World"  // x is assigned the string "Hello World”
    ```
    OR
    ```
    var x string = "Hello World" 
    ```
- #### short declaration / walrus  operator (:=)
    ```
    myvariable3 := 200
    ```
- #### const
    - Constants `cannot` be declared using the `:=` syntax.
    - Constants can be character, string, boolean, or numeric values.
    ```
    const Pi = 3.14
    ```

- #### `var` vs `short`  declaration operator `(:=)`
    
    -   `var` - used to declare and initialize the variables `inside` and `outside` the functions. (`global` or `local level scope`)

    - `:=`  &nbsp; is used to declare and initialize the variables `only inside` the functions.
        
        - Declaration and initialization at the same time.

---
## Looping

### Enhanced for loop

a := []string{"Foo", "Bar"}
for idx, str := range a {
    fmt.Println(idx, str)
}

```
0 Foo
1 Bar
```
- The range expression, a, is evaluated once before beginning the loop.
- The iteration values are assigned to the respective iteration variables, i and s, as in an assignment statement.
- The second iteration variable is optional.

### Iterating each character of string
const s = "日本語"
for i := 0; i < len(s); i++ {
    fmt.Printf("%x ", s[i])
}
```
e6 97 a5 e6 9c ac e8 aa 9e
```
