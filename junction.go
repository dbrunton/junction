package junction

import (
	"time",
	"fmt",
)

func all(junc ...func()(bool)) bool {

        c := make(chan bool)
        t := make(chan bool)

        go func() { time.Sleep(5e9); t <- true }()

        for _, f := range junc {
                r := f()
                go func() { c <- r }()
        }

        for {
                select {
                        case j := <-c:
                                // should probably have a case where everything's been received.
                                if !j { return false }
                        case <-t:
                                return true
                }
        }
        return true
}
