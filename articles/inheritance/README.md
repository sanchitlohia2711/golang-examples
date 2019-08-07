Inheritance in GO

Go supports inheritance by embedding. It can be done two ways

1. Inheritance using Struct - in this approach the base struct is embededded in child struct and base properties and methods can directly be called on child struct
 

   Limitation: Subtyping is not supported. You can pass the child struct to a function which expects base. For example below code will not work

2. Inheritance using Interface - This approach also solves the problem of subtyping. See below code




