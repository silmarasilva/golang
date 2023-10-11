// se eu trocar o nome da função 'escrever' com letra minuscula, como o go não e uma linguagem orientada a objetos, não existe nada parecido com public private protect, se a funçaõ é pública vamos identificar pela primeira letra dela.
// se uma função começa com letra minúscula, ela só vai ser visível dentro do pacote que ela está.
// se começa com letra maiúscula, ela poderá ser exportada por outros pacotes. Isso é bem particular do Go.
// é uma boa prática ter um comentário sobre uma função exportadas explicando o que ela faz. Se não for feito o comentário ficará gerando um warning, mas funcionará normalmente.
// dentro do próprio pacote eu consigo referenciar as funções em letra minúscula, mas fora dele não.
// O modo não se atualiza automaticamente, para atualizá-lo, rodar o comando <go build> novamente.
// o comando <go install> salva o arquivo na raiz onde vc instalou o go ao invez de ficar dentro do projeto. Então usamos normalmente o <go build>

package auxiliar

// Escrever registra uma mensagem na tela.
import "fmt"

func Escrever() {
	fmt.Println("Escrevendo do pacote auxiliar")
	escrever2()
}
