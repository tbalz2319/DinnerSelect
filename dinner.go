package main

import (
	"fmt"
	"io"
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
	//10 variables below for 5 recipes for Monday-Friday week 1 and 5 for week 2
	var recipe1, recipe2, recipe3, recipe4, recipe5, recipe6, recipe7, recipe8, recipe9, recipe10 Recipe //Declare recipe1 of Type Recipe

	//Test of reading file and outputting array
	file, err := os.Open("file")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var perline int
	var nums []int

	for {
		_, err := fmt.Fscanf(file, "%d\n", &perline) // give a patter to scan

		if err != nil {
			if err == io.EOF {
				break // stop reading the file
			}
			fmt.Println(err)
			os.Exit(1)
		}
		nums = append(nums, perline)
	}

	// print out the nums array content
	fmt.Println(nums)
	fmt.Printf("This is the nums line: %d\n", nums)
	fmt.Printf("This is the first number of nums array: %d\n", nums[0])
	fmt.Printf("This is the second number of the nums array: %d\n", nums[1])
	//Test function for printing variables for week 2
	//using the variable k because r has been taken in
	//another part of the program

	if nums[0] >= 6 {
		fmt.Printf("Looks like it is time for week2 because of %d\n", nums[0])
		//Kick off week1 sequence
		//choose random number for recipe
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		i := r.Perm(5)
		fmt.Printf("%v\n", i)
		fmt.Printf("%d\n", i[0])
		fmt.Printf("%d\n", i[1])
		fmt.Printf("%d\n", i[2])

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

		//REcipes have to go before call to print recipe of the day function
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

		//this creates "file" which identifies the recipes that were picked
		file, fileErr := os.Create("file")
		if fileErr != nil {
			fmt.Println(fileErr)
			return
		}
		fmt.Fprintf(file, "%d\n", monday)
		fmt.Fprintf(file, "%d\n", tuesday)
		fmt.Fprintf(file, "%d\n", wednesday)
		fmt.Fprintf(file, "%d\n", thursday)
		fmt.Fprintf(file, "%d\n", friday)

	} else if nums[0] < 6 {
		fmt.Printf("Looks like it is time for week2 %d\n", nums[0])
		//Kick off random number sequence for 6-10 below
		k := rand.New(rand.NewSource(time.Now().UnixNano()))
		min, max := 6, 10
		p := k.Perm(max - min + 1)
		//fmt.Println(p)
		for i := range p {
			p[i] += min
		}
		//below has been take out in order to not print week 1
		//numbers
		//fmt.Println(p)

		for _, k := range p {
			fmt.Println(k)
		}

		//assign slices of int from Perm to variables assigned days of the week
		var monday = p[0]
		var tuesday = p[1]
		var wednesday = p[2]
		var thursday = p[3]
		var friday = p[4]

		//testing printing of variables assigned to days
		fmt.Printf("This is for the day Monday %d\n", monday)
		fmt.Printf("This is for the day Tuesday %d\n", tuesday)
		fmt.Printf("This is for the day Wednesday %d\n", wednesday)
		fmt.Printf("This is for the day Thursday %d\n", thursday)
		fmt.Printf("This is for the day Friday %d\n", friday)

		fmt.Fprintf(file, "%d\n", monday)
		fmt.Fprintf(file, "%d\n", tuesday)
		fmt.Fprintf(file, "%d\n", wednesday)
		fmt.Fprintf(file, "%d\n", thursday)
		fmt.Fprintf(file, "%d\n", friday)

		//REcipes have to go before call to print recipe of the day function
		/* recipe6 specifications */
		recipe6.name = "BBQ Pulled Chicken level 6"
		recipe6.prepTime = 25
		recipe6.cookTime = 5
		recipe6.Ingredients = append(
			recipe6.Ingredients,
			"1 8-ounce can reduced-sodium tomato sauce",
		)
		recipe6.Ingredients = append(
			recipe6.Ingredients,
			"1/2 medium onion (grated),",
		)
		recipe6.ID = 6
		recipe6.Yield = 8

		/* Recipe 7 specifications */

		recipe7.name = "Steak Tacos with Pineapple Level 7"
		recipe7.prepTime = 45
		recipe7.cookTime = 45
		recipe7.Ingredients = append(
			recipe7.Ingredients,
			"3 tablespoons soy sauce,",
		)
		recipe7.Ingredients = append(
			recipe7.Ingredients,
			"1 tablespoon finely grated garlic,",
		)
		recipe7.Ingredients = append(
			recipe7.Ingredients,
			"1 tablespoon finely grated peeled fresh ginger,",
		)
		recipe7.Ingredients = append(
			recipe7.Ingredients,
			"1 1/2 pounds skirt steak, cut into 5-inch lengths,",
		)
		recipe7.Ingredients = append(
			recipe7.Ingredients,
			"Salt",
		)
		recipe7.Ingredients = append(
			recipe7.Ingredients,
			"Pepper",
		)
		recipe7.ID = 7
		recipe7.Yield = 4

		recipe8.name = "Simple Lemon Herb Chicken level 8"
		recipe8.prepTime = 10
		recipe8.cookTime = 15
		recipe8.Ingredients = append(
			recipe8.Ingredients,
			"2 skinless boneless chicken breast halves,",
		)
		recipe8.Ingredients = append(
			recipe8.Ingredients,
			"1 Lemon,",
		)
		recipe8.Ingredients = append(
			recipe8.Ingredients,
			"Salt and Pepper to taste,",
		)
		recipe8.Ingredients = append(
			recipe8.Ingredients,
			"1 tablespoon olive oil,",
		)
		recipe8.Ingredients = append(
			recipe8.Ingredients,
			"2 sprigs fresh parsley (for garnish),",
		)
		recipe8.Ingredients = append(
			recipe8.Ingredients,
			"1 pinch dried oregano,",
		)
		recipe8.ID = 8
		recipe8.Yield = 2

		//recipe4 specifications
		recipe9.name = "Easy Meatloaf lvl 9"
		recipe9.prepTime = 10
		recipe9.cookTime = 60
		recipe9.Ingredients = append(
			recipe9.Ingredients,
			"1 onion (chopped),",
		)
		recipe9.Ingredients = append(
			recipe9.Ingredients,
			"1 cup milk,",
		)
		recipe9.Ingredients = append(
			recipe9.Ingredients,
			"1 cup dried bread crumbs,",
		)
		recipe9.ID = 9
		recipe9.Yield = 8

		//recipe 10 specifications
		recipe10.name = "Fast Salmon with a Ginger Glaze lvl 10"
		recipe10.prepTime = 5
		recipe10.cookTime = 20
		recipe10.Ingredients = append(
			recipe10.Ingredients,
			"salt to taste,",
		)
		recipe10.Ingredients = append(
			recipe10.Ingredients,
			"1/3 cup cold water,",
		)
		recipe10.Ingredients = append(
			recipe10.Ingredients,
			"1/4 cup seasoned rice vinegar,",
		)
		recipe10.ID = 10
		recipe10.Yield = 4

		//logic for Mondays Recipe week 2 algorithm
		if monday == 6 {
			fmt.Println(0)
			printRecipeOfTheDay(recipe6)
		} else if monday == 7 {
			fmt.Println(7)
			printRecipeOfTheDay(recipe7)
		} else if monday == 8 {
			fmt.Println(8)
			printRecipeOfTheDay(recipe8)
		} else if monday == 9 {
			fmt.Println(9)
			printRecipeOfTheDay(recipe9)
		} else if monday == 10 {
			fmt.Println(10)
			printRecipeOfTheDay(recipe10)
		}

		if tuesday == 6 {
			fmt.Println(0)
			printRecipeOfTheDay(recipe6)
		} else if tuesday == 7 {
			fmt.Println(7)
			printRecipeOfTheDay(recipe7)
		} else if tuesday == 8 {
			fmt.Println(8)
			printRecipeOfTheDay(recipe8)
		} else if tuesday == 9 {
			fmt.Println(9)
			printRecipeOfTheDay(recipe9)
		} else if tuesday == 10 {
			fmt.Println(10)
			printRecipeOfTheDay(recipe10)
		}

		if wednesday == 6 {
			fmt.Println(0)
			printRecipeOfTheDay(recipe6)
		} else if wednesday == 7 {
			fmt.Println(7)
			printRecipeOfTheDay(recipe7)
		} else if wednesday == 8 {
			fmt.Println(8)
			printRecipeOfTheDay(recipe8)
		} else if wednesday == 9 {
			fmt.Println(9)
			printRecipeOfTheDay(recipe9)
		} else if wednesday == 10 {
			fmt.Println(10)
			printRecipeOfTheDay(recipe10)
		}

		if thursday == 6 {
			fmt.Println(0)
			printRecipeOfTheDay(recipe6)
		} else if thursday == 7 {
			fmt.Println(7)
			printRecipeOfTheDay(recipe7)
		} else if thursday == 8 {
			fmt.Println(8)
			printRecipeOfTheDay(recipe8)
		} else if thursday == 9 {
			fmt.Println(9)
			printRecipeOfTheDay(recipe9)
		} else if thursday == 10 {
			fmt.Println(10)
			printRecipeOfTheDay(recipe10)
		}

		if friday == 6 {
			fmt.Println(0)
			printRecipeOfTheDay(recipe6)
		} else if friday == 7 {
			fmt.Println(7)
			printRecipeOfTheDay(recipe7)
		} else if friday == 8 {
			fmt.Println(8)
			printRecipeOfTheDay(recipe8)
		} else if friday == 9 {
			fmt.Println(9)
			printRecipeOfTheDay(recipe9)
		} else if friday == 10 {
			fmt.Println(10)
			printRecipeOfTheDay(recipe10)
		}
		//this creates "file" which identifies the recipes that were picked
		file, fileErr := os.Create("file")
		if fileErr != nil {
			fmt.Println(fileErr)
			return
		}
		fmt.Fprintf(file, "%d\n", monday)
		fmt.Fprintf(file, "%d\n", tuesday)
		fmt.Fprintf(file, "%d\n", wednesday)
		fmt.Fprintf(file, "%d\n", thursday)
		fmt.Fprintf(file, "%d\n", friday)

	}
	//file output testing for persistence
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
