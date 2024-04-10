package main
import(
    "fmt"
    "strings"
)
func isPalindrome(s string)bool{
    s=strings.ToLower(s)
    for i:=0;i<len(s)/2;i++{
        if s[i]!=s[len(s)-i-1]{
            return false
        }
    }
    return true
}
func main(){
    fmt.Println(isPalindrome("Malayalam"))
    fmt.Println(isPalindrome("Hello"))
}
