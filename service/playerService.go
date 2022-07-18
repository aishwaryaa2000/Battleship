package service

import (
	"battleShip/components/player"
	"os"
	"fmt"
	"strconv"
	"bufio"
	"strings"
	"unicode"
)

func MakePlayer() (*player.Player){
	in := bufio.NewReader(os.Stdin)
	nameLabel :
	fmt.Print("Hey player,enter your name : ")
	name, _ := in.ReadString('\n')
	name = strings.TrimRight(name, "\r\n")

	if !IsLetter(name){
		fmt.Println("INVALID INPUT.A-Z or a-z only")
		goto nameLabel
	}
	
	PlayerNow:= player.New(name)
	return PlayerNow
}

func inputFromPlayerToAttack(name string,rowSize,colSize uint8 )(uint8,uint8){
	var X,Y uint8
	begin1:
		fmt.Print(name," enter X cordinate where you wish to attack : ")
		userInput  := bufio.NewReader(os.Stdin)
		userVal, err := userInput.ReadString('\n')
		if err != nil {
			goto begin1
		}

		input := strings.TrimSpace(userVal)
		xCordinate, err := strconv.Atoi(input)

		if err!=nil || xCordinate>= int(rowSize) || xCordinate<0{
			fmt.Println("Please enter a positive integer for board lesser than ",rowSize)
			goto begin1
		}

		X=uint8(xCordinate)

		
	begin2:
		fmt.Print(name," enter Y cordinate where you wish to attack : ")
		userInput  = bufio.NewReader(os.Stdin)
		userVal, err = userInput.ReadString('\n')
		if err != nil {
			goto begin2
		}

		input = strings.TrimSpace(userVal)
		yCordinate, err := strconv.Atoi(input)

		if err!=nil || yCordinate>= int(colSize) || yCordinate<0{
			fmt.Println("Please enter a positive integer for board lesser than ",colSize)
			goto begin2
		}
		
		Y=uint8(yCordinate)
		
	return X,Y
}

func inputSizeFromPlayer() (uint8,uint8){

	var rowSize uint8
	var colSize uint8
	begin1 : 
	fmt.Print("Enter number of rows of the board : ")
	userInput  := bufio.NewReader(os.Stdin)
    userVal, err := userInput.ReadString('\n')
    if err != nil {
        goto begin1
    }

    input := strings.TrimSpace(userVal)
    noOfrow, err := strconv.Atoi(input)

	if err!=nil || noOfrow<5{
		fmt.Println("Please enter an integer greater than 4 for board")
		goto begin1
	}
	rowSize=uint8(noOfrow)


	begin2 : 
	fmt.Print("Enter number of columns of the board : ")
	userInput  = bufio.NewReader(os.Stdin)
    userVal, err = userInput.ReadString('\n')
    if err != nil {
        goto begin2
    }

    input = strings.TrimSpace(userVal)
    noOfCol, err := strconv.Atoi(input)

	if err!=nil || noOfCol<5{
		fmt.Println("Please enter an integer greater than 4 for board")
		goto begin2
	}
	colSize=uint8(noOfCol)

	return rowSize,colSize

}


func IsLetter(s string) bool {
    for _, r := range s {
        if !unicode.IsLetter(r) {
            return false
        }
    }
    return true
}