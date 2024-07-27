package consts

// TimeInForce ...
type TimeInForce string

const (
	// TimeInForceFAK ...
	TimeInForceFAK = TimeInForce("FAK")
	// TimeInForceFAS ...
	TimeInForceFAS = TimeInForce("FAS")
	// TimeInForceFOK ...
	TimeInForceFOK = TimeInForce("FOK")
	// TimeInForceSOK ...
	TimeInForceSOK = TimeInForce("SOK")
)
