/*



What This Program Does : This Program Is Designed To Calculate The Sum Of 1 To N, Even Or Odd Natural Numbers. 
Weather To Calculate Sum Of Even Numbets Or Odd Numbers Depends On User Input. 
This Program Asks To Enter 1 To Calculate Sum Of Odd Numbers Till N Inclusively And 0 To Calculate Sum Of Even Numbers. 

This Program Is Written On : 2026-03-21 



*/

package main
import "fmt"
func input() int {
    var num int = 0
    var evenOrOdd int
    var ptr *int = &num
    var eOrOPtr *int = &evenOrOdd
    fmt.Printf("Enter A Natural Number N : ")
    fmt.Scan(ptr)
    if *ptr < 1 {
        fmt.Printf("Firstly, %v Is Not A Natural Number!\n", *ptr)
        return 0
    }
    fmt.Printf("Enter 1 For Odd Or Enter 0 For Even : ")
    fmt.Scan(eOrOPtr)
    evenOrOddFun(ptr,eOrOPtr)
    return 0
}
func evenOrOddFun(ptr1, eOrOdptr *int) int {
    var i int
    var sum int = 0
    var sumPtr *int = &sum
    if *eOrOdptr == 1 {
        i = 1
    } else if *eOrOdptr == 0 {
        i = 2
    } else {
        fmt.Printf("Entered Invalid option!\n")
        return 0
    }
    var evenOrOddWord string
        if i == 1 {
            evenOrOddWord = "Odd"
        } else  {
            evenOrOddWord = "Even"
        }
        for  ; i<=(*ptr1) ; i=i+2 {
        (*sumPtr)=(*sumPtr)+i
        }
    fmt.Printf("Sum Of 1 To %d %s Numbers Inclusive : %d\n",(*ptr1) ,evenOrOddWord , (*sumPtr))
    return 0
}
func main(){
    input()
}

/*

--------------------:    Test Cases    :--------------------

Test Case 1:

Enter A Natural Number N : 0
Firstly, 0 Is Not A Natural Number!

Test Case 2: 

Enter A Natural Number N : -28
Firstly, -28 Is Not A Natural Number!

Test Case 3: 

Enter A Natural Number N : 43
Enter 1 For Odd Or Enter 0 For Even : 37
Entered Invalid option!

Test Case 4: 

Enter A Natural Number N : 37
Enter 1 For Odd Or Enter 0 For Even : -27
Entered Invalid option!

Test Case 5:

Enter A Natural Number N : 20
Enter 1 For Odd Or Enter 0 For Even : 0
Sum Of 1 To 20 Even Numbers Inclusive : 110

Test Case 6: 

Enter A Natural Number N : 19
Enter 1 For Odd Or Enter 0 For Even : 1
Sum Of 1 To 19 Odd Numbers Inclusive : 100


*/
