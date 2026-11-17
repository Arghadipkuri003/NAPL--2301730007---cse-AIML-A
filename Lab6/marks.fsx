open System

printf "Enter marks for 4 subjects:\n"

let marks =
    [ for i in 1..4 ->
        printf "Subject %d: " i
        int (Console.ReadLine()) ]

let total = List.sum marks
let percentage = float total / 4.0

let division =
    match percentage with
    | p when p >= 60.0 -> "First Division"
    | p when p >= 50.0 -> "Second Division"
    | p when p >= 40.0 -> "Third Division"
    | _ -> "Fail"

printfn "Total Marks = %d" total
printfn "Percentage = %.2f%%" percentage
printfn "Division = %s" division

