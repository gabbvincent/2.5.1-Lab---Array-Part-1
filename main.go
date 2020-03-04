// Programmer name: Vincent G.
// Date completed:  3/3/2020
//Description: 2.5.1 Lab - Array Part 1

package main

import "fmt"

func main() {
    //Create a string array of at least 5 values.  It should hold 5 products to buy

    var products [5]string
    products[0] = "Fries"
    products[1] = "Hotdog"
    products[2] = "Soda"
    products[3] = "Burger"
    products[4] = "Icecream"

    //Create a float array of at least 5 values.  It should hold a price for each of the products

    var prices [5]float64
    prices[0] = 2.49
    prices[1] = 3.21
    prices[2] = 1.07
    prices[3] = 4.28
    prices[4] = 1.57

    //Iterate through the array and output the product and it's price (similar to a menu)

    for i :=0; i < 5; i++ {
      fmt.Println(products[i] ,"=", prices[i])
    }
}