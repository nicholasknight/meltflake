// Copyright 2018 Nicholas Knight
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package meltflake provides a simple Melt function to parse and extract
// information from Snowflake IDs.
//
//  import "github.com/nicholasknight/meltflake"
//  ...
//  flake := meltflake.Melt(381898139189116930, meltflake.Discord)
//  fmt.Print(flake)
package meltflake

import (
	"fmt"
	"time"
)

// Snowflake contains the parsed out fields of a snowflake ID.
//
// Snowflake conforms to the typical Stringer interface.
type Snowflake struct {
	Time      time.Time
	WorkerID  uint64
	ProcessID uint64
	Increment uint64
}

const timeFormat = "2006-01-02 15:04:05.999999999 Z07:00"

func (sf Snowflake) String() string {
	return fmt.Sprintf("%s, w %d, p %d, i %d",
		sf.Time.UTC().Format(timeFormat), sf.WorkerID, sf.ProcessID, sf.Increment)
}

// Variant holds information concerning specific services' Snowflake
// configurations. Currently we support variants that differ only in Epoch.
type Variant struct {
	Epoch uint64
}

// Twitter Snowflakes count from 2010-11-04 01:42:54.657 Z
var Twitter = Variant{
	Epoch: 1288834974657,
}

// Discord Snowflakes count from 2015-01-01 00:00:00 Z
var Discord = Variant{
	Epoch: 1420070400000,
}

// Melt parses the provided `snowflake` according to the rules of the `variant`.
func Melt(snowflake uint64, variant Variant) Snowflake {
	var sp Snowflake
	ts := ((snowflake >> 22) + variant.Epoch)
	sp.Time = time.Unix(int64(ts/1000), int64((ts%1000)*1000000))
	sp.WorkerID = (snowflake & 0x3E0000) >> 17
	sp.ProcessID = (snowflake & 0x1F000) >> 12
	sp.Increment = snowflake & 0xFFF
	return sp
}
