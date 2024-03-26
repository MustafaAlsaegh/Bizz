package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)
const tickets=50
var remainingTickets uint=50
//var bookings = make([]map[string]string,0) //empty list of slice maps
var bookings = make([]UserData,0)
type UserData struct{
	firstName string
	lastName string
	email string
	orderedTickets uint
}
var wg=sync.WaitGroup{}
func main(){
	
	//bookings  := []string{}

	//fmt.Print("Welcome to our conference booking application");
	//fmt.Println("Welcome to our conference booking application") //with new line

	//simple function
	GreetUsers()

	//var conferenceName="Go Conference"
	

	//Parametarized function
	GreetUsersTicket()
	fmt.Println("===========================")
	
	firstName,lastName,email,orderedTickets := GetUserInputs()

		isValidName,isValidEmail,isValidTicketNumber := helper.UserInputValidation(firstName,lastName,email,orderedTickets,remainingTickets)

		if isValidName&&isValidEmail&&isValidTicketNumber{
			
			BookTicket(remainingTickets,orderedTickets,firstName,lastName,email)
			wg.Add(1)
			go sendTicket(orderedTickets,firstName,lastName,email)//creating seperate threat for run this function with [go] keyword
			
			fmt.Printf("All booking First Names %v\n",PrintFirstNames())
			fmt.Println()
			isremainingTicketCheck := remainingTickets <= 0
			if  isremainingTicketCheck {
				fmt.Println("All tickets booked out come back tomorrow")
				//break
			}
		}else{
			if !isValidName{
				fmt.Println("You entered firstname or lastname too short")
			}
			if !isValidEmail{
				fmt.Println("You entered email address dosen't contain '@' sign")
			}
			if !isValidTicketNumber{
				fmt.Println("Number of tickets you entered incorrect.")
			}
			fmt.Println("Enter correct information.Try again..")
			fmt.Println("=======================")		
		}			
	wg.Wait()
	// city := "London"

	// switch city {
	// case "New York":
	// case "Singapore":
	// case "London":
	// case "Berlin":
	// case "Mexico City":
	// case "Hong Kong":
	// default:
	// 	fmt.Println("No valid city selected.")

	// }

	// fmt.Printf("The whole array: %v\n",bookings)
	// fmt.Printf("First Value: %v\n",bookings[0])
	// fmt.Printf("Array type: %T\n",bookings)
	// fmt.Printf("Array length: %v\n",len(bookings))

	// fmt.Printf("The whole slice: %v\n",bookings)
	// fmt.Printf("First Value of slice: %v\n",bookings[0])
	// fmt.Printf("slice type: %T\n",bookings)
	// fmt.Printf("slice length: %v\n",len(bookings))
}
func GreetUsers(){
	fmt.Println("Welcome to our conference booking application")
}
func GreetUsersTicket(){
	fmt.Printf("Total Tickets %v Remaining Tickets %v \n",tickets,remainingTickets)
	fmt.Print("Get Your tickets here..\n")
}
//Function with value return
func PrintFirstNames() [] string{
	firstNameList := []string{}
	for _,booking := range bookings{
		firstNameList= append(firstNameList,booking.firstName)
	}
	return firstNameList		
}
func GetUserInputs()(string,string,string,uint){
	var firstName string
	var lastName string
	var email string
	var orderedTickets uint
	fmt.Printf("Enter your name: ")
	fmt.Scan(&firstName)
	fmt.Println()
	fmt.Printf("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println()
	fmt.Printf("Enter your email: ")
	fmt.Scan(&email)
	fmt.Println()
	fmt.Printf("Enter tickets: ")
	fmt.Scan(&orderedTickets)
	return firstName,lastName,email,orderedTickets	
}
func BookTicket(remainingTickets uint,orderedTickets uint,firstName string,lastName string,email string){
	
	//creating a map for users
	//var userData=make(map[string]string)//empty map [key data type]map data type
	//creating a struct for users
	var userData=UserData{
		firstName: firstName,
		lastName: lastName,
		email: email,
		orderedTickets: orderedTickets,
	}
	

	//remainingTickets=remainingTickets-orderedTickets

	bookings=append(bookings,userData)
	fmt.Println("==============================")
	fmt.Printf("List of bookings: %v\n",bookings)
	// fmt.Printf("Last name is: %v\n",lastName)
	// fmt.Printf("Email is: %v\n",email)
	// fmt.Printf("Ordered tickets: %v\n",orderedTickets)
	// fmt.Printf("Remaining tickets: %v\n",remainingTickets)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep((10*time.Second))
	var ticket=fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#########################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n",ticket,email)
	fmt.Println("#########################")
	wg.Done()
}