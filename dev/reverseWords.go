package dev

func Reverse(str string) (string ,string){


chars := []rune(str)
words := strings.Fields(str)
for i,j :=0, len(words)-1; i<j;i, j=i+1, j-1{
chars[i],chars[j]= chars[j],chars[i]
words[i], words[j] = words[j], words[i]
}
 return strings.Join(words, " "), string(chars)

 }