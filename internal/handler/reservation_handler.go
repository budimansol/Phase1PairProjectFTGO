package handler

import (
	"fmt"
	"strconv"
	"time"

	"github.com/budimansol/pairproject/internal/model"
	"github.com/budimansol/pairproject/internal/service"
	"github.com/manifoldco/promptui"
)

type ReservationHandler struct {
	service *service.ReservationService
}

func NewReservationHandler(service *service.ReservationService) *ReservationHandler {
	return &ReservationHandler{service: service}
}

func (h *ReservationHandler) Menu() {
	for {
		prompt := promptui.Select{
			Label: "Reservation Management",
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
		// case "2. List":
		// 	h.list()
		// case "3. Update":
		// 	h.update()
		// case "4. Delete":
		// 	h.delete()
		case "5. Exit":
			return
		}
	}
}

func (h *ReservationHandler) create() {
	// Ambil daftar member dulu
	members, err := h.memberService.GetAllMembers()
	if err != nil {
		fmt.Println("❌ Gagal mengambil daftar member:", err)
		return
	}
	if len(members) == 0 {
		fmt.Println("⚠️ Belum ada member. Tambahkan member dulu.")
		return
	}

	// Pilih member
	memberSelect := promptui.Select{
		Label: "Pilih Member",
		Items: members,
		Templates: &promptui.SelectTemplates{
			Active:   "👉  {{ .Name | cyan }} - {{ .Phone | faint }}",
			Inactive: "   {{ .Name | cyan }} - {{ .Phone | faint }}",
			Selected: "✅ {{ .Name | green }} dipilih",
		},
	}

	index, _, err := memberSelect.Run()
	if err != nil {
		fmt.Println("❌ Dibatalkan:", err)
		return
	}
	selectedMember := members[index]

	// Input tanggal
	datePrompt := promptui.Prompt{
		Label: "Tanggal Reservasi (YYYY-MM-DD)",
	}
	dateStr, _ := datePrompt.Run()
	resDate, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		fmt.Println("❌ Format tanggal salah:", err)
		return
	}

	// Input time slot
	timeSlotPrompt := promptui.Prompt{
		Label: "Time Slot (contoh: 10:00-11:00)",
	}
	timeSlot, _ := timeSlotPrompt.Run()

	// Input jumlah orang
	totalPrompt := promptui.Prompt{
		Label: "Jumlah Orang",
	}
	totalStr, _ := totalPrompt.Run()
	totalPeople, err := strconv.Atoi(totalStr)
	if err != nil {
		fmt.Println("❌ Jumlah orang harus angka:", err)
		return
	}

	// Catatan opsional
	notePrompt := promptui.Prompt{
		Label: "Catatan (opsional)",
	}
	note, _ := notePrompt.Run()

	// Konfirmasi
	confirm := promptui.Select{
		Label: fmt.Sprintf("Yakin ingin menambahkan reservasi untuk %s?", selectedMember.Name),
		Items: []string{"Ya, simpan", "Tidak, batalkan"},
	}

	_, choice, _ := confirm.Run()
	if choice == "Tidak, batalkan" {
		fmt.Println("🚫 Reservasi dibatalkan.")
		return
	}

	// Buat struct reservation
	reservation := model.Reservation{
		MemberID:       selectedMember.ID,
		ReservationDate: resDate,
		TimeSlot:       timeSlot,
		TotalPeople:    totalPeople,
		Note:           note,
	}

	// Simpan ke database
	err = h.service.CreateReservation(reservation)
	if err != nil {
		fmt.Println("❌ Gagal menambahkan reservasi:", err)
		return
	}

	fmt.Println("✅ Reservasi berhasil ditambahkan!")
}
