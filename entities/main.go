package entities

import (
	"bufio"
	"encoding/json"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
)

type Entity struct {
	Type      string `json:"type"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	NetworkID uint64 `json:"network_id"`
}

func (e *Entity) CommonAddress() common.Address {
	return common.HexToAddress(e.Address)
}

func getEntitiesFromFile(fn string) []Entity {
	var result []Entity
	f, err := os.OpenFile(fn, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		var lineEntity Entity
		line := sc.Bytes() // GET the line string
		err = json.Unmarshal(line, &lineEntity)
		if err != nil {
			log.Fatalf("error unmarchalling routers: %v", err)
		}
		result = append(result, lineEntity)
	}
	return result
}

func GetRouters() []Entity {
	return getEntitiesFromFile("./entities/routers.txt")
}

func GetTokens() []Entity {
	return getEntitiesFromFile("./entities/tokens.txt")
}

func GetSourceTokens() []Entity {
	return getEntitiesFromFile("./entities/source_tokens.txt")
}

func GetRouterPairs() [][2]Entity {
	result := [][2]Entity{}
	routers := GetRouters()
	for _, r1 := range routers {
		for _, r2 := range routers {
			if r1.Address != r2.Address {
				result = append(result, [2]Entity{r1, r2})
			}
		}
	}
	return result
}

func GetTokenPairs() [][2]Entity {
	result := [][2]Entity{}

	for _, sourceTkn := range GetSourceTokens() {
		for _, tkn := range GetTokens() {
			result = append(result, [2]Entity{sourceTkn, tkn})
		}
	}
	return result
}

func GetDualDexPermutations() [][4]Entity {
	result := [][4]Entity{}
	for _, routerPair := range GetRouterPairs() {
		for _, tknPair := range GetTokenPairs() {
			result = append(result, [4]Entity{routerPair[0], routerPair[1], tknPair[0], tknPair[1]})
		}
	}
	return result
}
