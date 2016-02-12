// One Manager handles one farm
package manager
import(
    "fmt"
    "time"
    "fuzzy/farm"
    "fuzzy/models"
)
type Manager struct {
    Farm *farm.Farm
    FSocket models.FarmSocket
}

func NewManager(fs models.FarmSocket) (*Manager){
   return &Manager{
       FSocket: fs,
   }
}

func (m *Manager) StartFarm(){
    fmt.Println("Starting farm")
    m.Farm = &farm.Farm {
        NFuzzies: 10,
        Animals: make(map[string]*farm.Animal, 4),
    }
    m.Farm.Animals["cats"] = &farm.Animal {
        Name: "Kitties",
        Count: 0,
        DIncome: 0,
        DCount: 0,
        Embryos: 0,
    }
    m.Farm.Animals["dogs"] = &farm.Animal {
        Name: "Puppies",
        Count: 0,
        DIncome: 0,
        DCount: 0,
        Embryos: 0,
    }
    go m.ManageFarm()
}

func (m *Manager) ManageFarm(){
    c1 := time.Tick(1 * time.Millisecond)
    c1000 := time.Tick(1000 * time.Millisecond)
    go m.updateFarmLoop(c1)
    go m.outputFarmLoop(c1000)
}

func (m *Manager) updateFarmLoop(c1 <-chan time.Time){
    m.Farm.Animals["cats"].Spawn(20)
    m.Farm.Animals["dogs"].Spawn(10)
    for _ = range c1 {
        // 1 ms tick
        for key := range m.Farm.Animals {
            a := m.Farm.Animals[key]
            income, _ := a.Update()
            m.Farm.NFuzzies += income
            if(a.Count < 0){
                fmt.Fprintln(m.FSocket, "Animal Overflow")
                m.FSocket.Close()
            }
        }
    }    
}

func (m *Manager) outputFarmLoop(c1000 <-chan time.Time){
    for _ = range c1000 {
        // 1 s tick
        fmt.Printf("Cats: %v  Dogs: %v  Fuzzies: %.3f\n", m.Farm.Animals["cats"].Count, m.Farm.Animals["dogs"].Count, m.Farm.NFuzzies)
        fmt.Fprintln(m.FSocket, m.Farm.NFuzzies)
    }

}