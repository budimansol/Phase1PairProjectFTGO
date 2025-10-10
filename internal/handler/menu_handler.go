package handler

import (
	"fmt"
	"strconv"
	"time"

	"github.com/budimansol/pairproject/internal/model"
	"github.com/budimansol/pairproject/internal/service"
	"github.com/manifoldco/promptui"
)

type MenuHandler struct {
	service *service.MenuService
}

func NewMenuHandler(service *service.MenuService) *MenuHandler {
	return &MenuHandler{service: service}
}

func (h *MenuHandler) Menu() {
	for {
		prompt := promptui.Select{
			Label: "Menu Management",
			Items: []string{
				"1. Create", 
				"2. List", 
				"3. Update", 
				"4. Delete", 
				"5. Exit"},
		}

		_, result, _ := prompt.Run()

		switch result {
		case "1. Create":
			h.create()
		case "2. List":
			h.list()
		case "3. Update":
			h.update()
		case "4. Delete":
			h.delete()
		case "5. Exit":
			return
		}
	}
}



func (h *MenuHandler) create() {
	fmt.Println("\n=== 🍽️ Tambah Menu Baru ===")

	namePrompt := promptui.Prompt{Label: "Nama Menu"}
	name, _ := namePrompt.Run()

	categoryPrompt := promptui.Prompt{Label: "Kategori"}
	category, _ := categoryPrompt.Run()

	pricePrompt := promptui.Prompt{Label: "Harga (contoh: 25000)"}
	priceStr, _ := pricePrompt.Run()
	price, _ := strconv.ParseFloat(priceStr, 64)

	stockPrompt := promptui.Prompt{Label: "Stok Awal"}
	stockStr, _ := stockPrompt.Run()
	stock, _ := strconv.Atoi(stockStr)

	menu := model.Menu{
		Name:     name,
		Category: category,
		Price:    price,
		Stock:    stock,
	}

	err := h.service.AddMenu(menu)
	if err != nil {
		fmt.Println("❌ Gagal menambahkan menu:", err)
		return
	}

	fmt.Println("✅ Menu berhasil ditambahkan!")
}

func (h *MenuHandler) list() {
	menus, err := h.service.GetAllMenu()
	if err != nil {
		fmt.Println("❌ Gagal mengambil data menu:", err)
		return
	}

	if len(menus) == 0 {
		fmt.Println("⚠️ Belum ada menu yang terdaftar.")
		return
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "👉 {{ .ID | cyan }}  {{ .Name | cyan }} ({{ .Category | yellow }}) - Rp{{ .Price }} | Stock: {{ .Stock }}",
		Inactive: " {{ .ID | cyan }} {{ .Name | cyan }} ({{ .Category | yellow }}) - Rp{{ .Price }} | Stock: {{ .Stock }}",
		Selected: "✅  {{ .Name | green }} dipilih",
		Details: `
───────────────
ID			: {{ .ID }}
Nama        : {{ .Name }}
Kategori    : {{ .Category }}
Harga       : Rp{{ .Price }}
Stok        : {{ .Stock }}
Dibuat      : {{ .CreatedAt.Format "2006-01-02" }}
───────────────`,
	}

	prompt := promptui.Select{
		Label:     "=== 🍽️ Daftar Menu ===",
		Items:     menus,
		Templates: templates,
		Size:      10,
	}

	_, _, err = prompt.Run()
	if err != nil {
		fmt.Println("❌ Dibatalkan:", err)
		return
	}
}

func (h *MenuHandler) update() {
	menus, err := h.service.GetAllMenu()
	if err != nil {
		fmt.Println("❌ Gagal mengambil daftar menu:", err)
		return
	}
	if len(menus) == 0 {
		fmt.Println("⚠️ Belum ada menu yang terdaftar.")
		return
	}

	menuSelect := promptui.Select{
		Label: "Pilih menu yang ingin diupdate",
		Items: menus,
		Templates: &promptui.SelectTemplates{
			Active:   "👉 {{ .ID | cyan }} {{ .Name | cyan }} ({{ .Category | yellow }}) - Rp{{ .Price }} | Stock: {{ .Stock }}",
			Inactive: "  {{ .ID | cyan }} {{ .Name | cyan }} ({{ .Category | yellow }}) - Rp{{ .Price }} | Stock: {{ .Stock }}",
			Selected: "✅ {{ .Name | green }} dipilih",
			Details: `
───────────────
ID			: {{ .ID }}
Nama        : {{ .Name }}
Kategori    : {{ .Category }}
Harga       : Rp{{ .Price }}
Stok        : {{ .Stock }}
Dibuat      : {{ .CreatedAt.Format "2006-01-02" }}
───────────────`,
		},
		Size: 10,
	}

	index, _, err := menuSelect.Run()
	if err != nil {
		fmt.Println("❌ Dibatalkan:", err)
		return
	}

	selected := menus[index]

	namePrompt := promptui.Prompt{
		Label:   fmt.Sprintf("Nama Menu [%s]", selected.Name),
		Default: selected.Name,
	}
	name, _ := namePrompt.Run()

	categoryPrompt := promptui.Prompt{
		Label:   fmt.Sprintf("Kategori [%s]", selected.Category),
		Default: selected.Category,
	}
	category, _ := categoryPrompt.Run()

	pricePrompt := promptui.Prompt{
		Label:   fmt.Sprintf("Harga (Rp) [%.2f]", selected.Price),
		Default: fmt.Sprintf("%.2f", selected.Price),
	}
	priceStr, _ := pricePrompt.Run()
	price, _ := strconv.ParseFloat(priceStr, 64)

	stockPrompt := promptui.Prompt{
		Label:   fmt.Sprintf("Stok [%d]", selected.Stock),
		Default: fmt.Sprintf("%d", selected.Stock),
	}
	stockStr, _ := stockPrompt.Run()
	stock, _ := strconv.Atoi(stockStr)

	confirm := promptui.Select{
		Label: "Yakin ingin menyimpan perubahan?",
		Items: []string{"Ya, simpan", "Tidak, batalkan"},
	}
	_, choice, _ := confirm.Run()
	if choice == "Tidak, batalkan" {
		fmt.Println("🚫 Update dibatalkan.")
		return
	}

	updated := model.Menu{
		ID:        selected.ID,
		Name:      name,
		Category:  category,
		Price:     price,
		Stock:     stock,
		CreatedAt: time.Now(),
	}

	err = h.service.UpdateMenu(updated)
	if err != nil {
		fmt.Println("❌ Gagal mengupdate menu:", err)
		return
	}

	fmt.Println("✅ Menu berhasil diperbarui!")
}

func (h *MenuHandler) delete() {
	menus, err := h.service.GetAllMenu()
	if err != nil {
		fmt.Println("❌ Gagal mengambil daftar menu:", err)
		return
	}
	if len(menus) == 0 {
		fmt.Println("⚠️ Tidak ada menu untuk dihapus.")
		return
	}

	// Pilih menu yang akan dihapus
	menuSelect := promptui.Select{
		Label: "Pilih menu yang ingin dihapus",
		Items: menus,
		Templates: &promptui.SelectTemplates{
			Active:   "👉 {{ .ID | cyan }}  {{ .Name | red }} ({{ .Category | yellow }}) - Rp{{ .Price }} | Stock: {{ .Stock }}",
			Inactive: "  {{ .ID | cyan }} {{ .Name | red }} ({{ .Category | yellow }}) - Rp{{ .Price }} | Stock: {{ .Stock }}",
			Selected: "🗑️  {{ .Name | red }} dipilih untuk dihapus",
			Details: `
───────────────
ID			: {{ .ID }}
Nama        : {{ .Name }}
Kategori    : {{ .Category }}
Harga       : Rp{{ .Price }}
Stok        : {{ .Stock }}
Dibuat      : {{ .CreatedAt.Format "2006-01-02" }}
───────────────`,
		},
		Size: 10,
	}

	index, _, err := menuSelect.Run()
	if err != nil {
		fmt.Println("❌ Dibatalkan:", err)
		return
	}

	selected := menus[index]

	// Konfirmasi penghapusan
	confirmPrompt := promptui.Select{
		Label: fmt.Sprintf("Yakin ingin menghapus menu '%s'?", selected.Name),
		Items: []string{"Ya, hapus", "Tidak, batalkan"},
	}

	_, confirm, err := confirmPrompt.Run()
	if err != nil {
		fmt.Println("❌ Dibatalkan:", err)
		return
	}

	if confirm == "Tidak, batalkan" {
		fmt.Println("🚫 Penghapusan dibatalkan.")
		return
	}

	err = h.service.DeleteMenu(selected.ID)
	if err != nil {
		fmt.Println("❌ Gagal menghapus menu:", err)
		return
	}

	fmt.Printf("✅ Menu '%s' berhasil dihapus!\n", selected.Name)
}
