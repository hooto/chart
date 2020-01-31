// Copyright 2020 Eryx <evorui аt gmail dοt com>, All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hcapi

const (
	ChartTypeLine      = "line"
	ChartTypeBar       = "bar"
	ChartTypeHistogram = "histogram"
)

type ChartItem struct {
	Type     string       `json:"type"`
	Options  ChartOptions `json:"options"`
	Labels   []string     `json:"labels,omitempty"`
	Datasets []*DataItem  `json:"datasets,omitempty"`
}

func (it *ChartItem) Valid() error {
	return nil
}

type ChartOptions struct {
	Title  string      `json:"title,omitempty"`
	Width  string      `json:"width,omitempty"`
	Height string      `json:"height,omitempty"`
	X      AxisOptions `json:"x,omitempty"`
	Y      AxisOptions `json:"y,omitempty"`
}

type AxisOptions struct {
	Title string `json:"title,omitempty"`
}

func (it *ChartOptions) WidthLength() float64 {
	return 800
}

func (it *ChartOptions) HeightLength() float64 {
	return 400
}

type ChartRenderOptions struct {
	Name      string `json:"name"`
	SvgEnable bool   `json:"svg_enable"`
	PngEnable bool   `json:"png_enable"`
}
