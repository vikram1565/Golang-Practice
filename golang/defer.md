# **Defer**

## *Explain 'defer'*

A defer statement defers the execution of a function until the surrounding function returns.

A defer statement pushes a function call onto a list. The list of saved calls is executed after the surrounding function returns. Defer is commonly used to simplify functions that perform various clean-up actions.

A deferred function's arguments are evaluated when the defer statement is evaluated.

Deferred function calls are executed in Last In First Out order after the surrounding function returns.

Deferred functions may read and assign to the returning function's named return values.

ex. In this example, a deferred function increments the return value i after the surrounding function returns. Thus, this function returns 2:

```go
func c() (i int) {
    defer func() { i++ }()
    return 1
}
```

## *What are your comments on Go compiler*

Answer to question

## *In how many ways we can import packages*

Answer to question
