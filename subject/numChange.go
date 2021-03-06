package subject

import (
	"fmt"
	"strconv"
)

/**1.一个n位数字，找出一个比他大的最小数，数字组合保持不变，比如12345，下一个就是12354***/
/**2.一个随机数生成器，出现0的概率是n，出现1的概率是1-n，设计这个生成器，让n=0.5***/
/**3.一个2g内存的服务器，如何从1个亿数据中找出指定的几个数据**/
func NextMinNum(num int) int {
	numStr := strconv.Itoa(num)
	n := len(numStr)
	if n <= 1 {
		return num
	}

	var numSplit [100]byte
	for i := 0; i < n; i++ {
		numSplit[i] = numStr[i] - 48
		// fmt.Println(numSplit[i])
	}

	numArray := numSplit[:n]
	//倒序循环数组找出需要改变的位n和位h，交换两者的值，然后将数组分为3部分,1:n的左边所有的位（顺序不变），2:第n位，3:第n位右边的所有位
	//然后将第三部分升序排列，在将数字重新拼接
	// fmt.Println(numArray)
	numArray, position := changeThePosition(numArray)
	// fmt.Println(position)

	if position == 0 {
		return num
	}

	sliceLeft := numArray[:position+1]

	sliceRight := numSplit[position+1 : n]
	// fmt.Println("sliceLeft", sliceLeft)
	// fmt.Println("sliceRight", sliceRight)

	if len(sliceRight) == 1 {
		sliceLeft = append(sliceLeft, sliceRight...)
	} else {
		//右边的部分升序排列
		sliceRight = orderTheSlice(sliceRight)
		sliceLeft = append(sliceLeft, sliceRight...)
	}

	fmt.Println(sliceLeft)
	for iter, value := range sliceLeft {
		sliceLeft[iter] = value + 48
	}

	resultStr := string(sliceLeft)
	// fmt.Println(resultStr)
	result, _ := strconv.Atoi(resultStr)
	// fmt.Println(result)
	return result
}

func orderTheSlice(sliceRight []byte) []byte {
	rightLen := len(sliceRight)
	for outLoop := 0; outLoop < rightLen; outLoop++ {
		for inLoop := outLoop; inLoop < rightLen; inLoop++ {
			if sliceRight[outLoop] > sliceRight[inLoop] {
				sliceRight[outLoop], sliceRight[inLoop] = sliceRight[inLoop], sliceRight[outLoop]
			}
		}
	}
	return sliceRight
}

func changeThePosition(numArray []byte) ([]byte, int) {
	n := len(numArray)
	for pos := len(numArray) - 1; pos >= 0; pos-- {
		for j := n - 1; j > pos; j-- {
			if numArray[j] > numArray[pos] {
				numArray[pos], numArray[j] = numArray[j], numArray[pos]
				return numArray, pos
			}
		}
	}

	return numArray, 0
}
