package error

// 错误类型
import (
	"fmt"
	"strconv"
)

// 自定义错误类型
type DivError struct{
	a,b float64
	des string
	Err error
}
// 实现 Error 接口
func (e *DivError) Error() string {
	return strconv.FormatFloat(e.a,'f',6,64) + " / "+ strconv.FormatFloat(e.b,'f',6,64)+" = Nil. "+ e.des +e.Err.Error()
}


// 运算
func div(a,b float64)(res interface{},err error){
	if b == 0{
		err = fmt.Errorf("div not is 0 \n")
		return "this is error",&DivError{a,b," b is 0. \n",err}
	}
	return a/b,nil
}

func main(){
	res,err := div(1,2)
	if err != nil{
		// 打印错误
		fmt.Print(err.Error())
		// 错误处理
		if e,ok := err.(*DivError); ok && e.Err!= nil{
			// 获取 DivError 变量中的其它信息处理
			fmt.Println(e.a)
			fmt.Println(e.b)
			fmt.Println(res.(string))
		}
	}
	res,ok := res.(float64)
	if ok{
		fmt.Println(res.(float64))
	}
}
