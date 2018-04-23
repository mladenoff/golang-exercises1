## Make Change

Solve the make change problem. You can declare a list of coins as a
global variable:

    var coins = []{25, 10, 5, 1}

My `makeChange` function had this signature:

    makeChange(amount int, purse []int, minIdx int) []int

The `purse` contained the set of coins used thus far. `minIdx` says that
I will use only coins at index `>= minIdx`. In the initial call, I pass
an empty slice as `purse` and `minIdx = 0`.

The return value is either the set of coins to make the change, or nil.

Benchmark your code against the Ruby code!
