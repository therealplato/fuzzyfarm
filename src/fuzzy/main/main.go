package main
import(
    //"fuzzy"
    "time"
    "fmt"
    "math"
)

const CostPerAnimal float64 = 0.001
const IncomePerAnimal float64 = 0.000001   // per ms per animal
const AnimalBreedRate float64 = 0.0000001  // per ms per animal

type Farm struct {
  nFuzzies float64            // currency amount
  Animals map[string] *Animal
}


type Animal struct {
    Name string           // what animal
    Count int             // how many animals
    DIncome float64      // change in Fuzzies per millisecond
    DCount float64            // change in animal count per millisecond
    Embryos float64        // Keeps track of breeding animals before birth
}

func(a *Animal) Spawn(n int) {
    a.Count += n
    a.DIncome += (float64(n)*IncomePerAnimal)
    a.DCount += (float64(n)*AnimalBreedRate)
}

func(a *Animal) Update() (float64, bool) {
    // Breed new animals
    // Return income and error
    babyAnimalsFloat := float64(a.Count) * a.DCount
    a.Embryos += babyAnimalsFloat
    newborns := math.Floor(a.Embryos)
    a.Embryos -= newborns
    a.Spawn(int(newborns))
    return a.DIncome, true
}


func main(){
    TheFarm := &Farm {
        nFuzzies: 10,
        Animals: make(map[string]*Animal, 4),
    }
    TheFarm.Animals["cats"] = &Animal {
        Name: "Kitties",
        Count: 0,
        DIncome: 0,
        DCount: 0,
        Embryos: 0,
    }
    TheFarm.Animals["dogs"] = &Animal {
        Name: "Puppies",
        Count: 0,
        DIncome: 0,
        DCount: 0,
        Embryos: 0,
    }
    
    fmt.Println(TheFarm)
    c1 := time.Tick(1 * time.Millisecond)
    c1000 := time.Tick(1000 * time.Millisecond)
    quit := make(chan bool)
    go run(TheFarm, c1, quit)
    go output(TheFarm, c1000)
    _ = <-quit
}

func run(TheFarm *Farm, c <-chan time.Time, quit chan bool){
    TheFarm.Animals["cats"].Spawn(200)
    TheFarm.Animals["dogs"].Spawn(10)
    for _ = range c {
        // 1 ms tick
        for a := range TheFarm.Animals {
            income, _ := TheFarm.Animals[a].Update()
            TheFarm.nFuzzies += income
        }
    }
}

func output(TheFarm *Farm, c <-chan time.Time){
    for _ = range c {
        // 1 s tick
        fmt.Printf("Cats: %v  Dogs: %v  Fuzzies: %.3f\n", TheFarm.Animals["cats"].Count, TheFarm.Animals["dogs"].Count, TheFarm.nFuzzies)
    }
}