// Base class
[<AbstractClass>]
type Shape() =
    abstract member Area: unit -> float

// Derived class: Circle
type Circle(radius: float) =
    inherit Shape()
    override this.Area() = 3.14 * radius * radius

// Derived class: Rectangle
type Rectangle(length: float, width: float) =
    inherit Shape()
    override this.Area() = length * width

let c = Circle(5.0)
let r = Rectangle(4.0, 6.0)

printfn "Area of Circle = %.2f" (c.Area())
printfn "Area of Rectangle = %.2f" (r.Area())

let shapes: Shape list = [c; r]
shapes |> List.iter (fun s -> printfn "Area = %.2f" (s.Area()))
