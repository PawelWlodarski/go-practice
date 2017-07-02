package main

import "fmt"

//EXERCISE1
type DockerImage struct {
	name string
	memory float64
}

func display(image DockerImage) string  {
	return fmt.Sprintf("{name=%s,memory=%.1f}",image.name,image.memory)
}

func main() {
	//EXERCISE1
	//Create function to display structure to json (display string)
	fmt.Println("Exercise1")
	image := DockerImage{"phpmyadmin", 200.0}
	comparison := display(image) == "{name=phpmyadmin,memory=200.0}"
	fmt.Println("docker struct comparison : ",comparison)


	//EXERCISE2
	//complex structure conversions
	type PhoneNumber string
	type City  string

	type Address struct {
		city City
		street string
		number int
	}

	type Employee struct {

		address Address
		name string
		phoneNumber PhoneNumber
	}

	//this must compile
	var c City = "Lodz"

	address := Address{city:c,street:"Piotrkowska",number:7}

	var pn PhoneNumber = "123456"

	employee := Employee{address,"Roman",pn}

	fmt.Println(employee)
}
