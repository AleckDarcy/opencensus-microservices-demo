package hipstershop

import (
	"github.com/AleckDarcy/reload/core/tracer"
)

// Cart Service
func (m *AddItemRequest) SetFI_Trace(trace *tracer.Trace) {
	m.FI_Trace = trace
}

func (m *AddItemRequest) GetFI_Name() string {
	return "AddItemRequest"
}

func (m *AddItemRequest) GetMessageType() tracer.MessageType {
	return tracer.MessageType_Message_Request
}

func (m *EmptyCartRequest) SetFI_Trace(trace *tracer.Trace) {
	m.FI_Trace = trace
}

func (m *EmptyCartRequest) GetFI_Name() string {
	return "EmptyCartRequest"
}

func (m *EmptyCartRequest) GetMessageType() tracer.MessageType {
	return tracer.MessageType_Message_Request
}

func (m *GetCartRequest) SetFI_Trace(trace *tracer.Trace) {
	m.FI_Trace = trace
}

func (m *GetCartRequest) GetFI_Name() string {
	return "GetCartRequest"
}

func (m *GetCartRequest) GetMessageType() tracer.MessageType {
	return tracer.MessageType_Message_Request
}

func (m *Cart) SetFI_Trace(trace *tracer.Trace) {
	m.FI_Trace = trace
}

func (m *Cart) GetFI_Name() string {
	return "Cart"
}

func (m *Cart) GetMessageType() tracer.MessageType {
	return tracer.MessageType_Message_Response
}

// Recommendation Service
func (m *ListRecommendationsRequest) SetFI_Trace(trace *tracer.Trace) {
	m.FI_Trace = trace
}

func (m *ListRecommendationsRequest) GetFI_Name() string {
	return "ListRecommendationsRequest"
}

func (m *ListRecommendationsRequest) GetMessageType() tracer.MessageType {
	return tracer.MessageType_Message_Request
}

func (m *ListRecommendationsResponse) SetFI_Trace(trace *tracer.Trace) {
	m.FI_Trace = trace
}

func (m *ListRecommendationsResponse) GetFI_Name() string {
	return "ListRecommendationsResponse"
}

func (m *ListRecommendationsResponse) GetMessageType() tracer.MessageType {
	return tracer.MessageType_Message_Response
}

// Product Catalog Service
func (m *ListProductsRequest) SetFI_Trace(trace *tracer.Trace) {
	m.FI_Trace = trace
}

func (m *ListProductsRequest) GetFI_Name() string {
	return "ListProductsRequest"
}

func (m *ListProductsRequest) GetMessageType() tracer.MessageType {
	return tracer.MessageType_Message_Request
}

func (m *Product) SetFI_Trace(trace *tracer.Trace) {
	m.FI_Trace = trace
}

func (m *Product) GetFI_Name() string {
	return "Product"
}

func (m *Product) GetMessageType() tracer.MessageType {
	return tracer.MessageType_Message_Response
}

func (m *ListProductsResponse) SetFI_Trace(trace *tracer.Trace) {
	m.FI_Trace = trace
}

func (m *ListProductsResponse) GetFI_Name() string {
	return "ListProductsResponse"
}

func (m *ListProductsResponse) GetMessageType() tracer.MessageType {
	return tracer.MessageType_Message_Response
}

func (m *GetProductRequest) SetFI_Trace(trace *tracer.Trace) {
	m.FI_Trace = trace
}

func (m *GetProductRequest) GetFI_Name() string {
	return "GetProductRequest"
}

func (m *GetProductRequest) GetMessageType() tracer.MessageType {
	return tracer.MessageType_Message_Request
}

func (m *SearchProductsRequest) SetFI_Trace(trace *tracer.Trace) {
	m.FI_Trace = trace
}

func (m *SearchProductsRequest) GetFI_Name() string {
	return "SearchProductsRequest"
}

func (m *SearchProductsRequest) GetMessageType() tracer.MessageType {
	return tracer.MessageType_Message_Request
}

func (m *SearchProductsResponse) SetFI_Trace(trace *tracer.Trace) {
	m.FI_Trace = trace
}

func (m *SearchProductsResponse) GetFI_Name() string {
	return "SearchProductsResponse"
}

func (m *SearchProductsResponse) GetMessageType() tracer.MessageType {
	return tracer.MessageType_Message_Response
}

// Shipping Service
func (m *GetQuoteRequest) SetFI_Trace(trace *tracer.Trace) {
	m.FI_Trace = trace
}

func (m *GetQuoteRequest) GetFI_Name() string {
	return "GetQuoteRequest"
}

func (m *GetQuoteRequest) GetMessageType() tracer.MessageType {
	return tracer.MessageType_Message_Request
}

func (m *GetQuoteResponse) SetFI_Trace(trace *tracer.Trace) {
	m.FI_Trace = trace
}

func (m *GetQuoteResponse) GetFI_Name() string {
	return "GetQuoteResponse"
}

func (m *GetQuoteResponse) GetMessageType() tracer.MessageType {
	return tracer.MessageType_Message_Response
}
func (m *ShipOrderRequest) SetFI_Trace(trace *tracer.Trace) {
	m.FI_Trace = trace
}

func (m *ShipOrderRequest) GetFI_Name() string {
	return "ShipOrderRequest"
}

func (m *ShipOrderRequest) GetMessageType() tracer.MessageType {
	return tracer.MessageType_Message_Request
}

func (m *ShipOrderResponse) SetFI_Trace(trace *tracer.Trace) {
	m.FI_Trace = trace
}

func (m *ShipOrderResponse) GetFI_Name() string {
	return "ShipOrderResponse"
}

func (m *ShipOrderResponse) GetMessageType() tracer.MessageType {
	return tracer.MessageType_Message_Response
}

// Currency Service
func (m *GetSupportedCurrenciesRequest) SetFI_Trace(trace *tracer.Trace) {
	m.FI_Trace = trace
}

func (m *GetSupportedCurrenciesRequest) GetFI_Name() string {
	return "GetSupportedCurrenciesRequest"
}

func (m *GetSupportedCurrenciesRequest) GetMessageType() tracer.MessageType {
	return tracer.MessageType_Message_Request
}

func (m *Money) SetFI_Trace(trace *tracer.Trace) {
	m.FI_Trace = trace
}

func (m *Money) GetFI_Name() string {
	return "Money"
}

func (m *Money) GetMessageType() tracer.MessageType {
	return tracer.MessageType_Message_Response
}

func (m *GetSupportedCurrenciesResponse) SetFI_Trace(trace *tracer.Trace) {
	m.FI_Trace = trace
}

func (m *GetSupportedCurrenciesResponse) GetFI_Name() string {
	return "GetSupportedCurrenciesResponse"
}

func (m *GetSupportedCurrenciesResponse) GetMessageType() tracer.MessageType {
	return tracer.MessageType_Message_Response
}

func (m *CurrencyConversionRequest) SetFI_Trace(trace *tracer.Trace) {
	m.FI_Trace = trace
}

func (m *CurrencyConversionRequest) GetFI_Name() string {
	return "CurrencyConversionRequest"
}

func (m *CurrencyConversionRequest) GetMessageType() tracer.MessageType {
	return tracer.MessageType_Message_Request
}

// Payment Service
func (m *ChargeRequest) SetFI_Trace(trace *tracer.Trace) {
	m.FI_Trace = trace
}

func (m *ChargeRequest) GetFI_Name() string {
	return "ChargeRequest"
}

func (m *ChargeRequest) GetMessageType() tracer.MessageType {
	return tracer.MessageType_Message_Request
}

func (m *ChargeResponse) SetFI_Trace(trace *tracer.Trace) {
	m.FI_Trace = trace
}

func (m *ChargeResponse) GetFI_Name() string {
	return "ChargeResponse"
}

func (m *ChargeResponse) GetMessageType() tracer.MessageType {
	return tracer.MessageType_Message_Response
}

// Email Service
func (m *SendOrderConfirmationRequest) SetFI_Trace(trace *tracer.Trace) {
	m.FI_Trace = trace
}

func (m *SendOrderConfirmationRequest) GetFI_Name() string {
	return "SendOrderConfirmationRequest"
}

func (m *SendOrderConfirmationRequest) GetMessageType() tracer.MessageType {
	return tracer.MessageType_Message_Request
}

// Checkout Service
func (m *PlaceOrderRequest) SetFI_Trace(trace *tracer.Trace) {
	m.FI_Trace = trace
}

func (m *PlaceOrderRequest) GetFI_Name() string {
	return "PlaceOrderRequest"
}

func (m *PlaceOrderRequest) GetMessageType() tracer.MessageType {
	return tracer.MessageType_Message_Request
}

func (m *PlaceOrderResponse) SetFI_Trace(trace *tracer.Trace) {
	m.FI_Trace = trace
}

func (m *PlaceOrderResponse) GetFI_Name() string {
	return "PlaceOrderResponse"
}

func (m *PlaceOrderResponse) GetMessageType() tracer.MessageType {
	return tracer.MessageType_Message_Response
}

// Ad Service
func (m *AdRequest) SetFI_Trace(trace *tracer.Trace) {
	m.FI_Trace = trace
}

func (m *AdRequest) GetFI_Name() string {
	return "AdRequest"
}

func (m *AdRequest) GetMessageType() tracer.MessageType {
	return tracer.MessageType_Message_Request
}

func (m *AdResponse) SetFI_Trace(trace *tracer.Trace) {
	m.FI_Trace = trace
}

func (m *AdResponse) GetFI_Name() string {
	return "AdResponse"
}

func (m *AdResponse) GetMessageType() tracer.MessageType {
	return tracer.MessageType_Message_Response
}
