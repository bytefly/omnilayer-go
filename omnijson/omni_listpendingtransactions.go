package omnijson

/*
Result
[                       (array of tx)
	{
    "txid": "0e37d459c7a777bf46278b054573bf5df131e09426ff4eb560b483b9271e141c",
    "fee": "0.00036720",
    "sendingaddress": "1Hw7fAJpzR84WCi2CyG6w5r3i9ztS4wh3s",
    "referenceaddress": "16XXgty328LTStJtzY8ut26CqhZtJApzM7",
    "ismine": false,
    "version": 0,
    "type_int": 0,
    "type": "Simple Send",
    "propertyid": 31,
    "divisible": true,
    "amount": "742.00000000",
    "confirmations": 0
	}
]
*/

type OmniListPendingTransactionsResult []struct {
	ID            string `json:"txid"`
	Fee           string `json:"fee"`
	From          string `json:"sendingaddress"`
	To            string `json:"referenceaddress"`
	Mine          bool   `json:"ismine"`
	Version       int32  `json:"version"`
	TypeInt       int32  `json:"type_int"`
	Type          string `json:"type"`
	PropertyID    uint32 `json:"propertyid"`
	Divisible     bool   `json:"divisible"`
	Amount        string `json:"amount"`
	Confirmations uint32 `json:"confirmations"`
}

type OmniListPendingTransactionsCommand struct {
	Address string
}

func (OmniListPendingTransactionsCommand) Method() string {
	return "omni_listpendingtransactions"
}

func (OmniListPendingTransactionsCommand) ID() string {
	return "1"
}

func (cmd OmniListPendingTransactionsCommand) Params() []interface{} {
	return []interface{}{cmd.Address}
}
