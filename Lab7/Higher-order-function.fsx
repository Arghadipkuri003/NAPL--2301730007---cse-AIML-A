let list1 = [1; 2; 3; 4; 5]
let list2 = [10; 15; 20; 25; 30]

let squares1 = List.map (fun x -> x * x) list1
let squares2 = List.map (fun x -> x * x) list2

printfn "List.map Results:"
printfn "Squares of list1: %A" squares1
printfn "Squares of list2: %A" squares2

let evens1 = List.filter (fun x -> x % 2 = 0) list1
let evens2 = List.filter (fun x -> x % 2 = 0) list2

printfn "\nList.filter Results:"
printfn "Even numbers in list1: %A" evens1
printfn "Even numbers in list2: %A" evens2

let sum1 = List.fold (fun acc x -> acc + x) 0 list1
let sum2 = List.fold (fun acc x -> acc + x) 0 list2

printfn "\nList.fold Results:"
printfn "Sum of list1: %d" sum1
printfn "Sum of list2: %d" sum2
