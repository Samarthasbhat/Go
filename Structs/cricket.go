package main

import ("fmt")

type Batsman struct{
	Name string
	Runs int
	Matches int
	Balls int
}

// Method to calculate the batting average
func (b Batsman) BattingAverage() float64{
	if b.Matches == 0{
		return 0.0
	}
	return float64(b.Runs) / float64(b.Matches)
}

// Method to calculate the strike rate
func(b Batsman) StrikeRate() float64{
	if b.Balls == 0{
		return 0.0
	}
	return (float64(b.Runs) / float64(b.Balls)) * 100
}

type Bowler struct{
	Name  string
	Wickets int
	RunsGiven int
	Overs int
}

// Method to calculate the bowling economy
func (bo Bowler) EconomyRate() float64{
	if bo.Overs == 0{
		return 0.0
	}
	return float64(bo.RunsGiven) / float64(bo.Overs)
}

// Method to calculate average wickets per match
func (bo Bowler) AverageWickets(matches int) float64 {
    if matches == 0 {
        return 0.0
    }
    return float64(bo.Wickets) / float64(matches)
}


func main(){
	
	// Create a Batsman instance
	batsman := Batsman{
		Name: "Virat Kohli",
		Runs: 12000,
		Matches: 250,
		Balls: 900,
	}

	    // Create a Bowler instance
		bowler := Bowler{
			Name:     "Jasprit Bumrah",
			Wickets:  320,
			RunsGiven: 8500,
			Overs:    1500,
		}
	

	 // Display batsman's stats
	fmt.Println("-----BATSMAN-----")
	fmt.Printf("Batsman: %s\n", batsman.Name)
	fmt.Printf("Batting Average: %.2f\n", batsman.BattingAverage())
	fmt.Printf("Strike Rate: %.2f\n", batsman.StrikeRate())
	println()

	 // Display bowler's stats
	 fmt.Println("----BOWLER-----")
	 fmt.Printf("Bowler: %s\n", bowler.Name)
	 fmt.Printf("Economy Rate: %.2f\n", bowler.EconomyRate())
	 fmt.Printf("Average Wickets per Match (for 200 matches): %.2f\n", bowler.AverageWickets(200))
 
}