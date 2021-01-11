package main
import ("fmt"
        "os"
        "io/ioutil"
        "strings"
        "regexp"
        "strconv"
        "github.com/wcharczuk/go-chart"
       )

var path = "/Users/ben/Desktop/remote#1"

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
  max := 0;
  for _, v2 := range(list){
    if(v2 > max){
      max = v2;
    }
  }
  var x_axis = make([]int, max);
  var y_axis []int;

  for i := 0 ; i < max ; i++ {
    x_axis[i] = i;
  }

  y_axis = make([]int, max);
  for i := 0; i < max; i++ {
    for j:= 0; j < len(list); j++{
      if(list[j] == i){
        y_axis[i]++;
      }
    }
  }


  lineChart(x_axis, y_axis, max);
  barChart(x_axis, y_axis, max);
}

func lineChart(x_axis []int, y_axis []int, max int){
  var x_axis_f []float64;
  x_axis_f = make([]float64, max);
  var y_axis_f []float64;
  y_axis_f = make([]float64, max);
  for i:=0 ; i < max ; i++ {
    x_axis_f[i] = float64(x_axis[i]);
    y_axis_f[i] = float64(y_axis[i]);
  }
  graph := chart.Chart{
    Series: []chart.Series{
        chart.ContinuousSeries{
            XValues: x_axis_f,
            YValues: y_axis_f,
        },
    },
  }
  f, _ := os.Create("line_output.png");
  defer f.Close();
  graph.Render(chart.PNG, f)
}
func barChart(x_axis []int, y_axis []int, max int){
  graph := chart.BarChart{
		Title: "Remote#1",
		Background: chart.Style{
		Padding: chart.Box{
		Top: 40,
			},
		},
		Height:   512,
		BarWidth: 60,
		Bars: []chart.Value{
      {Value: float64(y_axis[0]), Label:strconv.Itoa(x_axis[0])},
      {Value: float64(y_axis[1]), Label:strconv.Itoa(x_axis[1])},
      {Value: float64(y_axis[2]), Label:strconv.Itoa(x_axis[2])},
      {Value: float64(y_axis[3]), Label:strconv.Itoa(x_axis[3])},
      {Value: float64(y_axis[4]), Label:strconv.Itoa(x_axis[4])},
      {Value: float64(y_axis[5]), Label:strconv.Itoa(x_axis[5])},
      {Value: float64(y_axis[6]), Label:strconv.Itoa(x_axis[6])},
      {Value: float64(y_axis[7]), Label:strconv.Itoa(x_axis[7])},
      {Value: float64(y_axis[8]), Label:strconv.Itoa(x_axis[8])},
      {Value: float64(y_axis[9]), Label:strconv.Itoa(x_axis[9])},
      {Value: float64(y_axis[10]), Label:strconv.Itoa(x_axis[10])},
      {Value: float64(y_axis[11]), Label:strconv.Itoa(x_axis[11])},
      {Value: float64(y_axis[12]), Label:strconv.Itoa(x_axis[12])},
      {Value: float64(y_axis[13]), Label:strconv.Itoa(x_axis[13])},
      {Value: float64(y_axis[14]), Label:strconv.Itoa(x_axis[14])},
      {Value: float64(y_axis[15]), Label:strconv.Itoa(x_axis[15])},
      {Value: float64(y_axis[16]), Label:strconv.Itoa(x_axis[16])},
    },
	}

	f, _ := os.Create("bar_output.png")
	defer f.Close()
	graph.Render(chart.PNG, f)
}
