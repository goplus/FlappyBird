// Code generated by gop (Go+); DO NOT EDIT.

package main

import "github.com/goplus/spx"

const _ = true
const (
	gameScrollX    = -3.5
	gameScrollTime = 0.05
)

type Calf struct {
	spx.Sprite
	*index
}
type GameOver struct {
	spx.Sprite
	*index
}
type Ground struct {
	spx.Sprite
	*index
}
type Logo struct {
	spx.Sprite
	*index
}
type Pipe struct {
	spx.Sprite
	*index
}
type Reset struct {
	spx.Sprite
	*index
}
type Start struct {
	spx.Sprite
	*index
}
type TapInfo struct {
	spx.Sprite
	*index
}
type index struct {
	spx.Game
	Logo      Logo
	GameOver  GameOver
	Reset     Reset
	Ground    Ground
	Calf      Calf
	Pipe      Pipe
	Start     Start
	TapInfo   TapInfo
	mdie      spx.Sound
	mwing     spx.Sound
	mhit      spx.Sound
	mpoint    spx.Sound
	userScore int
}

var calfPlay = false
var calfDie = false
var calfGravity = 0.0
//line index.gmx:30:1
func (this *index) reset() {
//line index.gmx:31:1
	this.userScore = 0
//line index.gmx:32:1
	calfPlay = false
//line index.gmx:33:1
	calfDie = false
//line index.gmx:34:1
	calfGravity = 0.0
}
//line index.gmx:37
func (this *index) MainEntry() {
//line index.gmx:37:1
	spx.Gopt_Game_Run(this, "assets", &spx.Config{Title: "FlappyCalf (by Go+)"})
}
func (this *index) Main() {
	spx.Gopt_Game_Main(this, new(Calf), new(GameOver), new(Ground), new(Logo), new(Pipe), new(Reset), new(Start), new(TapInfo))
}
//line Calf.spx:1
func (this *Calf) Main() {
//line Calf.spx:1:1
	this.OnMsg__1("start", func() {
//line Calf.spx:2:1
		this.SetXYpos(-80, 0)
//line Calf.spx:3:1
		this.TurnTo(90)
//line Calf.spx:4:1
		this.Show()
	})
//line Calf.spx:7:1
	this.OnMsg__1("menu", func() {
//line Calf.spx:8:1
		this.Hide()
	})
//line Calf.spx:11:1
	this.OnMsg__1("menu", func() {
//line Calf.spx:12:1
		for !calfPlay {
			spx.Sched()
//line Calf.spx:13:1
			if calfDie == false {
//line Calf.spx:14:1
				for
//line Calf.spx:14:1
				i := 0; i < 10;
//line Calf.spx:14:1
				i++ {
					spx.Sched()
//line Calf.spx:15:1
					if calfPlay == false {
//line Calf.spx:16:1
						this.ChangeYpos(-0.5)
					}
//line Calf.spx:18:1
					this.Wait(0.05)
				}
//line Calf.spx:20:1
				for
//line Calf.spx:20:1
				i := 0; i < 10;
//line Calf.spx:20:1
				i++ {
					spx.Sched()
//line Calf.spx:21:1
					if calfPlay == false {
//line Calf.spx:22:1
						this.ChangeYpos(0.5)
					}
//line Calf.spx:24:1
					this.Wait(0.05)
				}
			}
		}
	})
//line Calf.spx:30:1
	this.OnMsg__1("tap", func() {
//line Calf.spx:31:1
		this.SetCostume(0)
//line Calf.spx:32:1
		for calfPlay {
			spx.Sched()
//line Calf.spx:33:1
			calfGravity = calfGravity - 0.25
//line Calf.spx:34:1
			this.Wait(0.05)
		}
	})
//line Calf.spx:38:1
	this.OnMsg__1("tap", func() {
//line Calf.spx:39:1
		for calfPlay {
			spx.Sched()
//line Calf.spx:40:1
			for !(this.KeyPressed(spx.KeySpace) || this.MousePressed()) {
				spx.Sched()
//line Calf.spx:41:1
				this.Wait(0.01)
			}
//line Calf.spx:43:1
			calfGravity = 0.8
//line Calf.spx:44:1
			for
//line Calf.spx:44:1
			i := 0; i < 10;
//line Calf.spx:44:1
			i++ {
				spx.Sched()
//line Calf.spx:45:1
				this.ChangeYpos(3.5)
//line Calf.spx:46:1
				this.Wait(0.03)
			}
//line Calf.spx:48:1
			this.Wait(0.03)
		}
	})
//line Calf.spx:52:1
	this.OnMsg__1("tap", func() {
//line Calf.spx:53:1
		for calfPlay {
			spx.Sched()
//line Calf.spx:54:1
			for !(this.KeyPressed(spx.KeySpace) || this.MousePressed()) {
				spx.Sched()
//line Calf.spx:55:1
				this.Wait(0.01)
			}
//line Calf.spx:57:1
			this.Play__1(this.mwing, true)
//line Calf.spx:58:1
			this.Wait(0.04)
		}
	})
//line Calf.spx:62:1
	this.OnMsg__1("tap", func() {
//line Calf.spx:63:1
		for calfPlay {
			spx.Sched()
//line Calf.spx:64:1
			if this.Touching("Pipe") || this.Touching(spx.EdgeTop) || this.Ypos() < -123.9 {
//line Calf.spx:66:1
				calfPlay = false
//line Calf.spx:67:1
				calfDie = true
//line Calf.spx:68:1
				calfGravity = 0
//line Calf.spx:69:1
				this.Broadcast__0("game over")
//line Calf.spx:70:1
				this.GotoFront()
//line Calf.spx:71:1
				break
			}
//line Calf.spx:73:1
			this.Wait(0.03)
		}
	})
//line Calf.spx:77:1
	this.OnMsg__1("tap", func() {
//line Calf.spx:78:1
		flag := 0
//line Calf.spx:79:1
		for calfPlay {
			spx.Sched()
//line Calf.spx:80:1
			if calfGravity != 0.0 {
//line Calf.spx:82:1
				if this.Heading() < 150 && flag == 0 {
//line Calf.spx:83:1
					this.Turn(calfGravity * -0.75)
				} else {
//line Calf.spx:85:1
					if this.Heading() < 90 {
//line Calf.spx:86:1
						flag = 0
					} else {
//line Calf.spx:88:1
						flag = 1
					}
//line Calf.spx:90:1
					this.Turn(calfGravity * 0.75)
				}
//line Calf.spx:92:1
				this.ChangeYpos(calfGravity)
			}
//line Calf.spx:94:1
			this.Wait(0.03)
		}
	})
//line Calf.spx:98:1
	this.OnMsg__1("game over", func() {
//line Calf.spx:99:1
		if this.Touching("Pipe") {
//line Calf.spx:100:1
			this.Play__1(this.mhit, true)
		}
//line Calf.spx:102:1
		this.Play__1(this.mdie, false)
	})
//line Calf.spx:105:1
	this.OnMsg__1("game over", func() {
//line Calf.spx:106:1
		this.Stop(spx.OtherScriptsInSprite)
//line Calf.spx:107:1
		this.Wait(0.1)
//line Calf.spx:108:1
		for !this.Touching("Ground") {
			spx.Sched()
//line Calf.spx:109:1
			this.ChangeYpos(-10)
//line Calf.spx:110:1
			this.Wait(0.04)
		}
//line Calf.spx:112:1
		this.SetYpos(-124)
//line Calf.spx:113:1
		this.TurnTo(180)
//line Calf.spx:114:1
		for !calfPlay {
			spx.Sched()
//line Calf.spx:116:1
			this.NextCostume()
//line Calf.spx:117:1
			this.Wait(0.1)
		}
	})
}
func (this *Calf) Classfname() string {
	return "Calf"
}
//line GameOver.spx:1
func (this *GameOver) Main() {
//line GameOver.spx:1:1
	this.OnMsg__1("start", func() {
//line GameOver.spx:2:1
		this.Hide()
//line GameOver.spx:3:1
		this.SetXYpos(0, -220)
	})
//line GameOver.spx:6:1
	this.OnMsg__1("game over", func() {
//line GameOver.spx:7:1
		this.Wait(0.5)
//line GameOver.spx:8:1
		this.GotoFront()
//line GameOver.spx:9:1
		this.Show()
//line GameOver.spx:10:1
		for this.Ypos() < 0 {
			spx.Sched()
//line GameOver.spx:11:1
			this.ChangeYpos(10.0)
//line GameOver.spx:12:1
			this.Wait(0.02)
		}
	})
//line GameOver.spx:16:1
	this.OnMsg__1("menu", func() {
//line GameOver.spx:17:1
		this.Hide()
//line GameOver.spx:18:1
		this.SetXYpos(0, -220)
	})
}
func (this *GameOver) Classfname() string {
	return "GameOver"
}
//line Ground.spx:1
func (this *Ground) Main() {
//line Ground.spx:1:1
	this.OnStart(func() {
//line Ground.spx:2:1
		this.SetXYpos(0, -130)
//line Ground.spx:3:1
		this.Show()
	})
}
func (this *Ground) Classfname() string {
	return "Ground"
}
//line Logo.spx:1
func (this *Logo) Main() {
//line Logo.spx:1:1
	this.OnMsg__1("menu", func() {
//line Logo.spx:2:1
		for {
			spx.Sched()
//line Logo.spx:3:1
			if calfDie == false {
//line Logo.spx:4:1
				for
//line Logo.spx:4:1
				i := 0; i < 10;
//line Logo.spx:4:1
				i++ {
					spx.Sched()
//line Logo.spx:5:1
					if calfPlay == false {
//line Logo.spx:6:1
						this.ChangeYpos(-0.5)
					}
//line Logo.spx:8:1
					this.Wait(0.05)
				}
//line Logo.spx:10:1
				for
//line Logo.spx:10:1
				i := 0; i < 10;
//line Logo.spx:10:1
				i++ {
					spx.Sched()
//line Logo.spx:11:1
					if calfPlay == false {
//line Logo.spx:12:1
						this.ChangeYpos(0.5)
					}
//line Logo.spx:14:1
					this.Wait(0.05)
				}
			}
		}
	})
//line Logo.spx:20:1
	this.OnMsg__1("menu", func() {
//line Logo.spx:21:1
		this.SetXYpos(0, 55)
//line Logo.spx:22:1
		this.GotoFront()
//line Logo.spx:23:1
		this.Show()
	})
//line Logo.spx:26:1
	this.OnMsg__1("start", func() {
//line Logo.spx:27:1
		this.Hide()
	})
//line Logo.spx:30:1
	this.OnMsg__1("tap", func() {
//line Logo.spx:31:1
		this.Hide()
	})
//line Logo.spx:34:1
	this.OnMsg__1("game over", func() {
//line Logo.spx:35:1
		this.Stop(spx.OtherScriptsInSprite)
	})
}
func (this *Logo) Classfname() string {
	return "Logo"
}
//line Pipe.spx:1
func (this *Pipe) Main() {
//line Pipe.spx:1:1
	this.OnMsg__1("tap", func() {
//line Pipe.spx:2:1
		this.Wait(1.5)
//line Pipe.spx:3:1
		for !calfDie {
			spx.Sched()
//line Pipe.spx:4:1
			spx.Gopt_Sprite_Clone__0(this)
//line Pipe.spx:5:1
			this.Wait(2.8)
		}
	})
//line Pipe.spx:9:1
	this.OnCloned__1(func() {
//line Pipe.spx:10:1
		this.SetCostume(spx.Rand__0(0, 4))
//line Pipe.spx:11:1
		this.SetXYpos(250, 17)
//line Pipe.spx:12:1
		this.Show()
//line Pipe.spx:13:1
		for this.Xpos() >= -250 {
			spx.Sched()
//line Pipe.spx:14:1
			this.ChangeXpos(gameScrollX)
//line Pipe.spx:15:1
			this.Wait(gameScrollTime)
		}
//line Pipe.spx:17:1
		this.Die()
	})
//line Pipe.spx:20:1
	this.OnCloned__1(func() {
//line Pipe.spx:21:1
		for {
			spx.Sched()
//line Pipe.spx:22:1
			if this.Xpos() < -80.0 && calfPlay == true {
//line Pipe.spx:23:1
				this.userScore++
//line Pipe.spx:24:1
				this.Play__1(this.mpoint, true)
//line Pipe.spx:25:1
				break
			}
//line Pipe.spx:27:1
			this.Wait(0.03)
		}
	})
//line Pipe.spx:31:1
	this.OnMsg__1("game over", func() {
//line Pipe.spx:32:1
		this.Stop(spx.OtherScriptsInSprite)
	})
//line Pipe.spx:35:1
	this.OnMsg__1("menu", func() {
//line Pipe.spx:36:1
		this.Destroy()
	})
}
func (this *Pipe) Classfname() string {
	return "Pipe"
}
//line Reset.spx:1
func (this *Reset) Main() {
//line Reset.spx:1:1
	this.OnMsg__1("game over", func() {
//line Reset.spx:2:1
		this.Wait(0.5)
//line Reset.spx:3:1
		this.GotoFront()
//line Reset.spx:4:1
		this.Show()
//line Reset.spx:5:1
		for {
			spx.Sched()
//line Reset.spx:6:1
			if this.Touching(spx.Mouse) {
//line Reset.spx:7:1
				this.SetYpos(-90)
			} else {
//line Reset.spx:9:1
				this.SetYpos(-95)
			}
//line Reset.spx:11:1
			this.Wait(0.1)
		}
	})
//line Reset.spx:15:1
	this.OnMsg__1("menu", func() {
//line Reset.spx:16:1
		this.Hide()
//line Reset.spx:17:1
		this.SetXYpos(0, -100)
	})
//line Reset.spx:20:1
	this.OnClick(func() {
//line Reset.spx:21:1
		this.Stop(spx.OtherScriptsInSprite)
//line Reset.spx:22:1
		this.Hide()
//line Reset.spx:23:1
		this.reset()
//line Reset.spx:24:1
		this.Broadcast__0("menu")
	})
}
func (this *Reset) Classfname() string {
	return "Reset"
}
//line Start.spx:1
func (this *Start) Main() {
//line Start.spx:1:1
	this.OnStart(func() {
//line Start.spx:2:1
		this.Broadcast__0("menu")
//line Start.spx:3:1
		this.SetXYpos(0, -95)
	})
//line Start.spx:6:1
	this.OnClick(func() {
//line Start.spx:7:1
		this.Broadcast__0("start")
//line Start.spx:8:1
		this.Hide()
	})
//line Start.spx:11:1
	this.OnMsg__1("start", func() {
//line Start.spx:12:1
		this.Hide()
	})
//line Start.spx:15:1
	this.OnMsg__1("menu", func() {
//line Start.spx:16:1
		this.GotoFront()
//line Start.spx:17:1
		this.Show()
//line Start.spx:18:1
		for this.Visible() {
			spx.Sched()
//line Start.spx:19:1
			if this.Touching(spx.Mouse) {
//line Start.spx:20:1
				this.SetYpos(-90)
			} else {
//line Start.spx:22:1
				this.SetYpos(-95)
			}
//line Start.spx:24:1
			this.Wait(0.1)
		}
	})
//line Start.spx:27:1
	this.OnMsg__1("game over", func() {
//line Start.spx:28:1
		this.Stop(spx.OtherScriptsInSprite)
	})
}
func (this *Start) Classfname() string {
	return "Start"
}
//line TapInfo.spx:1
func (this *TapInfo) Main() {
//line TapInfo.spx:1:1
	this.OnStart(func() {
//line TapInfo.spx:2:1
		this.SetXYpos(0, 0)
//line TapInfo.spx:3:1
		this.Hide()
	})
//line TapInfo.spx:6:1
	this.OnMsg__1("start", func() {
//line TapInfo.spx:7:1
		this.Show()
//line TapInfo.spx:8:1
		this.Wait(0.25)
//line TapInfo.spx:9:1
		for !(this.KeyPressed(spx.KeySpace) || this.MousePressed()) {
			spx.Sched()
//line TapInfo.spx:10:1
			this.Wait(0.03)
		}
//line TapInfo.spx:12:1
		calfPlay = true
//line TapInfo.spx:13:1
		this.Broadcast__0("tap")
//line TapInfo.spx:14:1
		this.Hide()
	})
//line TapInfo.spx:17:1
	this.OnMsg__1("menu", func() {
//line TapInfo.spx:18:1
		this.Hide()
	})
//line TapInfo.spx:20:1
	this.OnMsg__1("game over", func() {
//line TapInfo.spx:21:1
		this.Stop(spx.OtherScriptsInSprite)
	})
}
func (this *TapInfo) Classfname() string {
	return "TapInfo"
}
func main() {
	new(index).Main()
}
