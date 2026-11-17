open System
open System.Text.RegularExpressions

// (1) Reversing a string and checking for palindrome
let reverseString (str: string) =
    new String(str.ToCharArray() |> Array.rev)

let isPalindrome (str: string) =
    let reversed = reverseString str
    str.Equals(reversed, StringComparison.OrdinalIgnoreCase)

// (2) Transforming a string based on its contents
let transformString (str: string) =
    if str.Length > 0 && str |> Seq.forall Char.IsLetter then
        str.ToUpper()
    elif str.Length > 0 && str |> Seq.forall Char.IsDigit then
        str + str
    else
        str

// (3) Validating email using Regex
let isValidEmail (email: string) =
    let pattern = @"^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$"
    Regex.IsMatch(email, pattern)

// Main program

// Task 1: Reverse and Palindrome
printfn "----- STRING REVERSAL AND PALINDROME -----"

let strings = ["madam"; "hello"; "Level"]

for s in strings do
    printfn "Original: %s" s
    printfn "Reversed: %s" (reverseString s)
    printfn "Palindrome: %b\n" (isPalindrome s)

// Task 2: String Transformation
printfn "----- STRING TRANSFORMATION -----"

let inputs = ["hello"; "12345"; "Hello123"]

for s in inputs do
    printfn "Original: %s" s
    printfn "Transformed: %s\n" (transformString s)

// Task 3: Email Validation
printfn "----- EMAIL VALIDATION -----"

let emails = [
    "student@gmail.com"
    "admin@yahoo.co.in"
    "invalid-email"
    "test@domain"
]

for email in emails do
    if isValidEmail email then
        printfn "%s -> Valid Email" email
    else
        printfn "%s -> Invalid Email" email
