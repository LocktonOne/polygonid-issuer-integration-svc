package types

type RequestMetadata struct {
	Id   string `json:"id"`
	Typ  string `json:"typ"`
	Type string `json:"type"`
	Thid string `json:"thid"`
	Body struct {
		Reason          string `json:"reason"`
		TransactionData struct {
			ContractAddress string `json:"contract_address"`
			MethodId        string `json:"method_id"`
			ChainId         int64  `json:"chain_id"`
			Network         string `json:"network"`
		} `json:"transaction_data"`
		Scope []struct {
			Id        int64  `json:"id"`
			CircuitId string `json:"circuitId"`
			Query     struct {
				AllowedIssuers    []string    `json:"allowedIssuers"`
				Context           string      `json:"context"`
				CredentialSubject interface{} `json:"credentialSubject"`
				Type              string      `json:"type"`
			} `json:"query"`
		} `json:"scope"`
	} `json:"body"`
}
