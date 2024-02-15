package handler

import (
	"crypto/sha256"
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
	key := r.URL.Path
	log.Print(key)
	id := hashToIndices(key)

	imagePaths := getImages(
		id[0],
		id[1],
		id[2],
		id[3],
		id[4],
		id[5],
	)

	img, err := combineImages(imagePaths)
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
}

/*
body
eyes
lips
shirs
glasses
hats
*/

func hashToIndices(key string) [6]int {
	hash := sha256.Sum224([]byte(key))
	return [6]int{
		1 + int(hash[0]%20),
		1 + int(hash[4]%6),
		1 + int(hash[8]%4),
		1 + int(hash[12]%16),
		1 + int(hash[16]%8),
		1 + int(hash[20]%10),
	}
}

func getImages(body int, eyes int, lips int, shirts int, glasses int, hats int) []string {
	// prefix := os.Getenv("STATIC_PATH")
	prefix := "/usr/share/static"
	return []string{
		fmt.Sprintf("%s/body/%d.png", prefix, body),
		fmt.Sprintf("%s/eyes/%d.png", prefix, eyes),
		fmt.Sprintf("%s/lips/%d.png", prefix, lips),
		fmt.Sprintf("%s/shirts/%d.png", prefix, shirts),
		fmt.Sprintf("%s/glasses/%d.png", prefix, glasses),
		fmt.Sprintf("%s/hats/%d.png", prefix, hats),
	}
}

func combineImages(imgPaths []string) (*image.RGBA, error) {
	log.Print("combining", imgPaths)
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
