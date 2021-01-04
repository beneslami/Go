package main
import ("fmt"
        "io/ioutil"
        "strings"
       )

var path = "/Users/ben/Desktop/Memory.rtf"

func isError(err error) bool{
  if err != nil {
    fmt.Println(err.Error());
  }
  return (err != nil)
}
func main() {
  content, err := ioutil.ReadFile(path);
  if isError(err){
    return;
  }
  lines := strings.Split(string(content), "\n");
  for _, v := range(lines) {
    if strings.Contains(v, "KAIN Total READ byte") {
      fmt.Println(v);
    }
  }
}
