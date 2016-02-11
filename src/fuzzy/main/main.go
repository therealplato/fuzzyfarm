package main
import(
    //"fuzzy"
    "time"
    "fmt"
    "math"
)

const CostPerAnimal float64 = 0.001
const AnimalBreedRate float64 = 0.00001

type Animal struct {
    Name string           // what animal
    Count int             // how many animals
    DFuzzies float64      // change in Fuzzies per millisecond
    DCount float64            // change in animal count per millisecond
    Embryos float64        // Keeps track of breeding animals before birth
}

func(a *Animal) Buy(n int) {
    a.Count += n
    a.DFuzzies += (float64(n)*CostPerAnimal)
    a.DCount += (float64(n)*AnimalBreedRate)
}

func(a *Animal) Update() (int, bool) {
    // Breed new animals
    // Return income and error
    babyAnimalFloat := float64(a.Count) * a.DCount
    a.Embryos += babyAnimalFloat
    newborns := math.Floor(a.Embryos)
    a.Embryos -= newborns
    a.Count += int(newborns)
    floatIncome := float64(a.Count) * a.DFuzzies
    intIncome := int(floatIncome)
    return intIncome, true
}

type Farm struct {
  nFuzzies int            // currency amount
  Animals map[string] *Animal
}


func main(){
    TheFarm := &Farm {
        nFuzzies: 10,
        Animals: make(map[string]*Animal, 4),
    }
    TheFarm.Animals["cats"] = &Animal {
        Name: "Kitties",
        Count: 0,
        DFuzzies: 0,
        DCount: 0,
        Embryos: 0,
    }
    TheFarm.Animals["dogs"] = &Animal {
        Name: "Puppies",
        Count: 0,
        DFuzzies: 0,
        DCount: 0,
        Embryos: 0,
    }
    
    fmt.Println(TheFarm)
    c := time.Tick(1000 * time.Millisecond)
    
    quit := make(chan bool)
    go run(*TheFarm, c, quit)
    _ = <-quit
}

func run(TheFarm Farm, c <-chan time.Time, quit chan bool){
    TheFarm.Animals["cats"].Buy(200)
    for _ = range c {
        // 1 ms tick
        for a := range TheFarm.Animals {
            income, _ := TheFarm.Animals[a].Update()
            TheFarm.nFuzzies += income
        }
        fmt.Printf("Cats: %v  Dogs: %v  Fuzzies: %v\n", TheFarm.Animals["cats"].Count, TheFarm.Animals["dogs"].Count, TheFarm.nFuzzies)
    }
}