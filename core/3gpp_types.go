package core

type Ng_plmn_id_t struct {
    mcc2 uint8
    mcc1 uint8
    mnc1 uint8
    mcc3 uint8
    mnc3 uint8
    mnc2 uint8
}

type Ng_s_nssai_t struct {
    sst uint8
    sd uint24
}
