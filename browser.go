/*
Copyright © 2021 Left Shift Logical, LLC. <support@leftshift.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package smock

import (
	"image"
	"image/color"

	"github.com/fogleman/gg"
)

const headerHeight = 32

var headerColor = &color.RGBA{
	R: 232,
	G: 232,
	B: 232,
	A: 255,
}

type BrowserOptions struct {
	VerticalPadding   int
	HorizontalPadding int
}

var defaultBrowserOpts = &BrowserOptions{
	VerticalPadding:   75,
	HorizontalPadding: 100,
}

func Browser(img image.Image, opts *BrowserOptions) (image.Image, error) {
	if opts == nil {
		opts = defaultBrowserOpts
	}

	inBounds := img.Bounds()

	dc := gg.NewContext(inBounds.Dx()+(opts.HorizontalPadding*2), inBounds.Dy()+(opts.VerticalPadding*2)+headerHeight)

	dc.SetColor(color.White)
	dc.Clear()

	radiusMask := gg.NewContext(inBounds.Dx()+(opts.HorizontalPadding*2), inBounds.Dy()+(opts.VerticalPadding*2)+headerHeight)
	radiusMask.SetColor(color.Black)
	radiusMask.DrawRoundedRectangle(float64(opts.HorizontalPadding), float64(opts.VerticalPadding), float64(inBounds.Dx()), float64(inBounds.Dy()+headerHeight), 12)
	radiusMask.Fill()

	err := dc.SetMask(radiusMask.AsMask())
	if err != nil {
		return nil, err
	}

	dc.DrawImage(img, opts.HorizontalPadding, opts.VerticalPadding+headerHeight)

	dc.SetColor(headerColor)
	dc.DrawRectangle(float64(opts.HorizontalPadding), float64(opts.VerticalPadding), float64(inBounds.Dx()), headerHeight)
	dc.Fill()

	radius := 7
	dc.SetFillRule(gg.FillRuleEvenOdd)
	dc.DrawCircle(float64(opts.HorizontalPadding+13+radius), float64(opts.VerticalPadding+10+radius), float64(radius))
	dc.SetColor(color.RGBA{R: 238, G: 106, B: 93, A: 255})
	dc.Fill()

	dc.DrawCircle(float64(opts.HorizontalPadding+13+radius+13+radius), float64(opts.VerticalPadding+10+radius), float64(radius))
	dc.SetColor(color.RGBA{R: 231, G: 175, B: 89, A: 255})
	dc.Fill()

	dc.DrawCircle(float64(opts.HorizontalPadding+13+radius+13+radius+13+radius), float64(opts.VerticalPadding+10+radius), float64(radius))
	dc.SetColor(color.RGBA{R: 90, G: 180, B: 90, A: 255})
	dc.Fill()

	return dc.Image(), nil
}
