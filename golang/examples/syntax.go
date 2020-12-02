// memes and syntax, enjoying myself as I learn the language :)

// has to start with pkg declaration
package main

// very python much wow
import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"sort"
	// "strconv"
	"strings"
)

// comments

/* multi-line
comments */

func main() {
	fmt.Println("very based")

	// I'm looking at you c, this is how you do declarations
	var age int = 999
	var pi float64 = 3.14
	rand_num := 666
	fmt.Println(age, pi, rand_num)

	// overflows be everywhere
	var num_one = 1.000
	var num_nien_nien = 0.999 // by design :)
	fmt.Println(num_one - num_nien_nien)

	//can do arithmetic with it as well....
	// ain't doing that here

	// can delare constants like this
	const pi_thicc float64 = 3.14159265

	// fancy variable naming
	// var (
	// varA = 2
	// varB = 3
	// )

	var wow_string string = "how based"
	fmt.Println(len(wow_string))

	// concatenate strings
	fmt.Println(wow_string + " of you")

	// booleans
	// var over9000 bool = true

	fmt.Printf("Var = %f\nType = %T\n", pi, pi)

	// loops
	//don't try to trick me sir, that's a while loop
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// now this right here is a for loop yes sir
	for j := 0; j < 5; j++ {
		fmt.Println(j)
	}

	yourAge := 20 // I like feeling young

	// if things
	if yourAge >= 16 {
		fmt.Println("Zoomer")
	} else if yourAge >= 18 && yourAge <= 30 {
		fmt.Println("Based")
	} else {
		fmt.Println("Boomer")
	}

	// better than 'if' things
	switch yourAge {
	case 16:
		fmt.Println("cool")
	case 22:
		fmt.Println("very cool")
	default:
		fmt.Println("uncool")
	}

	// arrays be weird

	var numbers [3]float64

	numbers[0] = 125
	numbers[1] = 666
	numbers[2] = 25 // yes, that's a spongebob reference

	fmt.Println(numbers[2])

	// another way of declaring

	more_nums := [5]float64{1, 2, 3, 4, 5}

	// iterating through it, how very python of you go
	for _, value := range more_nums {
		fmt.Println(value)
	}

	// did someone say slices!
	num_slice := []int{5, 4, 3, 2, 1} // I know, I'm very creative

	another_slice := num_slice[3:5]
	fmt.Println("another_slice[1] =", another_slice[1])

	// empty slices
	empty_slice := make([]int, 5, 10)
	copy(empty_slice, num_slice)
	fmt.Println(empty_slice)
	empty_slice = append(num_slice, -1, 3, 4)

	// maps, not geography ones
	items := make(map[string]string)

	items["Tablet"] = "iPad Pro"
	items["Laptop"] = "HP Envy 13"
	items["Keyboard"] = "Epomaker EP 84"

	fmt.Println(items["Tablet"])
	delete(items, "Laptop")

	listNums := []float64{1, 2, 3, 4, 5}
	fmt.Println(addThings(listNums))

	num1, num2 := returnTwoValues(10)
	fmt.Println(num1, num2)

	fmt.Println(manyInputs(1, 2, 3, 4, 5, 6, 7))

	// closures, OwO
	a_num := 3

	doubleNum := func() int {
		a_num *= 2

		return a_num
	}

	fmt.Println(doubleNum())
	fmt.Println(factorial(5))

	defer printTwo()
	printOne()

	fmt.Println(safeDiv(0, 0))
	demoPanic()

	x := 5
	changeValInMem(&x)
	fmt.Println("x =", x)
	fmt.Println("Memory address of x =", &x)

	yPtr := new(int)
	changeValInMem(yPtr)
	fmt.Println("y =", *yPtr)

	rect1 := Rectangle{width: 100, height: 100}
	circle := Circle{radius: 2}
	fmt.Println("Area of rect is", getArea(rect1))
	fmt.Println("Area of circle is", getArea(circle))

	// ymm python, in c, very cool
	samp_string := "Hello World"
	fmt.Println(strings.Contains(samp_string, "lo"))
	fmt.Println(strings.Index(samp_string, "lo"))
	fmt.Println(strings.Count(samp_string, "l"))
	fmt.Println(strings.Replace(samp_string, "l", "x", 3))

	// such string functions, much cool
	csvString := "1,2,3,4,5,6"
	fmt.Println(strings.Split(csvString, ","))
	list_of_chars := []string{"c", "a", "b"}
	sort.Strings(list_of_chars)
	fmt.Println("Letters: ", list_of_chars)

	// working with files
	file, err := os.Create("samp.txt")

	if err != nil {
		log.Fatal(err)
	}

	file.WriteString("Such python, much wow")
	file.Close()

	stream, err := ioutil.ReadFile("samp.txt")

	if err != nil {
		log.Fatal(err)
	}

	read_string := string(stream)
	fmt.Println(read_string)
	// lots of steps honestly, I'd much rather use something like `with open`

}

func addThings(numbers []float64) float64 {
	sum := 0.0

	for _, value := range numbers {
		sum += value
	}

	return sum

}

func returnTwoValues(number int) (int, int) {
	return number + 1, number + 2
}

// "variadic function", yeah, big brain vocabulary
func manyInputs(args ...int) int {

	sum := 0
	for _, value := range args {
		sum += value
	}

	return sum
}

// well, recursion, you know what that means, make a factorial func
// best advice I've ever gotten about them, *clears throat*
// "Stare at them for a little bit and you'll get them"
func factorial(num int) int {
	if num == 0 {
		return 1
	}

	return num * factorial(num-1)
}

// defering things to others, like a master procrastinator
// useful for cleanup functions
func printOne() { fmt.Println(1) }
func printTwo() { fmt.Println(2) }

func safeDiv(num1, num2 int) int {
	// you don't want your programming stopping do you, DO YOU?
	// well then do error handeling goddamnit
	defer func() {
		fmt.Println(recover())
	}()

	solution := num1 / num2
	return solution
}

func demoPanic() {
	defer func() {
		fmt.Println(recover())
	}()

	panic("I'M HAVING A SEIZURE")
}

// here we go, pointers, let me break my brain
func changeValInMem(x *int) {
	*x = 25
}

// interfaces, hmm, by brain is reaching the limit of making jokes
// if you can call them that
// if you do, thanks
type Shape interface {
	area() float64
}

// structs, we're getting to the good stuff now people
type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

// attach methods to structures like legos, not that I've ever owned one
func (rect Rectangle) area() float64 {
	return rect.height * rect.width
}

func (circle Circle) area() float64 {
	return math.Pi * math.Pow(circle.radius, 2)
}

// now this is pretty cool
func getArea(shape Shape) float64 {
	return shape.area()
}
