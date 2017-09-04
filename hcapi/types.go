// Copyright 2017 The hchart Authors, All rights reserved.
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

import (
	"sync"
)

var (
	chart_ls_mu       sync.RWMutex
	chart_datasets_mu sync.RWMutex
)

const (
	ChartTypeBar           = "bar"
	ChartTypeBarHorizontal = "bar-h"
	ChartTypeLine          = "line"
	ChartTypePie           = "pie"
)

type ChartOptions struct {
	Title  string `json:"title,omitempty"`
	Width  string `json:"width,omitempty"`
	Height string `json:"height,omitempty"`
}

type ChartEntry struct {
	Type    string       `json:"type"`
	Options ChartOptions `json:"options"`
	Data    ChartData    `json:"data"`
}

type ChartData struct {
	Labels   []string       `json:"labels,omitempty"`
	Datasets []ChartDataset `json:"datasets,omitempty"`
}

type ChartDataset struct {
	Label string  `json:"label,omitempty"`
	Data  []int64 `json:"data,omitempty"`
}

func (it *ChartData) Sync(d_label, ds_label string, ds_data int64) {

	chart_datasets_mu.Lock()
	defer chart_datasets_mu.Unlock()

	for k, v := range it.Datasets {
		if v.Label == ds_label {
			it.Datasets[k].Data = append(v.Data, ds_data)
			if len(it.Datasets[k].Data) > len(it.Labels) {
				it.Labels = append(it.Labels, d_label)
			}
			return
		}
	}

	it.Datasets = append(it.Datasets, ChartDataset{
		Label: ds_label,
		Data:  []int64{ds_data},
	})
	if len(it.Labels) < 1 {
		it.Labels = append(it.Labels, d_label)
	}
}

type ChartList struct {
	Items []ChartEntry `json:"items"`
}

func (it *ChartList) Sync(c_type, c_title, d_label, ds_label string, ds_data int64) {

	chart_ls_mu.Lock()
	defer chart_ls_mu.Unlock()

	for k, v := range it.Items {
		if v.Type == c_type && v.Options.Title == c_title {
			it.Items[k].Data.Sync(d_label, ds_label, ds_data)
			return
		}
	}

	it.Items = append(it.Items, ChartEntry{
		Type: c_type,
		Options: ChartOptions{
			Title: c_title,
		},
		Data: ChartData{
			Labels: []string{d_label},
			Datasets: []ChartDataset{
				{
					Label: ds_label,
					Data:  []int64{ds_data},
				},
			},
		},
	})
}
