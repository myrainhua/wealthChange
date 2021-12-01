package main

import "fmt"
import "math/rand"
import "time"
import "flag"

/*import "sort"

//要对golang map按照value进行排序，思路是直接不用map，用struct存放key和value，实现sort接口，就可以调用sort.Sort进行排序了。
func listSort(){
    type Pair struct {
        Key   int64
        Value int64
    }

    type PairList []Pair

    p := make(PairList, len(wealthList))
    for k, v := range wealthList {
        p[k] = Pair{k, v}
    }
    sort.Sort(p)
} */

// 定义命令行参数
var people int64
var initWealth int64
var rounds int64
func Init() {
    flag.Int64Var(&people, "p", 100, "Input number of people")
    flag.Int64Var(&initWealth, "w", 100, "Input initial wealth of each people")
    flag.Int64Var(&rounds, "r", 100, "Input wealth transation rounds")
}

var wealthList = make(map[int64]int64)
var effortPeople int64
var rich2People int64

//初始化大家的原始资本为相同
func wealthInit() {  
    var i int64
    for i = 0; i < people; i++ {
        wealthList[i] = initWealth
    }
}

//输出交易后的结果，努力的人序号从0开始，富二代的序号跟在努力的人后面
func listOutput(){
    for k,_ := range wealthList {
        if k < effortPeople {
            fmt.Printf("p=**%d**\t%d\n", k, wealthList[k])//努力的人
        }else if k >= effortPeople && k < (effortPeople+rich2People){
            fmt.Printf("p=$$%d$$\t%d\n", k, wealthList[k])//富二代
        }else {
            fmt.Printf("p=%d\t%d\n", k, wealthList[k])
        }
    }
}

//基于原始资本进行随机交易
func wealthTrans(){
    rand.Seed(time.Now().UnixNano())
    for k,_ := range wealthList {
        num := rand.Int63n(people)
        wealthList[k] = wealthList[k]-1
        wealthList[num] = wealthList[num]+1
    }
}

//少量人较他人努力的情况
func wealthTransEffort(){
    var enlargeTimes int64 = people * people
    var randBase int64 = enlargeTimes+effortPeople //放大随机基数，实现100人的情况下，少量人较他人努力1%
    rand.Seed(time.Now().UnixNano())
    for k,_ := range wealthList {
        num := rand.Int63n(randBase) 
        if num >= enlargeTimes {
            num -= enlargeTimes
        }else{
            num = num/people
        }
        wealthList[k] = wealthList[k]-1
        wealthList[num] = wealthList[num]+1
    }
}

//给少量富二代加原始资本，努力的人序号从0开始，富二代的序号跟在努力的人后面
func Rich2AddWealth(addWealth int64){
    for i := effortPeople; i < (effortPeople+rich2People); i++ {
        wealthList[i]+=addWealth
    }
}

//少量努力的人+少量富二代 都在集合中 
func wealthTransEffortAndRich2(){
    var enlargeTimes int64 = people * people
    var randBase int64 = enlargeTimes+effortPeople //放大随机基数，实现100人的情况下，少量人较他人努力1%

    rand.Seed(time.Now().UnixNano())
    for k,_ := range wealthList {
        num := rand.Int63n(randBase) 
        if num >= enlargeTimes { //增加少量命中努力人的几率
            num -= enlargeTimes
        }else{
            num = num/people
        }
        wealthList[k] = wealthList[k]-1
        wealthList[num] = wealthList[num]+1
    }
}


func main() {
    Init()
    flag.Parse()
    
    //5%的富二代和5%的更努力的人
    effortPeople = people*5/100
    rich2People = people*5/100
    //fmt.Printf("Wealth init ...\n")
    wealthInit()
    Rich2AddWealth(400) //富二代增加原始资本

    //fmt.Printf("Begin wealth Trans ...\n")
    var i int64
    for i = 0; i < rounds; i++ {
        //wealthTransEffort()
        //wealthTransRich2()
        wealthTransEffortAndRich2()
    }

    //fmt.Printf("Wealth after trans:\n")
    listOutput()
}



