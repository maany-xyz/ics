package types

const (
    ModuleName   = "consumergov"
    StoreKey     = ModuleName
    RouterKey    = ModuleName
    QuerierRoute = ModuleName
)

// Param Store Keys
var (
    ParamStoreKeyAllowedConsumerChainID = []byte("AllowedConsumerChainID") // Allowed consumer chain
)
