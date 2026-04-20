package lasagnamaster

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, timePerLayer int) int {
	if timePerLayer == 0 {
		timePerLayer = 2
	}
	return len(layers) * timePerLayer
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	noodles := 0
	sauce := 0
	for i := 0; i < len(layers); i++ {
		if layers[i] == "noodles" {
			noodles++
		}
		if layers[i] == "sauce" {
			sauce++
		}
	}
	return noodles * 50, float64(sauce) * 0.2
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList []string, myList []string) {
	secretSauce := friendList[len(friendList)-1]
	myList[len(myList)-1] = secretSauce
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amounts []float64, portions int) []float64 {
	scaledQuantities := []float64{}
	for i := 0; i < len(amounts); i++ {
		scaledQuantities = append(scaledQuantities, amounts[i]*float64(portions)/2)
	}
	return scaledQuantities
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
