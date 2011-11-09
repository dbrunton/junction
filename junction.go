package junction

import "time"

func all(junc ...func()(bool)) bool {

        c := make(chan bool)
        t := make(chan bool)

        go func() { time.Sleep(1e10); t <- true }()

        for _, f := range junc {
		go one(f, c)
        }

        for i := 0; i < cap(junc); i++ {
                select {
                        case j := <-c:
                                if j != true { return false }
                        case <-t:
                                return true
                }
        }
        return true
}

func one(junc func()(bool), ch chan bool)() {
	ch <- junc()
}
