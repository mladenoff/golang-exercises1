## 02 Kitten Array

* Define a struct for a Kitten that includes a name string and an age int.
* Start with an empty slice of kittens.
* Prompt the user for a name and an age and add that Kitten to the slice.
* Use `bufio` and `fmt` and `os` as before. But also use `strconv` to convert from string to integer.
* To do this, write a `readCat` method that takes in a **pointer** to a `bufio.Reader`.
* After each cat is created, show the entire slice of cats.
* Also use `cap` and reslicing to show the entire backing store of cats too. Observe what happens as you push more.
* Now, before asking for input to create a cat, ask whether they want to push or pop or shift.
