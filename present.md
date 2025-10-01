---
theme: 
  name: "catppuccin-mocha"
---
Go for JS Developers
---
<!-- pause -->
```python +no_background
██████╗ ██╗   ██╗████████╗          ██╗    ██╗██╗  ██╗██╗   ██╗
██╔══██╗██║   ██║╚══██╔══╝          ██║    ██║██║  ██║╚██╗ ██╔╝
██████╔╝██║   ██║   ██║             ██║ █╗ ██║███████║ ╚████╔╝ 
██╔══██╗██║   ██║   ██║             ██║███╗██║██╔══██║  ╚██╔╝  
██████╔╝╚██████╔╝   ██║██╗██╗██╗    ╚███╔███╔╝██║  ██║   ██║   
╚═════╝  ╚═════╝    ╚═╝╚═╝╚═╝╚═╝     ╚══╝╚══╝ ╚═╝  ╚═╝   ╚═╝   
                                                               
```
<!-- speaker_note: Obvously not for front end -->
<!-- pause -->
```markdown +no_background {all|1|3|5|7|all}
1. Better concurrency support

2. Faster single thread performance

3. Lower memory overheads

4. Lower running costs
```
<!-- speaker_note: These are reasons for many languages... -->
<!-- alignment: center -->
<!-- end_slide -->
## Simplicity...

--- 

You have a few options...
<!-- pause -->
<!-- alignment: center -->
<!-- column_layout: [1, 1] -->
<!-- column: 0 -->
Java
```java +line_numbers
public class HelloWorld {
    public static void main(String[] args) {
        System.out.println("Hello, World!");
    }
  }
```
<!-- pause -->
C#
```java +line_numbers
using System;

public class Program {
    public static void Main(string[] args) {
        Console.WriteLine("Hello, World!");
    }
}
```
<!-- pause -->
<!-- column: 1 -->
C++
```c +line_numbers
#include <iostream> 

int main() {
    std::cout << "Hello World!" << std::endl; 
    return 0;
}
```
<!-- pause -->
Rust
```rust +line_numbers
fn main() {
    println!("Hello, world!");
}
```
^ Borrow checker & async complexity overhead
<!-- pause  -->
<!-- alignment: center -->
<!-- end_slide -->
## Clear is better than Clever

---
![image:w:55%](images/meme2.jpg)
<!-- end_slide -->
Syntax
---
---
<!-- column_layout: [1, 1] -->
<!-- alignment: center -->
<!-- column: 0 -->
TS
```typescript +line_numbers
export function sayHello(name: string) {
  console.log("Hello", name);
}
```
<!-- column: 1 -->
GO
```go +line_numbers
func SayHello(name string) {
  fmt.Println("Hello", name)
}
```
<!-- pause -->
<!-- column: 0 -->
```typescript +line_numbers
export interface foo {
  numA: number;
  numB: number;
}

export function sumFoos(foos: foo[]) {
  let sum = 0

  for (let foo of foos) {
    sum += foo.numA
    sum += foo.numB
  }
  return sum
}
```
<!-- column: 1 -->
```go +line_numbers
package hello

type Foo struct {
	numA int
	numB int
}

func SumFoos(foos []Foo) int {
	sum := 0

	for _, foo := range foos {
		sum += foo.numA
		sum += foo.numB
	}
	return sum
}
```
<!-- end_slide -->
<!-- alignment: center -->
Standard Library and Tooling
---
<!-- alignment: left -->
## The Full Package
---
<!-- speaker_note: Go has a philosophy of no dependencies -->
<!-- column_layout: [1, 1] -->
<!-- column: 1 -->
![image:w:100%](images/meme4.png)
<!-- column: 0 -->
### Tooling
- go build: compiles to a single binary, no runtime needed.
- go's GC: Has garbage collection
- go fmt: highly opinionated (unconfigurable) formatting built in. 
- go mod: dependency management 
<!-- pause -->
### Networking
- Full featured http routing, middleware and file serving 
- TCP, UDP, SMTP, etc
<!-- pause -->
### Testing, Profiling & Benchmarking
- Comes with a full testing suite and coverage tools
- Performance and memory profiling with Pprof
<!-- pause -->
### Type Safety
- Simple and powerful type system
- Build right into the compiler and LSP 
<!-- pause -->
### Plus more...
- Powerful string/html templating
- Full cryptography suite
<!-- end_slide -->
<!-- alignment: center -->
Concurrency
---
<!-- alignment: left -->
---

<!-- column_layout: [1, 99] -->
<!-- column: 0 -->
#### Sequential
<!-- column: 1 -->
```go +line_numbers {all|1|6-8|all}
	files, err := os.ReadDir(inDir)
	if err != nil {
		panic(err)
	}

	for _, f := range files {
		resizeImg(f)
	}
```
<!-- pause -->
<!-- column: 0 -->
<!-- new_lines: 9 -->
#### Parallel.. However...

<!-- column: 1 -->
```go +line_numbers {all|2|all}
	for _, f := range files {
		go resizeImg(f)
	}
```
<!-- pause -->
<!-- column: 0 -->
<!-- new_lines: 5 -->
#### Parallel
<!-- column: 1 -->
```go +line_numbers {all|1|4|5|7|all}
	wg := sync.WaitGroup{}

	for _, f := range files {
		wg.Add(1)
		go resizeImg(f, &wg)
	}
	wg.Wait()
```
<!-- end_slide -->
<!-- alignment: center -->
Concurrency
---
<!-- alignment: left -->
---
##### Using A Mutex
```go +line_numbers {all|1-4|9-11|all}
type results struct {
	sizes []int64
	m     sync.Mutex
}

func resizeImgWithMutex(path os.DirEntry, wg *sync.WaitGroup, result *results) {
  // Process the image...

	result.m.Lock()
	result.sizes = append(result.sizes, outImgFileInfo.Size())
	result.m.Unlock()
}
```
<!-- end_slide -->
<!-- alignment: center -->
Concurrency
---
<!-- alignment: left -->
---
##### Don't communicate by sharing memory, share memory by communicating.
<!-- column_layout: [1,1] -->
<!-- column: 0 -->
```go +line_numbers {all|1|5|8-9|all}
	ch := make(chan int64)
    res := make([]int64, 0, len(files))

	for _, f := range files {
		go resizeImgWithChannel(f, ch)
	}

	for range len(files) {
		res = append(res, <-ch)
	}
```
<!-- column: 1 -->
```go +line_numbers {all|4|all}
func resizeImgWithChannel(path os.DirEntry, result chan int64) {
  // Process the image...

  result <- newImg.Size()
}
```
<!-- end_slide -->
<!-- alignment: center -->
Error Handling
---
<!-- alignment: left -->
## Errors Are Values
---
![](images/meme1.png)
```go +line_numbers
res, err := SomeDangerousFunction()
if err != nil {
  //... Handle the error
  log.Fatalf("failed to run dangerous: %v", err)
}
```
<!-- end_slide -->
<!-- alignment: center -->
Error Handling
---
<!-- alignment: left -->
## Errors Are Values
---
```go 
// The error built-in interface type is the conventional interface for
// representing an error condition, with the nil value representing no error.
type error interface {
	Error() string
}
```
<!-- pause -->
```go +line_numbers {all|5|all}
type CustomError struct {
  message string
}

func (e *CustomError) Error() string {
	return fmt.Println("this is a custom error: ", e.message)
}
```
<!-- end_slide -->
<!-- alignment: center -->
Error Handling
---
```go +line_numbers {all|6|all}
type HTTPError struct {
	Code int
	Msg  string
}

func (e *HTTPError) Error() string {
	return fmt.Println("http: ", e.Code, " message: " e.Msg)
}
```
<!-- pause -->
```go +line_numbers {all|1-2|3-4|5|all}
profile, err := GetProfile("badID")
if err != nil {
  var httpErr *HTTPError
  if errors.As(err, &httpErr) {
    log.Printf("got HTTP error: %d\n", httpErr.Code)
  } else {
    return err
  }
}
```
