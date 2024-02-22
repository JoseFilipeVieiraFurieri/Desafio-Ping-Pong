package main

import (
	"fmt"
	"time"
)

func main() {
	ping := make(chan string)
	pong := make(chan string)

	go func() {
		counter := 0
		for {

			<-time.After(1 * time.Second)

			if counter%2 == 0 {
				ping <- "ping"

			} else {
				pong <- "pong"
			}

			counter++
		}
	}()

	for i := 0; i < 10; i++ {
		select {
		case pingMsg := <-ping:
			fmt.Println(pingMsg)
		case pongMsg := <-pong:
			fmt.Println(pongMsg)
		}
	}

}

// etapas

/*
 - Primeiro criei dois canais
 - segundo criei uma função para chamada dos canais. o after foi usado para envio continuo de mensagem( senão o sleep so chamava uma vez e colocava a função em asleep)
 - O for foi usado para sempre ficar enviando para o canal. Porque senão ele so enviava uma vez mesmo com o after ( não achei com o time uma solução para sincronizar certinho
   o ping e o pong mas o contador serviu)
 - o contador foi para controle ele começa em zero e imprime "ping" porque zero é par e depois vai alternando com o else e enviando o pong, e incrementando o for a cada laço
   para o esquema funcionar
 - o for tradicional foi feito para controlar o numero de chamadas, porque o for dentro da função anonima é infinito.Foi colocado 10 ali para exemplificar so, mas pode ser colocado qualquer numero
 - e por final foi usado o select pra receber as mensagens, se for mensagens do canal ping, imprime o ping e vice-versa
 - foi pesquisado o time.after e o metodo com o uso do contador porque o time.sleep não estava de acordo com o que eu queria
*/
