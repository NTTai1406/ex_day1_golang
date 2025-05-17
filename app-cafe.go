package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MenuItem struct {
    Name  string
    Price float64
}

type CafeMenu map[string]*MenuItem

// Hàm thêm món mới
func AddItem(menu *CafeMenu, id string, name string, price float64) {
    (*menu)[id] = &MenuItem{Name: name, Price: price}
    fmt.Println("✅ Đã thêm món:", name)
}

// Hàm hiển thị menu
func DisplayMenu(menu *CafeMenu) {
    fmt.Println("\n📋 MENU:")
    for id, item := range *menu {
        fmt.Printf("ID: %s | %s - %.0fđ\n", id, item.Name, item.Price)
    }
    fmt.Println()
}

// Hàm cập nhật món
func UpdateItem(menu *CafeMenu, id string, newName string, newPrice float64) {
    if item, exists := (*menu)[id]; exists {
        item.Name = newName
        item.Price = newPrice
        fmt.Println("✅ Đã cập nhật món", id)
    } else {
        fmt.Println("❌ Không tìm thấy món với ID:", id)
    }
}

// Hàm xóa món
func DeleteItem(menu *CafeMenu, id string) {
    if _, exists := (*menu)[id]; exists {
        delete(*menu, id)
        fmt.Println("🗑️ Đã xóa món có ID:", id)
    } else {
        fmt.Println("❌ Không tìm thấy món với ID:", id)
    }
}

// Hàm điều khiển CLI
func main() {
    menu := make(CafeMenu)
    scanner := bufio.NewScanner(os.Stdin)

    for {
        fmt.Println("---- Quản lý Quán Cà Phê ----")
        fmt.Println("1. Thêm món")
        fmt.Println("2. Hiển thị menu")
        fmt.Println("3. Cập nhật món")
        fmt.Println("4. Xóa món")
        fmt.Println("5. Thoát")
        fmt.Print("Chọn chức năng: ")
        scanner.Scan()
        choice := scanner.Text()

        switch choice {
        case "1":
            fmt.Print("Nhập ID món: ")
            scanner.Scan()
            id := scanner.Text()

            fmt.Print("Tên món: ")
            scanner.Scan()
            name := scanner.Text()

            fmt.Print("Giá: ")
            scanner.Scan()
            priceStr := scanner.Text()
            price, err := strconv.ParseFloat(priceStr, 64)
            if err != nil {
                fmt.Println("❌ Giá không hợp lệ.")
                continue
            }

            AddItem(&menu, id, name, price)

        case "2":
            DisplayMenu(&menu)

        case "3":
            fmt.Print("Nhập ID món cần cập nhật: ")
            scanner.Scan()
            id := scanner.Text()

            fmt.Print("Tên mới: ")
            scanner.Scan()
            newName := scanner.Text()

            fmt.Print("Giá mới: ")
            scanner.Scan()
            newPriceStr := scanner.Text()
            newPrice, err := strconv.ParseFloat(strings.TrimSpace(newPriceStr), 64)
            if err != nil {
                fmt.Println("❌ Giá không hợp lệ.")
                continue
            }

            UpdateItem(&menu, id, newName, newPrice)

        case "4":
            fmt.Print("Nhập ID món cần xóa: ")
            scanner.Scan()
            id := scanner.Text()
            DeleteItem(&menu, id)

        case "5":
            fmt.Println("👋 Thoát chương trình.")
            return

        default:
            fmt.Println("❗ Lựa chọn không hợp lệ.")
        }

        fmt.Println()
    }
}
