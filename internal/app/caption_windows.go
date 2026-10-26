package app

import (
	"syscall"
	"unsafe"
)

var dwmSetWindowAttribute = syscall.NewLazyDLL("dwmapi.dll").NewProc("DwmSetWindowAttribute")

// DWMWA_CAPTION_COLOR, and DWMWA_COLOR_NONE: no color at all.
const (
	dwmwaCaptionColor = 35
	dwmColorNone      = 0xfffffffe
)

// plainCaption keeps the accent off the window's top. With the accent shown
// on title bars, Windows 11 paints the caption in it, and with the title
// bar hidden that is a red strip over the sidebar's top and the window's
// first row. Windows 10 knows no caption color, and ignores this.
func plainCaption(handle uintptr) {
	color := uint32(dwmColorNone)
	dwmSetWindowAttribute.Call(handle, dwmwaCaptionColor, uintptr(unsafe.Pointer(&color)), unsafe.Sizeof(color))
}
