package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/google/uuid"
)

// Recipe is a game state machine step abstracted out as a cooking terminology
type Recipe struct {
	CookbookID    string // the cookbook guid
	Name          string
	ID            string // the recipe guid
	CoinInputs    CoinInputList
	ItemInputs    ItemInputList
	Entries       WeightedParamList
	Description   string
	BlockInterval int64
	Sender        sdk.AccAddress
	Disabled      bool
}

// RecipeList is a list of cookbook
type RecipeList struct {
	Recipes []Recipe
}

func (cbl RecipeList) String() string {
	output := "RecipeList{"
	for _, cb := range cbl.Recipes {
		output += cb.String()
		output += ",\n"
	}
	output += "}"
	return output
}

// NewRecipe creates a new recipe
func NewRecipe(recipeName, cookbookID, description string,
	coinInputs CoinInputList, // coinOutputs CoinOutputList,
	itemInputs ItemInputList, // itemOutputs ItemOutputList,
	entries WeightedParamList, // newly created param instead of coinOutputs and itemOutputs
	execTime int64, sender sdk.AccAddress) Recipe {
	rcp := Recipe{
		Name:          recipeName,
		CookbookID:    cookbookID,
		CoinInputs:    coinInputs,
		ItemInputs:    itemInputs,
		Entries:       entries,
		BlockInterval: execTime,
		Description:   description,
		Sender:        sender,
	}

	rcp.ID = rcp.KeyGen()
	return rcp
}

// NewRecipeWithGUID creates a new recipe with GUID
func NewRecipeWithGUID(GUID, recipeName, cookbookID, description string,
	coinInputs CoinInputList, // coinOutputs CoinOutputList,
	itemInputs ItemInputList, // itemOutputs ItemOutputList,
	entries WeightedParamList, // newly created param instead of coinOutputs and itemOutputs
	execTime int64, sender sdk.AccAddress) Recipe {
	// TODO if user send same GUID what to do? fail or random GUID generate internally?
	rcp := Recipe{
		ID:            GUID,
		Name:          recipeName,
		CookbookID:    cookbookID,
		CoinInputs:    coinInputs,
		ItemInputs:    itemInputs,
		Entries:       entries,
		BlockInterval: execTime,
		Description:   description,
		Sender:        sender,
	}

	if len(GUID) == 0 {
		rcp.ID = rcp.KeyGen()
	}
	return rcp
}

func (rcp *Recipe) String() string {
	return fmt.Sprintf(`Recipe{
		Name: %s,
		CookbookID: %s,
		ID: %s,
		CoinInputs: %s,
		ItemInputs: %s,
		Entries: %s,
		ExecutionTime: %d,
	}`, rcp.Name, rcp.CookbookID, rcp.ID,
		rcp.CoinInputs.String(),
		rcp.ItemInputs.String(),
		rcp.Entries.String(),
		rcp.BlockInterval)
}

// KeyGen generates key for the store
func (rcp Recipe) KeyGen() string {
	id := uuid.New()
	return rcp.Sender.String() + id.String()
}
