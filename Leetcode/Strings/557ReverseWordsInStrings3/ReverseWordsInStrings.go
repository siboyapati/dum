package _57ReverseWordsInStrings3

import (
	"fmt"
	"strings"
)

//Input: s = "Let's take LeetCode contest"
//Output: "s'teL ekat edoCteeL tsetnoc"
func ReverseWords(s string) string {

	k := strings.Split(s," ");
	var result []string

	for _,v := range k{
		result= append(result, rev(v))
	}
	//result = append(result, "siva")
	fmt.Println(strings.Join(result, " "))
	return strings.Join(result, " ")
}

func rev(v string) string{
	vbyte := []byte(v)
	fmt.Println(vbyte)

	for i:=0; i<len(vbyte)/2; i++{
		vbyte[i], vbyte[len(v) -i -1] = vbyte[len(v) -i -1], vbyte[i]
	}
	return string(vbyte)
}


func ReverseWords2(s string) string {

	sByte:= []byte(s)

	for i:=0; i<len(sByte); i++{
		if sByte[i] != ' '{
			j := i
			for j+1 < len(sByte) && sByte[j+1] != ' '{
				j++
			}
			revrseWord(sByte, i, j)
			i = j
		}
	}
	return string(sByte)
}

func revrseWord(str []byte, low, high int){
	for low < high{
		str[low], str[high] = str[high], str[low]
		low++
		high--
	}
}

//Input: s = "Let's take LeetCode contest"
//Output: "s'teL ekat edoCteeL tsetnoc"