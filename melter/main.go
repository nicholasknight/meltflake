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

// Melter takes a command line consisting of an optional -variant parameter and
// any number of numerical strings representing Snowflake IDs, and prints out
// the time, worker ID, process ID, and increment corresponding to each provided
// Snowflake.
//
// If the variant is not specified, the Snowflake is parsed and displayed for
// every supported variant.
//
// Current supported variants are Twitter and Discord.
//
//  ~/gocode/bin$ ./melter -variant=Discord 381898139189116930
//  381898139189116930 Discord: 2017-11-19 20:06:51.707 Z, w 0, p 0, i 2
//  ~/gocode/bin$
//
package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"

	"github.com/nicholasknight/meltflake"
)

var knownVariants = map[string]meltflake.Variant{
	"Discord": meltflake.Discord,
	"Twitter": meltflake.Twitter,
}

func main() {
	variantFlag := flag.String("variant", "", "Name of specific Snowflake variant to melt.")
	flag.Parse()

	var variants []string

	if variantFlag == nil || *variantFlag == "" {
		for k := range knownVariants {
			variants = append(variants, k)
		}
	} else {
		if _, ok := knownVariants[*variantFlag]; ok {
			variants = append(variants, *variantFlag)
		} else {
			log.Fatalf("Unknown variant '%s'.", *variantFlag)
		}
	}

	snowflakes := flag.Args()

	for _, snowStr := range snowflakes {
		snowflake, err := strconv.ParseUint(snowStr, 10, 64)
		if err != nil {
			log.Fatalf("Invalid ID (must be digits) '%s': %s", snowStr, err)
		}
		for _, k := range variants {
			sp := meltflake.Melt(snowflake, knownVariants[k])
			fmt.Printf("%d %s: %s\n", snowflake, k, sp)
		}
	}
}
