package main
import "fmt"
import "bufio"
import "os"
import "strings"
func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
    var str = strings.Split(sc.Text(), " ")
    counts := make(map[string]int)

    for _, s := range str {
        counts[s] ++
    }

    // 結果の出力
    for s, n := range counts {
        fmt.Printf("%s %d\n", s, n)
    }

    // for i, v := range s {
    //     fmt.Println(v, i)
    // }

}
