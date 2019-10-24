# zhgtn
Game to learn Chinese numbers (0 to 100)


A game that will say a number in Chinese, then lets you write what it was. To practice your Chinese number listening skills.

Based on some nice libraries: [go-mp3](github.com/hajimehoshi/go-mp3) and [oto](github.com/hajimehoshi/oto) to play MP3 files,
and [packr](https://github.com/gobuffalo/packr) to pack the MP3s into the executable.

Tested on Linux and Windows with go version go1.13.3 linux/amd64

## compile

 - `go get -u github.com/gobuffalo/packr/packr`
 - `packr`
 - `go build`
 - `GOOS=windows go build` # make a Windows cross-compile build

## license

<pre>
Copyright 2019 Thomas van der Berg

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
 </pre>
 
 ## Note on go-mp3
 
As I wrote this, [go-mp3](github.com/hajimehoshi/go-mp3) only supports MP3s using MPEG-1 Audio layer III (v1). To test which version
your mp3 is, run: `file audio/000.mp3`

Output example:
```
audio/000.mp3: Audio file with ID3 version 2.4.0, contains:MPEG ADTS, layer III, v1,  40 kbps, 32 kHz, Monaural
```
Note the `v1`.

To convert an mp3 file to v1, you can use ffmpeg:
```
ffmpeg -i input.mp3 -ar 32000 -ab 32k -codec:a libmp3lame -ac 1 output.mp3
```
Values are taken from the [bitrate table for MPEG-1 from the Wikipedia article on mp3](https://en.wikipedia.org/wiki/MP3#Bit_rate).
