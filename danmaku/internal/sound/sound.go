package sound

import (
	"io/ioutil"

	"github.com/yohamta/godanmaku/danmaku/internal/resources/audios"

	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/mp3"
	"github.com/hajimehoshi/ebiten/audio/wav"
)

const (
	sampleRate = 22050
)

// BgmKind represents kind of se
type BgmKind int

const (
	BgmKindBattle BgmKind = iota
)

// SeKind represents kind of se
type SeKind int

const (
	SeKindHit SeKind = iota
	SeKindHit2
	SeKindShot
	SeKindBomb
	SeKindJump
)

var (
	audioContext *audio.Context
	seDic        = map[SeKind]*audio.Player{}
	bgmDic       = map[BgmKind]*audio.Player{}
	volume128    = 64
)

// Load loads audio files
func Load() {
	audioContext, _ = audio.NewContext(sampleRate)

	bgmDic[BgmKindBattle] = loadMp3(audioContext, &audios.BATTLE)

	seDic[SeKindHit] = loadWav(audioContext, &audios.HIT)
	seDic[SeKindHit2] = loadWav(audioContext, &audios.HIT2)
	seDic[SeKindShot] = loadWav(audioContext, &audios.SHOT)
	seDic[SeKindBomb] = loadWav(audioContext, &audios.BOMB)
	seDic[SeKindJump] = loadWav(audioContext, &audios.JUMP)

}

// PlayBgm playes SE
func PlayBgm(kind BgmKind) {
	bgmDic[kind].Rewind()
	bgmDic[kind].Play()
}

// PlaySe playes SE
func PlaySe(kind SeKind) {
	seDic[kind].Rewind()
	seDic[kind].Play()
}

func loadWav(c *audio.Context, wavBytes *[]byte) *audio.Player {
	s, _ := wav.Decode(c, audio.BytesReadSeekCloser(*wavBytes))
	b, _ := ioutil.ReadAll(s)
	player, _ := audio.NewPlayerFromBytes(audioContext, b)
	player.SetVolume(float64(volume128) / 128)
	return player
}

func loadMp3(c *audio.Context, mp3Bytes *[]byte) *audio.Player {
	s, _ := mp3.Decode(audioContext, audio.BytesReadSeekCloser(*mp3Bytes))
	player, _ := audio.NewPlayer(audioContext, s)
	player.SetVolume(float64(volume128) / 128)
	return player
}
