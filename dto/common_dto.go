package dto

import "errors"

const (
	MESSAGE_FAILED_PANIC_OCCURED   = "server panic occured"
	MESSAGE_FAILED_PARSE_TIME      = "failed to parse time"
	PESAN_DILUAR_MASA_REGISTRASI   = "Di luar masa registrasi"
	PESAN_WEB_MAINTENANCE          = "Website sedang dalam masa maintenance"
	PESAN_AKSI_TIDAK_DIPERBOLEHKAN = "Aksi tidak diperbolehkan"
	PESAN_WSN_SEDANG_DITUTUP       = "WSN ditutup untuk sementara. Tunggu yaa"
)

var (
	ErrGeneral = errors.New("something went wrong")
)
