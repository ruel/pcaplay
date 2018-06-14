# pcaplay

Replays application layer packets (enclosed in either tcp or udp) from pcap files to a connection

## Installation 

```sh
$ go install github.com/ruel/pcaplay
```

## Usage

```
pcaplay -file <pcap-file> [-port <listening-port>] [-proto tcp|udp] [-bpf <bpf-filter>] [-delay <packet-delay-ms>] [-wait <wait-flag-toggle>]
```

|Option|Description|
|----|----|
| **bpf** | Berkeley Packet Filter (BPF) string to isolate packets |
| **delay** | Delay between sent packets in milliseconds *(default 100)* |
| **file** | Location of PCAP file to replay |
| **port** | Layer 4 listening port *(default 8484)* |
| **proto** | Layer 4 protocol to use *(default "tcp")* |
| **wait** | Wait for first packet before starting replay *(default true)* |

## BPF Syntax

Pcap files can be prepared and trimmed with [Wireshark](https://www.wireshark.org/). In some cases, pcap files can be very large or unmanageable in such softwares
and warrants on the fly (while reading) filtering. This can be achieved by using **BPF** or **Berkley Packet Filter**.

**pcaplay** accepts a bpf filter option to localize packets to be sent to the connection. Most common filters can be:

* `host 1.1.1.1` - Isolates packets with the given host in either directions
* `port 1234` - Returns packets with connections that uses given port in either directions
* `dst port 1234` - Returns packets with a given destination port

More information about the syntax can be found at this link: http://biot.com/capstats/bpf.html

## License

MIT License

Copyright (c) 2018 Ruel Pagayon

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
