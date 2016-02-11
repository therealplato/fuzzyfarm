package farm
import(
    "math"
)

const CostPerAnimal float64 = 0.001
const IncomePerAnimal float64 = 0.000001   // per ms per animal
const AnimalBreedRate float64 = 0.0000001  // per ms per animal

type Farm struct {
  NFuzzies float64            // currency amount
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

