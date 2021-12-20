package foo

import (
	"errors"
	"github.com/Rhymond/go-money"
	gql "github.com/graphql-go/graphql"
	"golang.org/x/text/currency"
)

/*
	Using custom types requires implementing:
		<Name>FromArgs(field map[string]interface{}) *<Name>
		MoneyGraphqlInputType
		MoneyGraphqlType
*/

func MoneyFromArgs(field map[string]interface{}) *Money {
	mon := Money{
		CurrencyCode: field["currencyCode"].(string),
		Units:        field["currencyCode"].(int64),
	}
	return &mon
}

var MoneyGraphqlInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "MoneyInput",
	Fields: gql.InputObjectConfigFieldMap{
		"currencyCode": &gql.InputObjectFieldConfig{
			Type:         gql.String,
			Description:  "The three-letter currency code defined in ISO 4217.",
			DefaultValue: "GBP",
		},
		"units": &gql.InputObjectFieldConfig{
			Type:        gql.NewNonNull(gql.Int),
			Description: "The whole units of the amount. For example if `currencyCode` is `GBP`, then 1 unit is one UK penny.",
		},
	},
})

var MoneyGraphqlType = gql.NewObject(gql.ObjectConfig{
	Name: "Money",
	Fields: gql.Fields{
		"currencyCode": &gql.Field{
			Type:        gql.String,
			Description: "The three-letter currency code defined in ISO 4217.",
			Resolve: func(p gql.ResolveParams) (interface{}, error) {
				return p.Source.(*Money).CurrencyCode, nil
			},
		},
		"units": &gql.Field{
			Type:        gql.Int,
			Description: "The smallest unit for the given currency code. For example if `currencyCode` is `GBP`, then 1 unit is one UK penny.",
			Resolve: func(p gql.ResolveParams) (interface{}, error) {
				return p.Source.(*Money).Units, nil
			},
		},
		"symbol": &gql.Field{
			Type:        gql.String,
			Description: "The currency symbol associated with the currencyCode",
			Resolve: func(p gql.ResolveParams) (interface{}, error) {
				return currency.ParseISO(p.Source.(*Money).CurrencyCode)
			},
		},
		"format": &gql.Field{
			Description: `https://github.com/leekchan/accounting`,
			Type:        gql.String,
			Resolve: func(p gql.ResolveParams) (interface{}, error) {
				currencyCode := "GBP"
				if p.Source.(*Money).CurrencyCode != "GBP" {
					currencyCode = p.Source.(*Money).CurrencyCode
				}

				currency := money.GetCurrency(currencyCode)
				if currency == nil {
					return nil, errors.New("invalid currency code")
				}

				units := p.Source.(*Money).Units

				value := money.New(units, currency.Code)
				return value.Display(), nil
			},
		},
	},
})
