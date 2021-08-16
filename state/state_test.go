package state

import (
	"testing"
)

func TestMario(t *testing.T) {
	ctx := NewMarioContext()
	ctx.MeetMushroom()
	ctx.MeetFireFlower()
	ctx.MeetFireFlower()
	ctx.MeetMonster()
	ctx.MeetFireFlower()
	ctx.MeetMushroom()
	ctx.MeetMonster()
	ctx.MeetMushroom()
	ctx.MeetMushroom()
	ctx.MeetMonster()
}
