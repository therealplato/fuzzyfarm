package main
import(
    "fuzzy/farm"
    "time"
    "fmt"
)

func main(){
    TheFarm := &farm.Farm {
        NFuzzies: 10,
        Animals: make(map[string]*farm.Animal, 4),
    }
    TheFarm.Animals["cats"] = &farm.Animal {
        Name: "Kitties",
        Count: 0,
        DIncome: 0,
        DCount: 0,
        Embryos: 0,
    }
    TheFarm.Animals["dogs"] = &farm.Animal {
        Name: "Puppies",
        Count: 0,
        DIncome: 0,
        DCount: 0,
        Embryos: 0,
    }
    
    //fmt.Println(TheFarm)
    c1 := time.Tick(1 * time.Millisecond)
    c1000 := time.Tick(1000 * time.Millisecond)
    quit := make(chan bool)
    go run(TheFarm, c1, quit)
    go output(TheFarm, c1000)
    _ = <-quit
}

func run(TheFarm *farm.Farm, c <-chan time.Time, quit chan bool){
    TheFarm.Animals["cats"].Spawn(200)
    TheFarm.Animals["dogs"].Spawn(10)
    for _ = range c {
        // 1 ms tick
        for key := range TheFarm.Animals {
            a := TheFarm.Animals[key]
            income, _ := a.Update()
            TheFarm.NFuzzies += income
            if(a.Count < 0){
                quit <- true
            }
        }
    }
}

func output(TheFarm *farm.Farm, c <-chan time.Time){
    for _ = range c {
        // 1 s tick
        fmt.Printf("Cats: %v  Dogs: %v  Fuzzies: %.3f\n", TheFarm.Animals["cats"].Count, TheFarm.Animals["dogs"].Count, TheFarm.NFuzzies)
    }
}