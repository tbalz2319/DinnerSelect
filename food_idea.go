package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

//Recipe type
type Recipe struct { //Struct for recipe information
	name        string
	prepTime    int
	cookTime    int
	Ingredients []string //this is now a slice that will accept multiple elements
	ID          int
	Yield       int
}

//main method
func main() {
	//5 variables below for 5 recipes for Monday-Friday
	var recipe1 Recipe //Declare recipe1 of Type Recipe
	var recipe2 Recipe
	var recipe3 Recipe
	var recipe4 Recipe
	var recipe5 Recipe

	//choose random number for recipe
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	i := r.Perm(5)
	fmt.Printf("%v\n", i)
	fmt.Printf("%d\n", i[0])
	fmt.Printf("%d\n", i[1])

	//assign slices of int from Perm to variables assigned days of the week
	var monday = i[0]
	var tuesday = i[1]
	var wednesday = i[2]
	var thursday = i[3]
	var friday = i[4]

	//testing printing of variables assigned to days
	fmt.Printf("This is for the day Monday %d\n", monday)
	fmt.Printf("This is for the day Tuesday %d\n", tuesday)
	fmt.Printf("This is for the day Wednesday %d\n", wednesday)
	fmt.Printf("This is for the day Thursday %d\n", thursday)
	fmt.Printf("This is for the day Friday %d\n", friday)

	/* recipe1 specifications */
	recipe1.name = "BBQ Pulled Chicken"
	recipe1.prepTime = 25
	recipe1.cookTime = 5
	recipe1.Ingredients = append(
		recipe1.Ingredients,
		"1 8-ounce can reduced-sodium tomato sauce",
	)
	recipe1.Ingredients = append(
		recipe1.Ingredients,
		"1/2 medium onion (grated),",
	)
	recipe1.ID = 1
	recipe1.Yield = 8

	/* Recipe 2 specifications */

	recipe2.name = "Steak Tacos with Pineapple"
	recipe2.prepTime = 45
	recipe2.cookTime = 45
	recipe2.Ingredients = append(
		recipe2.Ingredients,
		"3 tablespoons soy sauce,",
	)
	recipe2.Ingredients = append(
		recipe2.Ingredients,
		"1 tablespoon finely grated garlic,",
	)
	recipe2.Ingredients = append(
		recipe2.Ingredients,
		"1 tablespoon finely grated peeled fresh ginger,",
	)
	recipe2.Ingredients = append(
		recipe2.Ingredients,
		"1 1/2 pounds skirt steak, cut into 5-inch lengths,",
	)
	recipe2.Ingredients = append(
		recipe2.Ingredients,
		"Salt",
	)
	recipe2.Ingredients = append(
		recipe2.Ingredients,
		"Pepper",
	)
	recipe2.ID = 2
	recipe2.Yield = 4

	recipe3.name = "Simple Lemon Herb Chicken"
	recipe3.prepTime = 10
	recipe3.cookTime = 15
	recipe3.Ingredients = append(
		recipe3.Ingredients,
		"2 skinless boneless chicken breast halves,",
	)
	recipe3.Ingredients = append(
		recipe3.Ingredients,
		"1 Lemon,",
	)
	recipe3.Ingredients = append(
		recipe3.Ingredients,
		"Salt and Pepper to taste,",
	)
	recipe3.Ingredients = append(
		recipe3.Ingredients,
		"1 tablespoon olive oil,",
	)
	recipe3.Ingredients = append(
		recipe3.Ingredients,
		"2 sprigs fresh parsley (for garnish),",
	)
	recipe3.Ingredients = append(
		recipe3.Ingredients,
		"1 pinch dried oregano,",
	)
	recipe3.ID = 3
	recipe3.Yield = 2

	//recipe4 specifications
	recipe4.name = "Easy Meatloaf"
	recipe4.prepTime = 10
	recipe4.cookTime = 60
	recipe4.Ingredients = append(
		recipe4.Ingredients,
		"1 onion (chopped),",
	)
	recipe4.Ingredients = append(
		recipe4.Ingredients,
		"1 cup milk,",
	)
	recipe4.Ingredients = append(
		recipe4.Ingredients,
		"1 cup dried bread crumbs,",
	)
	recipe4.ID = 4
	recipe4.Yield = 8

	//recipe 5 specifications
	recipe5.name = "Fast Salmon with a Ginger Glaze"
	recipe5.prepTime = 5
	recipe5.cookTime = 20
	recipe5.Ingredients = append(
		recipe5.Ingredients,
		"salt to taste,",
	)
	recipe5.Ingredients = append(
		recipe5.Ingredients,
		"1/3 cup cold water,",
	)
	recipe5.Ingredients = append(
		recipe5.Ingredients,
		"1/4 cup seasoned rice vinegar,",
	)
	recipe5.ID = 5
	recipe5.Yield = 4

	//call to printRecipe function below
	printRecipe(recipe1)
	totalTime(recipe1)
	printRecipe(recipe2)
	totalTime(recipe2)
	printRecipe(recipe3)
	totalTime(recipe3)
	printRecipe(recipe4)
	totalTime(recipe4)
	printRecipe(recipe5)
	totalTime(recipe5)

	//logic for Mondays Recipe
	if monday == 0 {
		fmt.Println(0)
		printRecipeOfTheDay(recipe1)
	} else if monday == 1 {
		fmt.Println(1)
		printRecipeOfTheDay(recipe2)
	} else if monday == 2 {
		fmt.Println(2)
		printRecipeOfTheDay(recipe3)
	} else if monday == 3 {
		fmt.Println(3)
		printRecipeOfTheDay(recipe4)
	} else if monday == 4 {
		fmt.Println(4)
		printRecipeOfTheDay(recipe5)
	}

	//logic for Tuesdays Recipe
	if tuesday == 0 {
		fmt.Println(0)
		printRecipeOfTheDay(recipe1)
	} else if tuesday == 1 {
		fmt.Println(1)
		printRecipeOfTheDay(recipe2)
	} else if tuesday == 2 {
		fmt.Println(2)
		printRecipeOfTheDay(recipe3)
	} else if tuesday == 3 {
		fmt.Println(3)
		printRecipeOfTheDay(recipe4)
	} else if tuesday == 4 {
		fmt.Println(4)
		printRecipeOfTheDay(recipe5)
	}

	if wednesday == 0 {
		fmt.Println(0)
		printRecipeOfTheDay(recipe1)
	} else if wednesday == 1 {
		fmt.Println(1)
		printRecipeOfTheDay(recipe2)
	} else if wednesday == 2 {
		fmt.Println(2)
		printRecipeOfTheDay(recipe3)
	} else if wednesday == 3 {
		fmt.Println(3)
		printRecipeOfTheDay(recipe4)
	} else if wednesday == 4 {
		fmt.Println(4)
		printRecipeOfTheDay(recipe5)
	}

	if thursday == 0 {
		fmt.Println(0)
		printRecipeOfTheDay(recipe1)
	} else if thursday == 1 {
		fmt.Println(1)
		printRecipeOfTheDay(recipe2)
	} else if thursday == 2 {
		fmt.Println(2)
		printRecipeOfTheDay(recipe3)
	} else if thursday == 3 {
		fmt.Println(3)
		printRecipeOfTheDay(recipe4)
	} else if thursday == 4 {
		fmt.Println(4)
		printRecipeOfTheDay(recipe5)
	}

	if friday == 0 {
		fmt.Println(0)
		printRecipeOfTheDay(recipe1)
	} else if friday == 1 {
		fmt.Println(1)
		printRecipeOfTheDay(recipe2)
	} else if friday == 2 {
		fmt.Println(2)
		printRecipeOfTheDay(recipe3)
	} else if friday == 3 {
		fmt.Println(3)
		printRecipeOfTheDay(recipe4)
	} else if friday == 4 {
		fmt.Println(4)
		printRecipeOfTheDay(recipe5)
	}

	//file output testing for persistence

	file, fileErr := os.Create("file")
	if fileErr != nil {
		fmt.Println(fileErr)
		return
	}
	//fmt.Fprintf(file, "%v\n", i)
	//The above line was a test to file to a file and is not needed
	fmt.Fprintf(file, "%d\n", monday)
	fmt.Fprintf(file, "%d\n", tuesday)
	fmt.Fprintf(file, "%d\n", wednesday)
	fmt.Fprintf(file, "%d\n", thursday)
	fmt.Fprintf(file, "%d\n", friday)
}

//function to print Recipe
func printRecipe(recipe Recipe) {
	fmt.Printf("Recipe Name : %s\n", recipe.name)
	fmt.Printf("Prep Time : %d\n", recipe.prepTime)
	fmt.Printf("Cook Time : %d\n", recipe.cookTime)
	fmt.Printf("Ingredients : %s\n", recipe.Ingredients)
	fmt.Printf("Recipe ID : %d\n", recipe.ID)
}

//Returns total time by addings cookTime and prepTime
func totalTime(recipe Recipe) {
	fmt.Printf("The total time for this recipe is %d\n", recipe.cookTime+recipe.prepTime)
}

//function to print the winner for recipe of the day to use
//for either lunch or dinner
func printRecipeOfTheDay(recipe Recipe) {
	fmt.Printf("The recipe of the day is : %s\n", recipe.name)
}
