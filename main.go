package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Product struct {
	Name  string
	Price int
}

func main() {
	// le a entrada
	reader := bufio.NewReader(os.Stdin)

	// cria dois slices para gravar os valores pre e pos bf
	var preBf []Product
	var afterBf []Product

	// aqui preenche os valores pre bf
	for {
		// le a linha e grava em text o valor sem espaco final
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		// verifica se o valor de leitura e o *. Se for sai do loop
		if strings.Compare("*", text) == 0 {
			break
		}

		// caso nao seja, quebra a string de input em dois elementos
		name := strings.Split(text, ",")[0]

		// aqui foi necessario fazer um parse pois o valor é uma string
		price, _ := strconv.Atoi(strings.Split(text, ",")[1])

		// adiciona no slice o struct montado
		preBf = append(preBf, Product{
			Name:  name,
			Price: price,
		})
	}

	// aqui é a mesma logica de cima porem para o slice com precos pos bf
	for {
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare("*", text) == 0 {
			break
		}

		name := strings.Split(text, ",")[0]
		price, _ := strconv.Atoi(strings.Split(text, ",")[1])

		afterBf = append(afterBf, Product{
			Name:  name,
			Price: price,
		})
	}

	// depois de montados os dois slices com o nome e valor pecorro eles para encontrar os nomes iguais e ai comprar os valores
	for i := range preBf {
		for j := range afterBf {
			if strings.EqualFold(preBf[i].Name, afterBf[j].Name) {
				pricePreFloat := float32(preBf[i].Price)
				priceAfterFloat := float32(afterBf[j].Price)
				percent := (pricePreFloat - priceAfterFloat) * 100 / pricePreFloat

				if percent > 30 {
					fmt.Printf("%s,%.2f%%", preBf[i].Name, percent)
					fmt.Println()
				}
			}
		}
	}
}
