Koi baat nahi! Der nahi karte, sidhe chalte hain **Topic 6: `sync.WaitGroup**` par.

---

## 6. `sync.WaitGroup`

### 1. Simple Explanation & Analogy

`sync.WaitGroup` ek synchronization primitive hai jiska use multiple Goroutines ke complete hone ka wait karne ke liye kiya jata hai. Ye ek counter ki tarah kaam karta hai.

* **Real-life Analogy (School Tour Teacher):**
* Socho ek school teacher apne 5 students ko museum tour par le gayi hai.
* Har student ek Goroutine hai.
* Teacher ke paas ek tally counter (WaitGroup) hai.
* **`Add(5)`**: Teacher counter ko `5` par set kar deti hai (5 students exploration par gaye).
* **`Done()`**: Jab ek student wapas bus me aata hai, wo attendance marker click karta hai (counter `-1` ho jata hai).
* **`Wait()`**: Teacher bus ke door par khadi hokar tab tak wait karti hai jab tak counter `0` na ho jaye. `0` hote hi bus aage badhti hai.



---

### 2. Code Example

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	// Decrement the counter when the goroutine completes
	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		// Increment the counter BEFORE launching the goroutine
		wg.Add(1)
		go worker(i, &wg)
	}

	// Block until the counter reaches 0
	wg.Wait()
	fmt.Println("All workers finished!")
}

```

**Output:**

```text
Worker 1 starting
Worker 3 starting
Worker 2 starting
Worker 1 done
Worker 2 done
Worker 3 done
All workers finished!

```

---

### 3. Common Interview Questions

#### Q1: `sync.WaitGroup` ko function me pass karte waqt by Value pass karna chahiye ya by Pointer?

**Answer:**
Hamesha **by Pointer (`*sync.WaitGroup`)** pass karna chahiye!
Go me functions ke parameters by value (copy hokar) pass hote hain. Agar tum `WaitGroup` ko by value pass karoge, toh function ke andar uski ek **naya copy** ban jayegi. Inside function call hone wala `Done()` local copy ko decrement karega, jabki main function me original `WaitGroup` hamesha `Wait()` par block reh kar **Deadlock** kar dega.

#### Q2: `wg.Add()` ko Goroutine ke andar call karna sahi hai ya bahar?

**Answer:**
Hamesha Goroutine ke **BAHAR (parent Goroutine me)** call karna chahiye, `go` keyword execute karne se thik pehle.

* **Kyun?** Agar tum `wg.Add()` ko Goroutine function body ke inside likhte ho, toh ho sakta hai parent thread ka `wg.Wait()` pehle execute ho jaye isse pehle ki child Goroutine start hoke `wg.Add()` run kar paaye. Counter `0` milega aur `Wait()` bypass ho jayega!

---

### 4. Gotcha / Common Mistake

#### Reusing WaitGroup before previous Wait() completes / Negative Counter Panic

1. **Negative Counter:** Agar `wg.Done()` gallery call ho kar counter ko `0` se niche (`-1`) le aaye, toh Go runtime immediately **Panic** throw karega: `panic: sync: negative WaitGroup counter`.
2. **Reuse Race Condition:** `WaitGroup` ko reuse kiya ja sakta hai, lekin ek batch ka `Wait()` completely finish hone se pehle dobara `Add()` call karna data race create karta hai.

---

### 5. Quick Quiz (Active Recall)

Ab in 2 questions ka jawab do:

1. **Bug Hunter:** Is code snippet me kya issue/bug hai?
```go
func process(wg sync.WaitGroup) { // Line A
    defer wg.Done()
    fmt.Println("Processing")
}

func main() {
    var wg sync.WaitGroup
    wg.Add(1)
    go process(wg)
    wg.Wait()
}

```


2. **Panic Analysis:** Agar maine `wg.Add(2)` kiya lekin sirf 1 Goroutine run hui aur `wg.Done()` call hua, toh main Goroutine ka `wg.Wait()` kya karega?
*(a) Instantly complete ho jayega (b) Panic karega (c) Permanently block (deadlock) ho jayega)*

*Answers likho, fir hum Next topic (**`sync.Mutex / sync.RWMutex`**) par chalenge!*