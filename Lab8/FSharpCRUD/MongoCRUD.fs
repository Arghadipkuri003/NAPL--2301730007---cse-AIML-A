module MongoCRUD

open MongoDB.Bson
open MongoDB.Driver

[<CLIMutable>]
type Student =
    { mutable Id: int
      mutable Name: string
      mutable Age: int }

// Change this connection string if your MongoDB runs elsewhere
let connectionString = "mongodb://localhost:27017"
let databaseName = "CollegeDB"
let collectionName = "Students"

let getCollection () =
    let client = new MongoClient(connectionString)
    let database = client.GetDatabase(databaseName)
    database.GetCollection<Student>(collectionName)

let printAllStudents (collection: IMongoCollection<Student>) =
    let students = collection.Find(Builders<Student>.Filter.Empty).ToList()
    if students.Count = 0 then
        printfn "  (no students)"
    else
        for s in students do
            printfn "  Id=%d, Name=%s, Age=%d" s.Id s.Name s.Age

let run () =
    printfn "========================================="
    printfn "   MongoDB CRUD Operations"
    printfn "========================================="

    let collection = getCollection ()

    // Clean up previous test data
    collection.DeleteMany(Builders<Student>.Filter.Empty) |> ignore

    // CREATE
    printfn "\n[CREATE] Adding students..."
    collection.InsertOne({ Id = 1; Name = "Rahul"; Age = 21 })
    collection.InsertOne({ Id = 2; Name = "Priya"; Age = 20 })
    collection.InsertOne({ Id = 3; Name = "Aman"; Age = 22 })
    printfn "Database state after CREATE:"
    printAllStudents collection

    // READ
    printfn "\n[READ] Reading student with Id=2..."
    let filter = Builders<Student>.Filter.Eq((fun s -> s.Id), 2)
    let found = collection.Find(filter).FirstOrDefault()
    if box found <> null then
        printfn "  Found: Id=%d, Name=%s, Age=%d" found.Id found.Name found.Age
    else
        printfn "  Student not found"

    // UPDATE
    printfn "\n[UPDATE] Updating student Id=2 -> Name='Sneha', Age=23..."
    let update =
        Builders<Student>.Update
            .Set((fun s -> s.Name), "Sneha")
            .Set((fun s -> s.Age), 23)
    collection.UpdateOne(filter, update) |> ignore
    printfn "Database state after UPDATE:"
    printAllStudents collection

    // DELETE
    printfn "\n[DELETE] Deleting student with Id=3..."
    let deleteFilter = Builders<Student>.Filter.Eq((fun s -> s.Id), 3)
    collection.DeleteOne(deleteFilter) |> ignore
    printfn "Database state after DELETE:"
    printAllStudents collection
