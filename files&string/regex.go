package main
import ("fmt"
        "io/ioutil"
        "strings"
        "regexp"
        "strconv"
       )

var path = "remote#1"

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
  var list []int;

  m := regexp.MustCompile(`[0-9]+`);
  lines := strings.Split(string(content), "\n");
  for _, v := range(lines){
    i, _ := strconv.Atoi(m.FindString(v));
    list = append(list, i);
  }
  fmt.Println(list);
}
