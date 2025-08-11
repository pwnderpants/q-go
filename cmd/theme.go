package cmd

import "github.com/gdamore/tcell/v2"

// Catppuccin Macchiato color palette
var (
	// Background/Base Colors
	MacchiatoBase    = tcell.NewHexColor(0x24273a)
	MacchiatoMantle  = tcell.NewHexColor(0x1e2030)
	MacchiatoCrust   = tcell.NewHexColor(0x181926)

	// Text Colors
	MacchiatoText     = tcell.NewHexColor(0xcad3f5)
	MacchiatoSubtext1 = tcell.NewHexColor(0xb8c0e0)
	MacchiatoSubtext0 = tcell.NewHexColor(0xa5adcb)

	// Accent Colors
	MacchiatoRosewater = tcell.NewHexColor(0xf4dbd6)
	MacchiatoFlamingo  = tcell.NewHexColor(0xf0c6c6)
	MacchiatoPink      = tcell.NewHexColor(0xf5bde6)
	MacchiatoMauve     = tcell.NewHexColor(0xc6a0f6)
	MacchiatoRed       = tcell.NewHexColor(0xed8796)
	MacchiatoMaroon    = tcell.NewHexColor(0xee99a0)
	MacchiatoPeach     = tcell.NewHexColor(0xf5a97f)
	MacchiatoYellow    = tcell.NewHexColor(0xeed49f)
	MacchiatoGreen     = tcell.NewHexColor(0xa6da95)
	MacchiatoTeal      = tcell.NewHexColor(0x8bd5ca)
	MacchiatoSky       = tcell.NewHexColor(0x91d7e3)
	MacchiatoSapphire  = tcell.NewHexColor(0x7dc4e4)
	MacchiatoBlue      = tcell.NewHexColor(0x8aadf4)
	MacchiatoLavender  = tcell.NewHexColor(0xb7bdf8)
)

// Theme styles
var (
	MacchiatoStyleDefault    = tcell.StyleDefault.Background(MacchiatoBase).Foreground(MacchiatoText)
	MacchiatoStyleAccent     = tcell.StyleDefault.Background(MacchiatoBase).Foreground(MacchiatoMauve)
	MacchiatoStyleHighlight  = tcell.StyleDefault.Background(MacchiatoMantle).Foreground(MacchiatoText)
	MacchiatoStyleSecondary  = tcell.StyleDefault.Background(MacchiatoBase).Foreground(MacchiatoSubtext1)
	MacchiatoStyleBorder     = tcell.StyleDefault.Background(MacchiatoBase).Foreground(MacchiatoSubtext0)
)