# tcpsend

A net wrapper to send data over tcp4


# How to use

```go
package main

import (
	"fmt"
	"time"

	"github.com/angelside/tcpsend"
)

func main() {

	//

	ip := "127.0.0.1"
	port := "9100"
	//data := "~ZPL CODE~"
	data := `^XA
^LT0
^LH0,0
^JMA
~SD15
^CI28
^PA0,1,1,0
^FO1,1^GB810,164,148^FS
^FT4,347^A0N,72,71^FB808,1,18,C^FH\^CI28^FD{{.Location}}^FS^CI27
^FT4,512^A0N,51,51^FB808,1,13,C^FH\^CI28^FD{{.Message}}^FS^CI27
^FT120,657^A0N,25,25^FH\^CI28^FD{{.IpAddress}}^FS^CI27
^FT120,723^A0N,25,25^FH\^CI28^FD{{.Time}}^FS^CI27
^XZ`

	err := tcpsend.Run(ip, port, data, time.Second*1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("[OK] Data sent")
	}
}

```
