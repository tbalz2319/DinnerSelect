package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Recipe struct { //Struct for recipe information
	name        string
	prepTime    int
	cookTime    int
	Ingredients []string //this is now a slice that will accept multiple elements
	ID          int
	Yield       int
}

func main() {
	var recipe1 Recipe //Declare recipe1 of Type Recipe
	var recipe2 Recipe
	var recipe3 Recipe

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

	/* print 1 info */
	fmt.Printf(" Recipe name : %s\n", recipe1.name)
	fmt.Printf(" Recipe prepTime : %d\n", recipe1.prepTime)
	fmt.Printf(" Recipe cookTime: %d\n", recipe1.cookTime)
	fmt.Printf(" Recipe Ingredients : %s\n", recipe1.Ingredients)
	fmt.Printf(" Recipe ID : %d\n", recipe1.ID)

	//call to printRecipe function below
	printRecipe(recipe1)
	totalTime(recipe1)
	printRecipe(recipe2)
	totalTime(recipe2)
	printRecipe(recipe3)
	totalTime(recipe3)

	//choose random number for recipe
	rand.Seed(time.Now().UTC().UnixNano())
	myrand := random(1, 6)
	fmt.Println(myrand)

	//logic for recipe to choose
	if myrand == 1 {
		fmt.Println(1)
		printRecipeOfTheDay(recipe1)
	} else if myrand == 2 {
		fmt.Println(2)
		printRecipeOfTheDay(recipe2)
	} else if myrand == 3 {
		fmt.Println(3)
		printRecipeOfTheDay(recipe3)
	} else if myrand == 4 {
		fmt.Println(4)
	}
}

//function to print Recipe
func printRecipe(recipe Recipe) {
	fmt.Printf("Recipe Name : %s\n", recipe.name)
	fmt.Printf("Prep Time : %d\n", recipe.prepTime)
	fmt.Printf("Cook Time : %d\n", recipe.cookTime)
	fmt.Printf("Ingredients : %s\n", recipe.Ingredients)
	fmt.Printf("Recipe ID : %d\n", recipe.ID)
}

//random number function
func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func printRecipeOfTheDay(recipe Recipe) {
	fmt.Printf("The recipe of the day is : %s\n", recipe.name)
}

//Returns total time by addings cookTime and prepTime
func totalTime(recipe Recipe) {
	fmt.Printf("The total time for this recipe is %d\n", recipe.cookTime+recipe.prepTime)
}
