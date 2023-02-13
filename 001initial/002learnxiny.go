/*

https://learnxinyminutes.com/docs/go/

*/

/* https://go.dev/doc/modules/gomod-ref
* go.mod defines modules and its properites */



// This is a single line comment
/* Multi-
line comment */

/* A build tag is a line comment strting with // +build and can be executed by
* go build -tags="foo bar" command. Build tags are placed before the package
* clause near or at the top of the file followed by a blank line or other line
* comments. */
// +build prod, dev, test


/* A package clase starts every source file.
* main is a specal name declaring an executable rather than a library */
package main

/* Import declared library packages */
import (
        "fmt"
)

func main() {
        fmt.Println("Hello, World!")
}

func beyondHello() {
        var  x int /* Declaration of a Variable */
        x = 3 /* assigning a variable */

        /* Shor declaration for declaring and assigning a variable */
        y := 4

        sum, prod := learnMultiple(x, y)                // Function returns 2 values & needs 2 inputs
        fmt.Println("sum: ", sum, "prod: ", prod)       // Simple output with vars
}

/* Functions can have parameters and multiple return values. */
func learnMultiple(x, y int) (sum, prod int) {
        return x + y, x * y /* Return two values */
}

/* Built-in types and literals. */
func learnTypes() {
        str := "Learn To Program In Go!"

        st := `A "raw" string literal can include line breaks.`
        /* Same String Type */

        /* Non-ASCII literal The Source UTF-8 */
        g := 'Σ' /* runt type, an alias for int32, holts unicode */

        f := 7.461262   /* float64, an 64-bit floating point */
        c := 3 + 4i     /* complex 128, represented internally with two float64's */

        /* Var syntax with initilizers */
        var u uint = 7
        var pi float32 = 22. / 7

        /* Conversion syntax with a short declaration */
        n := byte('\n') /* byte alias for uint8 */

        /* Arrays have size fixed compile times */
        var a4 [4]int   /* Array of 4ints, initialized to all 0 */
        a5 := [...]int{3, 1, 5, 10, 100}
        /* Array initialized fith a fixed size of five */

        /* Arrays have value semantics */
        a4_cpy := a4    /* this copy of a4, two separates instances */
        a4_cpy[0] = 25  /* Only a4_cpy is changed a4 stays the same */
        fmt.Println(a4_cpy[0] == a4[0])

        /* Slices have dynamic size 
        * Arrays and slices each have advantages but use cases for slices 
        * are much more common. */
        s3 := []int{4, 5, 9}   /* no ellipsis  */
        s4 := make([]int, 4)   /* slice 4 ints, init to all 0 */
        var d2 [][]float64     /* Declare only no allocated */
        bs := []byte("a slice")/* Type conversion */

        /* Slices have reference semantics */
        s3_cpy := s3   /* Both variables point to the same instance */
        s3_cpy[0] = 0  /*      both are then updated */
        fmt.Println(s3_cpy[0] == s3[0])        /* Prints true */

        /* Slices can be appended to on demand
        * use append() to append elements, first arg is a slice to append*/
        s := []int{1, 2, 3}   /* Result is a slice of len 3 */
        s = append(s, 4, 5, 6)        /* add 3 elements. slice now len 6 */
        fmt.Println(s)        /* Update slice */

        /* To append slices, we can pass a reference to a slice or a literal
        * slice, instead of using a list of atomic elements. */
        s = append(s, []int{7, 8, 9}...) /* Second arg is a literal slice. */
        fmt.Println(s) /* Updated slice is now [1 2 3 4 5 6 7 8 8] */

        p, q := learnMemory()  /* Declares p, q to be type pointer to int. */
        fmt.Println(*p, *q)    /* the * follows pointers, in this case 
        * it points to two intagers*/

        /* Mapes are a dynamicaly changeable array type, like a hashmap or
        * dictionary */
        m := map[string]int{"three": 3, "four": 4}
        m["one"] = 1

        /* Unused variables throw an error in go */
        /* To circumvent this you use an underscore to discard its value. */
        /* For Example: _ = 12 */
        /* Usually you use it to ignore one of the return values of a function,
        * For example, in a script you want to ignore the error value returned
        * from os.Crate, and expect tat the file will always be created. */
        file, _ := os.Create("output.txt")
        fmt.Fprint(file, "This is how you write to a file")
        file.Close()

        /* Output of course counts as using a variable. */
        fmt.Println(s, c, a4, s3, d2, m)

        //learnFlowControl() /* Back into the flow */
}

func learnMemory() (p, q *int) {
        /* Named return values p and q have type pointer to int. */
        p = new(int) /* Built-in function new allocates memory. */
        /* The allocated intager slice is initialize to 0, p is not nil. */
        s := make([]int, 20)    /* Aloocates 30 intagers in single block of memory */
        s[3] = 7                /* Assign one of them. */
        r := -2                 /* Assignes another local variable */
        return &s[3], &r        /* &: takes the address of another object */
}
