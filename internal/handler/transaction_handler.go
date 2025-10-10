package handler

import (
	"fmt"
	"strconv"
	"time"

	"github.com/budimansol/pairproject/internal/model"
	"github.com/budimansol/pairproject/internal/service"
	"github.com/manifoldco/promptui"
)

type TransactionHandler struct {
	service       *service.TransactionService
	memberService *service.MemberService
	menuService   *service.MenuService
	staffService  *service.StaffService
}

func NewTransactionHandler(
	service *service.TransactionService,
	memberService *service.MemberService,
	menuService *service.MenuService,
	staffService *service.StaffService,
) *TransactionHandler {
	return &TransactionHandler{
		service:       service,
		memberService: memberService,
		menuService:   menuService,
		staffService:  staffService,
	}
}

func (h *TransactionHandler) menu(staffID int) {
	for {
		prompt := promptui.Select{
			Label: "Transaction Management",
			Items: []string{"1. Create Transaction", "2. List Transactions", "3. Exit"},
		}

		_, result, _ := prompt.Run()

		switch result {
		case "1. Create Transaction":
			h.create(staffID)
		case "2. List Transactions":
			h.list()
		case "3. Exit":
			return
		}
	}
}

func (h *TransactionHandler) create(staffID int) {
	fmt.Println("\n=== 🛒 Tambah Transaksi Baru ===")

	// Pilih member (opsional)
	members, _ := h.memberService.GetAllMembers()
	memberItems := make([]string, len(members)+1)
	memberItems[0] = "Non-member"
	for i, m := range members {
		memberItems[i+1] = fmt.Sprintf("%s (%s)", m.Name, m.Email)
	}

	memberSelect := promptui.Select{
		Label: "Pilih Member (opsional)",
		Items: memberItems,
	}

	memberIndex, _, _ := memberSelect.Run()
	var memberID *int
	if memberIndex > 0 {
		memberID = &members[memberIndex-1].ID
	}

	// Pilih menu
	menus, _ := h.menuService.GetAllMenu()
	var items []model.TransactionItem
	for {
		menuSelect := promptui.Select{
			Label: "Pilih Menu untuk ditambahkan (Exit untuk selesai)",
			Items: append(menuNames(menus), "Selesai"),
		}
		menuIndex, _, _ := menuSelect.Run()
		if menuIndex == len(menus) {
			break
		}
		selected := menus[menuIndex]

		qtyPrompt := promptui.Prompt{
			Label:   fmt.Sprintf("Jumlah %s", selected.Name),
			Default: "1",
		}
		qtyStr, _ := qtyPrompt.Run()
		qty, _ := strconv.Atoi(qtyStr)

		items = append(items, model.TransactionItem{
			MenuID:   selected.ID,
			MenuName: selected.Name,
			Quantity: qty,
			Subtotal: float64(qty) * selected.Price,
		})
	}

	if len(items) == 0 {
		fmt.Println("⚠️ Tidak ada item ditambahkan. Transaksi dibatalkan.")
		return
	}

	total := 0.0
	for _, it := range items {
		total += it.Subtotal
	}

	transaction := model.Transaction{
		StaffID:        staffID,
		MemberID:       memberID,
		TransactionDate: time.Now(),
		TotalAmount:    total,
		Items:          items,
	}

	id, err := h.service.CreateTransaction(&transaction)
	if err != nil {
		fmt.Println("❌ Gagal menambahkan transaksi:", err)
		return
	}

	fmt.Printf("✅ Transaksi berhasil ditambahkan! ID: %d\n", id)
}

func menuNames(menus []model.Menu) []string {
	names := make([]string, len(menus))
	for i, m := range menus {
		names[i] = fmt.Sprintf("%s (Rp%.2f) | Stock: %d", m.Name, m.Price, m.Stock)
	}
	return names
}

func (h *TransactionHandler) list() {
	transactions, err := h.service.GetAllTransactions()
	if err != nil {
		fmt.Println("❌ Gagal mengambil data transaksi:", err)
		return
	}
	if len(transactions) == 0 {
		fmt.Println("⚠️ Belum ada transaksi.")
		return
	}

	for _, t := range transactions {
		fmt.Printf("ID: %d | Staff: %s | Member: %s | Total: Rp%.2f | Tanggal: %s\n",
			t.ID, t.StaffName, t.MemberName, t.TotalAmount, t.TransactionDate.Format("2006-01-02 15:04"))
		for _, item := range t.Items {
			fmt.Printf("  - %s x%d = Rp%.2f\n", item.MenuName, item.Quantity, item.Subtotal)
		}
		fmt.Println("───────────────")
	}
}
