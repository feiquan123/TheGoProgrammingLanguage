package player

import (
	"fmt"
	"time"
)

type MP3Player struct{
	stat int
	progress int
}

// 实现播放接口
func (p *MP3Player) Play(source string){
	fmt.Println("Playing MP3 music ",source)

	p.progress=0

	for p.progress<100 {
		time.Sleep(100 * time.Millisecond) // 播放
		fmt.Print(".")
		p.progress +=10
	}

	fmt.Println("\nFinished Playing ",source)
}
