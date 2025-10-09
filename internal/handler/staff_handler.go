package handler

import (
	"fmt"

	"github.com/budimansol/pairproject/internal/model"
	"github.com/budimansol/pairproject/internal/service"
	"github.com/manifoldco/promptui"
)

type StaffHandler struct {
	service *service.StaffService
}

func NewStaffHandler(s *service.StaffService) *StaffHandler {
	return &StaffHandler{service: s}
}

func (h *StaffHandler) Menu() {
	for {
		prompt := promptui.Select{
			Label: "Staff Management",
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

func (h *StaffHandler) create() {
	namePrompt := promptui.Prompt{Label: "Name"}
	emailPrompt := promptui.Prompt{Label: "Email"}
	rolePrompt := promptui.Prompt{Label: "Role"}
	passPrompt := promptui.Prompt{Label: "Password"}

	name, _ := namePrompt.Run()
	email, _ := emailPrompt.Run()
	role, _ := rolePrompt.Run()
	pass, _ := passPrompt.Run()

	staff := model.Staff{Name: name, Email: email, Role: role, Password: pass}
	err := h.service.CreateStaff(staff)
	if err != nil {
		fmt.Println("❌ Failed to create staff:", err)
		return
	}
	fmt.Println("✅ Staff created successfully!")
}

func (h *StaffHandler) list() {
	staffs, err := h.service.GetAllStaff()
	if err != nil {
		fmt.Println("❌ Gagal mengambil data staff:", err)
		return
	}

	if len(staffs) == 0 {
		fmt.Println("⚠️ Tidak ada staff yang terdaftar.")
		return
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "👉  {{ .Name | cyan }} ({{ .Role | yellow }}) - {{ .Email | faint }}",
		Inactive: "   {{ .Name | cyan }} ({{ .Role | yellow }}) - {{ .Email | faint }}",
		Selected: "✅  {{ .Name | green }} dipilih",
		Details: `
───────────────
📛 Nama:	{{ .Name }}
📧 Email:	{{ .Email }}
🎯 Role:	{{ .Role }}
📅 Dibuat:	{{ .CreatedAt.Format "2006-01-02" }}
🏠 Alamat:	{{ if .Profile }}{{ .Profile.Address }}{{ else }}(tidak ada){{ end }}
📞 Telepon:	{{ if .Profile }}{{ .Profile.Phone }}{{ else }}(tidak ada){{ end }}
🎂 Lahir:	{{ if .Profile }}{{ .Profile.DateOfBirth.Format "2006-01-02" }}{{ else }}(tidak ada){{ end }}
🆘 Kontak Darurat:	{{ if .Profile }}{{ .Profile.EmergencyContact }}{{ else }}(tidak ada){{ end }}
───────────────`,
	}

	prompt := promptui.Select{
		Label:     "=== 📋 Daftar Staff ===",
		Items:     staffs,
		Templates: templates,
		Size:      10,
	}

	_, _, err = prompt.Run()
	if err != nil {
		fmt.Println("❌ Dibatalkan:", err)
		return
	}
}

func (h *StaffHandler) update() {
	// Ambil semua staff dari service
	staffList, err := h.service.GetAllStaff()
	if err != nil {
		fmt.Println("❌ Gagal mengambil data staff:", err)
		return
	}

	// Jika kosong
	if len(staffList) == 0 {
		fmt.Println("⚠️ Tidak ada staff yang bisa diupdate.")
		return
	}

	// Tampilkan list staff untuk dipilih
	staffNames := []string{}
	for _, s := range staffList {
		staffNames = append(staffNames, fmt.Sprintf("%d - %s (%s)", s.ID, s.Name, s.Role))
	}

	selectPrompt := promptui.Select{
		Label: "Pilih Staff yang ingin diupdate",
		Items: staffNames,
	}

	index, _, err := selectPrompt.Run()
	if err != nil {
		fmt.Println("❌ Pilihan dibatalkan:", err)
		return
	}

	selectedStaff := staffList[index]

	// Prompt untuk input data baru (boleh kosong, jika tidak ingin diubah)
	namePrompt := promptui.Prompt{
		Label:   "Nama baru (kosongkan jika tidak diubah)",
		Default: selectedStaff.Name,
	}
	emailPrompt := promptui.Prompt{
		Label:   "Email baru (kosongkan jika tidak diubah)",
		Default: selectedStaff.Email,
	}
	rolePrompt := promptui.Prompt{
		Label:   "Role baru (kosongkan jika tidak diubah)",
		Default: selectedStaff.Role,
	}

	name, _ := namePrompt.Run()
	email, _ := emailPrompt.Run()
	role, _ := rolePrompt.Run()

	// Update data di struct
	staff := model.Staff{
		ID:    selectedStaff.ID,
		Name:  name,
		Email: email,
		Role:  role,
	}

	err = h.service.UpdateStaff(staff)
	if err != nil {
		fmt.Println("❌ Update gagal:", err)
		return
	}
	fmt.Println("✅ Staff berhasil diupdate!")
}

func (h *StaffHandler) delete() {
	// Ambil semua staff
	staffs, err := h.service.GetAllStaff()
	if err != nil {
		fmt.Println("❌ Gagal mengambil data staff:", err)
		return
	}

	if len(staffs) == 0 {
		fmt.Println("⚠️  Tidak ada staff yang dapat dihapus.")
		return
	}

	// Daftar tampilan staff
	staffItems := []string{}
	for _, s := range staffs {
		staffItems = append(staffItems, fmt.Sprintf("%d - %s (%s)", s.ID, s.Name, s.Role))
	}

	// Pilih staff
	selectPrompt := promptui.Select{
		Label: "Pilih Staff yang ingin dihapus",
		Items: staffItems,
		Size:  10,
	}

	index, _, err := selectPrompt.Run()
	if err != nil {
		fmt.Println("❌ Penghapusan dibatalkan:", err)
		return
	}

	selectedStaff := staffs[index]

	// Konfirmasi penghapusan dengan Select
	confirmPrompt := promptui.Select{
		Label: fmt.Sprintf("Apakah kamu yakin ingin menghapus '%s'?", selectedStaff.Name),
		Items: []string{"Ya, hapus", "Tidak, batalkan"},
	}

	confirmIndex, _, err := confirmPrompt.Run()
	if err != nil {
		fmt.Println("❌ Penghapusan dibatalkan:", err)
		return
	}

	if confirmIndex != 0 { // User memilih "Tidak, batalkan"
		fmt.Println("❎ Penghapusan dibatalkan.")
		return
	}

	// Hapus data staff
	err = h.service.DeleteStaff(selectedStaff.ID)
	if err != nil {
		fmt.Println("❌ Gagal menghapus staff:", err)
		return
	}

	fmt.Printf("✅ Staff '%s' berhasil dihapus!\n", selectedStaff.Name)
}