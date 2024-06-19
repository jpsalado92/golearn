# [Go Class: 12 Structs, Struct tags & JSON](https://www.youtube.com/watch?v=0m6iFd9N_CY&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6&index=14)

Declaring a struc type
Inspectinc a struct 
Declaring and initializing a struct var
    Literal whole without fnames
    Literal whole with fnames
    Literal partial by fnames
    Field by field

Struct vars refering to other structs
    Nested structs

Gotchas
You cannot take the address of a map value

Maps of strings to struct pointers vs maps of strings to structs
    Maps of strings to struct pointers
    Maps of strings to structs

Anonymous structs
    Anonymous struct fields
    Anonymous struct fields with same name
    Anonymous struct fields with different names


Struct tags

Structural compatibility
    Two `struct` types are compatible if:
    1. The fields have the same types and names.
    2. The fields are in the same order.
    3. Both have the same tags.
    A struct is comparable if all its fields are comparable
    Named struct types are *convertible* if they are compatible
    Copy works if fields are same, even if names are different and anonymous structs are involved
    Copy does not work if fields are same, but declared struct types
    Copy does work if fields are same, but declared struct types, when casting.

Other struct stuff
The zero value of a struct is "zero" for each field.

Gotcha! Passing down structs as parameters vs pointers
    Passing down structs as parameters, they are passed as copies.
    Passing down structs as pointers, that is what you intend generally.

Singleton empty struct


Tags

Struct to json
Struct lowercase not exported