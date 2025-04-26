package main

import "fmt"

// ----- Mock types -----

type Change struct {
	Description string
}

type Status struct {
	Changes []Change
}

type Offer struct {
	Status Status
}

type ChangeResolver struct {
	change *Change
}

type OfferResolver struct {
	d *Offer
}

// ----- Method you wrote -----

func (r *OfferResolver) Changes() []ChangeResolver {
	var result []ChangeResolver

	if r == nil || r.d == nil {
		return result
	}

	changes := r.d.Status.Changes
	result = make([]ChangeResolver, 0, len(changes)) // preallocate

	for _, c := range changes {
		change := c
		result = append(result, ChangeResolver{&change})
	}

	return result
}

// ----- Main function -----

func main() {
	// Mock some data
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
}
