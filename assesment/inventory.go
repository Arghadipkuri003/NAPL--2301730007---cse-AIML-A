package main

import (
	"errors"
	"fmt"
	"strings"
)

type Item struct {
	ID       int
	Name     string
	Quantity int
	Price    float64
}

type Inventory struct {
	Items []Item
}

func NewInventory() *Inventory {
	return &Inventory{Items: []Item{}}
}


func (inv *Inventory) AddItem(item Item) error {
	if item.ID <= 0 {
		return errors.New("invalid ID: must be a positive integer")
	}
	if strings.TrimSpace(item.Name) == "" {
		return errors.New("invalid Name: cannot be empty")
	}
	if item.Quantity < 0 {
		return errors.New("invalid Quantity: cannot be negative")
	}
	if item.Price < 0 {
		return errors.New("invalid Price: cannot be negative")
	}

	for _, existing := range inv.Items {
		if existing.ID == item.ID {
			return fmt.Errorf("duplicate ID: item with ID %d already exists", item.ID)
		}
	}

	inv.Items = append(inv.Items, item)
	return nil
}

func (inv *Inventory) UpdateItem(id int, name string, quantity int, price float64) error {
	if strings.TrimSpace(name) == "" {
		return errors.New("invalid Name: cannot be empty")
	}
	if quantity < 0 {
		return errors.New("invalid Quantity: cannot be negative")
	}
	if price < 0 {
		return errors.New("invalid Price: cannot be negative")
	}

	for i, item := range inv.Items {
		if item.ID == id {
			inv.Items[i].Name = name
			inv.Items[i].Quantity = quantity
			inv.Items[i].Price = price
			return nil
		}
	}

	return fmt.Errorf("item not found: no item with ID %d", id)
}

func (inv *Inventory) RemoveItem(id int) error {
	for i, item := range inv.Items {
		if item.ID == id {
			inv.Items = append(inv.Items[:i], inv.Items[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("item not found: no item with ID %d", id)
}

func (inv *Inventory) FindItem(id int) (*Item, error) {
	for i, item := range inv.Items {
		if item.ID == id {
			return &inv.Items[i], nil
		}
	}
	return nil, fmt.Errorf("item not found: no item with ID %d", id)
}

type Summary struct {
	ItemCount  int
	TotalValue float64
}

func (inv *Inventory) ComputeSummary() Summary {
	ch := make(chan Summary)

	go func() {
		var s Summary
		for _, item := range inv.Items {
			s.ItemCount++
			s.TotalValue += float64(item.Quantity) * item.Price
		}
		ch <- s
	}()

	return <-ch
}

func (inv *Inventory) PrintItems() {
	if len(inv.Items) == 0 {
		fmt.Println("  (inventory is empty)")
		return
	}
	fmt.Printf("  %-5s %-15s %-10s %-10s\n", "ID", "Name", "Qty", "Price")
	fmt.Println("  " + strings.Repeat("-", 42))
	for _, item := range inv.Items {
		fmt.Printf("  %-5d %-15s %-10d %-10.2f\n",
			item.ID, item.Name, item.Quantity, item.Price)
	}
}

