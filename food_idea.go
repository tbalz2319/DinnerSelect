package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Recipes struct { //Struct for recipe information
	name             string
	prepTime         int
	cookTime         int
	recipeIngredient string
	ID               int
	Yield            int
}

func main() {
	var Recipe1 Recipes //Declare Recipe1 of Type Recipes
	var Recipe2 Recipes
	var Recipe3 Recipes

	/* Recipe1 specifications */
	Recipe1.name = "BBQ Pulled Chicken"
	Recipe1.prepTime = 25
	Recipe1.cookTime = 5
	Recipe1.recipeIngredient = "1 8-ounce can reduced-sodium tomato sauce, two"
	Recipe1.ID = 1
	Recipe1.Yield = 8

	/* Recipe 2 specifications */

	Recipe2.name = "Steak Tacos with Pineapple"
	Recipe2.prepTime = 45
	Recipe2.cookTime = 45
	Recipe2.recipeIngredient = "3 tablespoons soy sauce, 3"
	Recipe2.ID = 2
	Recipe2.Yield = 4

	Recipe3.name = "Simple Lemon Herb Chicken"
	Recipe3.prepTime = 10
	Recipe3.cookTime = 15
	Recipe3.recipeIngredient = "2 skinless, boneless chicken breast halves, 1 Lemon"
	Recipe3.ID = 3
	Recipe3.Yield = 2

	/* print 1 info */
	fmt.Printf(" Recipe name : %s\n", Recipe1.name)
	fmt.Printf(" Recipe prepTime : %d\n", Recipe1.prepTime)
	fmt.Printf(" Recipe cookTime: %d\n", Recipe1.cookTime)
	fmt.Printf(" Recipe recipeIngredient : %s\n", Recipe1.recipeIngredient)
	fmt.Printf(" Recipe ID : %d\n", Recipe1.ID)

	//call to printRecipe function below
	printRecipe(Recipe1)
	totalTime(Recipe1)
	printRecipe(Recipe2)
	totalTime(Recipe2)
	printRecipe(Recipe3)
	totalTime(Recipe3)

	//choose random number for recipe
	rand.Seed(time.Now().UTC().UnixNano())
	myrand := random(1, 6)
	fmt.Println(myrand)

	//logic for recipe to choose
	if myrand == 1 {
		fmt.Println(1)
		printRecipeOfTheDay(Recipe1)
	} else if myrand == 2 {
		fmt.Println(2)
		printRecipeOfTheDay(Recipe2)
	} else if myrand == 3 {
		fmt.Println(3)
	} else if myrand == 4 {
		fmt.Println(4)
	}
}

//function to print recipes
func printRecipe(Recipe Recipes) {
	fmt.Printf("Recipe Name : %s\n", Recipe.name)
	fmt.Printf("Prep Time : %d\n", Recipe.prepTime)
	fmt.Printf("Cook Time : %d\n", Recipe.cookTime)
	fmt.Printf("Recipe ID : %d\n", Recipe.ID)
}

func printRecipeOfTheDay(Recipe Recipes) {
	fmt.Printf("The recipe of the day is : %s\n", Recipe.name)
}

//Returns total time by addings cookTime and prepTime
func totalTime(Recipe Recipes) {
	fmt.Printf("The total time for this recipe is %d\n", Recipe.cookTime+Recipe.prepTime)
}
