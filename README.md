## 00 Hello World

* But your code in `go/src/github.com/[username]/exercises1/00_hello_world/main.go`.
* Remember the package must be main.
* Use `fmt` to print out a message to the user.
* Then use `bufio.Reader` and `os.Stdin` to read in a user's name.
* Use `strings` to strip the newline off.
* Greet them. Use `fmt.Printf` to learn how to use that.
* Build your code with `go build`.

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

## 02 Kitten Array

* Define a struct for a Kitten that includes a name string and an age int.
* Start with an empty slice of kittens.
* Prompt the user for a name and an age and add that Kitten to the slice.
* Use `bufio` and `fmt` and `os` as before. But also use `strconv` to convert from string to integer.
* To do this, write a `readCat` method that takes in a **pointer** to a `bufio.Reader`.
* After each cat is created, show the entire slice of cats.
* Also use `cap` and reslicing to show the entire backing store of cats too. Observe what happens as you push more.
* Now, before asking for input to create a cat, ask whether they want to push or pop or shift.

## 03 Binary Search Tree

* Have we done this before? Make a struct treeNode that keeps track of a string, and a left and right child.
* Mind bend: a treeNode can't store literally other treeNodes in itself. What could it store though?
* Write a function `insert(root, value) that takes a pointer to a treeNode (the root) and a string value.
    * If the treeNode is `nil`, then this is an empty tree. Make a new treeNode with this   value.
    * If not, then look at `left` or `right` (as appropriate). If the value is nil, make    a new node for the value and *mutate* the root. Return that.
    * If the value is not nil, then recurse.
* Write a function called `traverse(root *treeNode, fn func(*treeNode)`.
    * Pass an anonymous function tht prints the values in order.