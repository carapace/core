package cellar_test

import (
	"fmt"
	"github.com/carapace/cellar"
)

func Example() {
	folder := "./db"
	writer, err := cellar.NewWriter(folder, 1000, []byte("testkeylongerthanbeforen"))
	if err != nil {
		panic(err)
	}

	for i := 0; i < 10; i++ {
		writer.Append([]byte(fmt.Sprintf("appending: %d", i)))
	}

	_, err = writer.Checkpoint()
	if err != nil {
		panic(err)
	}
	err = writer.SealTheBuffer()
	if err != nil {
		panic(err)
	}
	err = writer.Close()
	if err != nil {
		panic(err)
	}

	reader := cellar.NewReader(folder, []byte("testkeylongerthanbeforen"))

	err = reader.Scan(func(pos *cellar.ReaderInfo, data []byte) error {
		fmt.Println(string(data))
		return nil
	})

	if err != nil {
		panic(err)
	}
	// Output:
	// appending: 0
	// appending: 1
	// appending: 2
	// appending: 3
	// appending: 4
	// appending: 5
	// appending: 6
	// appending: 7
	// appending: 8
	// appending: 9
}
