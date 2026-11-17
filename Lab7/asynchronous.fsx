
open System

let task1 = async {
    printfn "Task 1 started"
    do! Async.Sleep(2000)
    printfn "Task 1 completed"
    return "Result from Task 1"
}
let task2 = async {
    printfn "Task 2 started"
    do! Async.Sleep(1000)
    printfn "Task 2 completed"
    return "Result from Task 2"
}
let task3 = async {
    printfn "Task 3 started"

    let mutable sum = 0L
    for i in 1L .. 1000000L do
        sum <- sum + i

    printfn "Task 3 completed"
    return sprintf "Result from Task 3: %d" sum
}
let parallelWork = async {
    let! results =
        [task1; task2; task3]
        |> Async.Parallel

    printfn "\nAll tasks completed!"
    results |> Array.iter (printfn "%s")
}

Async.RunSynchronously parallelWork
