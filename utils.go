package utils

import(
	"bufio"
	"os"
	"fmt"
	"reflect"
	"strings"
)


type Array struct{
	Ls []interface{}
	Length int
}

// SliceToArg converts []any to []interface{} for parsing
func SliceToArg(arg interface{})(output []interface{}, c int, ok bool){

	slice, suc := func ()(val reflect.Value, ok bool){
		val = reflect.ValueOf(arg)
		if val.Kind() == reflect.Slice{
			ok = true
		}
		return
	}()

	if !suc{
		ok = false
		return
	}

	c = slice.Len()
	out := make([]interface{}, c)
	for i:=0;i<c;i++{
		out[i] = slice.Index(i).Interface()
	}
	return out, c, true
}


func MakeArray()(Array){
	return Array{make([]interface{}, 0), 0}
}

func (a *Array) Len()(int){
	count := 0
	for range a.Ls{
		count++
	}
	return count
}

func (a *Array) Q(element interface{}){
	defer func(){
		err := recover()
		if err != nil{
			fmt.Println("Error on push!")
			fmt.Println(err)
		}
	}()

	a.Ls = append(a.Ls, element)
	a.Length = a.Length + 1
}

func (a *Array) Append(elements interface{}){
	defer func(){
		err := recover() 
		if err != nil{
			fmt.Println("Error on appending elements!")
			fmt.Println(err)
		}
	}()
	
	ele, c, ok := SliceToArg(elements)
	if !ok{
		panic("Not a slice of items! Try Q()!")
	}

	b := make([]interface{}, c+a.Length)
	for i := range a.Ls{
		b[i] = a.Ls[i]
	}
	for i := range ele{
		b[i+a.Length] = ele[i]
	}

	a.Ls = b
	a.Length = a.Len()
}


func (a *Array) Pop(index int)(temp interface{}){

	defer func(){
		err := recover()
		if err != nil{
			fmt.Println("Index error on pop!")
			fmt.Println(err)
		}
	}()



	length := a.Length

	if index == -1{
		temp = a.Ls[length-1]
	} else {
		temp = a.Ls[index]
	}

	if index == 0{
		a.Ls = a.Ls[1:]
	} else if index == length-1 || index == -1{
		a.Ls = a.Ls[:length-1]
	} else{
		a.Ls = append(a.Ls[:index], a.Ls[index+1:]...)
	}
	a.Length = a.Length - 1
	return 
}

func (a *Array) Acc(index int)(interface{}){
	defer func(){
		err := recover()
		if err != nil{
			fmt.Println("Index Error accessing array!")
			fmt.Println(err)
		}
	}()
	v := a.Ls[index]
	return v
}

func CleanString(s string)(string){
	o := strings.Replace(s, "\n", "", -1)
	o = strings.Replace(o, "\r", "", -1)
	return o
}


func ReadLinesArr(number int)([]string){
	scanner := bufio.NewScanner(os.Stdin)
	a := make([]string, number)
	for i := 0; i < number; i++{
		scanner.Scan()
		a[i] = CleanString(scanner.Text())
	}
	return a
}

func ReadLinesArrReader(number int, scanner *bufio.Reader)([]string){
	a := make([]string, number)
	
	defer func(){
		err := recover()
		if err != nil{
			print(err)
		}
	}()
	
	
	if number == 0{
		panic("Number is Zero")
	}
	for i := 0; i < number; i++{
		a[i], _ = scanner.ReadString('\n')
		a[i] = CleanString(a[i])
	}
	return a
}