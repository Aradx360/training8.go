package main
import "fmt"





/*
func main() {
  //fmt.Println("Hello World!")
/*
  var MyName string = "Arad"
  var AradAge int = 13
  var man     bool = true 
  var time float32 = 8.58

fmt.Println(MyName)
fmt.Println(AradAge)
fmt.Println(man)
fmt.Println(time)
  
  My2Name := "Arad"
  Arad2age := 13
  man2 := true
  time2 := 8.58
  
  fmt.Println(My2Name)
  fmt.Println(Arad2age)
  fmt.Println(man2)
  fmt.Println(time2)
*/


/*
var playername string = "Alex"
var playerage int = 27
var playerheight float32 = 1.85
var playerlevel bool = true

fmt.Println(playername)
fmt.Println(playerage)
fmt.Println(playerheight)
fmt.Println(playerlevel)

player2name := "pedro"
player2age := 32
player2height := 1.85
player2level := false

fmt.Println(player2name)
fmt.Println(player2age)
fmt.Println(player2height)
fmt.Println(player2level)


*/


/*

const Newc int = 19

fmt.Println(Newc)

fmt.Printf("your age is %v \n",Newc)



myagee := 14
mylength := 1.35
myheight := 65


fmt.Printf("i am %v , my height is %v , my length is %v\n",myagee ,myheight ,mylength)

var StudentsName = [5] string{"mmd, john, asghar, docy, json"}


fmt.Println(StudentsName)


SomeNumber := [...]int{1, 2, 3, 4}

fmt.Println(SomeNumber)
fmt.Println(SomeNumber[2])
*/


/*
n := 123


fmt.Printf("the answer is %v . \n", n)

p := 26
y := 34
c := 60


fmt.Printf("%v plus %v be %v \n", p, y, c)
fmt.Printf("%v + %v = %v \n", p, y, c)

var CarNames =[...]string{"Ford mustang , porche 919 , nissan gtr , Lamborginiaventado , BMW coupe"}


fmt.Println(CarNames)


BestCHaracters := [...]string{"kratos , Arthur morgan , Marcus , Master cheef , scorpion , John marston"}


fmt.Println(BestCHaracters)



const RDR2 string = "Arthus morgan"


fmt.Println(RDR2)


const Csname string = "hello world"

fmt.Println(Csname)

const a = "hello"
const x = 24

fmt.Println(x)

fmt.Printf("this is %o \n", x)
fmt.Printf("this is %v%% \n", x)
fmt.Printf("this is %b \n", x)
fmt.Printf("this is %d \n", x)
fmt.Printf("this is %4d \n", x)
fmt.Printf("this is %-4d \n", x)
fmt.Printf("this is %04d \n", x)
fmt.Printf("this is %s \n", a)
fmt.Printf("this is %q \n", a)
fmt.Printf("this is %8s \n", a)
fmt.Printf("this is %-8s \n", a)
fmt.Printf("this is %x \n", a)
fmt.Printf("this is % x \n", a)




var solosgames =[...]string{"cap price , soap bishop , ghost"}


fmt.Println(solosgames)






const ps bool = false
const xbox bool = true 


fmt.Println(ps)
fmt.Println(xbox)
fmt.Printf("%v is better than %v\n", xbox , ps)

/*
var insominac =[...]string{"marvelwolverine , marvalspiderman1 , marvelspiderman2 , marvelspidermanmilesmorales ,
marvelwolverine , marvelvenom"}



fmt.Println(insominac[3])
fmt.Println(insominac)

*/

/*
var NewGames =[...]string{"final fantsy VII , cod mw4 , GTA VI"}


fmt.Println(NewGames)




slc := [3]string{"slam" , "chetori" , "khoobi"}



myslice := slc[1:3]


fmt.Println(len(myslice))
fmt.Println(cap(myslice))
fmt.Printf("my slice : %v\n", myslice)
fmt.Printf("my slice capacity : %v\n", cap(myslice))
fmt.Printf("my slice length : %v\n", len(myslice))



var u = 30 + 15 

fmt.Println(u)
/*
var (
  sum1 = 100 + 120
  sum2 = 500 + 100
  sum3 = 200 + 350
)

fmt.Printf(sum3)
*/


/*
var sum1 = 125 + 25

fmt.Println(sum1)


var sum2 = sum1 + 50

fmt.Println(sum2)


var sum3 = sum2 * 5

fmt.Println(sum3)



var number = 11

fmt.Println(number)


var big = 10
var small = 6


fmt.Println(big<small)
fmt.Println(big>small)
fmt.Println(big==small)
fmt.Println(big!=small)


xnum := 12

ynum := 15

if ynum < xnum {

  fmt.Println("true")
}else {
  fmt.Println("false")

}


training := 20

if training < 5{

  fmt.Println("its not good")
} else if training >= 20{

  fmt.Println("very good")

} else {
fmt.Println("thats good")

*/
/*
	for x := 0; x < 1; x++ {
	
    fmt.Println(sum)
	}
  */

  /*
xa := 36
if 20 > 18 && xa < 30 {

fmt.Println("great")

}else if 20 > 18 && xa < 40{

fmt.Println("close")


}else{

fmt.Println("not good")

}

var color = 4

switch color{
case 1, 2, 3:
fmt.Println("Red")
case 4, 5 ,6:
  fmt.Println("blue")
case 7, 8, 9:
  fmt.Println("Yellow")
case 10, 11 ,12:
  fmt.Println("Green")
  default :
  fmt.Println("No team")



  




}



var team = 4

switch team{

case 1:
  fmt.Println("pioneers")
case 2:
  fmt.Println("cubes")
case 3:
  fmt.Println("Dragons")
case 4:
  fmt.Println("Lions")
}

*/

/*
func FamilyMmebers(names string)  {


  fmt.Println("my name is", names ,"marston")



}




func main() {

FamilyMmebers("John")
FamilyMmebers("Abigail")
FamilyMmebers("jack")

}

*/



/*
func math(x int, y int) (result int) {

  result = x * y

  return 
  
  
}


func main() {
  total := math(8, 9)

  fmt.Println(total)
}

*/



/*


func FirstFunc(y string , x int) (txt1 string , age int) {

  txt1 = y + "am"
  age = x
return




}

func main() {

  fmt.Println(FirstFunc("I",14))


}

*/


/*
func FuncFirst(y int) int{

if y == 1 {
  return 0
}
fmt.Println(y)
return FuncFirst(y + 1)
}



func main(){
  
  FuncFirst(2)
}
*/




//fibbonacci

/*
func fibbonacci(x int , y int) {

for i := 1 ; i < 10 ; i += 1{

z := x + y
fmt.Println(z)

x = y
y = z

}


}

func main() {

fibbonacci(1, 1)


}
*/

/*
func main() {

  
  var a = map[string]string{"brand" : "Ford" , "model" : "mustang" , "year" : "2015"}

fmt.Println(a)

}

*/


func main() {

x := 12


if (x > 5 && x >= 10) {

  fmt.Println("good")

}else if (x > 10 && x <+ 20) {

fmt.Println("very good")

}else {

fmt.Println("invalid input")

}




}
