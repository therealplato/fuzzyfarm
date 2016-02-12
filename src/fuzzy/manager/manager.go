// One Manager handles one farm
package manager
import(
    "fmt"
    "fuzzy/farm"
    "fuzzy/models"
)
type Manager struct {
    Farm *farm.Farm
}

func NewManager() (*Manager){
   return &Manager{}
}

func (m *Manager) StartFarm(models.FarmSocket){
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
}