package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var input = bufio.NewScanner(os.Stdin)

var variaveis = make(map[string]any)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("NEVER 0.3.0 - Henrique Becker Mendes")
		fmt.Println()
		a := bufio.NewScanner(os.Stdin)

		for {
			fmt.Print("> ")
			a.Scan()
			oi(a.Text())
			fmt.Println()
		}
	}

	if len(os.Args) == 2 {
		arquivo, _ := os.Open(os.Args[1])
		nota := bufio.NewScanner(arquivo)

		defer arquivo.Close()

		for nota.Scan() {
			oi(nota.Text())
		}
	}

}

func oi(sei_la string) {

	if strings.HasPrefix(sei_la, "print ") {
		print := sei_la[6:]

		if strings.HasPrefix(print, `"`) && strings.HasSuffix(print, `"`) {
			fmt.Println(print[1 : len(print)-1])
		} else if strings.HasPrefix(print, "n: ") {
			n := print[3:]
			valor_string_int, _ := strconv.Atoi(n)
			fmt.Println(valor_string_int)
		} else if strings.HasPrefix(print, "v: ") {
			v := print[3:]
			fmt.Println(variaveis[v])
		}
	}

	if strings.HasPrefix(sei_la, "var ") {
		var_never := strings.SplitN(sei_la[4:], " = ", 2)

		if strings.HasPrefix(var_never[1], `"`) && strings.HasSuffix(var_never[1], `"`) {
			variaveis[var_never[0]] = var_never[1][1 : len(var_never[1])-1]

		} else if strings.HasPrefix(var_never[1], "n: ") {
			n := var_never[1][3:]
			iint, _ := strconv.Atoi(n)
			variaveis[var_never[0]] = iint

		} else if strings.HasPrefix(var_never[1], "input") {
			input.Scan()
			variaveis[var_never[0]] = input.Text()
		} else if strings.HasPrefix(var_never[1], "v: ") {
			va := var_never[1][3:]
			variaveis[var_never[0]] = variaveis[va]
		}
	}
}
