package majority


 func majorityElement(A []int )  (int) {

    count := 0
    lastElement :=0

    for i := 0 ; i < len(A) ; i++ {
        if count == 0 {
            lastElement = A[i]
            count = 1
        }else {
            // count not 0
            if lastElement == A[i] {
                count++
            }else{
                count--
            }
        }


    }

    return lastElement
}
