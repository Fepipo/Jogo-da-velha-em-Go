package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"time"
)

var jogo bool = true
var tentativas = 0

var tabuleiro = []interface{}{
	1, 2, 3,
	4, 5, 6,
	7, 8, 9,
}

func clear() {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("clear")
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Erro ao limpar o terminal:", err)
	}
}

func msg_jogo() {
	fmt.Println("----------------•-----------------")
	fmt.Println("\t  Jogo da velha")
	fmt.Println("----------------•-----------------")
}

func printar_tabuleiro() {
	clear()
	msg_jogo()
	fmt.Println("\n", tabuleiro[0], "│", tabuleiro[1], "│", tabuleiro[2])
	fmt.Println("───┼───┼───")
	fmt.Println("", tabuleiro[3], "│", tabuleiro[4], "│", tabuleiro[5])
	fmt.Println("───┼───┼───")
	fmt.Println("", tabuleiro[6], "│", tabuleiro[7], "│", tabuleiro[8])
}

func verificar_posicao() int {
	var posicao int
	for {
		fmt.Println("Digite o numero da posição desejada:")
		fmt.Scanln(&posicao)

		if posicao < 1 || posicao > 9 {
			fmt.Println("\nValor invalido, coloque um numero entre 1 a 9!")
			time.Sleep(3 * time.Second)
			printar_tabuleiro()
			continue
		} else if tabuleiro[posicao-1] == "X" || tabuleiro[posicao-1] == "O" {
			fmt.Println("\nEsta posição ja esta ocupada, escolha outra!")
			time.Sleep(3 * time.Second)
			printar_tabuleiro()
			continue
		} else {
			return posicao
		}
	}
}

func verficar_vitoria() {
	//verifica na vertical
	for i := 0; i < 3; i++ {
		if tabuleiro[i] == tabuleiro[i+3] && tabuleiro[i+3] == tabuleiro[i+6] {
			printar_tabuleiro()
			fmt.Printf("\nPartida finalizada, jogador %s ganhou!", tabuleiro[i])
			jogo = false
			os.Exit(0)
		}
	}

	//verifica na horinzontal
	for j := 0; j < 9; j+=3 {
		if tabuleiro[j] == tabuleiro[j+1] && tabuleiro[j+1] == tabuleiro[j+2] {
			printar_tabuleiro()
			fmt.Printf("\nPartida finalizada, jogador %s ganhou!", tabuleiro[j])
			jogo = false
			os.Exit(0)
		}
	}

	//vefiica na diagonal
	if tabuleiro[0] == tabuleiro[4] && tabuleiro[4] == tabuleiro[8] || tabuleiro[2] == tabuleiro[4] && tabuleiro[4] == tabuleiro[6] {
		printar_tabuleiro()
		fmt.Printf("\nPartida finaliza, jogador %s ganhou!", tabuleiro[4])
		jogo = false
		os.Exit(0)
	}
}

func amigo() {
	printar_tabuleiro()

	for jogo == true {
		fmt.Println("\nVez do jogador X")
		tabuleiro[verificar_posicao() - 1] = "X"
		tentativas++
		verficar_vitoria()
		printar_tabuleiro()

		if !jogo {
			break
		}

		if tentativas == 9 {
			fmt.Println("\nPartida finalizada, empate!")
			os.Exit(0)
		}

		fmt.Println("\nVez do jogador O")
		tabuleiro[verificar_posicao() - 1] = "O"
		tentativas++
		verficar_vitoria()
		printar_tabuleiro()
	}
}

func escolher_posicao_comp() int {
	var posicao_comp int

	for {
		posicao_comp = rand.Intn(9)

		if tabuleiro[posicao_comp] == "X" || tabuleiro[posicao_comp] == "O" {
			continue
		} else {
			return posicao_comp
		}
	}
}

func computador() {
	printar_tabuleiro()

	for jogo == true {
		fmt.Println("\nSua vez")
		tabuleiro[verificar_posicao() - 1] = "X"
		tentativas++
		verficar_vitoria()
		printar_tabuleiro()
		
		if !jogo {
			break
		}

		if tentativas == 9 {
			fmt.Println("\nPartida finalizada, empate!")
			os.Exit(0)
		}
		
		fmt.Println("\nVez do computador")
		fmt.Println("Escolhendo...")
		time.Sleep(1 * time.Second)

		var_comp := escolher_posicao_comp()
		tabuleiro[var_comp] = "O"
		tentativas++
		
		printar_tabuleiro()
		fmt.Println("\nO computador escolheu a posição", var_comp + 1)
		time.Sleep(2 * time.Second)
		verficar_vitoria()
		printar_tabuleiro()
	}
}

func main() {
	var modo int

	msg_jogo()

	fmt.Println("\nEscolha o modo\n1 - Modo amigo\n2 - Modo computador")
	
	fmt.Println("\nDigite o numero do modo desejado:" )
	fmt.Scanln(&modo)
	if modo < 1 || modo > 2 {
		fmt.Println("\nValor invalido, coloque o numero 1 ou 2")
		time.Sleep(3 * time.Second)
		clear()
		main()
	}

	switch modo {
	case 1:
		amigo()
	case 2:
		computador()
	}
}