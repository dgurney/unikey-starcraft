package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/dgurney/unikey/generator"
	"github.com/dgurney/unikey/validator"
)

/*
   Copyright (C) 2023 Daniel Gurney
   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.
   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.
   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

const version = "0.5.4"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	bench := flag.Int("bench", 0, "Benchmark generation of N keys")
	ver := flag.Bool("ver", false, "Show version information and exit.")
	t := flag.Bool("t", false, "Show elapsed time")
	repeat := flag.Int("r", 1, "Repeat N times")
	validate := flag.String("v", "", "Validate a key.")
	flag.Parse()

	if *ver {
		fmt.Printf("unikey-starcraft v%s by Daniel Gurney\n", version)
		return
	}

	if *bench > 0 {
		generationBenchmark(*bench)
		return
	}

	if *repeat < 1 {
		*repeat = 1
	}

	var started time.Time
	if *t {
		started = time.Now()
	}

	if *validate != "" {
		k := validator.StarCraft{Key: *validate}
		err := k.Validate()
		switch {
		case err != nil:
			fmt.Printf("%s is invalid: %s\n", *validate, err)
		default:
			fmt.Printf("%s is valid\n", *validate)
		}
		return
	}

	for i := 0; i < *repeat; i++ {
		key := generator.StarCraft{}
		err := key.Generate()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(key.String())
	}

	if *t {
		var ended time.Duration
		switch {
		case time.Since(started).Round(time.Second) > 1:
			ended = time.Since(started).Round(time.Millisecond)
		default:
			ended = time.Since(started).Round(time.Microsecond)
		}
		if ended < 1 {
			// Oh Windows...
			fmt.Println("Could not display elapsed time correctly :(")
			return
		}
		switch {
		case *repeat > 1:
			fmt.Printf("Took %s to generate %d keys.\n", ended, *repeat)
			return
		case *repeat == 1:
			fmt.Printf("Took %s to generate %d key.\n", ended, *repeat)
			return
		}
	}
}
