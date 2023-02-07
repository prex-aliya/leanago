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
