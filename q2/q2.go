package q2

//Escreva uma função para encontrar o prefixo comum mais longo entre um array de strings.
//
//Se não houver um prefixo comum, retorne uma string vazia "".

func LongestCommonPrefix(strs []string) string {
	for _, palavra := range strs {
		if palavra == "" {
			return ""
		}
	}
	prefixo := strs[0]
	for i := 0; i < len(strs); i++ {
		for len(strs[i]) < len(prefixo) || strs[i][:len(prefixo)] != prefixo {
			prefixo = prefixo[:len(prefixo)-1]
			if prefixo == "" {
				return ""
			}
		}
	}
	return prefixo
}
