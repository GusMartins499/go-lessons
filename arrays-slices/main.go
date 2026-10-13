// ARRAYS E SLICES
//
// ARRAY: tamanho FIXO, que faz parte do tipo. [3]int e [4]int são tipos diferentes.
//
//	É um valor: atribuir/passar para função COPIA o array inteiro.
//	Pouco usado diretamente no dia a dia.
//
// SLICE: "janela" dinâmica sobre um array. Tipo: []int (sem tamanho).
//
//	É o que se usa 99% das vezes (equivalente ao array do JS).
//	- len(s) -> quantidade de elementos
//	- cap(s) -> capacidade do array por baixo
//	- append(s, v) -> adiciona e RETORNA o novo slice (sempre reatribuir!)
//	- s[i:j] -> fatia do índice i até j-1
//	- make([]T, len, cap) -> cria slice com tamanho/capacidade
//
// Rodar: go run ./09-arrays-e-slices
package main

import (
	"fmt"
	"slices"
)

func main() {
	// ---------- ARRAYS ----------
	var arr [3]int // [0 0 0]
	arr[0] = 10
	nums := [5]int{1, 2, 3, 4, 5}
	auto := [...]string{"a", "b", "c"} // tamanho inferido: [3]string
	fmt.Println("arrays:", arr, nums, auto, "len:", len(nums))

	copia := nums // cópia completa
	copia[0] = 99
	fmt.Println("original:", nums, "| cópia:", copia)

	// ---------- SLICES ----------
	frutas := []string{"maçã", "banana"}
	frutas = append(frutas, "uva", "manga")
	fmt.Println("slice:", frutas, "len:", len(frutas), "cap:", cap(frutas))

	// Fatiamento [início:fim) — o fim é exclusivo
	fmt.Println("frutas[1:3]:", frutas[1:3])
	fmt.Println("frutas[:2]:", frutas[:2])
	fmt.Println("frutas[2:]:", frutas[2:])

	// Slice a partir de array
	parte := nums[1:4]
	fmt.Println("parte do array:", parte)

	// make: pré-aloca (bom para performance quando se sabe o tamanho)
	s := make([]int, 0, 10)
	for i := range 5 {
		s = append(s, i*i)
	}
	fmt.Println("make + append:", s, "len:", len(s), "cap:", cap(s))

	// Slice nil vs vazio
	var nulo []int
	fmt.Println("slice nil:", nulo == nil, "len:", len(nulo))

	// Juntar slices
	a := []int{1, 2}
	b := []int{3, 4}
	a = append(a, b...)
	fmt.Println("juntos:", a)

	// copy: copia elementos para outro slice independente
	destino := make([]int, len(a))
	copy(destino, a)
	fmt.Println("copy:", destino)

	// Pacote slices (Go 1.21+) com utilitários
	fmt.Println("contém 3?", slices.Contains(a, 3), "| índice do 4:", slices.Index(a, 4))

	// Percorrer
	for i, f := range frutas {
		fmt.Println(i, f)
	}

	// Slice de slices (matriz)
	matriz := [][]int{{1, 2}, {3, 4}}
	fmt.Println("matriz:", matriz, "elemento [1][0]:", matriz[1][0])
}
