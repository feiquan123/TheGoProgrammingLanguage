package bubblesort

// 冒泡排序
func BubbleSort(values []int){
	flag := true  // 判断这个数组在，初始时是否为有序。

	for i := 0; i<len(values)-1;i++{
		flag = true

		for j:=0;j<len(values)-i-1;j++{
			if values[j]>values[j+1]{
				values[j],values[j+1] = values[j+1],values[j]
				flag = false
			}
		}

		if flag{
			break
		}
	}
}
