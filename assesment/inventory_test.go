package main

import (
	"testing"
)

func TestAddItem(t *testing.T) {
	t.Run("add valid item", func(t *testing.T) {
		inv := NewInventory()
		err := inv.AddItem(Item{ID: 1, Name: "Laptop", Quantity: 10, Price: 55000})
		if err != nil {
			t.Fatalf("expected no error, got: %v", err)
		}
		if len(inv.Items) != 1 {
			t.Fatalf("expected 1 item, got %d", len(inv.Items))
		}
		if inv.Items[0].Name != "Laptop" {
			t.Errorf("expected name 'Laptop', got '%s'", inv.Items[0].Name)
		}
	})

	t.Run("reject duplicate ID", func(t *testing.T) {
		inv := NewInventory()
		inv.AddItem(Item{ID: 1, Name: "Laptop", Quantity: 10, Price: 55000})
		err := inv.AddItem(Item{ID: 1, Name: "Mouse", Quantity: 5, Price: 500})
		if err == nil {
			t.Fatal("expected duplicate ID error, got nil")
		}
	})

	t.Run("reject negative price", func(t *testing.T) {
		inv := NewInventory()
		err := inv.AddItem(Item{ID: 1, Name: "Bad", Quantity: 5, Price: -100})
		if err == nil {
			t.Fatal("expected negative price error, got nil")
		}
	})
t.Run("reject empty name", func(t *testing.T) {
		inv := NewInventory()
		err := inv.AddItem(Item{ID: 1, Name: "   ", Quantity: 5, Price: 100})
		if err == nil {
			t.Fatal("expected empty name error, got nil")
		}
	})

	t.Run("reject negative quantity", func(t *testing.T) {
		inv := NewInventory()
		err := inv.AddItem(Item{ID: 1, Name: "Item", Quantity: -1, Price: 100})
		if err == nil {
			t.Fatal("expected negative quantity error, got nil")
		}
	})
}


func TestUpdateItem(t *testing.T) {
	t.Run("update existing item", func(t *testing.T) {
		inv := NewInventory()
		inv.AddItem(Item{ID: 1, Name: "Laptop", Quantity: 10, Price: 55000})

		err := inv.UpdateItem(1, "Laptop Pro", 20, 65000)
		if err != nil {
			t.Fatalf("expected no error, got: %v", err)
		}
		if inv.Items[0].Name != "Laptop Pro" {
			t.Errorf("expected 'Laptop Pro', got '%s'", inv.Items[0].Name)
		}
		if inv.Items[0].Quantity != 20 {
			t.Errorf("expected quantity 20, got %d", inv.Items[0].Quantity)
		}
		if inv.Items[0].Price != 65000 {
			t.Errorf("expected price 65000, got %.2f", inv.Items[0].Price)
		}
	})

	t.Run("update non-existent item", func(t *testing.T) {
		inv := NewInventory()
		err := inv.UpdateItem(99, "Ghost", 1, 10)
		if err == nil {
			t.Fatal("expected not-found error, got nil")
		}
	})

	t.Run("reject invalid update values", func(t *testing.T) {
		inv := NewInventory()
		inv.AddItem(Item{ID: 1, Name: "Laptop", Quantity: 10, Price: 55000})

		err := inv.UpdateItem(1, "", 10, 55000)
		if err == nil {
			t.Fatal("expected empty name error, got nil")
		}
	})
}
func TestRemoveItem(t *testing.T) {
	t.Run("remove existing item", func(t *testing.T) {
		inv := NewInventory()
		inv.AddItem(Item{ID: 1, Name: "Laptop", Quantity: 10, Price: 55000})
		inv.AddItem(Item{ID: 2, Name: "Mouse", Quantity: 50, Price: 500})

		err := inv.RemoveItem(1)
		if err != nil {
			t.Fatalf("expected no error, got: %v", err)
		}
		if len(inv.Items) != 1 {
			t.Fatalf("expected 1 item remaining, got %d", len(inv.Items))
		}
		if inv.Items[0].ID != 2 {
			t.Errorf("expected remaining item ID=2, got ID=%d", inv.Items[0].ID)
		}
	})
	t.Run("remove non-existent item", func(t *testing.T) {
		inv := NewInventory()
		err := inv.RemoveItem(99)
		if err == nil {
			t.Fatal("expected not-found error, got nil")
		}
	})
}


func TestComputeSummary(t *testing.T) {
	inv := NewInventory()
	inv.AddItem(Item{ID: 1, Name: "Laptop", Quantity: 10, Price: 55000})
	inv.AddItem(Item{ID: 2, Name: "Mouse", Quantity: 50, Price: 500})
	inv.AddItem(Item{ID: 3, Name: "Keyboard", Quantity: 30, Price: 1200})

	summary := inv.ComputeSummary()

	if summary.ItemCount != 3 {
		t.Errorf("expected ItemCount=3, got %d", summary.ItemCount)
	}

	expectedValue := 611000.0
	if summary.TotalValue != expectedValue {
		t.Errorf("expected TotalValue=%.2f, got %.2f", expectedValue, summary.TotalValue)
	}
}
