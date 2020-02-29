package main

import (
	// "bufio"
	// "encoding/json"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	// "io/ioutil"
	// "os"
	// "strings"
)

type ABC struct {
	X []string
	Y int
}

func (abc *ABC) UnmarshalJSON(data []byte) error {
	mystring := string(data)
	if !strings.HasPrefix(mystring, `"`) {
		return fmt.Errorf("invalid value")
	}
	mystring = strings.Trim(mystring, `"`)
	s := strings.Split(mystring, "|")
	x, y := s[0], s[1]
	abc.X = strings.Split(x, ",")
	i, err := strconv.Atoi(y)
	if err != nil {
		return err
	}
	abc.Y = i
	return nil

}

func (abc *ABC) MarshalJSON() ([]byte, error) {
	x := strings.Join(abc.X, ",")
	y := strconv.Itoa(abc.Y)
	xy := fmt.Sprintf(`"%s|%s"`, x, y)
	return []byte(xy), nil
}

func (a ABC) String() string {
	return fmt.Sprintf("I have %d items in X and the value of Y is %d", len(a.X), a.Y)
}

type abclist []ABC

func (a abclist) Len() int           { return len(a) }
func (a abclist) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a abclist) Less(i, j int) bool { return len(a[i].X) < len(a[j].X) }

// {
// 	"mydata": "1,2,3,4,5|10"
// }

func main() {

	a := ABC{[]string{"ahmad", "abdullah"}, 24}

	// str := `"1,2,3,4|9"`
	// b := ABC{}
	// if err := json.Unmarshal([]byte(str), &b); err != nil {
	// 	panic(err)
	// }
	// spew.Dump(b)
	// return

	buf, err := json.Marshal(&a)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf))

	// mylist := []ABC{{[]string{"hey", "you"}, 35}, {[]string{"a", "b", "c"}, 39}, {[]string{"alif"}, 36}}

	// fmt.Println(mylist)

	// sort.Sort(abclist(mylist))

	// fmt.Println(mylist)

}
