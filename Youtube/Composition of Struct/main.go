package main

import (
	"fmt"
	"path/filepath"
	"sort"
)

// Embedding struct in another struct
// A struct can embed a pointer to another type

type Pair struct{
	Path string
	Hash string
}


func (p Pair) String() string{
	return fmt.Sprintf("Path: %s, Hash: %s", p.Path, p.Hash)
}

type PairWithLength struct {
	Pair
	Length int
}

func (p PairWithLength) String() string{
	return fmt.Sprintf("Path: %s, Hash: %s: Length: %d", p.Path, p.Hash, p.Length)
}

func (p Pair)Filename () string {
	return filepath.Base(p.Path)
}

type Filenamer interface{
	Filename() string
}

func main() {
	p := Pair{"path/to/file", "0xfdfe"}
	// pl := PairWithLength{Pair{"usr/lib", "0xdead"}, 133}


	var fn Filenamer = PairWithLength{Pair{"usr/lib", "0xdead"}, 133}

	fmt.Println(p)
	
	fmt.Println(fn)

	// fmt.Println(Filename(pl))
	// We cannot access the Pair method directly from PairWithLength

	fmt.Println(fn.Filename()) // This will work because PairWithLength embeds Pair
	
	// Organs

	s := []Organ{{"Heart", 300}, {"Kidney", 3000}, {"Liver", 100}}

	sort.Sort(ByWeight{s})

	fmt.Println(s)

	sort.Sort(ByName{s})
	fmt.Println(s)
	

}



// sort.Interface is a good example of embedding structs in Go
// defined in the sort package
// type StringSlice []string

type Organ struct {
	Name string
	Weight int
}

type Organs []Organ 

func (s Organs) Len() int { return len(s)}

func (s Organs) Swap(i, j int) { s[i], s[j] = s[j], s[i] }


type ByName struct { Organs }
type ByWeight struct { Organs }

func (s ByName) Less(i, j int) bool { return s.Organs[i].Name < s.Organs[j].Name }

func (s ByWeight) Less(i, j int) bool { return s.Organs[i].Weight < s.Organs[j].Weight }
