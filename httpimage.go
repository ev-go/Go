package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/math/fixed"
)

func hexToRGBA(hex string) (color.RGBA, error) {
	var (
		rgba             color.RGBA
		err              error
		errInvalidFormat = fmt.Errorf("invalid")
	)
	rgba.A = 0xff
	if hex[0] != '#' {
		return rgba, errInvalidFormat
	}
	hexToByte := func(b byte) byte {
		switch {
		case b >= '0' && b <= '9':
			return b - '0'
		case b >= 'a' && b <= 'f':
			return b - 'a' + 10
		case b >= 'A' && b <= 'F':
			return b - 'A' + 10
		}
		err = errInvalidFormat
		return 0
	}
	switch len(hex) {
	case 7:
		rgba.R = hexToByte(hex[1])<<4 + hexToByte(hex[2])
		rgba.G = hexToByte(hex[3])<<4 + hexToByte(hex[4])
		rgba.B = hexToByte(hex[5])<<4 + hexToByte(hex[6])
	case 4:
		rgba.R = hexToByte(hex[1]) * 17
		rgba.G = hexToByte(hex[2]) * 17
		rgba.B = hexToByte(hex[3]) * 17
	default:
		err = errInvalidFormat
	}
	return rgba, err
}

func drawText(canvas *image.RGBA, text string) error {
	var (
		fgColor  image.Image
		fontFace *truetype.Font
		err      error
		fontSize = 150.0
	)
	fgColor = image.White
	fontFace, err = freetype.ParseFont(goregular.TTF)
	fontDrawer := &font.Drawer{
		Dst: canvas,
		Src: fgColor,
		Face: truetype.NewFace(fontFace, &truetype.Options{
			Size:    fontSize,
			Hinting: font.HintingFull,
		}),
	}
	textBounds, _ := fontDrawer.BoundString(text)
	xPosition := (fixed.I(canvas.Rect.Max.X) - fontDrawer.MeasureString(text)) / 2
	textHeight := textBounds.Max.Y - textBounds.Min.Y
	yPosition := fixed.I((canvas.Rect.Max.Y)-textHeight.Ceil())/2 + fixed.I(textHeight.Ceil())
	fontDrawer.Dot = fixed.Point26_6{
		X: xPosition,
		Y: yPosition,
	}
	fontDrawer.DrawString(text)
	return err
}

func main() {
	// initials := "LR"
	// size := 200
	// avatar, err := createAvatar(size, initials)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// filename := fmt.Sprintf("out-%d.png", time.Now().Unix())
	// file, err := os.Create(filename)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// png.Encode(file, avatar)
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/avatar", func(w http.ResponseWriter, r *http.Request) {
		//t := time.Now()
		initials := r.FormValue("initials") //t.String() //r.FormValue("initials")
		size, err := strconv.Atoi(r.FormValue("size"))
		if err != nil {
			size = 200
		}
		avatar, err := createAvatar(size, initials)
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "image/png")
		png.Encode(w, avatar)
	})
	http.ListenAndServe(":3000", router)
}
func createAvatar(size int, initials string) (*image.RGBA, error) {
	width, height := size, size
	bgColor, err := hexToRGBA("#764abc")
	if err != nil {
		log.Fatal(err)
	}
	background := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(background, background.Bounds(), &image.Uniform{C: bgColor},
		image.Point{}, draw.Src)
	drawText(background, initials)
	return background, err
}

// func GetTime(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Current time is %s\n", time.Now())
// }
