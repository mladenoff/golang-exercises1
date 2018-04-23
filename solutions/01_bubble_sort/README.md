## 01 Bubble Sort

* Write a bubbleSort method that takes in a slice of ints. It should sort in place (doesn't need a pointer).
* You may want to use the `for idx, val := range nums` syntax. If you don't know about that, look it up.
* You can use multiple asignment for sawpping.
* Create a literal slice of five or so numbers in jumbled order and sort it.
* After this, write a function to generate a bunch of random ints and put them in the array.
* Check out `math/rand`. Maybe generate ints in the range 0 to 1000.
* Now write a `swap` function that takes in two numbers and swaps them in place. Do
  not return anything from `swap`; use pointers.
* Test this function out with two local variables in your main function.
* Then use this in your `bubbleSort` method. Take pointers to the locations in the
  slice. Note that this is mildly crazy, but we're just experimenting.
