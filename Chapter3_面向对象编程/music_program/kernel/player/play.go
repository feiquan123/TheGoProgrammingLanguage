package player

import (
	"fmt"
	"strings"
)

// 创建一个播放接口，方便后期新增各种播放器
type Player interface {
	Play(source string)
}

// 创建指定类型的播放器
func Play(source , mtype string){
	var p Player

	switch strings.ToUpper(mtype) {
	case "MP3":
		p = &MP3Player{}
	case "WAV":
		p = &WAVPlayer{}
	default:
		fmt.Println("Unsupported music type ",mtype)
	}

	p.Play(source)
}
