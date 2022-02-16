package sbi

import (
   core "github.com/micjerry/nglib/core"
)

const (
    OGS_SBI_HTTP_PORT                          int    = 80
    OGS_SBI_HTTPS_PORT                         int    = 443
    OGS_SBI_HTTP_SCHEME                        string = "http"
    OGS_SBI_HTTPS_SCHEME                       string = "https"
    OGS_SBI_HTTP_STATUS_OK                     int    = 200
    OGS_SBI_HTTP_STATUS_CREATED                int    = 201
    OGS_SBI_HTTP_STATUS_ACCEPTED               int    = 202
    OGS_SBI_HTTP_STATUS_NO_CONTENT             int    = 204
    OGS_SBI_HTTP_STATUS_SEE_OTHER              int    = 303
    OGS_SBI_HTTP_STATUS_TEMPORARY_REDIRECT     int    = 307
)

type Http_t struct {
    accept string
    content_encoding string
    content_type string
    location string
    cache_control string
}

type Ng_sbi_message_t struct {
    http Http_t
    target_nf_type    NfType
    requester_nf_type NfType
    nf_id string
    nf_type NfType
    limit int
    dnn string
    
    plmn_id core.Ng_plmn_id_t
    s_nssai core.Ng_s_nssai_t
    plmn_id_presence bool
    single_nssai_presence bool
    snssai_presence bool
    slice_info_request_for_pdu_session_presence bool
}
