package sbi

type NfType int32

const (
    NfType_NULL NfType = iota
    NfType_NRF
    NfType_UDM
    NfType_AMF
    NfType_SMF
    NfType_AUSF
    NfType_NEF
    NfType_PCF
    NfType_NSSF
    NfType_UDR
    NfType_LMF
    NfType_GMLC
    NfType_EIR
    NfType_UPF
    NfType_AF
    NfType_BSF
)
