package armstrong
import "testing"
//import "fmt"
func TestArmstrong(t *testing.T) {
	var a int
	a = Armstrong(153)
	//fmt.Println(a)
	if a != 153{
		t.Error("Expected 153 but got : ",a)
	}
}
/*func TestArmstrong1(t *testing.T){
	var a int
	var i int
	for i=145; i<155; i++{
		a = Armstrong(i)
		if a != i{
			t.Error("Expected ", i, " but got : ",a)
		}
	}
}*/
