open System
type Department =
    | CSE
    | ECE
    | ME
    | CE
type Grade =
    | Distinction
    | FirstClass
    | SecondClass
    | ThirdClass
    | Fail
type Student =
    { Id: int
      Name: string
      Age: int
      Department: Department
      Marks: float list }
type StudentManager() =
    let mutable students: Student list = []

    member this.AddStudent(student: Student) =
        students <- student :: students

    member this.GetAll() = students

    member this.RemoveById(id: int) =
        students <- students |> List.filter (fun s -> s.Id <> id)

    member this.FindById(id: int) =
        students |> List.tryFind (fun s -> s.Id = id)

    member this.Count = students.Length

let deptToString dept =
    match dept with
    | CSE -> "CSE (Computer Science)"
    | ECE -> "ECE (Electronics)"
    | ME  -> "ME (Mechanical)"
    | CE  -> "CE (Civil)"

let average marks =
    if List.isEmpty marks then 0.0
    else marks |> List.sum |> fun total -> total / float (List.length marks)
let classifyGrade avg =
    match avg with
    | a when a >= 75.0 -> Distinction
    | a when a >= 60.0 -> FirstClass
    | a when a >= 50.0 -> SecondClass
    | a when a >= 40.0 -> ThirdClass
    | _ -> Fail

let gradeToString grade =
    match grade with
    | Distinction -> "Distinction (≥75%)"
    | FirstClass  -> "First Class (≥60%)"
    | SecondClass -> "Second Class (≥50%)"
    | ThirdClass  -> "Third Class (≥40%)"
    | Fail        -> "Fail (<40%)"

let printStudent (s: Student) =
    let avg = average s.Marks
    let grade = classifyGrade avg
    printfn "  ID=%-3d | %-12s | Age=%-2d | %-26s | Avg=%.1f | %s"
        s.Id s.Name s.Age (deptToString s.Department) avg (gradeToString grade)

printfn "   F# Student Management System"

let manager = StudentManager()

let sampleStudents =
    [ { Id = 1; Name = "Rahul";  Age = 21; Department = CSE; Marks = [85.0; 78.0; 92.0; 88.0; 76.0] }
      { Id = 2; Name = "Priya";  Age = 20; Department = ECE; Marks = [62.0; 55.0; 70.0; 65.0; 58.0] }
      { Id = 3; Name = "Aman";   Age = 22; Department = ME;  Marks = [45.0; 38.0; 52.0; 41.0; 44.0] }
      { Id = 4; Name = "Sneha";  Age = 21; Department = CSE; Marks = [90.0; 95.0; 88.0; 92.0; 85.0] }
      { Id = 5; Name = "Vikram"; Age = 23; Department = CE;  Marks = [30.0; 25.0; 35.0; 28.0; 32.0] }
      { Id = 6; Name = "Neha";   Age = 20; Department = ECE; Marks = [72.0; 68.0; 75.0; 80.0; 65.0] }
      { Id = 7; Name = "Rohan";  Age = 22; Department = ME;  Marks = [55.0; 48.0; 60.0; 52.0; 50.0] } ]

sampleStudents |> List.iter manager.AddStudent
printfn "  Added %d students" manager.Count
printfn "\n[ALL STUDENTS]"
manager.GetAll()
|> List.sortBy (fun s -> s.Id)
|> List.iter printStudent

printfn "\n[FILTER] Students with Distinction (avg ≥ 75%%):"
manager.GetAll()
|> List.filter (fun s -> average s.Marks >= 75.0)
|> List.sortBy (fun s -> s.Id)
|> List.iter printStudent

printfn "\n[FILTER] Students who Failed (avg < 40%%):"
manager.GetAll()
|> List.filter (fun s -> average s.Marks < 40.0)
|> List.sortBy (fun s -> s.Id)
|> List.iter printStudent

printfn "\n[FILTER] CSE Department students:"
manager.GetAll()
|> List.filter (fun s -> s.Department = CSE)
|> List.sortBy (fun s -> s.Id)
|> List.iter printStudent

printfn "\n[MAP] Student Name → Grade mapping:"
manager.GetAll()
|> List.sortBy (fun s -> s.Id)
|> List.map (fun s -> (s.Name, classifyGrade (average s.Marks)))
|> List.iter (fun (name, grade) ->
    printfn "  %-12s → %s" name (gradeToString grade))

printfn "\n[MAP] Per-student statistics:"
manager.GetAll()
|> List.sortBy (fun s -> s.Id)
|> List.map (fun s ->
    let avg = average s.Marks
    let highest = List.max s.Marks
    let lowest = List.min s.Marks
    (s.Name, avg, highest, lowest))
|> List.iter (fun (name, avg, high, low) ->
    printfn "  %-12s | Avg=%.1f | Highest=%.0f | Lowest=%.0f" name avg high low)

printfn "\n[SUMMARY] Department-wise student count:"
[CSE; ECE; ME; CE]
|> List.iter (fun dept ->
    let count =
        manager.GetAll()
        |> List.filter (fun s -> s.Department = dept)
        |> List.length
    let deptAvg =
        manager.GetAll()
        |> List.filter (fun s -> s.Department = dept)
        |> List.map (fun s -> average s.Marks)
        |> fun avgs ->
            if List.isEmpty avgs then 0.0
            else List.average avgs
    printfn "  %-26s | Students: %d | Dept Avg: %.1f" (deptToString dept) count deptAvg)

printfn "\n[REMOVE] Removing student Vikram (ID=5)..."
manager.RemoveById(5)
printfn "  Students remaining: %d" manager.Count

printfn "\n[FIND] Searching for student ID=4..."
match manager.FindById(4) with
| Some student ->
    printfn "  Found:"
    printStudent student
| None ->
    printfn "  Student not found"

printfn "\n[FIND] Searching for removed student ID=5..."
match manager.FindById(5) with
| Some student ->
    printStudent student
| None ->
    printfn "  Student not found (removed)"

printfn "\n[OVERALL SUMMARY]"
let allAvgs =
    manager.GetAll()
    |> List.map (fun s -> average s.Marks)
let classAvg = List.average allAvgs
let topStudent =
    manager.GetAll()
    |> List.maxBy (fun s -> average s.Marks)
let gradeDistribution =
    manager.GetAll()
    |> List.map (fun s -> classifyGrade (average s.Marks))
    |> List.countBy id

printfn "  Class Average  : %.1f" classAvg
printfn "  Top Student    : %s (Avg=%.1f)" topStudent.Name (average topStudent.Marks)
printfn "  Grade Distribution:"
gradeDistribution
|> List.iter (fun (grade, count) ->
    printfn "    %-25s : %d student(s)" (gradeToString grade) count)

printfn "   All operations completed successfully!"

