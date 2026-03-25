package main

import (
	"fmt"
	"study/feature1"
	"study/feature2"
	"study/feature_postgres/simple_connection"
)

//"postgres://YourUserName:YourPassword@YourHostName:5432/YourDatabaseName"

func main() {
	fmt.Println("Hello Git")
	feature1.Feature1()
	feature2.Feature2()
	simple_connection.CheckConnection()

}
