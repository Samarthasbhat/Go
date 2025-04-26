// Pointers : shared, not copied (Concept of null)
// Values : Copied, not shared
// Stack allocation is more efficient
// Accessing a dense sequence of data is more efficient than sparse data (array is faster than linked list)
// Any struct with a mutex must be 
// Use the index if you need to mutate the element.
// Anytime a func mutates a slice that;s passed in, we must return a copy

package main

import (
	"fmt"
)


type Change struct{
	Description string
}

type Status struct{
	Change []Change
}

type offer struct{
	Status Status
}

type ChangeResolver struct{
	d *offer
}

func main(){
	items := [][2]byte{{1,2}, {3,4}, {5,6}}
	a := [][]byte{}

	for _, item := range items {
		i := make([]byte, len(item))

		copy(i, item[:]) // Make unique 
		a = append(a, i) // : slice the whole array
	}

	fmt.Println(items)
	fmt.Println(a)

	offer := &Offer{
		Status: Status{
			Changes: []Change{
				{Description: "Price updated"},
				{Description: "Terms updated"},
				{Description: "Discount applied"},
			},
		},
	}

	// Create the resolver
	resolver := &OfferResolver{d: offer}

	// Get changes
	changes := resolver.Changes()

	// Print them
	for i, changeResolver := range changes {
		fmt.Printf("Change %d: %s\n", i+1, changeResolver.change.Description)

}

func (r *OfferResolver) Changes() []ChangeResolver {
	var result []ChangeResolver

	if r == nil || r.d == nil {
		return result
	}

	changes := r.d.Status.Changes
	result = make([]ChangeResolver, 0, len(changes)) 

	for _, c := range r.d.Status.Changes{
		change := c

		result = append(result, ChangeResolver{&change})
	}
	return result
}