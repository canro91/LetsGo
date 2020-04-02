package main

import "fmt"

const spanish = "Spanish"
const french = "French"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
    if name == "" {
        name = "world"
    }

    prefix := englishHelloPrefix
    if language == spanish {
        prefix = spanishHelloPrefix
    } else if language == french {
        prefix = frenchHelloPrefix
    }

    return prefix + name + "!"
}

func main(){
    fmt.Println(Hello("world", ""))
}
