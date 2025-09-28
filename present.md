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
<!-- pause -->
```markdown +no_background {all|1}
1. Simplicity Over Abstraction

2. Standard Library

3. Testing & Profiling

4. Error Handling

5. Concurrency
```
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

public class Program
{
    public static void Main(string[] args)
    {
        Console.WriteLine("Hello, World!");
    }
}
```
<!-- pause -->
<!-- column: 1 -->
C++
```c +line_numbers
#include <iostream> 

int main() 
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
^ Borrow checker complexity
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
Simplicity
---
<!-- alignment: left -->
## Some more stuff about simplicity...

---
<!-- alignment: center -->

<!-- end_slide -->

<!-- alignment: center -->
Standard Library and Tooling
---
<!-- alignment: left -->
## The Full Package
---
![image:w:55%](images/meme3.gif)
