package sorter

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

// 从文件中按行读取值
func ReadValues(infile string) (values []int, err error) {
	file, err := os.Open(infile)
	if err != nil {
		return nil,fmt.Errorf("Falied to open the input file ", infile," Error:",err)
	}

	defer file.Close()

	br := bufio.NewReader(file)
	values = make([]int, 0)

	for {
		line, isPreFix, err1 := br.ReadLine()

		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}

		if isPreFix {
			fmt.Println("A too long line, seems unexpected.")
			return
		}

		str := string(line) // 转化字符数组为字符串

		value, err1 := strconv.Atoi(str)
		if err1 != nil {
			err = err1
			return
		}

		values = append(values, value)
	}
	return
}

// 写出排序后的文件
func WriteValues(values []int, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		return fmt.Errorf("Failed to created the output file ", outfile," Error:",err)
	}
	defer file.Close()

	for _, value := range values {
		str := strconv.Itoa(value)
		if _, err := file.WriteString(str + "\n"); err != nil {
			return err
		}

	}
	return nil
}
