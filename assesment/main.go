package main

import "fmt"

func main() {
	inv := NewInventory()

	fmt.Println("  Inventory Management Module")
	fmt.Println("\n Adding items to inventory...")

	items := []Item{
		{ID: 1, Name: "Laptop", Quantity: 10, Price: 55000.00},
		{ID: 2, Name: "Mouse", Quantity: 50, Price: 500.00},
		{ID: 3, Name: "Keyboard", Quantity: 30, Price: 1200.00},
		{ID: 4, Name: "Monitor", Quantity: 15, Price: 18000.00},
	}

	for _, item := range items {
		if err := inv.AddItem(item); err != nil {
			fmt.Println("  Error:", err)
		} else {
			fmt.Printf("  Added: %s (ID=%d)\n", item.Name, item.ID)
		}
	}

	fmt.Println("\nInventory after ADD:")
	inv.PrintItems()

	fmt.Println("\n[ADD] Testing input validation...")

	badItems := []Item{
		{ID: -1, Name: "Invalid", Quantity: 5, Price: 100},
		{ID: 5, Name: "", Quantity: 5, Price: 100},
		{ID: 6, Name: "Negative Qty", Quantity: -3, Price: 100},
		{ID: 7, Name: "Negative Price", Quantity: 5, Price: -50},
		{ID: 1, Name: "Duplicate", Quantity: 1, Price: 10},
	}

	for _, item := range badItems {
		if err := inv.AddItem(item); err != nil {
			fmt.Println("  Expected error:", err)
		}
	}

	fmt.Println("\n[UPDATE] Updating Mouse (ID=2): Qty=100, Price=450.00...")

	if err := inv.UpdateItem(2, "pen drive)", 100, 450.00); err != nil {
		fmt.Println("  Error:", err)
	} else {
		fmt.Println("  Updated successfully")
	}

	fmt.Println("\nInventory after UPDATE:")
	inv.PrintItems()

	fmt.Println("\n[UPDATE] Updating non-existent item (ID=99)...")
	if err := inv.UpdateItem(99, "Ghost", 1, 10); err != nil {
		fmt.Println("  Expected error:", err)
	}

	fmt.Println("\n[REMOVE] Removing Keyboard (ID=3)...")

	if err := inv.RemoveItem(3); err != nil {
		fmt.Println("  Error:", err)
	} else {
		fmt.Println("  Removed successfully")
	}

	fmt.Println("\nInventory after REMOVE:")
	inv.PrintItems()

	fmt.Println("\n[REMOVE] Removing non-existent item (ID=3 again)...")
	if err := inv.RemoveItem(3); err != nil {
		fmt.Println("  Expected error:", err)
	}

	fmt.Println("\n[SUMMARY] Computing summary via goroutine + channel...")

	summary := inv.ComputeSummary()
	fmt.Printf("  Total Items : %d\n", summary.ItemCount)
	fmt.Printf("  Total Value : ₹%.2f\n", summary.TotalValue)

	fmt.Println("  All operations completed successfully!")
}
