package routing

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
	"text/template"
	"io/ioutil"
	"encoding/json"
	"time"
	. "../interfaces"
)

func universalRouter(w http.ResponseWriter, page string, data interface{}) {
	t, err := template.ParseFiles(
		"./templates/"+page+".html",
		"./templates/header.html",
		"./templates/footer.html")

	if err != nil {
		fmt.Println(err.Error())
	}

	t.ExecuteTemplate(w, page, data)
}

func IndexHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) { universalRouter(w, "index", nil) }
func BlogHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params)  { universalRouter(w, "blog", nil) }
func ShopHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params)  { universalRouter(w, "shop", nil) }

func BankHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	apiURL := "https://api.privatbank.ua/p24api/exchange_rates?json&date="+time.Now().AddDate(0, 0, -7).Format("02.01.2006")

	data := new(PrivatBankExchangeRates)
	getJson(apiURL, &data)

	EURO := Rates{In: data.ExchangeRate[17].SaleRateNB, Out: data.ExchangeRate[17].PurchaseRateNB}
	USD := Rates{In: data.ExchangeRate[15].SaleRateNB, Out: data.ExchangeRate[15].PurchaseRateNB}
	RUB := Rates{In: data.ExchangeRate[13].SaleRateNB, Out: data.ExchangeRate[13].PurchaseRateNB}

	ExchangeRates := Currency{Usd: USD, Euro: EURO, Rub: RUB}

	universalRouter(w, "bank", ExchangeRates)
}

func KriptHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) { universalRouter(w, "kript", nil) }
func ErrorHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) { universalRouter(w, "error", nil) }
func SearchInDb(w http.ResponseWriter, r *http.Request, ps httprouter.Params) { universalRouter(w, "search", nil) }

func getJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}

	defer r.Body.Close()

	data, errIoutil := ioutil.ReadAll(r.Body)
	if errIoutil != nil {
		return errIoutil
	}

	errUnmarshal := json.Unmarshal(data, target)
	if errUnmarshal != nil {
		return errUnmarshal
	}

	return nil
}