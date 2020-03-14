/*
 * Application Manager API
 *
 * Application Manager APIs to control Apache Flink jobs
 *
 * API version: 2.1.0
 * Contact: platform@ververica.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package appmanagerapi

type ObjectNode struct {
	Array               bool   `json:"array,omitempty"`
	BigDecimal          bool   `json:"bigDecimal,omitempty"`
	BigInteger          bool   `json:"bigInteger,omitempty"`
	Binary              bool   `json:"binary,omitempty"`
	Boolean             bool   `json:"boolean,omitempty"`
	ContainerNode       bool   `json:"containerNode,omitempty"`
	Double              bool   `json:"double,omitempty"`
	Empty               bool   `json:"empty,omitempty"`
	Float               bool   `json:"float,omitempty"`
	FloatingPointNumber bool   `json:"floatingPointNumber,omitempty"`
	Int_                bool   `json:"int,omitempty"`
	IntegralNumber      bool   `json:"integralNumber,omitempty"`
	Long                bool   `json:"long,omitempty"`
	MissingNode         bool   `json:"missingNode,omitempty"`
	NodeType            string `json:"nodeType,omitempty"`
	Null                bool   `json:"null,omitempty"`
	Number              bool   `json:"number,omitempty"`
	Object              bool   `json:"object,omitempty"`
	Pojo                bool   `json:"pojo,omitempty"`
	Short               bool   `json:"short,omitempty"`
	Textual             bool   `json:"textual,omitempty"`
	ValueNode           bool   `json:"valueNode,omitempty"`
}
