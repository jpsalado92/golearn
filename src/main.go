package main

func do() (a int) {
    defer func() {
        a = 2
    }()
    a = 1
    return a
}

func main() {
    a := do()
    println(a)
}