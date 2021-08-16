package state

import "fmt"

const (
	SMALL = "small"
	SUPER = "super"
	FIRE  = "fire"
	DIE   = "die"
)

// Mario is the game role
type Mario interface {
	State() string
	MeetMushroom(*MarioContext)
	MeetFireFlower(*MarioContext)
	MeetMonster(*MarioContext)
}

type defaultMario struct{}

func (*defaultMario) State() string                { return "" }
func (*defaultMario) MeetMushroom(*MarioContext)   {}
func (*defaultMario) MeetFireFlower(*MarioContext) {}
func (*defaultMario) MeetMonster(*MarioContext)    {}

// MarioContext is the state context of mario
type MarioContext struct {
	mario Mario
}

// NewMarioContext init the object with small mario
func NewMarioContext() *MarioContext {
	return &MarioContext{mario: &SmallMario{}}
}

// State return the current mario state
func (ctx *MarioContext) State() string {
	return ctx.mario.State()
}

// MeetMushroom change the current mario's state
func (ctx *MarioContext) MeetMushroom() {
	fmt.Printf("%s meet mushroom ", ctx.mario.State())
	ctx.mario.MeetMushroom(ctx)
	fmt.Printf(" --> %s\n", ctx.mario.State())
}

// MeetFireFlower change the current mario's state
func (ctx *MarioContext) MeetFireFlower() {
	fmt.Printf("%s meet fire flower ", ctx.mario.State())
	ctx.mario.MeetFireFlower(ctx)
	fmt.Printf(" --> %s\n", ctx.mario.State())
}

// MeetMonster change the current mario's state
func (ctx *MarioContext) MeetMonster() {
	fmt.Printf("%s meet monster ", ctx.mario.State())
	ctx.mario.MeetMonster(ctx)
	fmt.Printf(" --> %s\n", ctx.mario.State())
}

// SmallMario is the normal state of the game role
type SmallMario struct {
	*defaultMario
}

// NewSmallMario ...
func NewSmallMario() Mario {
	return &SmallMario{&defaultMario{}}
}

// State ...
func (small *SmallMario) State() string {
	return SMALL
}

// MeetMushroom change state to super
func (small *SmallMario) MeetMushroom(ctx *MarioContext) {
	ctx.mario = NewSuperMario()
}

// MeetFireFlower change state to fire
func (*SmallMario) MeetFireFlower(ctx *MarioContext) {
	ctx.mario = NewFireFlowerMario()
}

// MeetMonster change state to die
func (*SmallMario) MeetMonster(ctx *MarioContext) {
	ctx.mario = NewDieMario()
}

// SuperMario is the super state mario
type SuperMario struct {
	Mario
}

// NewSuperMario ...
func NewSuperMario() Mario {
	return &SuperMario{&defaultMario{}}
}

// State return the state of super mario
func (*SuperMario) State() string {
	return SUPER
}

// MeetFireFlower change the state to fire
func (*SuperMario) MeetFireFlower(ctx *MarioContext) {
	ctx.mario = NewFireFlowerMario()
}

// MeetMonster change the state to small
func (*SuperMario) MeetMonster(ctx *MarioContext) {
	ctx.mario = NewSmallMario()
}

// FireFlowerMario is the fire state mario
type FireFlowerMario struct {
	Mario
}

// NewFireFlowerMario ...
func NewFireFlowerMario() Mario {
	return &FireFlowerMario{&defaultMario{}}
}

// State return the state of fire mario
func (*FireFlowerMario) State() string {
	return FIRE
}

// MeetMonster change the state to small
func (*FireFlowerMario) MeetMonster(ctx *MarioContext) {
	ctx.mario = NewSmallMario()
}

// DieMario is the die state mario
type DieMario struct {
	Mario
}

// NewDieMario ...
func NewDieMario() Mario {
	return &DieMario{&defaultMario{}}
}

// State return the state of die mario
func (*DieMario) State() string {
	return DIE
}
