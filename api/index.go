package handler

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"net/http"
	"os"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	img, err := combineImages([]string{"static/body/1.png", "static/lips/1.png"})
	if err != nil {
		fmt.Fprint(w, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	err = png.Encode(w, img)
	if err != nil {
		fmt.Fprint(w, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	key := r.URL.Query().Get("key")
	log.Println(key)
	fmt.Fprintf(w, "Key: {%s}", key)
}

/*
body
eyes
lips
shirs
glasses
hats
*/

func combineImages(imgPaths []string) (*image.RGBA, error) {
	var rgbaImages []*image.RGBA
	width := 0
	height := 0

	for _, imagePath := range imgPaths {
		rgba, err := openImageAsRGBA(imagePath)
		if err != nil {
			return nil, err
		}
		w := rgba.Bounds().Dx()
		h := rgba.Bounds().Dy()
		if w > width {
			width = w
		}
		if h > height {
			height = h
		}
		rgbaImages = append(rgbaImages, rgba)
	}

	result := image.NewRGBA(image.Rect(0, 0, width, height))

	for _, rgba := range rgbaImages {
		mask := createAlphaMask(rgba)
		draw.DrawMask(result, rgba.Bounds(), rgba, image.Point{}, mask, image.Point{}, draw.Over)
	}

	return result, nil
}

func createAlphaMask(img *image.RGBA) *image.Alpha {
	mask := image.NewAlpha(img.Bounds())
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			_, _, _, a := img.At(x, y).RGBA()
			if a == 0 {
				mask.SetAlpha(x, y, color.Alpha{0})
			} else {
				mask.SetAlpha(x, y, color.Alpha{0xff})
			}
		}
	}
	return mask
}

func openImageAsRGBA(path string) (*image.RGBA, error) {
	imageFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer imageFile.Close()

	imageDecoded, _, err := image.Decode(imageFile)
	if err != nil {
		return nil, err
	}

	rgba := image.NewRGBA(imageDecoded.Bounds())

	draw.Draw(rgba, imageDecoded.Bounds(), imageDecoded, image.Point{}, draw.Src)
	return rgba, nil
}
