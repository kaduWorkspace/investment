package external

type BcbI interface {
    GetSelic() (float64, error)
    GetMediaIpca() (float64, error)
}
