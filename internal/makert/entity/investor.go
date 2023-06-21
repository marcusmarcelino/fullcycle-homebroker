package entity

// Classe investidor
type Investor struct {
	ID            string
	Name          string
	AssetPosition []*InvestorAssetPosition
}

// Classe posição de ativo do investidor
type InvestorAssetPosition struct {
	AssetID string
	Shares  int
}

// Função que cria um novo objeto do tipo Investor/Investidor
func NewInvestor(id string) *Investor {
	return &Investor{
		ID:            id,
		AssetPosition: []*InvestorAssetPosition{},
	}
}

// funcção que adiciona um assetPosition dentro do AssetPosition do Investor, como o ponteiro está sendo usado
// Ele irá alterar diretamente o objeto na memória e não numa referência, sendo assim, em todos os lugares que estiverem usando essa entidade
// eles receberam essa alteração
func (i *Investor) AddAssetPosition(assetPosition *InvestorAssetPosition) {
	i.AssetPosition = append(i.AssetPosition, assetPosition)
}

// Função para alteração de uma AssetPosition
func (i *Investor) UpdateAssetPosition(assetID string, qtdShares int) {
	// Acessou o método get para pegar a assetPosition específica
	assetPosition := i.GetAssetPosition(assetID)

	// verifica se essa ação já existe dentro da carteira do investidor
	// caso não escista ela adicionar um novo objeto da ação e a quantidade
	if assetPosition == nil {
		i.AssetPosition = append(i.AssetPosition, NewInvestorAssetPosition(assetID, qtdShares))
	} else {
		// do contrário ela apenas aumenta sua posição/quantidade atual
		assetPosition.Shares += qtdShares
	}
}

// Função que acessa o objeto Investor e retorna a assetPositoin que será alterada, de acordo com o assetid que foi passado
// ele acessa o objeto Investor e entra no array de AssetPosition e percorre o array em busca do asset com o id especificado
func (i *Investor) GetAssetPosition(assetID string) *InvestorAssetPosition {
	// key e value d o laço retornado pelo range
	for _, assetPosition := range i.AssetPosition {
		if assetPosition.AssetID == assetID {
			return assetPosition
		}
	}
	return nil
}

// Função que retorna a criação de um nova posição de ação
func NewInvestorAssetPosition(assetID string, shares int) *InvestorAssetPosition {
	return &InvestorAssetPosition{
		AssetID: assetID,
		Shares:  shares,
	}
}
