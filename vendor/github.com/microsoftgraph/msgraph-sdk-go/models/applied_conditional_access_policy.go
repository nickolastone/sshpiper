package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AppliedConditionalAccessPolicy 
type AppliedConditionalAccessPolicy struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Name of the conditional access policy.
    displayName *string
    // Refers to the grant controls enforced by the conditional access policy (example: 'Require multi-factor authentication').
    enforcedGrantControls []string
    // Refers to the session controls enforced by the conditional access policy (example: 'Require app enforced controls').
    enforcedSessionControls []string
    // Identifier of the conditional access policy.
    id *string
    // Indicates the result of the CA policy that was triggered. Possible values are: success, failure, notApplied (Policy isn't applied because policy conditions were not met),notEnabled (This is due to the policy in disabled state), unknown, unknownFutureValue, reportOnlySuccess, reportOnlyFailure, reportOnlyNotApplied, reportOnlyInterrupted. Note that you must use the Prefer: include-unknown-enum-members request header to get the following values in this evolvable enum: reportOnlySuccess, reportOnlyFailure, reportOnlyNotApplied, reportOnlyInterrupted.
    result *AppliedConditionalAccessPolicyResult
}
// NewAppliedConditionalAccessPolicy instantiates a new appliedConditionalAccessPolicy and sets the default values.
func NewAppliedConditionalAccessPolicy()(*AppliedConditionalAccessPolicy) {
    m := &AppliedConditionalAccessPolicy{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAppliedConditionalAccessPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAppliedConditionalAccessPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAppliedConditionalAccessPolicy(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AppliedConditionalAccessPolicy) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDisplayName gets the displayName property value. Name of the conditional access policy.
func (m *AppliedConditionalAccessPolicy) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetEnforcedGrantControls gets the enforcedGrantControls property value. Refers to the grant controls enforced by the conditional access policy (example: 'Require multi-factor authentication').
func (m *AppliedConditionalAccessPolicy) GetEnforcedGrantControls()([]string) {
    if m == nil {
        return nil
    } else {
        return m.enforcedGrantControls
    }
}
// GetEnforcedSessionControls gets the enforcedSessionControls property value. Refers to the session controls enforced by the conditional access policy (example: 'Require app enforced controls').
func (m *AppliedConditionalAccessPolicy) GetEnforcedSessionControls()([]string) {
    if m == nil {
        return nil
    } else {
        return m.enforcedSessionControls
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AppliedConditionalAccessPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["enforcedGrantControls"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetEnforcedGrantControls(res)
        }
        return nil
    }
    res["enforcedSessionControls"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetEnforcedSessionControls(res)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["result"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppliedConditionalAccessPolicyResult)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResult(val.(*AppliedConditionalAccessPolicyResult))
        }
        return nil
    }
    return res
}
// GetId gets the id property value. Identifier of the conditional access policy.
func (m *AppliedConditionalAccessPolicy) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
// GetResult gets the result property value. Indicates the result of the CA policy that was triggered. Possible values are: success, failure, notApplied (Policy isn't applied because policy conditions were not met),notEnabled (This is due to the policy in disabled state), unknown, unknownFutureValue, reportOnlySuccess, reportOnlyFailure, reportOnlyNotApplied, reportOnlyInterrupted. Note that you must use the Prefer: include-unknown-enum-members request header to get the following values in this evolvable enum: reportOnlySuccess, reportOnlyFailure, reportOnlyNotApplied, reportOnlyInterrupted.
func (m *AppliedConditionalAccessPolicy) GetResult()(*AppliedConditionalAccessPolicyResult) {
    if m == nil {
        return nil
    } else {
        return m.result
    }
}
// Serialize serializes information the current object
func (m *AppliedConditionalAccessPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetEnforcedGrantControls() != nil {
        err := writer.WriteCollectionOfStringValues("enforcedGrantControls", m.GetEnforcedGrantControls())
        if err != nil {
            return err
        }
    }
    if m.GetEnforcedSessionControls() != nil {
        err := writer.WriteCollectionOfStringValues("enforcedSessionControls", m.GetEnforcedSessionControls())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    if m.GetResult() != nil {
        cast := (*m.GetResult()).String()
        err := writer.WriteStringValue("result", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AppliedConditionalAccessPolicy) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDisplayName sets the displayName property value. Name of the conditional access policy.
func (m *AppliedConditionalAccessPolicy) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetEnforcedGrantControls sets the enforcedGrantControls property value. Refers to the grant controls enforced by the conditional access policy (example: 'Require multi-factor authentication').
func (m *AppliedConditionalAccessPolicy) SetEnforcedGrantControls(value []string)() {
    if m != nil {
        m.enforcedGrantControls = value
    }
}
// SetEnforcedSessionControls sets the enforcedSessionControls property value. Refers to the session controls enforced by the conditional access policy (example: 'Require app enforced controls').
func (m *AppliedConditionalAccessPolicy) SetEnforcedSessionControls(value []string)() {
    if m != nil {
        m.enforcedSessionControls = value
    }
}
// SetId sets the id property value. Identifier of the conditional access policy.
func (m *AppliedConditionalAccessPolicy) SetId(value *string)() {
    if m != nil {
        m.id = value
    }
}
// SetResult sets the result property value. Indicates the result of the CA policy that was triggered. Possible values are: success, failure, notApplied (Policy isn't applied because policy conditions were not met),notEnabled (This is due to the policy in disabled state), unknown, unknownFutureValue, reportOnlySuccess, reportOnlyFailure, reportOnlyNotApplied, reportOnlyInterrupted. Note that you must use the Prefer: include-unknown-enum-members request header to get the following values in this evolvable enum: reportOnlySuccess, reportOnlyFailure, reportOnlyNotApplied, reportOnlyInterrupted.
func (m *AppliedConditionalAccessPolicy) SetResult(value *AppliedConditionalAccessPolicyResult)() {
    if m != nil {
        m.result = value
    }
}
