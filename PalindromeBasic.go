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
    for{
        fmt.Print("Enter String(enter 'exit' for exiting:")
        var input string
        fmt.Scanln(&input)
        if(input=="exit"){
            break
        }
        fmt.Println(isPalindrome(input))
    }
}
