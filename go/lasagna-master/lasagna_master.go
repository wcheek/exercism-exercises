package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, prepTimePerLayer int) int {
	if prepTimePerLayer == 0 {
		prepTimePerLayer = 2
	}
	return len(layers) * prepTimePerLayer
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodlesQuant int, sauceQuant float64) {
	numNoodles := 0
	numSauce := 0
	for i := 0; i < len(layers); i++ {
		switch layers[i] {
		case "noodles":
			numNoodles++
		case "sauce":
			numSauce++
		}
	}
	return numNoodles * 50, float64(numSauce) * 0.2
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList []string, myList []string) {
	secretIngredient := friendList[len(friendList)-1]
	myList[len(myList)-1] = secretIngredient
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amountsNeeded []float64, numPortions int) []float64 {
	scale := float64(numPortions) / 2
	newRecipe := make([]float64, len(amountsNeeded))

	for i := 0; i < len(newRecipe); i++ {
		newRecipe[i] = float64(scale) * amountsNeeded[i]
	}
	return newRecipe
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
