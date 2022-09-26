package queenattack

import (
	"errors"
	"math"
)

type position struct {
	x int
	y int
}

func CanQueenAttack(whitePosition, blackPosition string) (canAttack bool, err error) {
	defer func() {
		if r := recover(); r != nil {
			canAttack = false
			err = errors.New("Not able to deduce the position")
		}
	}()
	whitePos := getPosFromStr(whitePosition)
	blackPos := getPosFromStr(blackPosition)

	if whitePos.x == blackPos.x && whitePos.y == blackPos.y{
		return false,errors.New("Queens can't be at the same position")
	}
	
	return (whitePos.x == blackPos.x) ||
				 (whitePos.y == blackPos.y) || 
				 (math.Abs(float64(whitePos.x -blackPos.x)) == math.Abs(float64(whitePos.y - blackPos.y))), nil
}

func getPosFromStr(pos string) position {
	if len(pos) < 2 {
		panic("invalid position")
	}
	p := position{}
	switch pos[0] {
	case 'a':
		p.x = 1
	case 'b':
		p.x = 2
	case 'c':
		p.x = 3
	case 'd':
		p.x = 4
	case 'e':
		p.x = 5
	case 'f':
		p.x = 6
	case 'g':
		p.x = 7
	case 'h':
		p.x = 8
	default:
		panic("invalid position")
	}
	switch pos[1] {
	case '1':
		p.y = 1
	case '2':
		p.y = 2
	case '3':
		p.y = 3
	case '4':
		p.y = 4
	case '5':
		p.y = 5
	case '6':
		p.y = 6
	case '7':
		p.y = 7
	case '8':
		p.y = 8
	default:
		panic("invalid position")
	}
	return p
}
