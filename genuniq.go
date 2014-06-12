package genuniq

//全局变量 num
var (
  num = make(chan int)
)

//包初始化函数
func init() {
  go func() {
    for i:= 0; ; i++ {
      num <- i
    }
  }()
}

//获取唯一的值
func GetUnique() int{
  return <-num
}
