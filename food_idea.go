package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Struct for Recipe below
type Recipe struct { //Struct for recipe information
	name        string
	prepTime    int
	cookTime    int
	Ingredients []string //this is now a slice that will accept multiple elements
	ID          int
	Yield       int
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

func main() {
	var recipe1 Recipe //Declare recipe1 of Type Recipe
	var recipe2 Recipe
	var recipe3 Recipe

	//choose random number for recipe
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	i := r.Perm(5)
	fmt.Printf("%v\n", i)
	fmt.Printf("%d\n", i[0])

	//assign slices of int from Perm to variables assigned days of the week
	var monday = i[0]
	var tuesday = i[1]
	var wednesday = i[2]
	var thursday = i[3]
	var friday = i[4]

	//testing printing of variables assigned to days
	fmt.Printf("This is for the day Monday %d\n", monday)
	fmt.Printf("%d\n", tuesday)
	fmt.Printf("%d\n", wednesday)
	fmt.Printf("%d\n", thursday)
	fmt.Printf("%d\n", friday)

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
	}
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

	//call to printRecipe function below
	printRecipe(recipe1)
	totalTime(recipe1)
	printRecipe(recipe2)
	totalTime(recipe2)
	printRecipe(recipe3)
	totalTime(recipe3)
}

//function to print the winner for recipe of the day to use
//for either lunch or dinner
func printRecipeOfTheDay(recipe Recipe) {
	fmt.Printf("The recipe of the day is : %s\n", recipe.name)
}
