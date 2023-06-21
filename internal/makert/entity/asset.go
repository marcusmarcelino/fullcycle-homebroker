package entity

type Asset struct {
	ID           string
	Name         string
	MarketVolume int
}

// Podemos dizer que está agindo como um construtor para gerar a entidade Asset
// Mas é um método que precisa ser chamado para que a criação ocorra
func NewAsset(id string, name string, marketVolume int) *Asset {
	return &Asset{
		ID:           id,
		Name:         name,
		MarketVolume: marketVolume,
	}
}
