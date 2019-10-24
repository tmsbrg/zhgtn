// Copyright 2019 Thomas van der Berg
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

// Chinese listening game: Type the right number after hearing it

// to compile: run packr (from go get -u github.com/gobuffalo/packr/packr) then go build

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/gobuffalo/packr"

	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"
)

func run() error {

	player, err := oto.NewPlayer(32000, 2, 2, 8192)
	if err != nil {
		return err
	}
	defer player.Close()

	audio := packr.NewBox("./audio")

	rand.Seed(int64(time.Now().Nanosecond()))

	correct := 0
	times := 10
	for i := 0; i < times; i++ {

		n := rand.Intn(101)

		track, err := audio.Find(fmt.Sprintf("%03d.mp3", n))
		if err != nil {
			return err
		}

		decoder, err := mp3.NewDecoder(bytes.NewReader(track))
		if err != nil {
			return err
		}

		if _, err := io.Copy(player, decoder); err != nil {
			return err
		}

		input := bufio.NewScanner(os.Stdin)
		for {
			fmt.Print("Enter the number you heard: ")

			if input.Scan() == false {
				err := input.Err()
				if err != nil {
					return err
				} else {
					fmt.Println("Exited.")
					return nil
				}
			}

			guess, err := strconv.Atoi(input.Text())
			if err != nil {
				fmt.Println("That's not a number!")
				continue
			}
			if guess == n {
				fmt.Println("Good! 很好")
				correct += 1
				break
			}
			fmt.Println("Wrong! 不对! The number was ", n)
			break
		}
	}
	fmt.Printf("You got %d/%d (%d%%)\n", correct, times, correct * 100 / times)
	time.Sleep(2*time.Second)

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
