# pacmaneffect

A simple Go library implementing a functionality similar to [Python's Slice Notation](https://railsware.com/blog/python-for-machine-learning-indexing-and-slicing-for-lists-tuples-strings-and-other-sequential-types/).
The library is named after the [Pacman](https://en.wikipedia.org/wiki/Pac-Man) game, where the main character can exit the maze on one side and reappear on the other side. 
The Slice notation implements something similar, making it possible to select the elements at the end of the array using negative indexes.

The slice notation uses three params, called *start*, *end*, and *step*, to define the elements that must be retrieved from the original array. 
1. **start**: The first index to be considered by the slice. If only the start element is defined, then it is the index of the retrieved element.
1. **end**: The first index to be ignored by the slice. To be defined it requires at least a colon in the slice string. 
1. **step**: The step param is used to define the next element to be retrieved (if the current index is "i", the next will be "i + step"). If the step is 0, this library returns an empty array (while python returns an error).

The default values are defined as follows:
* The default value for *start* is 0 if step > 0, (length - 1) if step < 0.
* The default value for *end* is length if step > 0, -1 if step < 0.
* The default value for *step* is 1.

This library uses a *colon-separated string* to define the three params. The three params can be omitted by leaving their position empty; in this case, the default value will be used for the omitted param.
The slice notation returns a single element when only the *start* param is retrieved, while it returns a slice if at least the *end* parameter is available (also as default, e.g. if the string contains at least one colon character).

Here you can find some examples written in Python:
```python
arr = [1, 2, 3, 4, 5]
arr[2]       # 3
arr[-1]      # 5
arr[]        # Invalid sintax, there's no default while using only start
arr[2:3]     # [3]
arr[2:]      # [3, 4, 5]
arr[:2]      # [1, 2]
arr[:]       # [1, 2, 3, 4, 5]
arr[::-1]    # [5, 4, 3, 2, 1]
arr[::]      # [1, 2, 3, 4, 5]
arr[0:5:2]   # [1, 3, 5]
```

## setup
To use this library you need to import 'github.com/Valmoz/pacmaneffect'

## functionality
To use this library, you must define a **Pacman** instance as follows:
```go
arr := []int{1, 2, 3, 4, 5}
p, _ := NewPacman(arr)
```
Pacman accepts *interface{}* as input, so you can use every type of slice.
To retrieve the result, you need to provide an **Effect** as follows:
```go
e := NewEffect("0:5:2")
result, _ := p.Apply(e)
```

This library implements two main functionalities called *Apply* and *ApplyUnbounded*.

The *Apply* method implements the same behavior defined by Python's slice notation. The accepted range of accepted indexes is [-length, length - 1], where "length" is the length of the array.
If the provided parameters try to retrieve and index outside these bounds, the index is just ignored.
Only in the case of single element retrieval (e.g. a string without colons), if an index outside these borders is used as input for Apply then an error is returned.
For example:
```go
p, _ := NewPacman([]int{1, 2, 3, 4, 5})
result, _ := p.Apply(NewEffect("-10:10")) // result = []int{1, 2, 3, 4, 5}
_, err := p.Apply(NewEffect("-10"))       // An error is returned
```

The *ApplyUnbounded* method is slightly different because it accepts even indexes outside these borders: for every index, the result of *Modulus* (i % length) is used. 
In this way, the array becomes theoretically infinite on both sides. This feature also makes it possible to take every element of the original array multiple times, appending the array to itself.
Obviously, ApplyUnbounded doesn't return an error if *start* is outside borders in the single element retrieval.
For example:
```go
p, _ := NewPacman([]int{1, 2, 3, 4, 5})
result, _ := p.ApplyUnbounded(NewEffect("-10:10")) // result = []int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
result2, _ := p.Apply(NewEffect("-10"))            // result2 = 1
```

The two methods described above can be used with every type of slice. Since Go doesn't support Generics (yet), *reflect* was used to make it possible; this introduces a slight deterioration of the functionality performances.
If performance is a key factor for your application, then you can use one of the following methods designed for the basic slice types:
* []string: *ApplyString* and *ApplyUnboundedString*
* []int: *ApplyInt* and *ApplyUnboundedInt*
* []uint: *ApplyUint* and *ApplyUnboundedUint*
* []bool: *ApplyBool* and *ApplyUnboundedBool*
* []byte: *ApplyByte* and *ApplyUnboundedByte*
* []rune: *ApplyRune* and *ApplyUnboundedRune*
* []float32: *ApplyFloat32* and *ApplyFloat32*

These methods can be used exactly as described above. The result is always of type *interface{}*, to make it possible to return a single element or a slice based on the provided input.
For example:
```go
p, _ := NewPacman([]string{"1", "2", "3", "4", "5"})
result, _ := p.ApplyString(NewEffect("3:")) 
resultStr := result.([]string)              // resultStr = []string{"4", "5"}
result2, _ := p.ApplyString(NewEffect("2"))       
result2Str := result2.(string)              // result2Str = "2"
```

# author
Mauro Valota
