open System

[<EntryPoint>]
let main argv =
    printfn "F# CRUD Application — Student Management"
    printfn "Choose database:"
    printfn "  1. SQLite"
    printfn "  2. MongoDB"
    printfn "  3. Both"
    printf "Enter choice (1/2/3): "

    let choice = Console.ReadLine().Trim()

    match choice with
    | "1" ->
        SQLiteCRUD.run ()
    | "2" ->
        MongoCRUD.run ()
    | "3" ->
        SQLiteCRUD.run ()
        printfn ""
        MongoCRUD.run ()
    | _ ->
        printfn "Invalid choice. Running SQLite demo..."
        SQLiteCRUD.run ()

    0
