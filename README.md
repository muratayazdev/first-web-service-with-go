# GO

The aim of the document is taking notes while learning Go for to have documentation for looking up when I need

[TOC]

## Problems

![image-20200327182819636](/Users/muratayaz/Library/Application Support/typora-user-images/image-20200327182819636.png)

Go  =>

* fast compilation time
* Fully compiled
* Strongly typed
* Concurrent by default
* Garbage collected
* Simplicity as a core value

###What is go good at?

**Web service + Web application** + Task automation + GUI/Thick Client + Machine learning

## Introduction to Go

**package** => where it sits in your application

**func main** => starting point of application.

**//** => One line comment

/* */ => multiple line comment

## Primitive Types

###var

```go
var i int
i =42 // i => 42

var  f float32 = 3.14 // f => 3,14
firstName := "asdf" //(implicit initializing) firstname => "asdf"
b := true // b => true
c := complex(3, 4) (like scala tuple2) // (3 + 4i)
r, im := real(c), imag(c) // r => 3, im => 4  real part and imaginary component
```

### Pointer

```go
var firstName *string //Nil
firstName = "test" // does not compile
*firstName = "test" // invalid memory address or 

var firstName *string = new String
firstName = "test" 
firstName // 0x40c128(some memory address)
*firstName// test(value)

firstName := "Test" // firstName => Test
ptr := &firstName
ptr => the address of the firstName
*ptr => the value of the ptr (in this case Test)

firstName := "Test" // firstName = test
ptr := &firstName // 0x40c128
firstName = "SomethingElse"// ptr => 0x40c128 *ptr => "SomethingElse" (address same value updated)
```

### Constant

```go
const pi = 3.1415 // pi => 3.1415

pi = 1.2 //can not assign to pi
```

### Collections

#### Arrays

```go
var arr [3]int 
arr[0] = 1
arr[1] = 2
arr[2] = 3 
arr //[1, 2, 3]

arr := [3]int{1,2,4}  // arr => [1,2,3]
```

#### Slices

```go
arr := [3]int{1,2,4}
slice := arr[:] // slice => [1,2,4]
arr[1] = 42
slice[2]=27
arr // arr => 1,42, 27
slice // slice => 1,42, 27 (Slice is not pointer but it works as kinda pointer. we did not change the second element but the change in the array effected it and changing the slice third element efected the array third element)

slice :=[]int{1,2,3} // slice => [1,2,3]

slice = append(slice, 4) // slice => [1,2,3,4]

slice = append(slice, 5, 6) // slice => [1,2,3,4,5,6]

slice :=[]int{1,2,3, 4, 47, 27} // slice => [1,2,3]
s2 := slice[1:] // create new slice start from 1 slice instance
s3 := slice[:2] // create new slice end with index 2 of slice instance
s4 := slice[1:2] // start from 1 and end with 2

s2 // s2 => [2 3 4 42 27]
s3 // s3 =>[1 2]
s4 // s4 =>[2]
```

#### Maps

```go
m := map[string]int{"foo":42} // m => map[foo:42]

m["foo"] // 42

m["foo"] = 27

m["foo"] // 27

delete(m, "foo")

m // m =>map[]
```



#### Structs

```go
type user struct {
  ID int
	FirstName string
	LasttName string
}

var u user // u => {0  }
u.ID = 1
u.FirstName = "FirstName"
u.LastName = "lastName"

u // u => {1, FirstName, LastName}

u.FirstName // FirstName

u2 := user{ ID: 1, FirstName: "FirstName", LastName: "LastName"} //{ 1, FirstName, LastName} 
```



## Functions

* `func functionName(parameterName parameterType, ...) returnType { code block}`

* A complex function definition and usage

  ```go
  func functionName(parameter1 int, parameter2 int) (int, error) {
    //do something
    return result, nil
  }
  
  parameter1 := 1
  result, err := startWebServer(parameter1, 2)//caling a function
  
  ```

## Controlling Program Flow

### Looping

#### Loop till condition

```go
var i int
for i < 5 {
  fmt.Printf("%d ", i)
  i++
	if i == 3 {
  	//do something
    continue
	}
  if i == 4 {
    break
  } 
  fmt.Printf("continuing ")
} 
//0 continuing 1 continuing 2 3
```



#### Loop till condition with post clause

```go
for i := 0; i < 5; i++ {
  fmt.Printf("%d, ",i)
}
//1, 2, 3, 4
```

#### Infinite Loops

```Go
var i int
for {
  if i == 5 {
    break
  }
}
```



#### Loop over collections

```go
slice := []int{1, 2, 3}
for i:= 0; i<len(slice); i++ {
  fmt.Printf("%d, "slice[i])
}
// 1, 2, 3

for i,v := range slice {
  fmt.Printf("%d => %d, ", i, v)
}
//0 => 1, 1 => 2, 2 => 3

m := map[string]int{"key1":1, "key2": 2, "key3":3}
for i,v := range m {
  fmt.Printf("%s => %d, ", i, v)
}
//key1 => 1, key2 => 2, key3 => 3,

for i := range m { // if we need just keys not values
  fmt.Printf("%s, ", i)
}
// key1, key2, key3

for _, v := range m { //if we need just values not keys
  fmt.Printf("%d, ", v)
}
//1, 2, 3

```

### Branching

###Panics

```go
//I need to learn it deeply.
panic ("Something happened bad!")
```

###If Statements

```go
if i == 3 {
  //do something
} else if i == 4 {
  //do something else
} else {
  //do something else
}
```

### Swithces

```go
i := 3
switch i {
  case 1:
  	//do something
  case 2:
  	//do something
  default: // => not mandotary
  	//do someting else
} 
```







## Notes

If a package is imported and not being used, It will cause compiling error,


