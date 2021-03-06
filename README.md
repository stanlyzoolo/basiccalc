### Basic calculator

Basic calculator provides a primitive implementation of a calculator for evaluating simple expressions consisting of single digits and mathematical addition and subtraction operators.

You can give input string representing an expression to `Eval()` function and
receive expected result.

### Installation

Use the command below to install module:

```golang
go get -u github.com/stanlyzoolo/basiccalc
```

### Usage

Given a string `s` representing an expression, implement a basic calculator to evaluate it.  
Expression consists of single digits (like `'2'`, `'7'` etc.), `'+'`, `'-'` and `' '` (space).

Example 1:

```golang
    // Input
    s := "1+1"

    fmt.Println(Eval(s))
    // Output: 2`
```

Example 2 (more complex expression with spaces `' '`):

```golang
    // Input
    s := "2-1 + 2"

    fmt.Println(Eval(s))
    // Output: 3
```

### Contributors

Thanks to [doublefint](https://github.com/doublefint)!


