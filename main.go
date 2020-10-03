package main

import (
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"time"
	"xxoo/config"
	"xxoo/initial"
	"xxoo/router"
)

func initServer() {
	fmt.Println("starting server")
	var Router = gin.Default()
	RouterGroup := Router.Group("")
	router.InitSysRouter(RouterGroup)
	//
	s := endless.NewServer(":8888", Router)
	s.ReadHeaderTimeout = 10 * time.Millisecond
	s.WriteTimeout = 18 * time.Second
	s.MaxHeaderBytes = 1 << 20
	time.Sleep(10 * time.Microsecond)
	s.ListenAndServe()
	fmt.Println("server started...")
}

func main() {
	initial.InitializeAll()

	//defer config.DB.Close()
	defer config.REDIS.Close()
	initServer()
}

type AAA struct {
	Username string
	Password string
	age      int
}

//func main() {
//	aaa := AAA{"弗兰克的撒封建时代", "fjkdlsjfsdla", 1}
//	typ := reflect.TypeOf(aaa)
//	val := reflect.ValueOf(aaa)
//	num := val.NumField()
//	fmt.Println("fkjdslajfldsjafs")
//	fmt.Println(typ)
//	fmt.Println(val)
//	fmt.Println(num)
//	kd := val.Kind()
//	fmt.Println(kd)
//
//	tagVal := typ.Field(0)
//	fmt.Println(tagVal)
//	fmt.Println(tagVal.Name)
//	fmt.Println(tagVal.Type)
//	fmt.Println(tagVal.Index)
//	fmt.Println(tagVal.Offset)
//	fmt.Println(tagVal.Anonymous)
//	vss := val.Field(0)
//	fmt.Println(vss)
//
//	c1 := Circle{10}
//	fmt.Println(getArea(c1))
//
//
//	bbb:= AAA{"1","2",3}
//	ccc:= &AAA{"1","2",3}
//
//	changeA(&bbb)
//	fmt.Println(bbb)
//
//	changeA(ccc)
//	fmt.Println(*ccc)
//}
//
//type Circle struct {
//	radius float64
//}
//
//func  getArea(c Circle) float64 {
//	//c.radius 即为 Circle 类型对象中的属性
//	return 3.14 * c.radius * c.radius
//}
//
//func changeA(aaa *AAA){
//	aaa.Username="fdjklsafkdsa1231fesgdfsadfg2121"
//}
