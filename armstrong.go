package armstrong


func Armstrong(a int) int {
        var num int
        n := a //storing the value of a in another variable
        for {
                if n != 0 {
                        rem := n % 10          //rem is for the remainder. % is used for modulus of a number. It returns the last digit
                        num += rem * rem * rem //multiplying the remainder thrice and then adding
                        n = n / 10             //reduce n to 2 digits
                } else {
                        break
                }
        }
	return num
}

