package handlers

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	core "github.com/iden3/go-iden3-core"
	"github.com/iden3/go-iden3-crypto/poseidon"
	"github.com/iden3/go-schema-processor/merklize"
	"github.com/iden3/go-schema-processor/utils"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokene/polygonid-issuer-integration/internal/service/helpers"
	"gitlab.com/tokene/polygonid-issuer-integration/internal/service/helpers/converter"
	"gitlab.com/tokene/polygonid-issuer-integration/internal/service/requests"
	"gitlab.com/tokene/polygonid-issuer-integration/internal/service/types"
	"gitlab.com/tokene/polygonid-issuer-integration/solidity/generated/zkpverifier"
	"io"
	"math/big"
	"net/http"
	"strings"
)

const PathToSubjectType = "https://www.w3.org/2018/credentials#credentialSubject"

func SetZKPRequest(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewSetZKPRequest(r)
	if err != nil {
		helpers.Log(r).WithError(err).Info("wrong request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	resp, err := http.Get(request.Data.Attributes.SchemaUrl)
	if err != nil {
		helpers.Log(r).WithError(err).Info("failed to get schema url content")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	schemaBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		helpers.Log(r).WithError(err).Info("failed to parse schema url content")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	schemaID, err := merklize.TypeIDFromContext(schemaBytes, request.Data.Attributes.SchemaType)
	if err != nil {
		helpers.Log(r).WithError(err).Info("failed to get type ID from schema content")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	querySchema := utils.CreateSchemaHash([]byte(schemaID)).BigInt().String()
	path, err := merklize.NewFieldPathFromContext(schemaBytes, request.Data.Attributes.SchemaType, request.Data.Attributes.FieldName)
	if err != nil {
		helpers.Log(r).WithError(err).Info("failed to get path from schema")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	if err = path.Prepend(PathToSubjectType); err != nil {
		helpers.Log(r).WithError(err).Info("failed to prepend path's parts from the beginning")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	mkPath, err := path.MtEntry()
	if err != nil {
		helpers.Log(r).WithError(err).Info("failed to get claim path key")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	queryHash, err := calculateQueryHash(
		request.Data.Attributes.Value,
		querySchema,
		request.Data.Attributes.Operator,
		request.Data.Attributes.SlotIndex,
		mkPath.String(),
		request.Data.Attributes.ClaimPathNotExists,
	)
	if err != nil {
		helpers.Log(r).WithError(err).Info("failed to calculate query hash")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	id := uuid.New().String()
	metadata, err := invokeRequestMetadata(r, *request, id)
	if err != nil {
		helpers.Log(r).WithError(err).Info("failed to invoke request metadata")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	data, err := getData(*request, querySchema, mkPath.String(), queryHash)
	if err != nil {
		helpers.Log(r).WithError(err).Info("failed to get data")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	auth, client, err := helpers.GetEthClient(r)
	if err != nil {
		helpers.Log(r).WithError(err).Info("failed to get eth client")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	erc20verifierContract, err := zkpverifier.NewZkpverifier(common.HexToAddress(request.Data.Attributes.VerifierAddress), client)
	tx, err := erc20verifierContract.SetZKPRequest(auth, uint64(request.Data.Attributes.RequestId), zkpverifier.IZKPVerifierZKPRequest{
		Metadata:  metadata,
		Validator: common.HexToAddress(*request.Data.Attributes.ValidatorAddress),
		Data:      data,
	})
	if err != nil {
		helpers.Log(r).WithError(err).Info("failed to set zkp request")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, converter.ToTransactionResource(id, request.Data.Attributes.VerifierAddress, tx.Hash().Hex()))
}

func calculateQueryHash(valuesRaw []int64, schema, operator string, slotIndex int64, claimPathKey string, claimPathNotExists bool) (string, error) {
	values := make([]*big.Int, 64)
	for i := 0; i < 64; i++ {
		if len(valuesRaw) > i {
			values[i] = big.NewInt(valuesRaw[i])
			continue
		}
		values[i] = big.NewInt(0)
	}

	valueHash, err := poseidon.SpongeHashX(values, 6)
	if err != nil {
		return "", err
	}

	schemaInt, ok := big.NewInt(0).SetString(schema, 10)
	if !ok {
		return "", errors.New("failed to set schema to big.Int")
	}
	schemaHash := core.NewSchemaHashFromInt(schemaInt)

	claimPathKeyBigInt, ok := big.NewInt(0).SetString(claimPathKey, 10)
	if !ok {
		return "", errors.New("failed to set schema to big.Int")
	}
	arr := []*big.Int{
		schemaHash.BigInt(),
		big.NewInt(slotIndex),
		big.NewInt(int64(types.StringOperator(operator))),
		claimPathKeyBigInt,
		big.NewInt(cast.ToInt64(claimPathNotExists)),
		valueHash,
	}

	queryHash, err := poseidon.Hash(arr)
	if err != nil {
		return "", err
	}
	return queryHash.String(), nil
}

func invokeRequestMetadata(r *http.Request, request requests.SetZKPRequest, id string) (string, error) {
	subject := make(map[string]interface{})
	subData := make(map[string]interface{})
	if types.StringOperator(request.Data.Attributes.Operator) == types.IN || types.StringOperator(request.Data.Attributes.Operator) == types.NIN {
		subData["$"+strings.ToLower(request.Data.Attributes.Operator)] = request.Data.Attributes.Value
	} else {
		subData["$"+strings.ToLower(request.Data.Attributes.Operator)] = request.Data.Attributes.Value[0]
	}
	subject[request.Data.Attributes.FieldName] = subData

	var allowedIssuers []string
	if request.Data.Attributes.AllowedIssuers != nil {
		allowedIssuers = *request.Data.Attributes.AllowedIssuers
	} else {
		allowedIssuers = append(allowedIssuers, "*")
	}

	metadata := types.RequestMetadata{
		Id:   id,
		Typ:  "application/iden3comm-plain-json",
		Type: "https://iden3-communication.io/proofs/1.0/contract-invoke-request",
		Thid: id,
		Body: struct {
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
		}{
			Reason: request.Data.Attributes.Reason,
			TransactionData: struct {
				ContractAddress string `json:"contract_address"`
				MethodId        string `json:"method_id"`
				ChainId         int64  `json:"chain_id"`
				Network         string `json:"network"`
			}{
				ContractAddress: request.Data.Attributes.VerifierAddress,
				MethodId:        "b68967e2",
				ChainId:         helpers.NetworkConfig(r).ChainId,
				Network:         helpers.NetworkConfig(r).Network,
			},
			Scope: []struct {
				Id        int64  `json:"id"`
				CircuitId string `json:"circuitId"`
				Query     struct {
					AllowedIssuers    []string    `json:"allowedIssuers"`
					Context           string      `json:"context"`
					CredentialSubject interface{} `json:"credentialSubject"`
					Type              string      `json:"type"`
				} `json:"query"`
			}{{
				Id:        request.Data.Attributes.RequestId,
				CircuitId: request.Data.Attributes.CircuitId,
				Query: struct {
					AllowedIssuers    []string    `json:"allowedIssuers"`
					Context           string      `json:"context"`
					CredentialSubject interface{} `json:"credentialSubject"`
					Type              string      `json:"type"`
				}{
					AllowedIssuers:    allowedIssuers,
					Context:           request.Data.Attributes.SchemaUrl,
					CredentialSubject: subject,
					Type:              request.Data.Attributes.SchemaType,
				},
			}},
		},
	}
	metadataJSON, err := json.Marshal(metadata)
	if err != nil {
		return "", errors.Wrap(err, "failed to marshal metadata")
	}
	return string(metadataJSON), nil
}

func getData(request requests.SetZKPRequest, schema, claimPathKey, queryHash string) ([]byte, error) {
	values := make([]*big.Int, 64)
	for i := 0; i < 64; i++ {
		if len(request.Data.Attributes.Value) > i {
			values[i] = big.NewInt(request.Data.Attributes.Value[i])
			continue
		}
		values[i] = big.NewInt(0)
	}
	credentialAtomicQuery, _ := abi.NewType("tuple", "CredentialAtomicQuery", []abi.ArgumentMarshaling{
		{Name: "schema", Type: "uint256"},
		{Name: "claimPathKey", Type: "uint256"},
		{Name: "operator", Type: "uint256"},
		{Name: "slotIndex", Type: "uint256"},
		{Name: "value", Type: "uint256[]"},
		{Name: "queryHash", Type: "uint256"},
		{Name: "allowedIssuers", Type: "uint256[]"},
		{Name: "circuitIds", Type: "string[]"},
		{Name: "skipClaimRevocationCheck", Type: "bool"},
		{Name: "claimPathNotExists", Type: "uint256"},
	})
	args := abi.Arguments{
		{Name: "CredentialAtomicQuery", Type: credentialAtomicQuery},
	}
	var allowedIssuers []*big.Int
	if request.Data.Attributes.AllowedIssuers != nil {
		for _, issuer := range *request.Data.Attributes.AllowedIssuers {
			did := new(core.DID)
			if err := did.SetString(issuer); err != nil {
				return nil, errors.Wrap(err, "failed to convert allowed issuer")
			}
			allowedIssuers = append(allowedIssuers, did.ID.BigInt())
		}
	}

	schemaBigInt, _ := big.NewInt(0).SetString(schema, 10)
	claimPathKeyBigInt, _ := big.NewInt(0).SetString(claimPathKey, 10)
	queryHashBigInt, _ := big.NewInt(0).SetString(queryHash, 10)
	record := types.CredentialAtomicQuery{
		Schema:                   schemaBigInt,
		ClaimPathKey:             claimPathKeyBigInt,
		Operator:                 big.NewInt(int64(types.StringOperator(request.Data.Attributes.Operator))),
		SlotIndex:                big.NewInt(request.Data.Attributes.SlotIndex),
		Value:                    values,
		QueryHash:                queryHashBigInt,
		AllowedIssuers:           allowedIssuers,
		CircuitIds:               []string{request.Data.Attributes.CircuitId},
		SkipClaimRevocationCheck: request.Data.Attributes.SkipClaimRevocationCheck,
		ClaimPathNotExists:       big.NewInt(cast.ToInt64(request.Data.Attributes.ClaimPathNotExists)),
	}
	//spew.Dump(record)
	packed, err := args.Pack(&record)
	//for len(packed) > 0 {
	//	r, size := utf8.DecodeRune(packed)
	//	fmt.Printf("%c", r)
	//	packed = packed[size:]
	//}
	if err != nil {
		return nil, errors.Wrap(err, "failed to pack data args")
	}
	return packed, nil
}
