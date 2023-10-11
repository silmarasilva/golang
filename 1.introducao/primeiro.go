// um pacote é um grupo de arquivos que estão no mesmo diretório e são compilados juntos.
// uma variável que é declarada em um arquivo firacarão visíveis para todos os outros arquivos dentro do mesmo pacote.
// para o pacote ser executável, ele precisa ser necessariamente main
// a lista de importe é onde você referencia os pacotes que serão usados nesse arquivo.
// para rodar use run e o nome do arquivo.
// a função e o pacote precisam ser main.

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Olá Mundo!")
}
