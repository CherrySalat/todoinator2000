package model

var LocalData = MemStore{
	data: map[string]Todo{
		"1": {
			Title:       "Buy some bread",
			Description: "1. Go to market\n2.take some bread\n3.Buy it",
		},
		"2": {
			Title:       "Be cool",
			Description: "1. Put on black glasses 2. Kick some asses 3. Drink beer.",
		},
	},
}
