package ConstantsTransactions

type ConstantsTransactions struct {
	SamplePathFile string
}

func Constants() ConstantsTransactions {
	constants := ConstantsTransactions{
		SamplePathFile: "./transactions.csv",
	}
	return constants
}
