// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package NotifyMonitoringReportRequest

import "fmt"
import "encoding/json"
import "reflect"

// UnmarshalJSON implements json.Unmarshaler.
func (j *VariableType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["name"]; !ok || v == nil {
		return fmt.Errorf("field name in VariableType: required")
	}
	type Plain VariableType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = VariableType(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId in CustomDataType: required")
	}
	type Plain CustomDataType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MonitoringDataType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["component"]; !ok || v == nil {
		return fmt.Errorf("field component in MonitoringDataType: required")
	}
	if v, ok := raw["variable"]; !ok || v == nil {
		return fmt.Errorf("field variable in MonitoringDataType: required")
	}
	if v, ok := raw["variableMonitoring"]; !ok || v == nil {
		return fmt.Errorf("field variableMonitoring in MonitoringDataType: required")
	}
	type Plain MonitoringDataType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if len(plain.VariableMonitoring) < 1 {
		return fmt.Errorf("field %s length: must be >= %d", "variableMonitoring", 1)
	}
	*j = MonitoringDataType(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *EVSEType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["id"]; !ok || v == nil {
		return fmt.Errorf("field id in EVSEType: required")
	}
	type Plain EVSEType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = EVSEType(plain)
	return nil
}

// A physical or logical component
//
type ComponentType struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty"`

	// Evse corresponds to the JSON schema field "evse".
	Evse *EVSEType `json:"evse,omitempty" yaml:"evse,omitempty"`

	// Name of instance in case the component exists as multiple instances. Case
	// Insensitive. strongly advised to use Camel Case.
	//
	Instance *string `json:"instance,omitempty" yaml:"instance,omitempty"`

	// Name of the component. Name should be taken from the list of standardized
	// component names whenever possible. Case Insensitive. strongly advised to use
	// Camel Case.
	//
	Name string `json:"name" yaml:"name"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ComponentType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["name"]; !ok || v == nil {
		return fmt.Errorf("field name in ComponentType: required")
	}
	type Plain ComponentType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ComponentType(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *VariableMonitoringType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["id"]; !ok || v == nil {
		return fmt.Errorf("field id in VariableMonitoringType: required")
	}
	if v, ok := raw["severity"]; !ok || v == nil {
		return fmt.Errorf("field severity in VariableMonitoringType: required")
	}
	if v, ok := raw["transaction"]; !ok || v == nil {
		return fmt.Errorf("field transaction in VariableMonitoringType: required")
	}
	if v, ok := raw["type"]; !ok || v == nil {
		return fmt.Errorf("field type in VariableMonitoringType: required")
	}
	if v, ok := raw["value"]; !ok || v == nil {
		return fmt.Errorf("field value in VariableMonitoringType: required")
	}
	type Plain VariableMonitoringType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = VariableMonitoringType(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MonitorEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_MonitorEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_MonitorEnumType_1, v)
	}
	*j = MonitorEnumType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MonitorEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_MonitorEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_MonitorEnumType, v)
	}
	*j = MonitorEnumType(v)
	return nil
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId" yaml:"vendorId"`
}

// EVSE
// urn:x-oca:ocpp:uid:2:233123
// Electric Vehicle Supply Equipment
//
type EVSEType struct {
	// An id to designate a specific connector (on an EVSE) by connector index number.
	//
	ConnectorId *int `json:"connectorId,omitempty" yaml:"connectorId,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty"`

	// Identified_ Object. MRID. Numeric_ Identifier
	// urn:x-enexis:ecdm:uid:1:569198
	// EVSE Identifier. This contains a number (&gt; 0) designating an EVSE of the
	// Charging Station.
	//
	Id int `json:"id" yaml:"id"`
}

type MonitorEnumType string

const MonitorEnumTypeDelta MonitorEnumType = "Delta"
const MonitorEnumTypeLowerThreshold MonitorEnumType = "LowerThreshold"
const MonitorEnumTypePeriodic MonitorEnumType = "Periodic"
const MonitorEnumTypePeriodicClockAligned MonitorEnumType = "PeriodicClockAligned"
const MonitorEnumTypeUpperThreshold MonitorEnumType = "UpperThreshold"

type MonitorEnumType_1 string

const MonitorEnumType_1_Delta MonitorEnumType_1 = "Delta"
const MonitorEnumType_1_LowerThreshold MonitorEnumType_1 = "LowerThreshold"
const MonitorEnumType_1_Periodic MonitorEnumType_1 = "Periodic"
const MonitorEnumType_1_PeriodicClockAligned MonitorEnumType_1 = "PeriodicClockAligned"
const MonitorEnumType_1_UpperThreshold MonitorEnumType_1 = "UpperThreshold"

// Class to hold parameters of SetVariableMonitoring request.
//
type MonitoringDataType struct {
	// Component corresponds to the JSON schema field "component".
	Component ComponentType `json:"component" yaml:"component"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty"`

	// Variable corresponds to the JSON schema field "variable".
	Variable VariableType `json:"variable" yaml:"variable"`

	// VariableMonitoring corresponds to the JSON schema field "variableMonitoring".
	VariableMonitoring []VariableMonitoringType `json:"variableMonitoring" yaml:"variableMonitoring"`
}

type NotifyMonitoringReportRequestJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty"`

	// Timestamp of the moment this message was generated at the Charging Station.
	//
	GeneratedAt string `json:"generatedAt" yaml:"generatedAt"`

	// Monitor corresponds to the JSON schema field "monitor".
	Monitor []MonitoringDataType `json:"monitor,omitempty" yaml:"monitor,omitempty"`

	// The id of the GetMonitoringRequest that requested this report.
	//
	//
	RequestId int `json:"requestId" yaml:"requestId"`

	// Sequence number of this message. First message starts at 0.
	//
	SeqNo int `json:"seqNo" yaml:"seqNo"`

	// “to be continued” indicator. Indicates whether another part of the
	// monitoringData follows in an upcoming notifyMonitoringReportRequest message.
	// Default value when omitted is false.
	//
	Tbc bool `json:"tbc,omitempty" yaml:"tbc,omitempty"`
}

// A monitoring setting for a variable.
//
type VariableMonitoringType struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty"`

	// Identifies the monitor.
	//
	Id int `json:"id" yaml:"id"`

	// The severity that will be assigned to an event that is triggered by this
	// monitor. The severity range is 0-9, with 0 as the highest and 9 as the lowest
	// severity level.
	//
	// The severity levels have the following meaning: +
	// *0-Danger* +
	// Indicates lives are potentially in danger. Urgent attention is needed and
	// action should be taken immediately. +
	// *1-Hardware Failure* +
	// Indicates that the Charging Station is unable to continue regular operations
	// due to Hardware issues. Action is required. +
	// *2-System Failure* +
	// Indicates that the Charging Station is unable to continue regular operations
	// due to software or minor hardware issues. Action is required. +
	// *3-Critical* +
	// Indicates a critical error. Action is required. +
	// *4-Error* +
	// Indicates a non-urgent error. Action is required. +
	// *5-Alert* +
	// Indicates an alert event. Default severity for any type of monitoring event.  +
	// *6-Warning* +
	// Indicates a warning event. Action may be required. +
	// *7-Notice* +
	// Indicates an unusual event. No immediate action is required. +
	// *8-Informational* +
	// Indicates a regular operational event. May be used for reporting, measuring
	// throughput, etc. No action is required. +
	// *9-Debug* +
	// Indicates information useful to developers for debugging, not useful during
	// operations.
	//
	Severity int `json:"severity" yaml:"severity"`

	// Monitor only active when a transaction is ongoing on a component relevant to
	// this transaction.
	//
	Transaction bool `json:"transaction" yaml:"transaction"`

	// Type corresponds to the JSON schema field "type".
	Type MonitorEnumType_1 `json:"type" yaml:"type"`

	// Value for threshold or delta monitoring.
	// For Periodic or PeriodicClockAligned this is the interval in seconds.
	//
	Value float64 `json:"value" yaml:"value"`
}

// Reference key to a component-variable.
//
type VariableType struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty" yaml:"customData,omitempty"`

	// Name of instance in case the variable exists as multiple instances. Case
	// Insensitive. strongly advised to use Camel Case.
	//
	Instance *string `json:"instance,omitempty" yaml:"instance,omitempty"`

	// Name of the variable. Name should be taken from the list of standardized
	// variable names whenever possible. Case Insensitive. strongly advised to use
	// Camel Case.
	//
	Name string `json:"name" yaml:"name"`
}

var enumValues_MonitorEnumType = []interface{}{
	"UpperThreshold",
	"LowerThreshold",
	"Delta",
	"Periodic",
	"PeriodicClockAligned",
}
var enumValues_MonitorEnumType_1 = []interface{}{
	"UpperThreshold",
	"LowerThreshold",
	"Delta",
	"Periodic",
	"PeriodicClockAligned",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *NotifyMonitoringReportRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["generatedAt"]; !ok || v == nil {
		return fmt.Errorf("field generatedAt in NotifyMonitoringReportRequestJson: required")
	}
	if v, ok := raw["requestId"]; !ok || v == nil {
		return fmt.Errorf("field requestId in NotifyMonitoringReportRequestJson: required")
	}
	if v, ok := raw["seqNo"]; !ok || v == nil {
		return fmt.Errorf("field seqNo in NotifyMonitoringReportRequestJson: required")
	}
	type Plain NotifyMonitoringReportRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if len(plain.Monitor) < 1 {
		return fmt.Errorf("field %s length: must be >= %d", "monitor", 1)
	}
	if v, ok := raw["tbc"]; !ok || v == nil {
		plain.Tbc = false
	}
	*j = NotifyMonitoringReportRequestJson(plain)
	return nil
}
