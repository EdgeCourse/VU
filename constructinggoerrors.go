/*
Error handling in Go is a little different than other mainstream programming languages like Java, JavaScript, or Python. Go's built-in errors don't contain stack traces, nor do they support conventional try/catch methods to handle them. Instead, errors in Go are just values returned by functions, and they can be treated in much the same way as any other datatype - leading to a surprisingly lightweight and simple design.
*/


//The Error Type

//The error type in Go is implemented as the following interface:

/*
type error interface {
    Error() string
}
*/

//an error is anything that implements the Error() method, which returns an error message as a string. 

//Constructing Errors

/*
Errors can be constructed on the fly using Go's built-in errors or fmt packages. For example, the following function uses the errors package to return a new error with a static error message:
*/

/*
import "errors"
func DoSomething() error {
    return errors.New("something didn't work")
}
*/

//Similarly, the fmt package can be used to add dynamic data to the error, such as an int, string, or another error. For example:



import "fmt"

func Divide(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("can't divide '%d' by zero", a)
    }
    return a / b, nil
}

//Note that fmt.Errorf will prove extremely useful when used to wrap another error with the %w format verb 


/*
Errors can be returned as nil, and in fact, it's the default, or "zero", value of on error in Go. This is important since checking if err != nil is the idiomatic way to determine if an error was encountered (replacing the try/catch statements you may be familiar with in other programming languages).
Errors are typically returned as the last argument in a function. Hence in our example above, we return an int and an error, in that order.
When we do return an error, the other arguments returned by the function are typically returned as their default "zero" value. A user of a function may expect that if a non-nil error is returned, then the other arguments returned are not relevant.
Lastly, error messages are usually written in lower-case and don't end in punctuation. Exceptions can be made though, for example when including a proper noun, a function name that begins with a capital letter, etc.
*/

/*
Defining Expected Errors
Another important technique in Go is defining expected Errors so they can be checked for explicitly in other parts of the code. This becomes useful when you need to execute a different branch of code if a certain kind of error is encountered.
*/

/*
Defining Sentinel Errors
Building on the Divide function from earlier, we can improve the error signaling by pre-defining a "Sentinel" error. Calling functions can explicitly check for this error using errors.Is:
*/


import (
    "errors"
    "fmt"
)

var ErrDivideByZero = errors.New("divide by zero")

func Divide(a, b int) (int, error) {
    if b == 0 {
        return 0, ErrDivideByZero
    }
    return a / b, nil
}

func main() {
    a, b := 10, 0
    result, err := Divide(a, b)
    if err != nil {
        switch {
        case errors.Is(err, ErrDivideByZero):
            fmt.Println("divide by zero error")
        default:
            fmt.Printf("unexpected division error: %s\n", err)
        }
        return
    }

    fmt.Printf("%d / %d = %d\n", a, b, result)
}

