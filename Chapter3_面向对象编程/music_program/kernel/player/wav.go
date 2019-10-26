package player

import (
	"fmt"
	"time"
)

type WAVPlayer struct{
	stat int
	progress int
}

// 实现播放接口
func (p *WAVPlayer) Play(source string){
	fmt.Println("Playing wav music ",source)

	p.progress=0

	for p.progress<100 {
		time.Sleep(100 * time.Millisecond) // 播放
		fmt.Print(".")
		p.progress +=10
	}

	fmt.Println("\nFinished Playing ",source)
}
