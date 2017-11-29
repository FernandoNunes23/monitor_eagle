package parser

import (
	"bufio"
	"fmt"
	"os"
	"io"
	"bytes"
	"strconv"
	"types"
)

var buf    bytes.Buffer

func ParseMemory(file io.Reader) types.Memory {
	mem := types.Memory{}

	scanner := bufio.NewScanner(file)
	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)
	// Count the words.
	count := 0

	for scanner.Scan() {
		
		if count == 1  { mem.Total 		= strconv.Itoa(ConvertKbToMb(scanner.Text()))  }
		if count == 4  { mem.Free  		= strconv.Itoa(ConvertKbToMb(scanner.Text()))  }
		if count == 7  { mem.Available 	= strconv.Itoa(ConvertKbToMb(scanner.Text()))  }
		if count == 10 { mem.Buffers 	= strconv.Itoa(ConvertKbToMb(scanner.Text()))  }
		if count == 13 { mem.Cache 		= strconv.Itoa(ConvertKbToMb(scanner.Text()))  }
		count++
	}

	mem.Used = strconv.Itoa(GetUsedAmount(mem.Total, mem.Free))

	fmt.Print(&buf)

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	return mem
}

func ParseSwap(file io.Reader) types.Swap {
	swap := types.Swap{}

	scanner := bufio.NewScanner(file)
	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)
	// Count the words.
	count := 0

	for scanner.Scan() {
		if count == 43 { swap.Total = strconv.Itoa(ConvertKbToMb(scanner.Text()))  }
		if count == 46 { swap.Free	= strconv.Itoa(ConvertKbToMb(scanner.Text()))  }
		count++
	}

	swap.Used =  strconv.Itoa(GetUsedAmount(swap.Total, swap.Free))

	fmt.Print(&buf)

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	return swap
}