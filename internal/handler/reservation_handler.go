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
	memberService *service.MemberService
}

func NewReservationHandler(resService *service.ReservationService, memberService *service.MemberService) *ReservationHandler {
	return &ReservationHandler{
		service:       resService,
		memberService: memberService,
	}
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

func (h *ReservationHandler) list() {
	reservations, err := h.service.GetAllReservations()
	if err != nil {
		fmt.Println("❌ Gagal mengambil data reservasi:", err)
		return
	}

	if len(reservations) == 0 {
		fmt.Println("⚠️ Belum ada reservasi yang terdaftar.")
		return
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "👉 {{ .ID | cyan }}  {{ .MemberName | yellow }} | Tanggal: {{ .ReservationDate.Format \"2006-01-02\" }} | Slot: {{ .TimeSlot }}",
		Inactive: "   {{ .ID | cyan }} {{ .MemberName | yellow }} | Tanggal: {{ .ReservationDate.Format \"2006-01-02\" }} | Slot: {{ .TimeSlot }}",
		Selected: "✅ Reservasi ID {{ .ID | green }} dipilih",
		Details: `
───────────────
ID Reservasi    : {{ .ID }}
Nama Member     : {{ .MemberName }}
Member ID       : {{ .MemberID }}
Tanggal         : {{ .ReservationDate.Format "2006-01-02" }}
Time Slot       : {{ .TimeSlot }}
Jumlah Orang    : {{ .TotalPeople }}
Catatan         : {{ .Note }}
Dibuat          : {{ .CreatedAt.Format "2006-01-02 15:04:05" }}
───────────────`,
	}

	prompt := promptui.Select{
		Label:     "=== 📋 Daftar Reservasi ===",
		Items:     reservations,
		Templates: templates,
		Size:      10,
	}

	_, _, err = prompt.Run()
	if err != nil {
		fmt.Println("❌ Dibatalkan:", err)
		return
	}
}

func (h *ReservationHandler) update() {
	// Ambil semua reservasi
	reservations, err := h.service.GetAllReservations()
	if err != nil {
		fmt.Println("❌ Gagal mengambil data reservasi:", err)
		return
	}

	if len(reservations) == 0 {
		fmt.Println("⚠️ Belum ada reservasi yang terdaftar.")
		return
	}

	// Pilih reservasi yang ingin diupdate
	promptSelect := promptui.Select{
		Label: "Pilih Reservasi yang ingin diupdate",
		Items: reservations,
		Templates: &promptui.SelectTemplates{
			Active:   "👉 {{ .ID | cyan }}  {{ .MemberName | yellow }} | Tanggal: {{ .ReservationDate.Format \"2006-01-02\" }} | Slot: {{ .TimeSlot }}",
			Inactive: "   {{ .ID | cyan }} {{ .MemberName | yellow }} | Tanggal: {{ .ReservationDate.Format \"2006-01-02\" }} | Slot: {{ .TimeSlot }}",
			Selected: "✅ Reservasi ID {{ .ID | green }} dipilih",
		},
	}

	index, _, err := promptSelect.Run()
	if err != nil {
		fmt.Println("❌ Dibatalkan:", err)
		return
	}

	selectedRes := reservations[index]

	// Pilih member baru (opsional)
	members, _ := h.memberService.GetAllMembers()
	memberSelect := promptui.Select{
		Label: "Pilih Member",
		Items: members,
		Templates: &promptui.SelectTemplates{
			Active:   "👉  {{ .Name | cyan }} - {{ .Phone | faint }}",
			Inactive: "   {{ .Name | cyan }} - {{ .Phone | faint }}",
			Selected: "✅ {{ .Name | green }} dipilih",
		},
	}
	mIndex, _, _ := memberSelect.Run()
	selectedMember := members[mIndex]

	// Input tanggal baru
	datePrompt := promptui.Prompt{
		Label:   fmt.Sprintf("Tanggal Reservasi (YYYY-MM-DD) [%s]", selectedRes.ReservationDate.Format("2006-01-02")),
		Default: selectedRes.ReservationDate.Format("2006-01-02"),
	}
	dateStr, _ := datePrompt.Run()
	resDate, _ := time.Parse("2006-01-02", dateStr)

	// Input time slot
	timePrompt := promptui.Prompt{
		Label:   fmt.Sprintf("Time Slot [%s]", selectedRes.TimeSlot),
		Default: selectedRes.TimeSlot,
	}
	timeSlot, _ := timePrompt.Run()

	// Input total people
	totalPrompt := promptui.Prompt{
		Label:   fmt.Sprintf("Jumlah Orang [%d]", selectedRes.TotalPeople),
		Default: fmt.Sprintf("%d", selectedRes.TotalPeople),
	}
	totalStr, _ := totalPrompt.Run()
	totalPeople, _ := strconv.Atoi(totalStr)

	// Input note
	notePrompt := promptui.Prompt{
		Label:   fmt.Sprintf("Catatan [%s]", selectedRes.Note),
		Default: selectedRes.Note,
	}
	note, _ := notePrompt.Run()

	// Update struct
	selectedRes.MemberID = selectedMember.ID
	selectedRes.ReservationDate = resDate
	selectedRes.TimeSlot = timeSlot
	selectedRes.TotalPeople = totalPeople
	selectedRes.Note = note

	// Simpan ke DB
	err = h.service.UpdateReservation(&selectedRes)
	if err != nil {
		fmt.Println("❌ Gagal update reservasi:", err)
		return
	}

	fmt.Println("✅ Reservasi berhasil diupdate!")
}

func (h *ReservationHandler) delete() {
	// Ambil semua reservasi
	reservations, err := h.service.GetAllReservations()
	if err != nil {
		fmt.Println("❌ Gagal mengambil data reservasi:", err)
		return
	}

	if len(reservations) == 0 {
		fmt.Println("⚠️ Belum ada reservasi yang terdaftar.")
		return
	}

	// Pilih reservasi yang ingin dihapus
	promptSelect := promptui.Select{
		Label: "Pilih Reservasi yang ingin dihapus",
		Items: reservations,
		Templates: &promptui.SelectTemplates{
			Active:   "👉 {{ .ID | cyan }}  {{ .MemberName | yellow }} | Tanggal: {{ .ReservationDate.Format \"2006-01-02\" }} | Slot: {{ .TimeSlot }}",
			Inactive: "   {{ .ID | cyan }} {{ .MemberName | yellow }} | Tanggal: {{ .ReservationDate.Format \"2006-01-02\" }} | Slot: {{ .TimeSlot }}",
			Selected: "✅ Reservasi ID {{ .ID | green }} dipilih",
		},
	}

	index, _, err := promptSelect.Run()
	if err != nil {
		fmt.Println("❌ Dibatalkan:", err)
		return
	}

	selectedRes := reservations[index]

	// Konfirmasi hapus
	confirmPrompt := promptui.Select{
		Label: "Apakah Anda yakin ingin menghapus reservasi ini?",
		Items: []string{"Tidak", "Ya"},
	}
	_, result, _ := confirmPrompt.Run()
	if result != "Ya" {
		fmt.Println("❌ Hapus dibatalkan.")
		return
	}

	// Hapus dari DB
	err = h.service.DeleteReservation(selectedRes.ID)
	if err != nil {
		fmt.Println("❌ Gagal menghapus reservasi:", err)
		return
	}

	fmt.Println("✅ Reservasi berhasil dihapus!")
}
