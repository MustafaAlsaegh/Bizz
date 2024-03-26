package helper

import "strings"

func UserInputValidation(firstName string,lastName string,email string,orderedTickets uint, remainingTickets uint)(bool,bool,bool){
	isValidName:=len(firstName)>=2 && len(lastName)>=2
	isValidEmail:=strings.Contains(email,"@")
	isValidTicketNumber:=orderedTickets>0 && orderedTickets<=remainingTickets

	return isValidName,isValidEmail,isValidTicketNumber
}