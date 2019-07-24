package businessFacades

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/stellar/go/build"
	"github.com/stellar/go/xdr"
	"github.com/zeemzo/jigsaw-gateway/api/apiModel"
	"github.com/zeemzo/jigsaw-gateway/dao"
	"github.com/zeemzo/jigsaw-gateway/model"
	"github.com/zeemzo/jigsaw-gateway/proofs/builder"
)

/*UserICOJIGXUHandler @desc Handles an incoming request and calls the BuildUserICOJIGXU
@author - Azeem Ashraf
@params - ResponseWriter,Request
*/
func UserICOJIGXUHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var userICOAPI model.UserICOAPI

	if r.Header == nil {
		w.WriteHeader(http.StatusBadRequest)
		result := apiModel.SubmitXDRSuccess{
			Status: "No Header present!",
		}
		json.NewEncoder(w).Encode(result)
		return
	}

	if r.Header.Get("Content-Type") == "" {
		w.WriteHeader(http.StatusBadRequest)
		result := apiModel.SubmitXDRSuccess{
			Status: "No Content-Type present!",
		}
		json.NewEncoder(w).Encode(result)

		return
	}

	err := json.NewDecoder(r.Body).Decode(&userICOAPI)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		result := apiModel.SubmitXDRSuccess{
			Status: "Error while Decoding the body",
		}
		json.NewEncoder(w).Encode(result)
		fmt.Println(err)
		return
	}
	fmt.Println(userICOAPI)

	display := &builder.AbstractICOBuilder{UserICOAPI: userICOAPI}
	display.BuildUserICOJIGXU(w, r)
	return
}

/*UserICOJIGXUHandler @desc Handles an incoming request and calls the BuildUserICOJIGXU
@author - Azeem Ashraf
@params - ResponseWriter,Request
*/
func UserICOXLMHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var userICOAPI model.UserICOAPI

	if r.Header == nil {
		w.WriteHeader(http.StatusBadRequest)
		result := apiModel.SubmitXDRSuccess{
			Status: "No Header present!",
		}
		json.NewEncoder(w).Encode(result)
		return
	}

	if r.Header.Get("Content-Type") == "" {
		w.WriteHeader(http.StatusBadRequest)
		result := apiModel.SubmitXDRSuccess{
			Status: "No Content-Type present!",
		}
		json.NewEncoder(w).Encode(result)

		return
	}

	err := json.NewDecoder(r.Body).Decode(&userICOAPI)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		result := apiModel.SubmitXDRSuccess{
			Status: "Error while Decoding the body",
		}
		json.NewEncoder(w).Encode(result)
		fmt.Println(err)
		return
	}
	fmt.Println(userICOAPI)

	display := &builder.AbstractICOBuilder{UserICOAPI: userICOAPI}
	display.BuildUserICOXLM(w, r)
	return
}



/*CreateKnowledge @desc Handles an incoming request and calls the CreateKnowledge
@author - Azeem Ashraf
@params - ResponseWriter,Request
*/
func CreateKnowledge(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var knowledgeAPI model.KnowledgeAPI

	if r.Header == nil {
		w.WriteHeader(http.StatusBadRequest)
		result := apiModel.SubmitXDRSuccess{
			Status: "No Header present!",
		}
		json.NewEncoder(w).Encode(result)
		return
	}

	if r.Header.Get("Content-Type") == "" {
		w.WriteHeader(http.StatusBadRequest)
		result := apiModel.SubmitXDRSuccess{
			Status: "No Content-Type present!",
		}
		json.NewEncoder(w).Encode(result)

		return
	}

	err := json.NewDecoder(r.Body).Decode(&knowledgeAPI)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		result := apiModel.SubmitXDRSuccess{
			Status: "Error while Decoding the body",
		}
		json.NewEncoder(w).Encode(result)
		fmt.Println(err)
		return
	}
	fmt.Println(knowledgeAPI)

	display := &builder.AbstractKnowledgeBuilder{knowledgeAPI}
	display.BuildCreateKnowledge(w, r)
	return
}

/*AddKnowledge @desc Handles an incoming request and calls the CreateKnowledge
@author - Azeem Ashraf
@params - ResponseWriter,Request
*/
func AddKnowledge(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var contributionAPI model.ContributionAPI

	if r.Header == nil {
		w.WriteHeader(http.StatusBadRequest)
		result := apiModel.SubmitXDRSuccess{
			Status: "No Header present!",
		}
		json.NewEncoder(w).Encode(result)
		return
	}

	if r.Header.Get("Content-Type") == "" {
		w.WriteHeader(http.StatusBadRequest)
		result := apiModel.SubmitXDRSuccess{
			Status: "No Content-Type present!",
		}
		json.NewEncoder(w).Encode(result)

		return
	}

	err := json.NewDecoder(r.Body).Decode(&contributionAPI)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		result := apiModel.SubmitXDRSuccess{
			Status: "Error while Decoding the body",
		}
		json.NewEncoder(w).Encode(result)
		fmt.Println(err)
		return
	}
	fmt.Println(contributionAPI)

	display := &builder.AbstractContributionBuilder{contributionAPI}
	display.BuildAddKnowledge(w, r)
	return
}
type Transuc struct {
	TXN string `json:"txn"`
}

type TranXDR struct {
	XDR string `json:"XDR"`
}

/*ConvertXDRToTXN - Test Endpoint @desc Handles an incoming request and Returns the TXN Hash for teh XDR Provided
@author - Azeem Ashraf
@params - ResponseWriter,Request
*/
func ConvertXDRToTXN(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var Trans xdr.Transaction
	// var lol string

	var TDP TranXDR
	// object := dao.Connection{}
	// var copy model.TransactionCollectionBody

	if r.Header == nil {
		w.WriteHeader(http.StatusBadRequest)
		result := apiModel.SubmitXDRSuccess{
			Status: "No Header present!",
		}
		json.NewEncoder(w).Encode(result)

		return
	}

	if r.Header.Get("Content-Type") == "" {
		w.WriteHeader(http.StatusBadRequest)
		result := apiModel.SubmitXDRSuccess{
			Status: "No Content-Type present!",
		}
		json.NewEncoder(w).Encode(result)

		return
	}

	// fmt.Println(TDP)
	err := json.NewDecoder(r.Body).Decode(&TDP)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		result := apiModel.SubmitXDRSuccess{
			Status: "Error while Decoding the body",
		}
		json.NewEncoder(w).Encode(result)
		fmt.Println(err)
		return
	}
	fmt.Println(TDP)

	err1 := xdr.SafeUnmarshalBase64(TDP.XDR, &Trans)
	if err1 != nil {
		fmt.Println(err1)
	}

	brr := build.TransactionBuilder{TX: &Trans, NetworkPassphrase: build.TestNetwork.Passphrase}
	fmt.Println(build.TestNetwork.Passphrase)
	// fmt.Println(brr.Hash())
	t, _ := brr.Hash()
	test := fmt.Sprintf("%x", t)

	w.WriteHeader(http.StatusOK)
	response := Transuc{TXN: test}
	json.NewEncoder(w).Encode(response)
	return

}



type TDP struct {
	TdpId string `json:"tdpId"`
}

/*TDPForTXN - Test Endpoint @desc Handles an incoming request and Returns the TDP ID for the TXN Provided.
@author - Azeem Ashraf
@params - ResponseWriter,Request
*/
func TDPForTXN(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	vars := mux.Vars(r)

	object := dao.Connection{}
	p := object.GetTdpIdForTransaction(vars["Txn"])
	p.Then(func(data interface{}) interface{} {

		result := data.(model.TransactionCollectionBody)

		res := TDP{TdpId: result.TdpId}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)
		return nil
	}).Catch(func(error error) error {
		w.WriteHeader(http.StatusBadRequest)
		response := model.Error{Message: "TdpId Not Found in Gateway DataStore"}
		json.NewEncoder(w).Encode(response)
		return error
	})
	p.Await()

}