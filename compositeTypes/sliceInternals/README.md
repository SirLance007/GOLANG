// copy() function ka use
Bilkul! `copy()` tab use hota hai jab tum **underlying array share nahi karna chahte**. Chalo step-by-step dekhte hain.

---

# Case 1: `copy()` ke bina

```go
package main

import "fmt"

func main() {
    original := []int{10, 20, 30, 40}

    sub := original[1:3] // [20 30]

    sub[0] = 99

    fmt.Println("Original:", original)
    fmt.Println("Sub:", sub)
}
```

### Memory

```
original
Index:   0    1    2    3
        +----+----+----+----+
        |10  |20  |30  |40  |
        +----+----+----+----+
              ^
              |
sub starts here
```

`sub` sirf isi array ko point kar raha hai.

Jab likhte ho

```go
sub[0] = 99
```

to actually

```
original[1] = 99
```

ho raha hai.

Memory ban jati hai

```
original
Index:   0    1    2    3
        +----+----+----+----+
        |10  |99  |30  |40  |
        +----+----+----+----+

sub
       +----+----+
       |99  |30  |
       +----+----+
```

Output

```
Original: [10 99 30 40]
Sub: [99 30]
```

Yaha dono same data use kar rahe hain.

---

# Case 2: `copy()` ke saath

```go
package main

import "fmt"

func main() {
    original := []int{10, 20, 30, 40}

    sub := original[1:3]

    copied := make([]int, len(sub))

    copy(copied, sub)

    copied[0] = 99

    fmt.Println("Original:", original)
    fmt.Println("Sub:", sub)
    fmt.Println("Copied:", copied)
}
```

Yaha

```go
copied := make([]int, len(sub))
```

ek **naya array** banata hai.

Memory

```
Original Array

Index:   0    1    2    3
        +----+----+----+----+
        |10  |20  |30  |40  |
        +----+----+----+----+

sub
         |
         v
        +----+----+
        |20  |30  |
        +----+----+


Copied Array (alag array)

        +----+----+
        |20  |30  |
        +----+----+
```

Ab

```go
copy(copied, sub)
```

sirf values copy karta hai.

Dono arrays alag ho gaye.

Ab jab

```go
copied[0] = 99
```

karte ho to

sirf copied array change hota hai.

Memory

```
Original Array

+----+----+----+----+
|10  |20  |30  |40  |
+----+----+----+----+

Copied Array

+----+----+
|99  |30  |
+----+----+
```

Output

```
Original: [10 20 30 40]
Sub: [20 30]
Copied: [99 30]
```

Original bilkul change nahi hua.

---

# `make()` ki zarurat kyun?

Agar tum seedha likh do

```go
var copied []int

copy(copied, sub)
```

to output hoga

```
[]
```

kyunki `copied` ki **length 0** hai.

`copy()` sirf utne hi elements copy karta hai jitni destination slice ki length hoti hai.

Isliye pehle destination banana padta hai:

```go
copied := make([]int, len(sub))
```

Ab `copied` ke paas 2 elements ki jagah hai.

```
Before copy()

copied

+----+----+
| 0  | 0  |
+----+----+

After copy()

+----+----+
|20  |30  |
+----+----+
```

---

## Ek line mein difference

**Without `copy()`**

```
original ----+
             |
             +----> same array <---- sub

Change in sub  ==> Original bhi change
```

**With `copy()`**

```
original ---> Array A

copied  ---> Array B

Change in copied ==> Original par koi effect nahi
```

**Yaad rakhne ka shortcut:**

* `sub := original[1:3]` → **Reference (same underlying array)**
* `copy(newSlice, sub)` → **Duplicate data (new underlying array)**
