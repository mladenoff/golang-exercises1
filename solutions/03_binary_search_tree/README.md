## 03 Binary Search Tree

* Have we done this before? Make a struct treeNode that keeps track of a string, and a left and right child.
* Mind bend: a treeNode can't store literally other treeNodes in itself. What could it store though?
* Write a function `insert(root, value) that takes a pointer to a treeNode (the root) and a string value.
    * If the treeNode is `nil`, then this is an empty tree. Make a new treeNode with this   value.
    * If not, then look at `left` or `right` (as appropriate). If the value is nil, make    a new node for the value and *mutate* the root. Return that.
    * If the value is not nil, then recurse.
* Write a function called `traverse(root *treeNode, fn func(*treeNode)`.
    * Pass an anonymous function tht prints the values in order.
