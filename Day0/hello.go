package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const es = "es"
const fr = "fr"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
    if name == "" {
        name = "world"
    }

    return helloPrefix(language) + name + "!"
}

func helloPrefix(language string) (prefix string) {
    switch language {
    case spanish, es:
            prefix = spanishHelloPrefix

    case french, fr:
        prefix = frenchHelloPrefix

    default:
        prefix = englishHelloPrefix
    }
    return
}

func main(){
    fmt.Println(Hello("world", ""))
}
