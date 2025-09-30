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
## Simplicity...

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
- go fmt: opinionated formatting built in. 
- go mod: dependency management built in (no npm/yarn/lockfile drama).
- go build: compiles to a single binary, no runtime needed.
<!-- pause -->
### Networking
- Full featured http routing, middleware and file serving baked in
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
## Simple
---

```go +line_numbers
	for _, f := range files {
		if f.Name() == ".DS_Store" {
			continue
		}
		resizeImg(f.Name())
	}
 
```
<!-- end_slide -->
<!-- alignment: center -->
Concurrency
---
<!-- alignment: left -->
## Simple
---

```go +line_numbers
	for _, f := range files {
		if f.Name() == ".DS_Store" {
			continue
		}
		resizeImg(f.Name())
	}
```
<!-- pause -->
<!-- column_layout: [1, 2] -->
<!-- column: 0 -->
#### Don't communicate by sharing memory, share memory by communicating.
