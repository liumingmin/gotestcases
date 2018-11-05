// @title app api
// @version 1.0
// @description  api server
// @termsOfService http://swagger.io/terms/

// @contact.name liumm
// @contact.url liumm
// @contact.email liumm

// @license.name liumm
// @license.url liumm

// @host 127.0.0.1:8000
// @BasePath /
package main

import (
	"fmt"
	"time"
	"os"
	"log"
	"runtime/pprof"
	"runtime"
	"sync"
	_ "golang.org/x/text/encoding/simplifiedchinese"
	//"container/list"
	//"unsafe"
	//"unsafe"
	//"unsafe"
	"gopkg.in/iconv.v1"
	//"crypto/hmac"
	//"crypto/sha1"
	//"strconv"
	//"net"
	//"encoding/json"
	//"reflect"
	"encoding/json"
	//"unsafe"
	//"unsafe"
	"unsafe"
	//"woh/platform/common/model/err828"
	//"gitlab.xsjcs.cn/techcenter/go/log4go"
	//"gitlab.xsjcs.cn/techcenter/go/ginweb/util/log"
	//"gitlab.xsjcs.cn/techcenter/go/ginweb/util/log"
	//"strconv"
	//"gitlab.xsjcs.cn/techcenter/go/golib/safego"
	//"strconv"
	//"math"
	//"net"
	//"sync/atomic"
	//"bufio"
	//"encoding/binary"
	//"bytes"
	//"strconv"
	//"math"
	"github.com/pkg/errors"

	_ "github.com/liumingmin/gotestcases/apackage"
	//"strconv"
	"net"
	"strconv"
	"math/rand"
	"github.com/gin-gonic/gin"
	"github.com/liumingmin/gotestcases/gtime"
	"github.com/liumingmin/lighttimer"
	"woh/platform/common/util/safego"
	"net/rpc"
)

import "C"

func ConvertUTF8BytesTOGBKBytes(utf8Bs []byte)([]byte) {
	//code_convertor, err := iconv.Open("utf-8", "gb2312") // convert from gbk to utf8
	code_convertor, err := iconv.Open("gb18030","utf-8" ) // convert from utf8 to gbk
	defer code_convertor.Close()
	var str string
	if err == nil {
		str = code_convertor.ConvString(string(utf8Bs))
		return []byte(str)
	}
	return nil
}

type UserChatRole struct{
	PersonID      string      `json:"personID,omitempty"` //自然人Id
	Passport string `json:"passport" binding:"required"`
	GlobalRoleID string `json:"globalRoleID"`
	RoleID string `json:"roleID"`
	Zone  string `json:"zone"`
	Server  string `json:"server"`
	Bodily string `json:"bodily"`//体型
	Camp string `json:"camp"`//阵营
	Force string `json:"force"`//门派
	Name string `json:"name"`//名称
	CenterID  string `json:"centerID"`
}

var connectingServerMap = new(sync.Map)

func testConcurrent(grId, serverId string){
	_, exist := connectingServerMap.LoadOrStore(serverId,1)
	if exist{
		fmt.Printf("repeat connect , serverId : %v ,return !!!!\n", serverId)
		return
	}

	fmt.Printf("now  %v connecting ,serverId : %v ,ok\n",grId, serverId)

	defer func() {
		connectingServerMap.Delete(serverId)
	}()
}

var gTestMap = make(map[int]string)

func testmap(idx int){
	v,isok := gTestMap[idx]

	if isok{
		fmt.Println(v)
	}


}

var gid int=0
type pobj struct {
	id int
}

func main(){
	//testchan1()
	////c:=make(chan string,1)
	//close1c(11)
	//var m int64 =1
	//xx := &m
	//p :=unsafe.Pointer(xx)
	//fmt.Println(unsafe.Sizeof(m))

	//for i := 0; i < 100; i++ {
	//	gTestMap[i]="xxxx"
	//}
	//
	//for i := 0; i < 100; i++ {
	//	tmpI := i
	//	//go testConcurrent(strconv.Itoa(tmpI),"s-1")
	//	go testmap(tmpI)
	//}
	//c := make(chan string,10)
	//fmt.Println(len(c))
	//
	//time.Sleep(time.Second*100)

	//p:=sync.Pool{
	//	New: func() interface{} {
	//		//gid++
	//		return &pobj{}
	//	},
	//}
	//
	//p1 := p.Get().(*pobj)
	//
	//fmt.Println(unsafe.Pointer(p1))
	//
	////p.Put(p1)
	//
	//p2 := p.Get().(*pobj)
	//
	//fmt.Println(unsafe.Pointer(p2))
	//
	//p3 := p.Get().(*pobj)
	//
	//var t *pobj = &pobj{}
	//
	//fmt.Println(unsafe.Sizeof(uintptr(0)))
	//fmt.Println(uintptr(unsafe.Pointer(t)))

	//bc :=make(chan []byte, 32)
	//bc <-[]byte("11111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111")
	//
	//fmt.Println(<-bc)
	//
	//fmt.Println(<-bc)

	//mut := &sync.Mutex{}
	//cond := sync.NewCond(mut)

	//for i := 0; i <200; i++ {
	//	//go func() {
	//		strs = append(strs,"11")
	//	//}()
	//}
	//
	//fmt.Println(cap(strs))
	//
	//
	//time.Sleep(time.Second*10)


	//fmt.Println(len(strs))
	//p.Put(p2)

	//fmt.Println(giveerr2())
	//go testtcpserver()
	//testtcp()

	//fmt.Println(time.Unix(1540476526,0))
	//testgoslice()
	//serverId:=1
	//testomitempty()

	//testgbk()

	//testatomic()

	//atomic.CompareAndSwapInt32()

	//testgtime()

	//testchanmap()


	var respInfos []string
	//result := AsyncInvokeWithTimeout(time.Second*1, func() {
	//	time.Sleep(time.Second*2)
	//	respInfos = []string{"we add1","we add2"}
	//	fmt.Println("1done")
	//},func() {
	//	time.Sleep(time.Second*3)
	//	//respInfos = append(respInfos,"we add3")
	//	fmt.Println("2done")
	//})
	//
	//fmt.Println("1alldone:",result)



	//result2 := AsyncInvokesWithTimeout(time.Second*3,[]func(){func() {
	//	time.Sleep(time.Second*1)
	//	respInfos = []string{"we add1","we add2"}
	//
	//	//panic(nil)
	//	fmt.Println("1done")
	//},func() {
	//	time.Sleep(time.Second*2)
	//	respInfos = append(respInfos,"we add3")
	//	//panic(nil)
	//	fmt.Println("2done")
	//
	//}} )
	//
	//fmt.Println("2alldone:", result2)
	//fmt.Println(respInfos)
	//time.Sleep(time.Hour)
	//
	////fmt.Println(testretarray())
	//var retstrs []string
	//retstrs = append(retstrs,"testretstr")
	//
	//fmt.Println(retstrs)

	//rpc.Call{}
}




func testchanmap(){
	c := make(chan map[string]string)
	go func() {
		var roleInfo map[string]string
		defer func() {
			c <- roleInfo
		}()
		roleInfo = prodMap()
	}()

	fmt.Println(<-c)



}


func testretarray() (retstrs []string){
	retstrs = append(retstrs,"testretstr")
	return
}

func prodMap()map[string]string{
	m:= make(map[string]string)
	m["222"]="3333"
	return m
}

func AsyncInvokesWithTimeout(timeout time.Duration, fs []func()) bool {
	return AsyncInvokeWithTimeout(timeout, fs...)
}

func AsyncInvokeWithTimeout(timeout time.Duration,args ...func()) bool {
	if len(args) == 0{
		return false
	}

	wg :=&sync.WaitGroup{}

	for _,arg := range args{
		f := arg
		wg.Add(1)
		safego.Go( func() {
			defer wg.Done()
			f()
		})
	}

	return waitInvokeTimeout(wg,timeout)
}

func waitInvokeTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	c := make(chan struct{})
	safego.Go(func() {
		defer close(c)
		wg.Wait()
	})
	select {
	case <-c:
		return true // completed normally
	case <-time.After(timeout):
		return false // timed out
	}
}


func testcalltimeout(){
	lighttimer.StartTicks(time.Millisecond)
	c := make(chan string,1)
	defer close(c)

	go func() {
		time.Sleep(time.Second*2)
		select {
		case c <- "111":
		default:
		}

		fmt.Println("run 1 ...")
	}()

	lighttimer.AddCallback(time.Second*3, func() {
		select {
		case c <- "111":
		default:
		}

		fmt.Println("run 2 ...")
	})

	<-c
	fmt.Println("run end ...")

}

func testgtime() {
	fmt.Println(gtime.Now())
	time.Sleep(time.Second)
	fmt.Println(gtime.Now())
	time.Sleep(time.Second)

	gtime.Sync(time.Second*5)
	t1 := time.Now()
	t2 := time.Now()
	if t1.After(t2) {
		fmt.Println("time out of order, %v > %v", t1, t2)
	}

	for{
		func (){
			defer func() {recover()}()

			fmt.Println(gtime.Now())
		}()

		time.Sleep(time.Second)
	}

}

func ppprof(){
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// automatically add routers for net/http/pprof
	// e.g. /debug/pprof, /debug/pprof/heap, etc.
	//ginpprof.Wrap(router)

	// ginpprof also plays well with *gin.RouterGroup
	// group := router.Group("/debug/pprof")
	// ginpprof.WrapGroup(group)

	router.Run(":8080")
}

func testgbk() {
	code_convertor, err := iconv.Open("gb18030", "utf-8") // convert gbk to utf-8
	if err == nil {
		defer code_convertor.Close()
		line_gbk := code_convertor.ConvString("杨晞琳")
		if line_gbk == "" {
			fmt.Println("failed")
		}
	}
}

type testot struct {
	Id string `json:"id"`
	ChatFrozen bool `json:"chatFrozen,omitempty"`
}

func testomitempty(){
	t1 := &testot{Id:"myid",ChatFrozen:true}

	jstr,_ := json.Marshal(t1)
	fmt.Println(string(jstr))
}

func testmapping(){
	groupInt:=2
	memberInt:=1

	serverId := uint32(byte(groupInt)<<5 | (byte(memberInt << 3 >> 3)))

	fmt.Println(serverId)
}

func testgoslice(){
	strslice := make([]string,0,100000)

	go func() {
		for i:=0;i<100000;i++{
			strslice  = append(strslice,strconv.Itoa(i))

			time.Sleep(time.Microsecond)
		}
	}()

	for j:=0;j<20;j++{
		go func() {
			tmp1 := strslice[:]
			lentmp := len(tmp1)

			if lentmp>0{
				fmt.Println("len:",lentmp,"last:",tmp1[lentmp-1])
			}

		}()

		time.Sleep(time.Millisecond*time.Duration(1000*(1+rand.Intn(10))))
	}

	time.Sleep(time.Hour)
}


func testtcpserver(){
	l,_ := net.Listen("tcp",":10000")
	c,_ := l.Accept()

	bs := make([]byte,10)
	c.Read(bs)

	fmt.Println("tcp recv",string(bs))
}

func testtcp() {
	conn,_ := net.Dial("tcp","127.0.0.1:10000")

	//b := make([]byte,128)

	conn.Write([]byte("get"))
	conn.Close()
	conn.Close()

	fmt.Println(111)
}

func giveerr1()error{
	return errors.New("a error")
}

func giveerr2() (err error){
	if err:= giveerr1();err!=nil{
		fmt.Println(1111)
	}
	return
}

func testcgo(){
	packageCount := uint32(0)
	fmt.Println((*C.ushort)(unsafe.Pointer(&packageCount)))
	fmt.Printf("%x",(uintptr)(unsafe.Pointer(&packageCount)))


	testarray := make([]string,10,20)
	testarray[9]="222"

	fmt.Println("-------------------------")
	fmt.Println((*C.uint)(unsafe.Pointer(&testarray)))
	fmt.Println((*C.uint)(unsafe.Pointer(&testarray[0])))
}

var strs []string = make([]string,0,100)

type myint int

func close1c(i myint){

}

func testchan1(){
	c:=make(chan string,1)

	//closec(c)
	//c<-""
	//fmt.Println("close 1 c")
	close(c)
	<-c
	//fmt.Println("close 2 c")
	select {
	case _, isok := <-c:
		if isok{
			close(c)
			fmt.Println("close c")
		}
		break
	default:
		close(c)
		fmt.Println("close c")
	}

	select {
	case _, isok := <-c:
		if isok{
			close(c)
			fmt.Println("close c")
		}
		break
	default:
		close(c)
		fmt.Println("close c")
	}
}

func closec(c chan interface{}){
	close(c)
}

func test21(){
	//testlog4go("array: %+v ,%s",[]UserChatRole{{},{},{},{},{},{},{},{},{},{},{},{},{}},"ddddd")
	//defer func() {
	//	e := recover()
	//	if e !=nil{
	//		fmt.Println(e)
	//	}
	//}()
	//
	//for i := 0; i < 1000; i++ {
	//	 go logutil.Info(nil,"Find chat roles %+v for user %s",[]UserChatRole{{},{},{},{},{},{},{},{},{},{},{},{},{}},"ddddd")
	//}
	//
	//for i := 0; i < 1000; i++ {
	//	go logutil.Info(nil,"Find chat roles %+v for user %s",[]UserChatRole{{},{},{},{},{},{},{},{},{},{},{},{},{}},"ddddd")
	//}
	//
	//for i := 0; i < 1000; i++ {
	//	go logutil.Info(nil,"Find chat roles %+v for user %s",[]UserChatRole{{},{},{},{},{},{},{},{},{},{},{},{},{}},"ddddd")
	//}
	//
	//time.Sleep(time.Second*10)
	//panic(1)

	//params :=[]interface{}{"111","222"}
	//testlog4go(1, params...)
	//go (func() {
	//	time.Sleep(time.Second)
	//	panic(11)
	//})()
	//
	//for{
	//	fmt.Println(strconv.Atoi("1000"))
	//	time.Sleep(time.Second)
	//}
	//var tag string
	//var num int
	//fmt.Sscanf("1-tag1-2","1-%3s-%d",&tag,&num)
	//
	//fmt.Println(tag,num)

	//for i := 0; i < 100; i++ {
	//	//idx := i
	//	//go testgor(i)
	//
	//	go func() {
	//		fmt.Println(i)
	//	}()
	//}
	//
	//time.Sleep(time.Second*5)

	//b := &B{}
	//b.AMutex.RLock()
	//fmt.Println(11111)
	//b.AMutex.RUnlock()
	//ticker := time.NewTicker(time.Second * time.Duration(3))
	//
	//for{
	//	select {
	//	case <-time.After():
	//		fmt.Println("3s ...")
	//		ticker.Stop()
	//	}
	//}

	//a:=make([]string,10)
	//b :=a[:5]
	//fmt.Println(len(b))
	//fmt.Println(b[:5+6])

	//base :=20000000
	//diglen := len(strconv.Itoa(base))
	//gRoleIdOffset := int(math.Pow10(diglen-3)) * 9
	//fmt.Println(base+gRoleIdOffset)

	c :=make (chan string)
	mut :=&sync.RWMutex{}
	//close(c)

	//c<-""

	go readlock(c,mut)

	time.Sleep(time.Second)

	writerlock(c,mut)

	time.Sleep(time.Second)
	fmt.Println("exit")
}

func writerlock(c chan string, mut *sync.RWMutex){
	fmt.Println("start get writer lock")
	mut.Lock()
	fmt.Println("end get writer lock")
	defer mut.Unlock()

	close(c)
}

func readlock(c chan string, mut *sync.RWMutex){
	mut.RLock()
	defer mut.RUnlock()

	c<-""
}

func graceClose(c chan string){
	close(c)
}

type A struct {
}

type B struct {
	AMutex sync.RWMutex
}

func testgor( i int){
	fmt.Println(i)
}

func testlog4go( arg0 interface{}, args ...interface{}){
	//switch first := arg0.(type) {
	//case string:
	//	// Use the string as a format string
	//	//testlog4go2(first, args...)
	//}
}

//func testlog4go2(format string, ar ...interface{}){
//	_, file, lineno, ok := runtime.Caller(2)
//
//	src := ""
//	if ok {
//		//src = fmt.Sprintf("%s:%d", runtime.FuncForPC(pc).Name(), lineno)
//		src = fmt.Sprintf("%s:%d", file, lineno)
//	}
//
//	msg := format
//	if len(ar)>0{
//		msg = fmt.Sprintf(format,ar...)
//	}
//	//log.Log(Level., src, msg)
//}

func test22() {
	fmt.Println("1111")
	pLimit:=make(chan struct{},1)

	pLimit <- struct{}{}

	go func() {
		time.Sleep(time.Second*2)
		<-pLimit
	}()

	select {
	case pLimit <- struct{}{}:
	case <-time.After(time.Second * 3):
		fmt.Println("2222222")
	}

	fmt.Println("3333")
}

func test11() {
	//test8()
	//fmt.Print(test3())
	//test9();

	//fmt.Println(string(ConvertUTF8BytesTOGBKBytes([]byte("我吃西瓜"))))
	//var bs []byte =nil
	//fmt.Printf("%04d",0)

	//h := hmac.New(sha1.New,nil)
	//h.Write([]byte("abcdefg"))
	//
	//fmt.Println((h.Sum(nil)))
	//
	////s1 := sha1.New()
	//fmt.Println(sha1.Sum([]byte("abcdefg")))
	//testBool()

	//defer fmt.Println("GetMatchSummaryIds", time.Now().UnixNano())
	//fmt.Println(time.Now().UnixNano())
	//time.Sleep(time.Second*4)

	//fmt.Println(strconv.FormatFloat(1111.11, 'f', 0, 64) + "%")
	//addrs,_ :=net.InterfaceAddrs()
	//for _,v := range addrs{
	//	fmt.Println(v)
	//}


	var err error
	//n := &none{}//非空，指针内容有值
	//err := json.Unmarshal([]byte(`{"Name":"123"}`),n)
	//fmt.Println(n,err)
	//fmt.Println(unsafe.Pointer(n))

	var n1 *none//空指针
	fmt.Println(unsafe.Pointer(&n1))
	err = json.Unmarshal([]byte(`{"Name":"123"}`),&n1)
	fmt.Println(n1,err)

	var in1 interface{} = n1
	fmt.Println(unsafe.Pointer(&in1))
	err = json.Unmarshal([]byte(`{"Name":"123"}`),&in1)
	fmt.Println(n1,err)

	var n2 *none//空指针
	var n2ptr  =&n2
	*n2ptr = &none{"456"}
	fmt.Println(n2)
	//v := reflect.ValueOf(n)
	//vv := reflect.Indirect(v)
	//fmt.Println(n)
	//fmt.Println(unsafe.Pointer(n),unsafe.Pointer(n))

	v,err := json.Marshal(nil)
	fmt.Println(string(v),err)


	//fmt.Println(unsafe.Pointer(n),unsafe.Pointer(&n))
}

var bVaule bool =true
func testBool(){
	go func() {
		for{
			bVaule = false
			//fmt.Println(bVaule)

		}
	}()

	go func() {
		for{
			bVaule = true
			//fmt.Println(bVaule)
			//time.Sleep(time.Millisecond)
		}
	}()

	time.Sleep(time.Second*100)
}

var chan0 = make(chan string,1)

func testchan0() {
	go func() {
		for{
			select {
			case chan0 <-"11111":
			default:

			}

			time.Sleep(time.Second)
		}
	}()

	go func() {
		for{
			select {
			case  v,isok:=<-chan0:
				fmt.Println(v,isok)
				break
			default:

			}

			time.Sleep(time.Second/2)
		}
	}()

	time.Sleep(time.Second*100)
}

func TestMaohaoImplict(){
	var i int
	 str,i := getItem()

	fmt.Println(str,i)
}

func getItem() (string,int){
	return "22",1
}

// @Summary 备注
// @Produce json
// @Param name query string true "Name"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/testchan [post]
func testchan(){
	//runtime.GOMAXPROCS(0)
	chan1 := make(chan string, 10)
	chan2 := make(chan string,10)

	wg := &sync.WaitGroup{}
	wg.Add(2)


	go func() {
		for i := 0; i < 10; i++ {
			chan1 <- fmt.Sprintf("setA:%d",i)
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			chan2 <- fmt.Sprintf("setB:%d",i)
		}
		wg.Done()
	}()

	wg.Wait()

	close(chan1)

	//chan1<-"close send"

	for {
		select {
			case v,isok:=<-chan1:
				if isok{
					fmt.Println(v)
				}else{
					time.Sleep(time.Second)
					fmt.Println("closed")
				}
				break;
			case v:=<-chan2:
				//if isok{
				//	fmt.Println(v)
				//}
				fmt.Println("22222",v)
				break;
		default:
			fmt.Println("no")
			time.Sleep(time.Second)
			//goto end
		}
	}
	//end:

	fmt.Println("output")
}

type none struct {
	Name string
}


// @Summary intro
// @Produce  json
// @Param name query string false "名称"
// @Param id path string true "id"
// @Param pwd formData string true "表单字段" minlength(6) maxlength(20)
// @Param age formData int false "表单字段" mininum(16) maxinum(120)
// @Param kind formData int false "表单字段" Enums(1, 2, 3)
// @Param bData body string true "json字段" default(json数据)
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Failure 400 {string} json "{"code":400,"data":{},"msg":"failure"}"
// @Router /api/user/{id} [post]
func test10(){
	//m:=make(map[string]string)
	s:=make([]string,0)
	changeName(&s)
	fmt.Println(s)

	{
		t:= none{}
		fmt.Println(t)
	}

	//unsafe.
}

// @Summary intro
// @Produce  json
// @Param name query string false "名称"
// @Param id path string true "id"
// @Param pwd formData string true "表单字段" minlength(6) maxlength(20)
// @Param age formData int false "表单字段" mininum(16) maxinum(120)
// @Param kind formData int false "表单字段" Enums(1, 2, 3)
// @Param bData body string true "json字段" default(json数据)
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Failure 400 {string} json "{"code":400,"data":{},"msg":"failure"}"
// @Router /api/user/{id} [post]
func changeName(i *[]string){
	m:=*i//*(i.(*[]string))
	//m["Name"] = "22222"
	m=append(m,"222")


	m2 :=m[:]

	m2 = append(m2,"33")
	fmt.Println(m)
	fmt.Println(m2)
}

func test9(){
	set := TsSet{s:[]string{"1","3","5"}}

	ch := set.Iter()
	for{
		value,isok:= <-ch
		if isok{
			fmt.Println(value,"rec")
		}


		time.Sleep(time.Second)
	}
}

type TsSet struct {
	s []string
	m sync.RWMutex
}

func (this*TsSet) Iter() <-chan string{
	ch := make(chan string)

	go func() {
		this.m.RLock()
		defer this.m.RUnlock()

		for _,elem:= range this.s{
			fmt.Println(elem," got")
			ch<-elem
		}

		close(ch)


	}()

	return ch
}

func test8(){
	s:=make([]float64,5)
	s=append(s,1,2,3)
	fmt.Println(s)

	str:=make([]string,5)
	str=append(str,"1","2","3")
	fmt.Println(s)

	//fmt.Println('A')
}

func calc(index string, a,b int)int{
	ret:=a+b
	fmt.Println(index,a,b,ret)
	return ret
}

func test7(){
	a:=1
	b:=2
	defer calc("1:",a,calc("10",a,b)) //1,3

	a=0

	defer calc("2:",a,calc("20",a,b)) // 0 2
	b=1
}
//随机触发异常
func test6() {
	runtime.GOMAXPROCS(1)
	int_chan := make(chan int,1)
	string_chan := make(chan string,1)


	int_chan<-1
	string_chan<- "hhh"

	select {
		case value:=<-int_chan:
			fmt.Println(value)
		case value:=<-string_chan:
			panic(value)
	}
}

func test5(){
	var t Human=&People{}//接口无法容纳引用
	t.ShowA()

	fmt.Println(getTech())
}

func getTech() Human{
	var stu *Teacher
	return stu
}

type Human interface {
	ShowA()
}

type People struct {

}

func (this*People) ShowA(){
	fmt.Println("showA")
	this.ShowB()
}

func (this*People) ShowB(){
	fmt.Println("showB")
}


type Teacher struct {
	People
}

func (this*Teacher) ShowA(){
	fmt.Println("teacher showA")
	this.ShowB()
}

func (this*Teacher) ShowB(){
	fmt.Println("teacher showB")
}


//最后一个携程刚好抢占时间片
func test4(){
	runtime.GOMAXPROCS(1)

	wg:=sync.WaitGroup{}

	wg.Add(30)

	//2个I不一样
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("i:",i)
			wg.Done()
		}()


	}

	for i := 0; i < 20; i++ {
		go func(i int) {
			fmt.Println("j:",i)
			wg.Done()
		}(i)


	}
	time.Sleep(time.Nanosecond)
	wg.Wait()

	//for{
	//	time.Sleep(time.Second)
	//}
}

func test3() interface{}{
	m:=make(map[string]*student)
	stus:=[]*student{
		&student{Name:"1",Age:1},
		&student{Name:"2",Age:2},
		&student{Name:"3",Age:3},
	}

	for _,stu := range stus{
		m[stu.Name] = stu
	}

	return m
}

type student struct {
	Name string
	Age int
}

func test2(){
	defer func() {fmt.Print("打印前")}()
	defer func() {fmt.Print("打印中")}()
	defer func() {fmt.Print("打印后")}()

	panic("exception")
}

func test1() {
	f, err := os.Create("profile.out")

	if err != nil {

		log.Fatal(err)

	}


	pprof.StartCPUProfile(f)

	defer pprof.StopCPUProfile()

	c1 := make(chan string,1)
	c2 := make(chan string,1)
	//c3 := make(chan string,1)
	go func() {
		for i:=0;i<10;i++{
			c1<-"A"
		}
	}()

	go func(){
		for i:=0;i<10;i++{
			v := <-c1
			c2<-v+"B"
		}
	}()

	for i:=0;i<10;i++{
		v,_ := <-c2
		fmt.Print(v+"C")

		time.Sleep(time.Duration(200)*time.Millisecond)
	}
}
