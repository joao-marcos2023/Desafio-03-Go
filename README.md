# Desafio Ping Pong com Go

## Desafio de projeto da formação Go Developer da DIO

### Tecnologias 💻

<img src = "https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" height="50">

### Observações 🔍

O desafio consiste em criar um programa que exiba simultaneamente as palavras "ping" e "pong" usando conhecimento em **Goroutines** e **Canais**.

### Solução ✔

Criei os canais **pi** e **po** de modo que as goroutines **ping** e **pong** se comunicassem através deles na qual cada um recebe a string "ping" e "pong" respectivamente.

```go


    func ping(pi chan string) { 
        for i := 0; ; i++ {
            pi <- "ping" //canal pi recebe string ping
        }
    }

    func pong(po chan string) {
        for i := 0; ; i++ {
            po <- "pong" //canal po recebe string pong
        }
    }

    func main() {
        var pi chan string = make(chan string) //canal pi
        var po chan string = make(chan string) //canal po

        go ping(pi) //goroutine ping
        go pong(po) //goroutine pong
        print(pi, po)

        var input string
        fmt.Scanln(&input)

    }

```

Posteriormente foi preciso um função para imprimi as string "ping" e "pong" simultaneamente num intervalo de tempo. Foi usado o pacote **time** do Go para tal.

```go

    func print(pi, po chan string) {
        for {
            msg1 := <-pi
            msg2 := <-po
            fmt.Println(msg1)
            fmt.Println(msg2)
            time.Sleep(time.Second * 1)
        }
    }

```





