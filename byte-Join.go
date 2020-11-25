package main
import (
    "bytes"
    "fmt"
    "encoding/binary"
)
func main()  {
    hello := "hello"
    helloBytes := []byte( hello )
    fmt.Println( helloBytes )

    world := "world"
    worldBytes := []byte( world )
    fmt.Println( worldBytes )

    helloWord := [][]byte{ helloBytes,worldBytes }
    fmt.Println( helloWord )


    helloWords := bytes.Join(helloWord,[]byte{})
    fmt.Println( helloWords )

        beBuf := make([]byte, 4)
        size := uint32(len(helloWords)) + 4
        binary.BigEndian.PutUint32(beBuf, size)

        beBuf2 := make([]byte, 4)
        binary.BigEndian.PutUint32(beBuf2, uint32(1))

        bbff := [][]byte{beBuf, beBuf2, helloWords}
    prefix_helloWords := bytes.Join(bbff,[]byte{})

    fmt.Println( prefix_helloWords )
}
