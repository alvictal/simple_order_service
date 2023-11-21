package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/* To execute the tests just insert in cmd "go test ./"
There are two ways to work with testing, the first its a old
way that required you work direct with the results expected and if statements
another way its using a packages that reduce the number of lines writed

See both examples below
*/

// Old way
func Test_If_Gets_An_Error_If_ID_Is_Blank(t *testing.T) {
	order := Order{}
	if order.Validate() == nil {
		t.Error("id is required")
	}
}

// Second way usinh assert package
func Test_if_Gets_An_Error_If_Price_Is_Blank(t *testing.T) {
	order := Order{ID: "123"}
	assert.Error(t, order.Validate(), "price must be greater than zero")
}

func Test_if_Gets_An_Error_If_Tax_Is_Blank(t *testing.T) {
	order := Order{ID: "123", Price: 10, Tax: -1}
	assert.Error(t, order.Validate(), "invalid tax")
}

func Test_Final_Price(t *testing.T) {
	order := Order{ID: "123", Price: 10, Tax: 2}
	assert.NoError(t, order.Validate())
	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 2.0, order.Tax)
	order.CalculateFinalPrice()
	assert.Equal(t, 12.0, order.FinalPrice)
}
