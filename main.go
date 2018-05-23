package main

import "fmt"

//---MAIN-FLOW---

func main() {
	fmt.Println("starting app\n")

	positiveTime1, positiveTime2, positiveTime3, negativeTime1, negativeTime2, negativeTime3 := getTimes()

	fmt.Printf("%+v\n", positiveTime1)
	fmt.Printf("%+v\n", positiveTime2)
	fmt.Printf("%+v\n", positiveTime3)
	fmt.Printf("%+v\n", negativeTime1)
	fmt.Printf("%+v\n", negativeTime2)
	fmt.Printf("%+v\n", negativeTime3)

	arrayOfInterfaces := []Time{}

	arrayOfInterfaces = append(arrayOfInterfaces, positiveTime1)
	arrayOfInterfaces = append(arrayOfInterfaces, positiveTime2)
	arrayOfInterfaces = append(arrayOfInterfaces, positiveTime3)
	arrayOfInterfaces = append(arrayOfInterfaces, negativeTime1)
	arrayOfInterfaces = append(arrayOfInterfaces, negativeTime2)
	arrayOfInterfaces = append(arrayOfInterfaces, negativeTime3)

	fmt.Println("\nArray length: ", len(arrayOfInterfaces), "\n")

	for _, time := range arrayOfInterfaces {
		if time.IsPositive() {
			fmt.Printf("Positive: %+v\n", time)
		} else {
			fmt.Printf("Negative: %+v\n", time)
		}
	}

	fmt.Println("\nclosing app")
}

//---GARBAGE---

func getTimes() (positiveTime, positiveTime, positiveTime, negativeTime, negativeTime, negativeTime) {
	positiveTime1 := positiveTime{
		duration: 8,
		creator:  "Franco",
	}

	positiveTime2 := positiveTime{
		duration: 10,
		creator:  "Lucas",
	}

	positiveTime3 := positiveTime{
		duration: 12,
		creator:  "Franco",
	}

	negativeTime1 := negativeTime{
		duration: 6,
		creator:  "Franco",
	}

	negativeTime2 := negativeTime{
		duration: 8,
		creator:  "Lucas",
	}

	negativeTime3 := negativeTime{
		duration: 6,
		creator:  "Franco",
	}

	return positiveTime1, positiveTime2, positiveTime3, negativeTime1, negativeTime2, negativeTime3
}
