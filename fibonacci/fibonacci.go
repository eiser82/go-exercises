package fibonacci

import "fmt"

// Recursive way
// func Fibonacci(n uint) uint {

//   if n <= 1 {
//     return n
//   }

//   return Fibonacci(n-1) + Fibonacci(n-2)
// }

func Fibonacci(n int) uint64 {
	if n <= 1 {
		return uint64(n)
	}

	var n2, n1 uint64 = 0, 1

	for i := 2; i < n; i++ {
		n2, n1 = n1, n1+n2
	}

	return n2 + n1
}

func main() {

	var n int
	var sum uint64
	// j = 0

	fmt.Print("Enter the number of terms: ")
	fmt.Scanln(&n)

	for i := 0; i <= n+1; i++ {
		fmt.Println(Fibonacci(i))
		sum += Fibonacci(i)
	}

	fmt.Println("Sum:", sum)

}
