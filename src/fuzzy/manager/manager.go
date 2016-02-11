// One Manager handles one farm
package manager
import(
    "fuzzy/farm"
)
type Manager struct {
    Farm *farm.Farm
}
func (m *Manager) StartFarm() (*Manager){
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
    return m;
}