module SQLiteCRUD

open Microsoft.Data.Sqlite

let connectionString = "Data Source=students.db"

let initializeDb () =
    use conn = new SqliteConnection(connectionString)
    conn.Open()
    use cmd = conn.CreateCommand()
    cmd.CommandText <-
        "CREATE TABLE IF NOT EXISTS Students (
            Id INTEGER PRIMARY KEY,
            Name TEXT NOT NULL,
            Age INTEGER NOT NULL
        )"
    cmd.ExecuteNonQuery() |> ignore

let printAllStudents () =
    use conn = new SqliteConnection(connectionString)
    conn.Open()
    use cmd = conn.CreateCommand()
    cmd.CommandText <- "SELECT Id, Name, Age FROM Students ORDER BY Id"
    use reader = cmd.ExecuteReader()
    let mutable found = false
    while reader.Read() do
        found <- true
        printfn "  Id=%d, Name=%s, Age=%d"
            (reader.GetInt32(0))
            (reader.GetString(1))
            (reader.GetInt32(2))
    if not found then
        printfn "  (no students)"

let createStudent (id: int) (name: string) (age: int) =
    use conn = new SqliteConnection(connectionString)
    conn.Open()
    use cmd = conn.CreateCommand()
    cmd.CommandText <- "INSERT INTO Students (Id, Name, Age) VALUES ($id, $name, $age)"
    cmd.Parameters.AddWithValue("$id", id) |> ignore
    cmd.Parameters.AddWithValue("$name", name) |> ignore
    cmd.Parameters.AddWithValue("$age", age) |> ignore
    cmd.ExecuteNonQuery() |> ignore

let readStudent (id: int) =
    use conn = new SqliteConnection(connectionString)
    conn.Open()
    use cmd = conn.CreateCommand()
    cmd.CommandText <- "SELECT Id, Name, Age FROM Students WHERE Id = $id"
    cmd.Parameters.AddWithValue("$id", id) |> ignore
    use reader = cmd.ExecuteReader()
    if reader.Read() then
        Some (reader.GetInt32(0), reader.GetString(1), reader.GetInt32(2))
    else
        None

let updateStudent (id: int) (newName: string) (newAge: int) =
    use conn = new SqliteConnection(connectionString)
    conn.Open()
    use cmd = conn.CreateCommand()
    cmd.CommandText <- "UPDATE Students SET Name = $name, Age = $age WHERE Id = $id"
    cmd.Parameters.AddWithValue("$id", id) |> ignore
    cmd.Parameters.AddWithValue("$name", newName) |> ignore
    cmd.Parameters.AddWithValue("$age", newAge) |> ignore
    cmd.ExecuteNonQuery() |> ignore

let deleteStudent (id: int) =
    use conn = new SqliteConnection(connectionString)
    conn.Open()
    use cmd = conn.CreateCommand()
    cmd.CommandText <- "DELETE FROM Students WHERE Id = $id"
    cmd.Parameters.AddWithValue("$id", id) |> ignore
    cmd.ExecuteNonQuery() |> ignore

let cleanUp () =
    use conn = new SqliteConnection(connectionString)
    conn.Open()
    use cmd = conn.CreateCommand()
    cmd.CommandText <- "DELETE FROM Students"
    cmd.ExecuteNonQuery() |> ignore

let run () =
    printfn "========================================="
    printfn "   SQLite CRUD Operations"
    printfn "========================================="

    initializeDb ()
    cleanUp ()

    // CREATE
    printfn "\n[CREATE] Adding students..."
    createStudent 1 "Rahul" 21
    createStudent 2 "Priya" 20
    createStudent 3 "Aman" 22
    printfn "Database state after CREATE:"
    printAllStudents ()

    // READ
    printfn "\n[READ] Reading student with Id=2..."
    match readStudent 2 with
    | Some (id, name, age) ->
        printfn "  Found: Id=%d, Name=%s, Age=%d" id name age
    | None ->
        printfn "  Student not found"

    // UPDATE
    printfn "\n[UPDATE] Updating student Id=2 -> Name='Sneha', Age=23..."
    updateStudent 2 "Sneha" 23
    printfn "Database state after UPDATE:"
    printAllStudents ()

    // DELETE
    printfn "\n[DELETE] Deleting student with Id=3..."
    deleteStudent 3
    printfn "Database state after DELETE:"
    printAllStudents ()

