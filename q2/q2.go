package q2

import "strings"

//Escreva uma função para encontrar o prefixo comum mais longo entre um array de strings.
//
//Se não houver um prefixo comum, retorne uma string vazia "".

func LongestCommonPrefix(strs []string) string {
	var prefix string
	firstString := strs[0]
	for i := 0; i < len(strs); i++ {
		for j := 0; j < len(strs); j++ {
			if strings.HasPrefix(strs[j], firstString) {
				break
			} else {
				prefix, _ = strings.CutSuffix(firstString, string(firstString[len(firstString)-1]))
			}
		}
	}
	return prefix
}
