package infra_external

import (
	"fmt"
	"kaduhod/fin_v3/core/domain/external"
	struct_utils "kaduhod/fin_v3/pkg/utils/struct"
	"strconv"
	"sync"
	"time"
)
type IpcaApiRespopnse struct {
    Data string `json:"data"`
    Valor string `json:"valor"`
}
type BcbService struct {
    mu sync.Mutex
    data struct{
        MediaIpca float64
        IpcaLastYearSearched int
        Selic float64
        SelicLastMonthSearched time.Month
    }
}
func NewBcbService() external.BcbI {
    return &BcbService{}
}
func (s *BcbService) GetSelic() (float64, error) {
    valueSelic := 13.25 // default
    s.mu.Lock()
    defer s.mu.Unlock()
    if s.data.Selic != 0.0 && time.Now().Month() != s.data.SelicLastMonthSearched {
        return s.data.Selic, nil
    }
    result, err := struct_utils.HttpRequest("https://www.bcb.gov.br/api/servico/sitebcb//taxaselic/ultima?withCredentials=true", "GET",
        map[string]string{"content-type":"text/plain"}, "")
    if err != nil {
        fmt.Println(err)
        return valueSelic, err
    }
    var response map[string]interface{}
    err, response = struct_utils.FromJson[map[string]interface{}]([]byte(result))
    if err != nil {
        fmt.Println(err)
        return valueSelic, err

    }
    content, ok := response["conteudo"].([]interface{})
    if !ok || len(content) == 0 {
        fmt.Println(err)
        return valueSelic, err
    }
    firstItem, ok := content[0].(map[string]interface{})
    if !ok {
        fmt.Println(err)
        return valueSelic, err
    }
    if metaSelic, ok := firstItem["MetaSelic"].(float64); ok {
        valueSelic = metaSelic
    }
    s.data.Selic = valueSelic
    s.data.SelicLastMonthSearched = time.Now().Month()
    fmt.Println("Salvando dado novo de media")
    return valueSelic, err
}
func (s *BcbService) GetMediaIpca() (float64, error) {
    s.mu.Lock()
    defer s.mu.Unlock()
    if s.data.MediaIpca != 0.0 && (s.data.IpcaLastYearSearched != time.Now().Year() && time.Now().Month() != time.December) {
        return s.data.MediaIpca, nil
    }
    url := "https://api.bcb.gov.br/dados/serie/bcdata.sgs.13522/dados?formato=json&dataInicial=01/01/2000"
    res, err := struct_utils.HttpRequest(url, "GET", map[string]string{"content-type":"text/plain"}, "")
    defaultIpca := 4.5
    if err != nil {
        fmt.Println(err)
        return defaultIpca, err
    }
    err, parsedResponse := struct_utils.FromJson[[]IpcaApiRespopnse]([]byte(res))
    if err != nil {
        fmt.Println(err)
        return defaultIpca, err
    }
    dateLayout := "02/01/2006"
    years := 0
    ipcaAccrued := 0.0
    ipcas := []float64{}
    lastYear := time.Now().Year()
    for _, data := range parsedResponse {
        date, err := time.Parse(dateLayout, data.Data)
        if err != nil {
            fmt.Println(err)
            return defaultIpca, err
        }
        if date.Month() == time.December {
            years++
            f, err := strconv.ParseFloat(data.Valor, 64)
            if err != nil {
                fmt.Println(err)
                return defaultIpca, err
            }
            ipcas = append(ipcas, f)
            ipcaAccrued += f
            lastYear = date.Year()
        }
    }
    resultMedia := ipcaAccrued / float64(years)
    s.data.MediaIpca = resultMedia
    s.data.IpcaLastYearSearched = lastYear
    fmt.Println("Salvando dado novo de media")
    return resultMedia, nil
}
