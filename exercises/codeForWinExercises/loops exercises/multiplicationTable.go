
/*

About The Program : To print multiplication table of given integer inputs.
Working of the program : This program needs two integer inputs as operand One and one and operant Two
based on the values of operands, logic block is evaluated. 

*/


package main
import "fmt"
func input(){
    var num1, num2 int
    var operand1Ptr *int = &num1
    var operand2Ptr *int = &num2
    fmt.Printf("Enter Operand One : ")
    fmt.Scan(operand1Ptr)
    fmt.Printf("Enter Operand Two : ")
    fmt.Scan(operand2Ptr)
    logicBlock(num1,num2)
}
func logicBlock(operand1, operand2 int){
    if(operand2 == 0) {
        fmt.Printf("%d x %d = %d\n", operand1, operand2, operand1 * operand2)
    } else if ((operand1 >= 0) && (operand2 > 0)) {
        for i := 1; i <= operand2; i++ {
            fmt.Printf("%d x %d = %d\n", operand1, i, (operand1) * i)
        }
    } else if ((operand1 >= 0) && (operand2 < 0)) {
        for i := -1; i >= operand2; i-- {
            fmt.Printf("%d x %d = %d\n",operand1, i, (operand1)*(i))
        }
    } else if ((operand1 <= 0) && (operand2 < 0)) {
        for i := -1; i >= operand2; i-- {
            fmt.Printf("%d x %d = %d\n", operand1, i, (operand1)*(i))
        }
    } else {
        for i := 1; i <= operand2; i++ {
            fmt.Printf("%d x %d = %d\n", operand1, i, (operand1)*(i))
        }
    }
}
func main(){
    input()
}
  
/*


Test Cases : 

Case 1 : 

Enter Operand one : -4
Enter Operand Two : 17
-4 x 1 = -4
-4 x 2 = -8
-4 x 3 = -12
-4 x 4 = -16
-4 x 5 = -20
-4 x 6 = -24
-4 x 7 = -28
-4 x 8 = -32
-4 x 9 = -36
-4 x 10 = -40
-4 x 11 = -44
-4 x 12 = -48
-4 x 13 = -52
-4 x 14 = -56
-4 x 15 = -60
-4 x 16 = -64
-4 x 17 = -68

Case 2 : 

Enter Operand one : 5
Enter Operand Two : 10
5 x 1 = 5
5 x 2 = 10
5 x 3 = 15
5 x 4 = 20
5 x 5 = 25
5 x 6 = 30
5 x 7 = 35
5 x 8 = 40
5 x 9 = 45
5 x 10 = 50

Case 3 : 

Enter Operand one : -9
Enter Operand Two : 9
-9 x 1 = -9
-9 x 2 = -18
-9 x 3 = -27
-9 x 4 = -36
-9 x 5 = -45
-9 x 6 = -54
-9 x 7 = -63
-9 x 8 = -72
-9 x 9 = -81

Case 4 : 

Enter Operand one : 0
Enter Operand Two : -6
0 x -1 = 0
0 x -2 = 0
0 x -3 = 0
0 x -4 = 0
0 x -5 = 0
0 x -6 = 0

Case 5 : 

Enter Operand one : 0
Enter Operand Two : 0
0 x 0 = 0


*/
